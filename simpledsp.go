package echotron

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// PollingUpdates is a wrapper function for PollingUpdatesOptions.
func PollingUpdates(token string) <-chan *Update {
	return PollingUpdatesOptions(token, true, UpdateOptions{Timeout: 120})
}

// PollingUpdatesOptions returns a read-only channel of incoming  updates from the Telegram API.
func PollingUpdatesOptions(token string, dropPendingUpdates bool, opts UpdateOptions) <-chan *Update {
	var updates = make(chan *Update)

	go func() {
		defer close(updates)

		var (
			api        = NewAPI(token)
			timeout    = opts.Timeout
			isFirstRun = true
		)

		// deletes webhook if present to run in long polling mode
		if _, err := api.DeleteWebhook(dropPendingUpdates); err != nil {
			log.Println("echotron.PollingUpdates", err)
		}

		for {
			if isFirstRun {
				opts.Timeout = 0
			}

			response, err := api.GetUpdates(&opts)
			if err != nil {
				log.Println("echotron.PollingUpdates", err)
				time.Sleep(5 * time.Second)
				continue
			}

			if !dropPendingUpdates || !isFirstRun {
				for _, u := range response.Result {
					updates <- u
				}
			}

			if l := len(response.Result); l > 0 {
				opts.Offset = response.Result[l-1].ID + 1
			}

			if isFirstRun {
				isFirstRun = false
				opts.Timeout = timeout
			}
		}
	}()

	return updates
}

// WebhookUpdates is a wrapper function for WebhookUpdatesOptions.
func WebhookUpdates(url, token string) <-chan *Update {
	return WebhookUpdatesOptions(url, token, false, nil)
}

// WebhookUpdatesOptions returns a read-only channel of incoming updates from the Telegram API.
// The webhookUrl should be provided in the following format: '<hostname>:<port>/<path>',
// eg: 'https://example.com:443/bot_token'.
// WebhookUpdatesOptions will then proceed to communicate the webhook url '<hostname>/<path>'
// to Telegram and run a webserver that listens to ':<port>' and handles the path.
func WebhookUpdatesOptions(whURL, token string, dropPendingUpdates bool, opts *WebhookOptions) <-chan *Update {
	u, err := url.Parse(whURL)
	if err != nil {
		panic(err)
	}

	wURL := u.Hostname() + u.EscapedPath()
	api := NewAPI(token)
	if _, err := api.SetWebhook(wURL, dropPendingUpdates, opts); err != nil {
		panic(err)
	}

	var updates = make(chan *Update)
	http.HandleFunc(u.EscapedPath(), func(w http.ResponseWriter, r *http.Request) {
		var update Update

		jsn, err := readRequest(r)
		if err != nil {
			log.Println("echotron.WebhookUpdates", err)
			return
		}

		if err := json.Unmarshal(jsn, &update); err != nil {
			log.Println("echotron.WebhookUpdates", err)
			return
		}

		updates <- &update
	})

	go func() {
		defer close(updates)
		port := fmt.Sprintf(":%s", u.Port())
		for {
			if err := http.ListenAndServe(port, nil); err != nil {
				log.Println("echotron.WebhookUpdates", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	return updates
}
