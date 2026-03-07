# Echotron examples

Each subdirectory is a self-contained Go module that can be cloned as a project
skeleton with [`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew):

```bash
gonew github.com/NicoNex/echotron/v3/example/<name> <your-module-path>
```

For example:

```bash
gonew github.com/NicoNex/echotron/v3/example/polling github.com/you/mybot
```

## Examples

| Name | What it shows |
|---|---|
| [`polling-simple`](polling-simple/) | Stateless bot driven by a plain `PollingUpdates` channel; no `Dispatcher`. Shortest possible starting point. |
| [`polling`](polling/) | Per-chat stateful bot using `Dispatcher` with long-polling. One struct instance per conversation. |
| [`polling-keyboard`](polling-keyboard/) | Inline keyboards attached to messages and callback query handling. |
| [`polling-inline`](polling-inline/) | Inline mode: the bot responds to `@botname <query>` from any chat. |
| [`polling-ratelimit`](polling-ratelimit/) | Configuring the built-in dual-level rate limiter (global and per-chat). |
| [`polling-fsm`](polling-fsm/) | Finite state machine pattern via a self-referential `stateFn` type. Multi-step conversations without external libraries. |
| [`polling-fsm-lifecycle`](polling-fsm-lifecycle/) | FSM extended with session self-destruction: inactive bot instances remove themselves from the `Dispatcher` after a configurable idle timeout. |
| [`polling-fsm-persistence`](polling-fsm-persistence/) | FSM extended with disk persistence using [katalis](https://github.com/NicoNex/katalis). Per-chat data survives process restarts. |
| [`webhook`](webhook/) | `Dispatcher` with webhook delivery instead of long-polling. |
| [`webhook-simple`](webhook-simple/) | Stateless bot driven by a plain `WebhookUpdates` channel; no `Dispatcher`. |

## Suggested reading order

1. **`polling-simple`** — understand the raw update channel
2. **`polling`** — add per-chat state with the `Dispatcher`
3. **`polling-keyboard`** — interactive keyboards and callbacks
4. **`polling-inline`** — inline mode and `PollOptions`
5. **`polling-ratelimit`** — rate limiter configuration
6. **`polling-fsm`** — state machines for multi-step conversations
7. **`polling-fsm-lifecycle`** — memory management with session self-destruction
8. **`polling-fsm-persistence`** — durable state across restarts
9. **`webhook`** / **`webhook-simple`** — webhook delivery

## Prerequisites

- Go 1.21 or later
- [`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew): `go install golang.org/x/tools/cmd/gonew@latest`
- A bot token from [@BotFather](https://t.me/BotFather), exported as `TELEGRAM_TOKEN`
