# echotron [![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/) [![PkgGoDev](https://pkg.go.dev/badge/github.com/NicoNex/echotron)](https://pkg.go.dev/github.com/NicoNex/echotron) [![Go Report Card](https://goreportcard.com/badge/github.com/NicoNex/echotron)](https://goreportcard.com/report/github.com/NicoNex/echotron) [![License](http://img.shields.io/badge/license-LGPL3.0-orange.svg?style=flat)](https://github.com/NicoNex/echotron/blob/master/LICENSE)

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
    echotron.Api
}

const TOKEN = "YOUR TELEGRAM TOKEN"

func newBot(chatId int64) echotron.Bot {
    return &bot{
        chatId,
        echotron.NewApi(TOKEN),
    }
}

func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello world", b.chatId)
    }
}

func main() {
    dsp := echotron.NewDispatcher(TOKEN, newBot)
    dsp.Run()
}
```


Also proof of concept with self destruction for low ram usage

```go
package main

import (
    "time"

    "github.com/NicoNex/echotron"
)

type bot struct {
    chatId int64
    echotron.Api
}

const TOKEN = "YOUR TELEGRAM TOKEN"

var dsp echotron.Dispatcher

func newBot(chatId int64) echotron.Bot {
    var bot = &bot{
        chatId,
        echotron.NewApi(TOKEN),
    }
    go bot.selfDestruct(time.After(time.Hour))
    return bot
}

func (b *bot) selfDestruct(timech <- chan time.Time) {
    select {
    case <-timech:
        b.SendMessage("goodbye", b.chatId)
        dsp.DelSession(b.chatId)
    }
}

func (b *bot) Update(update *echotron.Update) {
    if update.Message.Text == "/start" {
        b.SendMessage("Hello world", b.chatId)
    }
}

func main() {
    dsp = echotron.NewDispatcher(TOKEN, newBot)
    dsp.Run()
}
```
