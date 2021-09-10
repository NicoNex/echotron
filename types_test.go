package echotron

import "testing"

func TestInputMediaPhotoImplements(_ *testing.T) {
	i := InputMediaPhoto{}
	i.GetMedia()
	i.ImplementsInputMediaGroupable()
}

func TestInputMediaVideoImplements(_ *testing.T) {
	i := InputMediaVideo{}
	i.GetMedia()
	i.ImplementsInputMediaGroupable()
}

func TestInputMediaAnimationImplements(_ *testing.T) {
	i := InputMediaAnimation{}
	i.GetMedia()
}

func TestInputMediaAudioImplements(_ *testing.T) {
	i := InputMediaAudio{}
	i.GetMedia()
	i.ImplementsInputMediaGroupable()
}

func TestInputMediaDocumentImplements(_ *testing.T) {
	i := InputMediaDocument{}
	i.GetMedia()
	i.ImplementsInputMediaGroupable()
}
