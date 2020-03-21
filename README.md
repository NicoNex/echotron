# echotron [![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/) [![GoDoc](https://godoc.org/github.com/NicoNex/echotron?status.svg)](https://godoc.org/github.com/NicoNex/echotron) [![Go Report Card](https://goreportcard.com/badge/github.com/NicoNex/echotron)](https://goreportcard.com/report/github.com/NicoNex/echotron) [![License](http://img.shields.io/badge/license-LGPL3.0-orange.svg?style=flat)](https://github.com/NicoNex/echotron/blob/master/LICENSE)

Library for telegram bots written in pure go

Fetch with
```bash
go get -u github.com/NicoNex/echotron
```

### Usage

A very simple implementation:

```go
package main

import "github.com/NicoNex/echotron"

type bot struct {
    chatId int64
    echotron.Engine
}

func newBot(engine echotron.Engine, chatId int64) echotron.Bot {
    return &bot{
        chatId,
        engine,
    }
}

func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello world", b.chatId)
    }
}

func main() {
    dsp := echotron.NewDispatcher("TELEGRAM TOKEN", newBot)
    dsp.Run()
}
```


Also proof of concept with self destruction for low ram usage

```go
package main

import "github.com/NicoNex/echotron"

type bot struct {
    chatId int64
    echotron.Engine
}

var dsp echotron.Dispatcher

func newBot(engine echotron.Engine, chatId int64) echotron.Bot {
    var bot = &bot{
        chatId,
        engine,
    }
    echotron.AddTimer(bot.chatId, "selfDestruct", bot.selfDestruct, 60)
    return bot
}

func (b *bot) selfDestruct() {
    b.SendMessage("goodbye", b.chatId)
    echotron.DelTimer(b.chatId, "selfDestruct")
    dsp.DelSession(b.chatId)
}

func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello world", b.chatId)
    }
}

func main() {
    dsp = echotron.NewDispatcher("TELEGRAM TOKEN", newBot)
    dsp.Run()
}
```
