// polling-fsm-persistence extends polling-fsm with disk persistence using
// katalis as the storage backend. Per-chat data is grouped in a session
// struct that is written to disk whenever a field changes, so state survives
// process restarts. The database is opened at startup and shared as a
// package-level variable across all bot instances.
package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/NicoNex/echotron/v3"
	"github.com/NicoNex/katalis"
)

// stateFn is a function that handles one update and returns the next state.
// Because it references itself in its own return type, states form a chain:
// each handler decides at runtime which function should run next.
type stateFn func(*echotron.Update) stateFn

// session holds only the fields that need to survive a process restart.
// Keeping them separate from bot makes it clear what is persisted and
// what is transient (e.g. state, chatID, API).
// All fields must be exported so the gob codec can encode and decode them.
type session struct {
	Name string
}

type bot struct {
	chatID int64
	state  stateFn // current state; replaced after every update
	session
	echotron.API
}

// db is the shared katalis store, keyed by chat ID, opened once at startup.
var (
	db    katalis.DB[int64, session]
	token = os.Getenv("TELEGRAM_TOKEN")
)

func newBot(chatID int64) echotron.Bot {
	b := &bot{
		chatID: chatID,
		API:    echotron.NewAPI(token),
	}

	// Restore persisted fields for this chat. Get returns a zero-value
	// session when the key is missing, which is the correct default for a
	// new session, so no separate Has check is needed.
	if sess, err := db.Get(chatID); err != nil {
		log.Println(err)
	} else {
		b.session = sess
	}

	// Note: the current FSM state (i.e. which handler is active) is not
	// persisted. The bot always restarts from handleMessage regardless of
	// where the conversation was when the process exited. If you need
	// mid-conversation continuity across restarts, add a state field to
	// session and map it back to a stateFn in newBot.
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

	b.setName(u.Message.Text)
	b.SendMessage("Got it, "+b.Name+"!", b.chatID, nil)
	// Name has been recorded; go back to the default state.
	return b.handleMessage
}

// setName updates the name and immediately persists the session to disk.
// Persisting on mutation rather than on every update avoids unnecessary
// writes for updates that don't change any stored field.
func (b *bot) setName(newName string) {
	b.Name = newName
	if err := db.Put(b.chatID, b.session); err != nil {
		log.Println(err)
	}
}

// initDB opens the katalis database, creating it if it does not exist.
func initDB() (katalis.DB[int64, session], error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return katalis.DB[int64, session]{}, err
	}

	return katalis.OpenOptions(
		// Replace "my-bot-name" with a unique name for your bot.
		// Each bot on the same machine should use a different path to
		// avoid sharing or overwriting each other's data.
		filepath.Join(cacheDir, "my-bot-name"),
		katalis.Int64Codec,
		katalis.Gob[session](),
		&katalis.Options{
			// Flush dirty pages to disk every 500ms. This balances write
			// safety and throughput; lower values increase durability at
			// the cost of more frequent syscalls.
			BackgroundSyncInterval: 500 * time.Millisecond,
			// Compact the database once a day to reclaim space from
			// overwritten or deleted records.
			BackgroundCompactionInterval: 24 * time.Hour,
		},
	)
}

// handleSignals blocks until SIGINT or SIGTERM is received, then closes the
// database and exits. Run it in a goroutine so main can start the Dispatcher.
func handleSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh
	db.Close()
	os.Exit(0)
}

func main() {
	var err error

	if db, err = initDB(); err != nil {
		log.Fatalln(err)
	}

	// Gracefully handle SIGINT and SIGTERM.
	go handleSignals()

	dsp := echotron.NewDispatcher(token, newBot)
	for {
		// Poll blocks until a network error occurs, then returns it.
		// Sleeping before retrying avoids hammering the API on transient failures.
		log.Println(dsp.Poll())
		time.Sleep(5 * time.Second)
	}
}
