// polling-simple demonstrates the minimal stateless polling bot using echotron's
// functional API. There is no per-chat state and no Dispatcher: a single API
// instance handles every incoming update in sequence.
// This is the quickest way to get a bot running; use the polling or
// polling-fsm examples when you need per-chat state or concurrent handling.
package main

import (
	"os"

	"github.com/NicoNex/echotron/v3"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	api := echotron.NewAPI(token)

	// PollingUpdates returns a channel that yields one update at a time.
	// It handles long-polling internally and reconnects on transient errors.
	for u := range echotron.PollingUpdates(token) {
		if u.Message != nil && u.Message.Text == "/start" {
			api.SendMessage("Hello, world!", u.ChatID(), nil)
		}
	}
}
