<img src="assets/readme_banner.png" alt="Echotron" width="800">

<div align="center">

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NicoNex/echotron/v3)](https://pkg.go.dev/github.com/NicoNex/echotron/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/NicoNex/echotron/v3)](https://goreportcard.com/report/github.com/NicoNex/echotron/v3)
[![codecov](https://codecov.io/gh/NicoNex/echotron/graph/badge.svg?token=LVJGOEYL5M)](https://codecov.io/gh/NicoNex/echotron)
[![License](http://img.shields.io/badge/license-LGPL3.0-orange.svg?style=flat)](https://github.com/NicoNex/echotron/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Telegram](https://img.shields.io/badge/Echotron%20News-blue?logo=telegram&style=flat)](https://t.me/echotronnews)

**The idiomatic, concurrent Telegram Bot library for Go.**  
Zero boilerplate. Built-in rate limiting. One instance per chat, by design.

</div>

## Why Echotron?

Most Telegram bot libraries hand you a stream of updates and leave everything else to you. Echotron goes further: it ships a battle-tested concurrency model out of the box, so you can focus on bot logic instead of synchronisation primitives, rate limiters, and state machinery.

```bash
go get github.com/NicoNex/echotron/v3
```

## Five-line bot

```go
package main

import "github.com/NicoNex/echotron/v3"

func main() {
    api := echotron.NewAPI("MY_TOKEN")
    
    for u := range echotron.PollingUpdates("MY_TOKEN") {
        if u.Message.Text == "/start" {
            api.SendMessage("Hello, world!", u.ChatID(), nil)
        }
    }
}
```

No setup, no registration, no middleware stack. `PollingUpdates` returns a plain Go channel: range over it and you are done.

## The dispatcher pattern

For production bots, Echotron's `Dispatcher` automatically maintains one isolated bot instance per chat. Each instance gets its own goroutine; state is naturally scoped per conversation.

```go
package main

import (
    "log"
    "time"

    "github.com/NicoNex/echotron/v3"
)

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
    return &bot{chatID, echotron.NewAPI("MY_TOKEN")}
}

// Update is the only method required by the echotron.Bot interface.
// The Dispatcher calls it in a new goroutine for every incoming update,
// so each chat is handled concurrently without blocking the others.
func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello!", b.chatID, nil)
    }
}

func main() {
    dsp := echotron.NewDispatcher("MY_TOKEN", newBot)
    for {
        // Poll blocks until a network error occurs, then returns it.
        // Sleeping before retrying avoids hammering the API on transient failures.
        log.Println(dsp.Poll())
        time.Sleep(5 * time.Second)
    }
}
```

| Concern | Handled by |
|---|---|
| Routing updates to the right chat | `Dispatcher` |
| Creating state for first-time users | `newBot` factory |
| Calling `Update` concurrently | `Dispatcher` (goroutine per update) |
| Rate limiting API calls | built-in `lclient` |
| Deduplication and offset tracking | `Dispatcher.Poll()` |

## A library, not a framework

Echotron does not ask you to structure your application in any particular way. There is no command router to register, no middleware stack to assemble, no lifecycle hooks to implement, no configuration object to fill out before anything works.

The `Dispatcher` is entirely optional. If you do not need per-chat state management, you can drive updates yourself with a plain channel:

```go
for u := range echotron.PollingUpdates("MY_TOKEN") {
    // your logic here, completely vanilla Go
}
```

If you want webhooks without the Dispatcher, that is a channel too:

```go
for u := range echotron.WebhookUpdates("https://example.com:443/MY_TOKEN", "MY_TOKEN") {
    // handle u
}
```

If you need only a subset of the Telegram API for a quick script or a one-off tool, instantiate `NewAPI` directly and call whatever methods you need. Nothing forces you to go further.

The `Bot` interface itself requires a single method:

```go
type Bot interface {
    Update(*Update)
}
```

That one line is the entire contract. Every other piece of Echotron is additive: pick what fits your use case and ignore the rest.

## Feature highlights

### One instance per chat

The `Dispatcher` stores one `Bot` per `chatID` in a lock-free `sync.Map`. `Update()` is dispatched in a fresh goroutine for every incoming message:

- No chat ever blocks another: a deadlock in chat A has zero effect on chat B.
- State lives in struct fields: no need for maps, mutexes, or context keys to correlate users.
- Crashes are isolated: a panic in one goroutine does not bring down the whole bot.

### Built-in dual-level rate limiting

Echotron ships a transparent, dual-layer rate limiter that mirrors Telegram's own limits. It is always active with sensible defaults and requires no configuration to be correct from day one.

| Limiter | Default | How to change |
|---|---|---|
| Global (all chats) | 30 req/s, burst 30 | `api.SetGlobalRequestLimit(interval, burst)` |
| Per-chat | 20 req/min, burst 20 | `api.SetChatRequestLimit(interval, burst)` |

One shared `http.Client` per bot token means connection pools are reused across all chat instances, keeping resource usage proportional to the number of bots, not the number of users.

### Functional state machines

Multi-step conversations can be modelled as a self-referential function type that returns the next state. No string enums, no external table, no switch-on-state:

```go
// stateFn is a function that handles one update and returns the next state.
// Because it references itself in its own return type, states form a chain:
// each handler decides at runtime which function should run next.
type stateFn func(*echotron.Update) stateFn

type bot struct {
    chatID int64
    state  stateFn // current state; replaced after every update
    name   string
    echotron.API
}

func newBot(chatID int64) echotron.Bot {
    b := &bot{chatID: chatID, API: echotron.NewAPI("MY_TOKEN")}
    b.state = b.handleMessage // set the initial state
    return b
}

func (b *bot) Update(update *echotron.Update) {
    // Execute the current state and store whatever it returns as the next one.
    // A single assignment is all the state-machine machinery needed.
    b.state = b.state(update)
}

func (b *bot) handleMessage(u *echotron.Update) stateFn {
    if u.Message.Text == "/setname" {
        b.SendMessage("What should I call you?", b.chatID, nil)
        // The next update will be the user's reply, so transition to handleName.
        return b.handleName
    }
    // No relevant command: stay in the default state.
    return b.handleMessage
}

func (b *bot) handleName(u *echotron.Update) stateFn {
    b.name = u.Message.Text
    b.SendMessage("Got it, "+b.name+"!", b.chatID, nil)
    // Name has been recorded; go back to the default state.
    return b.handleMessage
}
```

### Session self-destruction

Inactive sessions can remove themselves from the dispatcher, keeping memory usage proportional to currently active conversations rather than all users who ever interacted with the bot:

```go
func newBot(chatID int64) echotron.Bot {
    b := &bot{chatID, echotron.NewAPI("MY_TOKEN")}
    // Launch the timer in a separate goroutine so newBot returns immediately.
    // The bot instance is fully functional while selfDestruct waits in the background.
    go b.selfDestruct(time.After(time.Hour))
    return b
}

func (b *bot) selfDestruct(ch <-chan time.Time) {
    <-ch // block until the timer fires
    b.SendMessage("Goodbye!", b.chatID, nil)
    // Remove this instance from the dispatcher's session map.
    // After this call the struct will be garbage-collected once no other
    // references remain, freeing all per-chat state automatically.
    dsp.DelSession(b.chatID)
}
```

### Webhook support, with or without a custom server

Minimal webhook:

```go
dsp := echotron.NewDispatcher("MY_TOKEN", newBot)
dsp.ListenWebhook("https://example.com:443/MY_TOKEN")
```

Integrated into an existing HTTP server without displacing your own routes:

```go
mux := http.NewServeMux()
mux.HandleFunc("/api/v1/login", loginHandler)

server := &http.Server{Addr: ":8080", Handler: mux}
dsp.SetHTTPServer(server)
dsp.ListenWebhook("https://example.com:8080/MY_TOKEN")
```

Gzip-compressed payloads from Telegram are handled transparently.

### Direct API parity

Echotron maps 1-to-1 to the [official Telegram Bot API](https://core.telegram.org/bots/api). Method names are identical, just capitalised as required by Go:

```
sendMessage         → SendMessage
sendPhoto           → SendPhoto
answerCallbackQuery → AnswerCallbackQuery
```

No invented abstractions sit between you and the Telegram docs.

### Type-safe options and enum constants

Optional parameters are typed structs, not variadic `interface{}` bags. Enum values are compile-time constants so the compiler catches typos:

```go
b.SendMessage("*bold*", b.chatID, &echotron.MessageOptions{
    ParseMode: echotron.MarkdownV2,
})

b.SendChatAction(b.chatID, echotron.Typing)
```

Pass `nil` for optional parameters when you do not need them.

### Structured API errors

Errors from Telegram are typed `*APIError` values, not raw strings, so you can inspect the error code and description separately:

```go
_, err := api.SendMessage("hello", chatID, nil)
var apiErr *echotron.APIError
if errors.As(err, &apiErr) {
    fmt.Println(apiErr.ErrorCode(), apiErr.Description())
}
```

### Local Bot API server support

Running a [Telegram Local Bot API](https://github.com/tdlib/telegram-bot-api) server for increased file size limits and upload throughput? One function call is all it takes:

```go
api := echotron.CustomAPI("http://localhost:8081/bot", "MY_TOKEN")
```

All methods route through your local server with no further changes.

## Design gems

The following are a few of the internal decisions that make Echotron pleasant to work with and correct under concurrent load.

### A generic, type-safe sync.Map wrapper

Session storage is backed by a generic type declared as:

```go
type smap[K, V any] sync.Map
```

This thin wrapper over `sync.Map` provides full type safety without any runtime overhead. There are no `interface{}` assertions scattered through the codebase and no risk of a type mismatch crashing the dispatcher at runtime.

### Three levels of embedding, zero indirection

`API` embeds `*lclient` as an anonymous field. `lclient` carries the rate-limiter methods `SetGlobalRequestLimit` and `SetChatRequestLimit`. Because the user's bot struct embeds `echotron.API`, those same methods are promoted all the way to the bot struct level. Calling `b.SetGlobalRequestLimit(...)` from inside your bot just works, with no explicit delegation code anywhere.

### Atomic client initialisation without a mutex

When a new bot token is first used, Echotron creates an `http.Client` and stores it in a package-level cache keyed by base URL. The initialisation uses `sync.Map.LoadOrStore` atomically: if two goroutines race to create the client at the same time, exactly one wins and both receive the same pointer. No mutex, no double-checked locking pattern, no `sync.Once` per token.

### `Update.ChatID()` absorbs all update types

The `ChatID()` method on `Update` inspects every possible update variant, from plain messages and callback queries to business connections, story interactions, chat boosts, and more, and always returns the correct `int64`. The dispatcher calls this one method to route any update to the right bot instance, regardless of its type.

### `InputFile` is a sealed type

`InputFile` has only unexported fields. The only way to create one is through:

```go
echotron.NewInputFileID("AgACAgI...")        // existing Telegram file
echotron.NewInputFilePath("photo.jpg")       // local file on disk
echotron.NewInputFileBytes("img.png", data)  // in-memory bytes
```

It is structurally impossible to create an `InputFile` in an invalid state from outside the package.

### A single dispatch choke point

Every API call, whether a plain GET, a form-encoded POST, or a multipart file upload, passes through the internal `lclient.dispatch()` function. Rate limiting, HTTP execution, JSON decoding, and error checking all happen exactly once, in one place. There is no duplicated error handling across the hundreds of API methods.

### The `stateFn` self-referential type

`type stateFn func(*Update) stateFn` is a function type that returns itself. This single line enables recursive, allocation-free state machines where the current state is simply the function that will handle the next update. No string-keyed state table, no `iota` enum, no external dependency.

## Echotron vs. other Go Telegram libraries

| Feature | **Echotron** | go-telegram-bot-api | telebot | gotgbot |
|---|:---:|:---:|:---:|:---:|
| Per-chat isolated state | ✅ built-in | ❌ manual | ❌ manual | ❌ manual |
| Concurrency model | ✅ built-in | ❌ manual | ⚠️ partial | ❌ manual |
| Built-in rate limiting | ✅ dual-level | ❌ | ❌ | ❌ |
| Shared HTTP client pool | ✅ per-token | ❌ | ❌ | ❌ |
| 1:1 Telegram API parity | ✅ | ✅ | ⚠️ | ✅ |
| Type-safe enums | ✅ | ⚠️ | ⚠️ | ⚠️ |
| Typed API errors | ✅ | ⚠️ | ⚠️ | ⚠️ |
| Functional state machines | ✅ first-class | ❌ | ❌ | ❌ |
| Session self-destruction | ✅ | ❌ | ❌ | ❌ |
| Custom HTTP server | ✅ | ⚠️ | ⚠️ | ⚠️ |
| Local API server support | ✅ | ✅ | ✅ | ✅ |
| No mandatory abstractions | ✅ | ✅ | ❌ | ✅ |
| External dependencies | 1 | 0 | several | 0 |
| License | LGPL-3.0 | MIT | MIT | MIT |

On the topic of dependencies: Echotron's single dependency is [`golang.org/x/time`](https://pkg.go.dev/golang.org/x/time), which lives in the `golang.org/x` namespace. That namespace is maintained by the Go team itself, under the same review standards and stability guarantees as the standard library. In practice, adding Echotron to your project means depending on the Go team's own code and nothing else.

**On the license:** LGPL-3.0 allows you to use Echotron in closed-source and commercial products without releasing your own code, provided you do not modify Echotron itself.

## More examples

### Sending files

```go
// From a local path
b.SendPhoto(b.chatID, echotron.NewInputFilePath("photo.jpg"), nil)

// By Telegram file_id, no re-upload
b.SendPhoto(b.chatID, echotron.NewInputFileID("AgACAgI..."), nil)

// From raw bytes already in memory
b.SendPhoto(b.chatID, echotron.NewInputFileBytes("photo.jpg", data), nil)
```

### Media groups

```go
b.SendMediaGroup(b.chatID, []echotron.GroupableInputMedia{
    echotron.InputMediaPhoto{Media: echotron.NewInputFilePath("a.jpg")},
    echotron.InputMediaPhoto{Media: echotron.NewInputFilePath("b.jpg")},
    echotron.InputMediaVideo{Media: echotron.NewInputFilePath("clip.mp4")},
}, nil)
```

### Inline keyboards

```go
b.SendMessage("Choose an option:", b.chatID, &echotron.MessageOptions{
    ReplyMarkup: echotron.InlineKeyboardMarkup{
        InlineKeyboard: [][]echotron.InlineKeyboardButton{
            {{Text: "Option A", CallbackData: "a"}},
            {{Text: "Option B", CallbackData: "b"}},
        },
    },
})
```

### Handling callback queries

```go
func (b *bot) Update(u *echotron.Update) {
    // Telegram guarantees that at most one field in an Update is non-nil,
    // so a switch on nil checks is the idiomatic way to route update types.
    switch {
    case u.CallbackQuery != nil:
        // AnswerCallbackQuery must be called to dismiss the loading indicator
        // shown by Telegram on the user's side after they tap a button.
        b.AnswerCallbackQuery(u.CallbackQuery.ID, nil)
        b.SendMessage("You chose: "+u.CallbackQuery.Data, b.chatID, nil)
    case u.Message != nil:
        b.SendMessage("Send me a button press!", b.chatID, nil)
    }
}
```

### Custom rate limits

```go
api := echotron.NewAPI("MY_TOKEN")

api.SetGlobalRequestLimit(time.Second/50, 50)  // 50 req/s globally
api.SetChatRequestLimit(time.Second, 1)        // 1 msg/s per chat
```

## Installation

```bash
go get github.com/NicoNex/echotron/v3
```

Go 1.21 or later is required.

## Starter templates

Every example in the [`example/`](./example/) directory is a self-contained Go
module usable as a project skeleton with
[`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew).

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

| Template | What it gives you |
|---|---|
| `polling-simple` | Minimal stateless bot, plain update channel, no `Dispatcher` |
| `polling` | Per-chat stateful bot with `Dispatcher` and long-polling |
| `polling-keyboard` | Inline keyboards and callback query handling |
| `polling-inline` | Inline mode (`@botname <query>` from any chat) |
| `polling-ratelimit` | Rate limiter configuration |
| `polling-fsm` | Multi-step conversations via functional state machines |
| `polling-fsm-lifecycle` | FSM + session self-destruction on idle timeout |
| `polling-fsm-persistence` | FSM + disk persistence with [katalis](https://github.com/NicoNex/katalis) |
| `webhook` | `Dispatcher` with webhook delivery |
| `webhook-simple` | Minimal stateless bot on webhooks |

Clone any template and rename the module in one command:

```bash
gonew github.com/NicoNex/echotron/v3/example/polling github.com/you/mybot
cd mybot
TELEGRAM_TOKEN=<your-token> go run .
```

See [`example/README.md`](example/README.md) for the full list and a suggested reading order.

## Links

- [pkg.go.dev documentation](https://pkg.go.dev/github.com/NicoNex/echotron/v3)
- [Telegram Bot API reference](https://core.telegram.org/bots/api)
- [Echotron News on Telegram](https://t.me/echotronnews)
- [Report an issue](https://github.com/NicoNex/echotron/issues)

## License

Echotron is free software released under the **GNU Lesser General Public License v3.0**.
See [LICENSE](LICENSE) for details.
