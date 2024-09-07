/*
 * Echotron
 * Copyright (C) 2018-2022 The Echotron Devs
 *
 * Echotron is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Echotron is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron

import (
	"time"
	"sync/atomic"
	"sync"
	"golang.org/x/sync/semaphore"
	"context"
	"fmt"
	"runtime"
	"errors"
)

type Bot interface {
	Update(*Update)
}

type NewBotFn func(chatId int64) Bot

type Dispatcher struct {
	api        API
	newBot     NewBotFn
	sessions   *sessionManager
	options    dispatcherOptions

	inShutdown atomic.Bool
	mu         sync.Mutex
	ctx           context.Context
	cancel        context.CancelFunc
	isRunning    atomic.Bool
	updateChan chan *Update
	wg         sync.WaitGroup
}

type dispatcherOptions struct {
	pollInterval     time.Duration
	updateBufferSize int
	sessionTTL       time.Duration
	allowedUpdates   []UpdateType
	errorHandler     func(error)
	workerSem  *semaphore.Weighted
}

type Option func(*dispatcherOptions) error

func WithPollInterval(interval time.Duration) Option {
	return func(o *dispatcherOptions) error {
		if interval <= 0 {
			return errors.New("poll interval must be positive")
		}
		o.pollInterval = interval
		return nil
	}
}

func WithUpdateBufferSize(size int) Option {
	return func(o *dispatcherOptions) error {
		if size <= 0 {
			return errors.New("update buffer size must be positive")
		}
		o.updateBufferSize = size
		return nil
	}
}

func WithSessionTTL(ttl time.Duration) Option {
	return func(o *dispatcherOptions) error {
		if ttl <= 0 {
			return errors.New("session TTL must be positive")
		}
		o.sessionTTL = ttl
		return nil
	}
}

func WithAllowedUpdates(types []UpdateType) Option {
	return func(o *dispatcherOptions) error {
		o.allowedUpdates = types
		return nil
	}
}

func WithErrorHandler(handler func(error)) Option {
	return func(o *dispatcherOptions) error {
		if handler == nil {
			return errors.New("error handler cannot be nil")
		}
		o.errorHandler = handler
		return nil
	}
}

func WithMaxConcurrentUpdates(max int64) Option {
	return func(o *dispatcherOptions) error {
		if max <= 0 {
			return errors.New("max concurrent updates must be positive")
		}
		o.workerSem = semaphore.NewWeighted(max)
		return nil
	}
}

func NewDispatcher(token string, newBot NewBotFn, opts ...Option) (*Dispatcher, error) {
	options := dispatcherOptions{
		pollInterval:     5 * time.Second,
		updateBufferSize: 1000,
		sessionTTL:       24 * time.Hour,
		allowedUpdates:   []UpdateType{},
		errorHandler:     func(err error) { fmt.Println("Dispatcher error:", err) },
		workerSem:  semaphore.NewWeighted(int64(runtime.NumCPU() * 2)),
	}
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}
	d := &Dispatcher{
		api:     NewAPI(token),
		newBot:  newBot,
		options: options,
		sessions: newSessionManager(options.sessionTTL),
		updateChan: make(chan *Update, options.updateBufferSize),
	}
	return d, nil
}

// Start begins the Dispatcher's operations.
func (d *Dispatcher) Start() error {

	d.mu.Lock()
	defer d.mu.Unlock()

	if d.inShutdown.Load() {
		return errors.New("dispatcher is shutting down")
	}

	if !d.isRunning.CompareAndSwap(false, true) {
		return errors.New("dispatcher is already running")
	}
	d.ctx, d.cancel = context.WithCancel(context.Background())

	// Start the update polling goroutine
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		if err := d.pollUpdates(); err != nil {
			d.options.errorHandler(fmt.Errorf("polling updates: %w", err))
		}
	}()

	// Start the single worker goroutine
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		d.worker()
	}()

	// Start the session cleaner
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		if err := d.sessions.clean(d.ctx); err != nil {
			d.options.errorHandler(fmt.Errorf("cleaning sessions: %w", err))

		}
	}()

	return nil
}

func (d *Dispatcher) close() error {
	d.mu.Lock()
	if d.inShutdown.Swap(true) {
		d.mu.Unlock()
		return errors.New("dispatcher is already shutting down")
	}
	d.cancel()
	close(d.updateChan)
	d.mu.Unlock()
	d.wg.Wait()
	return nil
}

func (d *Dispatcher) Stop(ctx context.Context) error {
	if !d.isRunning.CompareAndSwap(true, false) {
		return errors.New("dispatcher is not running")
	}

	closeErr := d.close()
	if closeErr != nil {
		return closeErr
	}

	select {
	case <-d.ctx.Done():
		return nil
	}
}

func (d *Dispatcher) worker() {
	for {
		select {
		case <-d.ctx.Done():
			return
		case update, ok := <-d.updateChan:
			if !ok {
				return
			}
			if d.inShutdown.Load() {
				// Discard updates during shutdown
				continue
			}
			if err := d.options.workerSem.Acquire(d.ctx, 1); err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				d.options.errorHandler(fmt.Errorf("failed to acquire semaphore: %w", err))
				continue
			}
			d.wg.Add(1)
			go func(update *Update) {
				defer d.options.workerSem.Release(1)
				defer d.wg.Done()
				if err := d.processUpdate(update); err != nil {
					d.options.errorHandler(fmt.Errorf("processing update: %w", err))
				}
			}(update)
		}
	}
}

func (d *Dispatcher) processUpdate(update *Update) error {
	bot, err := d.sessions.getSession(update.ChatID(), d.newBot)
	if err != nil {
		return fmt.Errorf("getting bot instance: %w", err)
	}
	bot.Update(update)
	return nil
}

// Polls Telegram for updates and queues them for processing.
func (d *Dispatcher) pollUpdates() error {
	if _, err := d.api.DeleteWebhook(true); err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}

	opts := UpdateOptions{
		Timeout:        int(d.options.pollInterval.Seconds()),
		Limit:          100,
		Offset:         0,
		AllowedUpdates: d.options.allowedUpdates,
	}

	ticker := time.NewTicker(d.options.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-d.ctx.Done():
			return nil
		case <-ticker.C:
			if d.inShutdown.Load() {
				return nil
			}
			updates, err := d.api.GetUpdates(&opts)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return nil
				}
				d.options.errorHandler(fmt.Errorf("failed to get updates: %w", err))
				continue
			}

			if len(updates.Result) > 0 {
				lastUpdateID := updates.Result[len(updates.Result)-1].ID
				opts.Offset = lastUpdateID + 1

				for _, update := range updates.Result {
					select {
					case <-d.ctx.Done():
						return nil
					case d.updateChan <- update:
						// Update sent successfully
					default:
						// Channel is full, log an error or handle it appropriately
						d.options.errorHandler(fmt.Errorf("update channel is full, discarding update"))
					}
				}
			}
		}
	}
}

// Session Manager
type sessionManager struct {
	sessions sync.Map
	locks    sync.Map
	ttl      time.Duration
}

type session struct {
	bot          Bot
	lastAccessed atomic.Int64
}

// Create session manager with ttl option.
func newSessionManager(ttl time.Duration) *sessionManager {
	return &sessionManager{ttl: ttl}
}

// GetSession retrieves an existing session or creates a new one if it doesn't exist.
func (sm *sessionManager) getSession(chatID int64, newBot NewBotFn) (Bot, error) {

	// Get or create a lock for this chatID
	lock, _ := sm.locks.LoadOrStore(chatID, &sync.Mutex{})
	mutex := lock.(*sync.Mutex)
	// Lock to ensure only one goroutine accesses or creates a session for this chatID
	mutex.Lock()
	defer mutex.Unlock()

	// Check if a session exists (now inside the lock)
	if value, loaded := sm.sessions.Load(chatID); loaded {
		if s, ok := value.(*session); ok {
			s.lastAccessed.Store(time.Now().UnixNano())
			return s.bot, nil
		}
		return nil, fmt.Errorf("invalid session type for chat ID %d", chatID)
	}

	// Create a new session
	newSession := &session{
		bot: newBot(chatID),
	}
	newSession.lastAccessed.Store(time.Now().UnixNano())

	// Store the new session
	sm.sessions.Store(chatID, newSession)

	return newSession.bot, nil
}

// DelSession deletes the Bot instance (session) for the specified chat ID.
func (d *Dispatcher) DelSession(chatID int64) {
	d.sessions.delete(chatID)
}

// AddSession creates a new Bot instance (session) for the specified chat ID.
func (d *Dispatcher) AddSession(chatID int64) error {
	return d.sessions.add(chatID, d.newBot)
}

func (sm *sessionManager) delete(chatID int64) {
	sm.sessions.Delete(chatID)
}

func (sm *sessionManager) add(chatID int64, newBot NewBotFn) error {
	now := time.Now().UnixNano()
	newSession := &session{
		bot: newBot(chatID),
	}
	newSession.lastAccessed.Store(now)

	_, loaded := sm.sessions.LoadOrStore(chatID, newSession)
	if loaded {
		return fmt.Errorf("session for chat ID %d already exists", chatID)
	}
	return nil
}


// Periodically cleans up inactive sessions based on the configured TTL.
func (sm *sessionManager) clean(ctx context.Context) error {
	ticker := time.NewTicker(sm.ttl / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Perform one last cleaning before exiting
			sm.forceCleanup(true)
			return nil
		case <-ticker.C:
			sm.forceCleanup(false)
		}
	}
}


// forceCleanup removes all expired sessions. If force is true, it removes all sessions regardless of expiration.
func (sm *sessionManager) forceCleanup(force bool) int {
	var cleaned int
	now := time.Now().UnixNano()
	ttlNano := sm.ttl.Nanoseconds()

	sm.sessions.Range(func(key, value interface{}) bool {
		s, ok := value.(*session)
		if !ok {
			fmt.Printf("Invalid session type for key %v", key)
			return true
		}

		lastAccessed := s.lastAccessed.Load()
		if force || (now - lastAccessed) > ttlNano {
			sm.sessions.Delete(key)
			cleaned++
		}
		return true
	})

	return cleaned
}

