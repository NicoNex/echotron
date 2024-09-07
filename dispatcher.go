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
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Bot interface {
	Update(*Update)
}

type NewBotFn func(chatId int64) Bot

type Dispatcher struct {
	api      API
	newBot   NewBotFn
	sessions *sessionManager
	options  dispatcherOptions

	inShutdown atomic.Bool
	mu         sync.Mutex
	ctx        context.Context
	cancel     context.CancelFunc
	isRunning  atomic.Bool
	updateChan chan *Update
	wg         sync.WaitGroup
}

type dispatcherOptions struct {
	pollInterval     time.Duration
	updateBufferSize int
	sessionTTL       time.Duration
	allowedUpdates   []UpdateType
	workerSem        *semaphore.Weighted
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
		workerSem:        semaphore.NewWeighted(int64(runtime.NumCPU() * 2)),
	}
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return nil, err
		}
	}
	d := &Dispatcher{
		api:        NewAPI(token),
		newBot:     newBot,
		options:    options,
		sessions:   newSessionManager(options.sessionTTL),
		updateChan: make(chan *Update, options.updateBufferSize),
	}
	return d, nil
}

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

	g, ctx := errgroup.WithContext(d.ctx)

	g.Go(func() error {
		return d.pollUpdates(ctx)
	})

	g.Go(func() error {
		return d.worker(ctx)
	})

	g.Go(func() error {
		return d.sessions.clean(ctx)
	})

	// Start a goroutine to wait for errors
	go func() {
		err := g.Wait()
		if err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("Dispatcher error: %v\n", err)
			if stopErr := d.Stop(context.Background()); stopErr != nil {
				fmt.Printf("Error stopping dispatcher: %v\n", stopErr)
			}
		}
	}()

	return nil
}

func (d *Dispatcher) Stop(ctx context.Context) error {
	if !d.isRunning.CompareAndSwap(true, false) {
		return errors.New("dispatcher is not running")
	}

	d.mu.Lock()
	if d.inShutdown.Swap(true) {
		d.mu.Unlock()
		return errors.New("dispatcher is already shutting down")
	}
	d.cancel()
	close(d.updateChan)
	d.mu.Unlock()

	// Use a channel to signal when all goroutines have finished
	done := make(chan struct{})
	go func() {
		d.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (d *Dispatcher) worker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case update, ok := <-d.updateChan:
			if !ok {
				return nil // Channel closed, exit gracefully
			}
			if d.inShutdown.Load() {
				continue // Discard updates during shutdown
			}
			if err := d.options.workerSem.Acquire(ctx, 1); err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}
				return fmt.Errorf("failed to acquire semaphore: %w", err)
			}
			d.wg.Add(1)
			go func(update *Update) {
				defer d.options.workerSem.Release(1)
				defer d.wg.Done()
				if err := d.processUpdate(update); err != nil {
					fmt.Printf("Error processing update: %v\n", err)
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
func (d *Dispatcher) pollUpdates(ctx context.Context) error {
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
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if d.inShutdown.Load() {
				return nil
			}
			updates, err := d.api.GetUpdates(&opts)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}
				fmt.Printf("Failed to get updates: %v\n", err)
				continue
			}

			if len(updates.Result) > 0 {
				lastUpdateID := updates.Result[len(updates.Result)-1].ID
				opts.Offset = lastUpdateID + 1

				for _, update := range updates.Result {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case d.updateChan <- update:
					default:
						fmt.Println("Update channel is full, discarding update")
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
func (sm *sessionManager) getSession(chatID int64, newBot NewBotFn) (bot Bot, err error) {
	lock, _ := sm.locks.LoadOrStore(chatID, &sync.Mutex{})
	mutex := lock.(*sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	if value, loaded := sm.sessions.Load(chatID); loaded {
		if s, ok := value.(*session); ok {
			s.lastAccessed.Store(time.Now().UnixNano())
			return s.bot, nil
		}
		return nil, fmt.Errorf("invalid session type for chat ID %d", chatID)
	}

	// Use a defer to handle potential panics from newBot
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic in newBot: %v", r)
		}
	}()

	bot = newBot(chatID)
	newSession := &session{
		bot: bot,
	}
	newSession.lastAccessed.Store(time.Now().UnixNano())
	sm.sessions.Store(chatID, newSession)

	return bot, nil
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
	sm.locks.Delete(chatID)
}

func (sm *sessionManager) add(chatID int64, newBot NewBotFn) error {
	lock, _ := sm.locks.LoadOrStore(chatID, &sync.Mutex{})
	mutex := lock.(*sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

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
			// Perform the last cleanup in a separate goroutine to not delay shutdown
			go sm.forceCleanup(true)
			return ctx.Err()
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
		chatID, ok := key.(int64)
		if !ok {
			fmt.Printf("Invalid key type for session: %v\n", key)
			sm.delete(chatID)
			cleaned++
			return true
		}

		s, ok := value.(*session)
		if !ok {
			fmt.Printf("Invalid session type for chat ID %d\n", chatID)
			sm.delete(chatID)
			cleaned++
			return true
		}

		lastAccessed := s.lastAccessed.Load()
		if force || (now-lastAccessed) > ttlNano {
			sm.delete(chatID)
			cleaned++
		}
		return true
	})

	return cleaned
}
