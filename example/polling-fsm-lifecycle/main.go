// polling-fsm-lifecycle extends the polling-fsm example with automatic session cleanup.
// A per-bot goroutine counts down 5 minutes of inactivity; every incoming
// update resets the timer. When the timer fires, the bot says goodbye and
// removes itself from the Dispatcher so its memory can be reclaimed.
// This pattern is useful for bots that serve many users and need to stay lean.
package main

import (
	"log"
	"os"
	"time"

	"github.com/NicoNex/echotron/v3"
)

// stateFn is a function that handles one update and returns the next state.
// Because it references itself in its own return type, states form a chain:
// each handler decides at runtime which function should run next.
type stateFn func(*echotron.Update) stateFn

type bot struct {
	chatID int64
	state  stateFn // current state; replaced after every update
	name   string
	reset  chan struct{} // signals the countdown goroutine to restart the timer
	echotron.API
}

var (
	dsp   *echotron.Dispatcher
	token = os.Getenv("TELEGRAM_TOKEN")
)

func newBot(chatID int64) echotron.Bot {
	b := &bot{
		chatID: chatID,
		// Buffered so resetCountdown never blocks even if the goroutine is
		// momentarily busy selecting on the other case.
		reset: make(chan struct{}, 1),
		API:   echotron.NewAPI(token),
	}
	go b.countdown()
	b.state = b.handleMessage // set the initial state
	return b
}

// countdown runs in its own goroutine. It waits for either an inactivity
// timeout or a reset signal from Update. On timeout it sends a farewell
// message and removes the session from the Dispatcher.
func (b bot) countdown() {
	for {
		select {
		case <-b.reset:
			// An update arrived; restart the inactivity timer.
		case <-time.After(5 * time.Minute):
			b.SendMessage("Bye bye!", b.chatID, nil)
			dsp.DelSession(b.chatID)
			return
		}
	}
}

// resetCountdown notifies the countdown goroutine that a new update arrived,
// resetting the inactivity timer without blocking the caller.
func (b bot) resetCountdown() {
	b.reset <- struct{}{}
}

func (b *bot) Update(update *echotron.Update) {
	b.resetCountdown()
	// Execute the current state and store whatever it returns as the next one.
	// A single assignment is all the state-machine machinery needed.
	b.state = b.state(update)
}

func (b *bot) handleMessage(u *echotron.Update) stateFn {
	if u.Message != nil && u.Message.Text == "/setname" {
		b.SendMessage("What should I call you?", b.chatID, nil)
		// The next update will be the user's reply, so transition to handleName.
		return b.handleName
	}
	// No relevant command: stay in the default state.
	return b.handleMessage
}

func (b *bot) handleName(u *echotron.Update) stateFn {
	if u.Message == nil {
		// The user has sent a wrong update type, so clarify the situation and
		// stay in the handleName state by returning it again.
		b.SendMessage("Please just send me your name as a normal message", b.chatID, nil)
		return b.handleName
	}

	b.name = u.Message.Text
	b.SendMessage("Got it, "+b.name+"!", b.chatID, nil)
	// Name has been recorded; go back to the default state.
	return b.handleMessage
}

func main() {
	dsp = echotron.NewDispatcher(token, newBot)
	for {
		// Poll blocks until a network error occurs, then returns it.
		// Sleeping before retrying avoids hammering the API on transient failures.
		log.Println(dsp.Poll())
		time.Sleep(5 * time.Second)
	}
}
