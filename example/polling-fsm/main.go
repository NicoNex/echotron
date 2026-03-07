// polling-fsm demonstrates the FSM (finite state machine) pattern built on top of the
// Dispatcher. Each bot instance owns a stateFn field that points to the
// handler for the current state; Update simply calls that function and stores
// whatever it returns as the next state. This makes multi-step conversations
// (wizards, forms, menus) easy to implement without any extra libraries.
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

var token = os.Getenv("TELEGRAM_TOKEN")

type bot struct {
	chatID int64
	state  stateFn // current state; replaced after every update
	name   string
	echotron.API
}

func newBot(chatID int64) echotron.Bot {
	b := &bot{chatID: chatID, API: echotron.NewAPI(token)}
	b.state = b.handleMessage // set the initial state
	return b
}

func (b *bot) Update(update *echotron.Update) {
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
	dsp := echotron.NewDispatcher(token, newBot)
	for {
		// Poll blocks until a network error occurs, then returns it.
		// Sleeping before retrying avoids hammering the API on transient failures.
		log.Println(dsp.Poll())
		time.Sleep(5 * time.Second)
	}
}
