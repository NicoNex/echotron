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

// Represents an active bot session. The Update method is called whenever the bot receives a new update from Telegram.
type Bot interface {
	// Update will be called upon receiving any update from Telegram.
	Update(*Update)
}

// A function type that is called to create a new Bot instance when an update is received from a new chat ID.
type NewBotFn func(chatId int64) Bot

// Dispatcher manages the distribution of Telegram updates to bot instances.
type Dispatcher struct {
	api             API 				// Manages communication with Telegram's API.
	newBot          NewBotFn 			// A function to create new bot instances for new chat IDs.
	sessions        *sessionManager 	// Manages bot sessions.
	updateTimeout   time.Duration 		// The maximum time allowed for processing each update.
	pollInterval    time.Duration 		// The interval between long polling requests.
	updateBufferSize int 				// The buffer size for the update channel.
	sessionTTL      time.Duration 		// The time-to-live for inactive sessions.
	maxConcurrentUpdates int64
	updateSemaphore      *semaphore.Weighted
	allowedUpdates  []UpdateType		// Specifies the types of updates the bot is allowed to receive.
	errorHandler    func(error)			// Handles errors that occur during the dispatch process.
	running         atomic.Bool 		// Indicates whether the dispatcher is currently running.
	ctx             context.Context 	// The context for controlling the dispatcher's operations.
	cancel          context.CancelFunc 	// A function to cancel the dispatcher's operations.
	group           *errgroup.Group
	shutdownTimeout time.Duration
	updateChan      chan *Update
	doneChan        chan struct{}
}

// A function type used to configure the Dispatcher.
type Option func(*Dispatcher) error

// WithUpdateTimeout sets the timeout for processing each update.
func WithUpdateTimeout(timeout time.Duration) Option {
	return func(d *Dispatcher) error {
		if timeout <= 0 {
			return errors.New("update timeout must be positive")
		}
		d.updateTimeout = timeout
		return nil
	}
}

// WithPollInterval sets the interval between long polling requests.
func WithPollInterval(interval time.Duration) Option {
	return func(d *Dispatcher) error {
		if interval <= 0 {
			return errors.New("poll interval must be positive")
		}
		d.pollInterval = interval
		return nil
	}
}

// WithUpdateBufferSize sets the size of the update buffer channel.
func WithUpdateBufferSize(size int) Option {
	return func(d *Dispatcher) error {
		if size <= 0 {
			return errors.New("update buffer size must be positive")
		}
		d.updateBufferSize = size
		return nil
	}
}

// WithSessionTTL sets the time-to-live for inactive sessions.
func WithSessionTTL(ttl time.Duration) Option {
	return func(d *Dispatcher) error {
		if ttl <= 0 {
			return errors.New("session TTL must be positive")
		}
		d.sessionTTL = ttl
		return nil
	}
}

// WithMaxConcurrentUpdates sets the maximum number of concurrent updates
func WithMaxConcurrentUpdates(max int64) Option {
	return func(d *Dispatcher) error {
		if max <= 0 {
			return errors.New("max concurrent updates must be positive")
		}
		d.maxConcurrentUpdates = max
		return nil
	}
}

// WithErrorHandler sets a custom error handler for the Dispatcher.
func WithErrorHandler(handler func(error)) Option {
	return func(d *Dispatcher) error {
		if handler == nil {
			return errors.New("error handler cannot be nil")
		}
		d.errorHandler = handler
		return nil
	}
}

// Specifies the types of updates that the bot is allowed to receive.
func WithAllowedUpdates(types []UpdateType) Option {
	return func(d *Dispatcher) error {
		d.allowedUpdates = types
		return nil
	}
}

// Add this option to set the shutdown timeout
func WithShutdownTimeout(timeout time.Duration) Option {
	return func(d *Dispatcher) error {
		if timeout <= 0 {
			return errors.New("shutdown timeout must be positive")
		}
		d.shutdownTimeout = timeout
		return nil
	}
}

// NewDispatcher creates a new Dispatcher instance with the given options.
func NewDispatcher(token string, newBot NewBotFn, opts ...Option) (*Dispatcher, error) {
	d := &Dispatcher{
		api:             NewAPI(token),
		newBot:          newBot,
		updateTimeout:   30 * time.Second,
		pollInterval:    5 * time.Second,
		updateBufferSize: 1000,
		sessionTTL:      24 * time.Hour,
		maxConcurrentUpdates: int64(runtime.NumCPU() * 2), // Default to 2x number of CPUs
		shutdownTimeout: 30 * time.Second,
		errorHandler:    func(err error) { fmt.Println("Dispatcher error:", err) },
	}

	for _, opt := range opts {
		if err := opt(d); err != nil {
			return nil, err
		}
	}
	d.updateSemaphore = semaphore.NewWeighted(d.maxConcurrentUpdates)
	d.sessions = newSessionManager(d.sessionTTL)
	return d, nil
}


// Start begins the Dispatcher's operations.
func (d *Dispatcher) Start() error {
	if !d.running.CompareAndSwap(false, true) {
		return errors.New("dispatcher is already running")
	}

	d.ctx, d.cancel = context.WithCancel(context.Background())
	d.updateChan = make(chan *Update, d.updateBufferSize)
	d.doneChan = make(chan struct{})

	// Start the update polling goroutine
	go func() {
		if err := d.pollUpdates(d.ctx); err != nil {
			d.errorHandler(fmt.Errorf("polling updates: %w", err))
		}
	}()

	// Start the single worker goroutine
	go d.worker(d.ctx)

	// Start the session cleaner
	go func() {
		if err := d.sessions.clean(d.ctx); err != nil {
			// Only log errors that aren't due to context cancellation
			if !errors.Is(err, context.Canceled) {
				d.errorHandler(fmt.Errorf("cleaning sessions: %w", err))
			}
		}
	}()

	return nil
}

