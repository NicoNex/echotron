# echotron

The true echotron repo

Fetch with
`env GIT_TERMINAL_PROMPT=1 go get -u gitlab.com/NicoNex/echotron`


### Usage

A very simple implementation:

```go
package main

import "gitlab.com/NicoNex/echotron"

type bot struct {
	chatId int64
	*echotron.Engine
}


func NewBot(token string, chatId int64) echotron.Bot {
	return &bot{
		chatId,
		echotron.NewEngine(token),
	}
}


func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello world", b.chatId)
    }
}


func main() {
	echotron.RunDispatcher("TELEGRAM TOKEN", NewBot)
}
```


Also proof of concept with self destruction for low ram usage

```go
package main

import "gitlab.com/NicoNex/echotron"

type bot struct {
    chatId int64
    *echotron.Engine
}


func NewBot(token string, chatId int64) echotron.Bot {
    var bot = &bot{
        chatId,
        echotron.NewEngine(token),
    }
    echotron.AddTimer(bot.chatId, "selfDestruct", bot.selfDestruct, 60)
    return bot
}


func (b *bot) selfDestruct() {
    b.SendMessage("goodbye", b.chatId)
    echotron.DelTimer(b.chatId, "selfDestruct")
    echotron.DelSession(b.chatId)
}


func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello world", b.chatId)
    }
}


func main() {
    echotron.RunDispatcher("TELEGRAM TOKEN", NewBot)
}
```


