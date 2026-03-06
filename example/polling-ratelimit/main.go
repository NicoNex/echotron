// polling-ratelimit demonstrates how to configure echotron's built-in
// dual-level rate limiter. Two independent limits mirror Telegram's guidelines:
//
//   - Global:   max requests per second across all chats (default: 30 req/s)
//   - Per-chat: max requests per chat per unit of time  (default: 20 req/min)
//
// The limiter is shared across every API instance for the same token, so
// configuring it once at startup is enough — all bot instances automatically
// respect the same limits.
package main

import (
	"log"
	"os"
	"time"

	"github.com/NicoNex/echotron/v3"
)

var token = os.Getenv("TELEGRAM_TOKEN")

// bot holds the per-chat state. Each conversation gets its own instance.
// Embedding echotron.API promotes all Telegram methods onto the struct,
// so b.SendMessage works directly without any explicit delegation.
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
	// Configure rate limits once before starting the Dispatcher.
	// NewAPI for the same token always returns a handle to the same underlying
	// client, so these settings apply to every API call made by any bot instance.
	api := echotron.NewAPI(token)

	// 30 requests/second globally across all chats — Telegram's hard limit.
	api.SetGlobalRequestLimit(time.Second/30, 30)

	// 1 request/second per chat — stricter than Telegram's default of 20/min,
	// but prevents bursts that could feel spammy to individual users.
	// Pass an interval of 0 to either method to disable that limiter entirely.
	api.SetChatRequestLimit(time.Second, 1)

	dsp := echotron.NewDispatcher(token, newBot)
	for {
		// Poll blocks until a network error occurs, then returns it.
		// Sleeping before retrying avoids hammering the API on transient failures.
		log.Println(dsp.Poll())
		time.Sleep(5 * time.Second)
	}
}
