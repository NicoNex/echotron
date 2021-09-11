package echotron

import "testing"

func TestInputMediaPhotoImplements(_ *testing.T) {
	i := InputMediaPhoto{}
	i.media()
	i.groupable()
}

func TestInputMediaVideoImplements(_ *testing.T) {
	i := InputMediaVideo{}
	i.media()
	i.groupable()
}

func TestInputMediaAnimationImplements(_ *testing.T) {
	i := InputMediaAnimation{}
	i.media()
}

func TestInputMediaAudioImplements(_ *testing.T) {
	i := InputMediaAudio{}
	i.media()
	i.groupable()
}

func TestInputMediaDocumentImplements(_ *testing.T) {
	i := InputMediaDocument{}
	i.media()
	i.groupable()
}
