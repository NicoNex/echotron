// polling-keyboard demonstrates how to send inline keyboards and handle
// callback queries using echotron's Dispatcher with long-polling.
// Inline keyboards attach interactive buttons directly to a message;
// pressing a button sends a CallbackQuery update back to the bot.
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
	switch {
	case update.Message != nil && update.Message.Text == "/start":
		// Build a two-button inline keyboard attached to the reply.
		keyboard := echotron.InlineKeyboardMarkup{
			InlineKeyboard: [][]echotron.InlineKeyboardButton{
				{
					{Text: "Option A", CallbackData: "a"},
					{Text: "Option B", CallbackData: "b"},
				},
			},
		}
		b.SendMessage("Pick an option:", b.chatID, &echotron.MessageOptions{
			ReplyMarkup: keyboard,
		})

	case update.CallbackQuery != nil:
		// Always answer the callback query to dismiss the loading indicator
		// on the client. Failing to do so leaves a spinner on the button.
		b.AnswerCallbackQuery(update.CallbackQuery.ID, nil)
		b.SendMessage("You picked: "+update.CallbackQuery.Data, b.chatID, nil)
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
