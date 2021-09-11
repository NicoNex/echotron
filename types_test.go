package echotron

import "testing"

func TestInputMediaPhoto(_ *testing.T) {
	i := InputMediaPhoto{}
	i.media()
	i.thumb()
	i.groupable()
}

func TestInputMediaVideo(_ *testing.T) {
	i := InputMediaVideo{}
	i.media()
	i.thumb()
	i.groupable()
}

func TestInputMediaAnimation(_ *testing.T) {
	i := InputMediaAnimation{}
	i.media()
	i.thumb()
}

func TestInputMediaAudio(_ *testing.T) {
	i := InputMediaAudio{}
	i.media()
	i.thumb()
	i.groupable()
}

func TestInputMediaDocument(_ *testing.T) {
	i := InputMediaDocument{}
	i.media()
	i.thumb()
	i.groupable()
}
