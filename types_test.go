package echotron

import "testing"

func TestInputMediaPhotoImplementsInputMedia(_ *testing.T) {
	i := InputMediaPhoto{}
	i.ImplementsInputMedia()
}

func TestInputMediaVideoImplementsInputMedia(_ *testing.T) {
	i := InputMediaVideo{}
	i.ImplementsInputMedia()
}

func TestInputMediaAnimationImplementsInputMedia(_ *testing.T) {
	i := InputMediaAnimation{}
	i.ImplementsInputMedia()
}

func TestInputMediaAudioImplementsInputMedia(_ *testing.T) {
	i := InputMediaAudio{}
	i.ImplementsInputMedia()
}

func TestInputMediaDocumentImplementsInputMedia(_ *testing.T) {
	i := InputMediaDocument{}
	i.ImplementsInputMedia()
}
