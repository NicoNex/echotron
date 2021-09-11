package echotron

import "testing"

func TestInputMediaPhotoImplements(_ *testing.T) {
	i := InputMediaPhoto{}
	i.getMedia()
	i.groupable()
}

func TestInputMediaVideoImplements(_ *testing.T) {
	i := InputMediaVideo{}
	i.getMedia()
	i.groupable()
}

func TestInputMediaAnimationImplements(_ *testing.T) {
	i := InputMediaAnimation{}
	i.getMedia()
}

func TestInputMediaAudioImplements(_ *testing.T) {
	i := InputMediaAudio{}
	i.getMedia()
	i.groupable()
}

func TestInputMediaDocumentImplements(_ *testing.T) {
	i := InputMediaDocument{}
	i.getMedia()
	i.groupable()
}
