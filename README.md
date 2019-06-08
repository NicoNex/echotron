# echotron

The true echotron repo

Fetch with
`env GIT_TERMINAL_PROMPT=1 go get -u gitlab.com/NicoNex/echotron`


### Usage
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
	b.SendMessage("Hello world", b.chatId)
}


func main() {
	echotron.RunDispatcher("568059758:AAFRN3Xg3dOkfe2n0gNlOWjlkM6dihommPQ", NewBot)
}
```

