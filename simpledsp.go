package echotron

import "time"

func PollingUpdates(token string) <-chan *Update {
	return PollingUpdatesOptions(token, true, UpdateOptions{Timeout: 120})
}

func PollingUpdatesOptions(token string, dropPendingUpdates bool, opts UpdateOptions) <-chan *Update {
	var (
		api     = NewAPI(token)
		updates = make(chan *Update)
	)

	go func() {
		defer close(updates)

		var (
			timeout    = opts.Timeout
			isFirstRun = true
		)

		// deletes webhook if present to run in long polling mode
		api.DeleteWebhook(dropPendingUpdates)

		for {
			if isFirstRun {
				opts.Timeout = 0
			}

			response, err := api.GetUpdates(&opts)
			if err != nil {
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