func (d *Dispatcher) Stop() error {
	if !d.running.CompareAndSwap(true, false) {
		return errors.New("dispatcher is not running")
	}

	d.cancel() // Signal all goroutines to stop
	close(d.updateChan) // Close update channel to stop workers

	// Wait for the worker to finish
	select {
	case <-d.doneChan:
		return nil
	case <-time.After(d.shutdownTimeout):
		return fmt.Errorf("shutdown timed out after %v", d.shutdownTimeout)
	}
}

// Modify the worker method
func (d *Dispatcher) worker(ctx context.Context) {
	defer close(d.doneChan)

	for {
		select {
		case <-ctx.Done():
			return
		case update, ok := <-d.updateChan:
			if !ok {
				return // Channel closed, exit the worker
			}
			if err := d.updateSemaphore.Acquire(ctx, 1); err != nil {
				if err != context.Canceled {
					d.errorHandler(fmt.Errorf("failed to acquire semaphore: %w", err))
				}
				continue
			}
			go func(update *Update) {
				defer d.updateSemaphore.Release(1)
				if err := d.processUpdate(ctx, update); err != nil {
					d.errorHandler(fmt.Errorf("processing update: %w", err))
				}
			}(update)
		}
	}
}

// Processes updates received from Telegram, distributing them to the appropriate bot instances.
func (d *Dispatcher) processUpdate(ctx context.Context, update *Update) error {
	bot, err := d.sessions.getOrCreate(update.ChatID(), d.newBot)
	if err != nil {
		return fmt.Errorf("getting bot instance: %w", err)
	}

	updateCtx, cancel := context.WithTimeout(ctx, d.updateTimeout)
	defer cancel()

	done := make(chan struct{})
	go func() {
		bot.Update(update)
		close(done)
	}()

	select {
	case <-updateCtx.Done():
		return fmt.Errorf("update timed out for chat ID %d", update.ChatID())
	case <-done:
		return nil
	}
}

// Polls Telegram for updates and queues them for processing.
func (d *Dispatcher) pollUpdates(ctx context.Context) error {
	if _, err := d.api.DeleteWebhook(true); err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}

	opts := UpdateOptions{
		Timeout:        int(d.pollInterval.Seconds()),
		Limit:          100,
		Offset:         0,
		AllowedUpdates: d.allowedUpdates,
	}

	ticker := time.NewTicker(d.pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			updates, err := d.api.GetUpdates(&opts)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return nil
				}
				d.errorHandler(fmt.Errorf("failed to get updates: %w", err))
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
						// Update sent successfully
					default:
						// Channel is full, log an error or handle it appropriately
						d.errorHandler(fmt.Errorf("update channel is full, discarding update"))
					}
				}
			}
		}
	}
}

type session struct {
	bot          Bot
	lastAccessed time.Time
}

type sessionManager struct {
	sessions sync.Map
	ttl      time.Duration
}

// DelSession deletes the Bot instance, seen as a session, from the map with all of them.
func (d *Dispatcher) DelSession(chatID int64) {
	d.sessions.delete(chatID)
}

// AddSession allows to arbitrarily create a new Bot instance.
func (d *Dispatcher) AddSession(chatID int64) error {
	return d.sessions.add(chatID, d.newBot)
}

// Create session manager with ttl option.
func newSessionManager(ttl time.Duration) *sessionManager {
	return &sessionManager{ttl: ttl}
}

// Retrieves an existing bot session or creates a new one if it doesn't exist.
func (sm *sessionManager) getOrCreate(chatID int64, newBot NewBotFn) (Bot, error) {
	bot, ok := sm.sessions.Load(chatID)
	if !ok {
		newBot := newBot(chatID)
		bot, _ = sm.sessions.LoadOrStore(chatID, &session{
			bot:          newBot,
			lastAccessed: time.Now(),
		})
	}
	s := bot.(*session)
	s.lastAccessed = time.Now()
	return s.bot, nil
}

// Deletes a bot session for the specified chat ID.
func (sm *sessionManager) delete(chatID int64) {
	sm.sessions.Delete(chatID)
}

// Adds a new bot session for the specified chat ID.
func (sm *sessionManager) add(chatID int64, newBot NewBotFn) error {
	_, loaded := sm.sessions.LoadOrStore(chatID, &session{
		bot:          newBot(chatID),
		lastAccessed: time.Now(),
	})
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
			sm.cleanOnce(context.Background())
			return nil // Exit without error on context cancellation
		case <-ticker.C:
			sm.cleanOnce(ctx)
		}
	}
}


func (sm *sessionManager) cleanOnce(ctx context.Context) {
	now := time.Now()
	var cleaned int
	sm.sessions.Range(func(key, value interface{}) bool {
		select {
		case <-ctx.Done():
			return false // Stop iteration if context is canceled
		default:
			s := value.(*session)
			if now.Sub(s.lastAccessed) > sm.ttl {
				sm.sessions.Delete(key)
				cleaned++
			}
			return true
		}
	})
	// Optionally log the number of cleaned sessions
	// fmt.Printf("Cleaned %d expired sessions", cleaned)
}
