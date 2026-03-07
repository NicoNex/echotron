// polling demonstrates a per-chat stateful bot using echotron's Dispatcher
// with long-polling. The Dispatcher creates one bot instance per chat ID and
// calls Update concurrently, so each conversation is fully independent.
// Use this as your starting point for any bot that needs to remember things
// about individual users across multiple messages.
package main

import (
	"log"
	"os"
	"time"

	"github.com/NicoNex/echotron/v3"
)

// bot holds the per-chat state. Each conversation gets its own instance.
// Embedding echotron.API promotes all Telegram methods onto the struct,
// so b.SendMessage works directly without any explicit delegation.
var token = os.Getenv("TELEGRAM_TOKEN")

type bot struct {
	chatID int64
	echotron.API
}

// newBot is the factory function called by the Dispatcher the first time
// a given chatID sends a message. It must return an echotron.Bot.
func newBot(chatID int64) echotron.Bot {
	return &bot{chatID, echotron.NewAPI(token)}
}

// Update is the only method required by the echotron.Bot interface.
// The Dispatcher calls it in a new goroutine for every incoming update,
// so each chat is handled concurrently without blocking the others.
func (b *bot) Update(update *echotron.Update) {
	if update.Message != nil && update.Message.Text == "/start" {
		b.SendMessage("Hello!", b.chatID, nil)
	}
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
