// polling-inline demonstrates how to handle inline queries using echotron's
// Dispatcher with long-polling. Inline mode lets users trigger the bot from
// any chat by typing "@botname <query>" without opening a private conversation.
// Before running, enable inline mode for your bot via @BotFather (/setinline).
package main

import (
	"log"
	"os"
	"time"

	"github.com/NicoNex/echotron/v3"
)

var token = os.Getenv("TELEGRAM_TOKEN")

// bot holds the per-user state. For inline queries ChatID() returns the
// sender's user ID, so each user gets their own Dispatcher-managed instance.
// Embedding echotron.API promotes all Telegram methods onto the struct.
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
// so each user is handled concurrently without blocking the others.
// AllowedUpdates in PollOptions guarantees update.InlineQuery is never nil here.
func (b *bot) Update(update *echotron.Update) {
	query := update.InlineQuery.Query
	if query == "" {
		// Telegram sends an empty string when the user has not typed anything
		// yet. Return a default result instead of an empty list.
		query = "hello"
	}

	// Return a single article result that echoes the query back to the chat.
	// InlineQueryResultArticle requires an InputMessageContent that specifies
	// what is actually sent when the user selects the result.
	b.AnswerInlineQuery(
		update.InlineQuery.ID,
		[]echotron.InlineQueryResult{
			echotron.InlineQueryResultArticle{
				Type:  echotron.InlineArticle,
				ID:    "1",
				Title: "Echo: " + query,
				InputMessageContent: echotron.InputTextMessageContent{
					MessageText: query,
				},
			},
		},
		nil,
	)
}

func main() {
	dsp := echotron.NewDispatcher(token, newBot)
	for {
		// PollOptions blocks until a network error occurs, then returns it.
		// Sleeping before retrying avoids hammering the API on transient failures.
		// AllowedUpdates restricts delivery to inline queries only, so Update
		// can safely access update.InlineQuery without a nil check.
		log.Println(dsp.PollOptions(false, echotron.UpdateOptions{
			AllowedUpdates: []echotron.UpdateType{echotron.InlineQueryUpdate},
			Timeout:        120,
		}))
		time.Sleep(5 * time.Second)
	}
}
