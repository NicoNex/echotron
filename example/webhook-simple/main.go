// webhook-simple demonstrates the minimal stateless webhook bot using echotron's
// functional API. There is no per-chat state and no Dispatcher: a single API
// instance handles every incoming update in sequence.
// This is the quickest way to get a webhook-based bot running; use the webhook
// or polling-fsm examples when you need per-chat state or concurrent handling.
//
// Before running, replace the URL passed to WebhookUpdates with your own
// publicly reachable HTTPS endpoint. Telegram requires a valid TLS certificate
// (self-signed certificates are accepted) on port 443, 80, 88, or 8443.
package main

import (
	"os"

	"github.com/NicoNex/echotron/v3"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	api := echotron.NewAPI(token)

	// WebhookUpdates registers the webhook URL with Telegram, then starts an
	// HTTP server that listens for incoming updates and forwards them to the
	// returned channel. Replace the URL with your own HTTPS endpoint.
	updates := echotron.WebhookUpdates("https://example.com:443/"+token, token)

	for u := range updates {
		if u.Message != nil && u.Message.Text == "/start" {
			api.SendMessage("Hello, world!", u.ChatID(), nil)
		}
	}
}
