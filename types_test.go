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

func TestChatMemberOwner(_ *testing.T) {
	i := ChatMemberOwner{}
	i.ImplementsChatMember()
}

func TestChatMemberAdministrator(_ *testing.T) {
	i := ChatMemberAdministrator{}
	i.ImplementsChatMember()
}

func TestChatMemberMember(_ *testing.T) {
	i := ChatMemberMember{}
	i.ImplementsChatMember()
}

func TestChatMemberRestricted(_ *testing.T) {
	i := ChatMemberRestricted{}
	i.ImplementsChatMember()
}

func TestChatMemberLeft(_ *testing.T) {
	i := ChatMemberLeft{}
	i.ImplementsChatMember()
}

func TestChatMemberBanned(_ *testing.T) {
	i := ChatMemberBanned{}
	i.ImplementsChatMember()
}
