package echotron

import (
	"testing"
	"time"
)

type test struct{}

func (t test) Update(_ *Update) {}

var dsp *Dispatcher

func TestNewDispatcher(t *testing.T) {
	if dsp = NewDispatcher("token", func(_ int64) Bot { return test{} }); dsp == nil {
		t.Fatal("dispatcher is nil")
	}
}

func TestAddSession(t *testing.T) {
	err := dsp.AddSession(0)
	if err != nil {
		t.Fatalf("could not add session: %v", err)
	}

	_, ok := dsp.sessionMap.Load(int64(0))
	if !ok {
		t.Fatal("could not find added session")
	}
}

func TestDelSession(t *testing.T) {
	dsp.DelSession(0)

	_, ok := dsp.sessionMap.Load(int64(0))
	if ok {
		t.Fatal("session was not deleted")
	}
}

func TestListenWebhook(_ *testing.T) {
	go func() {
		_ = dsp.ListenWebhook("http://example.com:8443/test")
	}()
	time.Sleep(time.Second)
}

func TestPoll(t *testing.T) {
	go func() {
		_ = dsp.Poll()
	}()

	// Wait a bit for polling to start
	time.Sleep(time.Millisecond * 100)

	updates := []Update{
		{ChatJoinRequest: &ChatJoinRequest{Chat: Chat{ID: 0}}},
		{ChatBoost: &ChatBoostUpdated{Chat: Chat{ID: 0}}},
		{RemovedChatBoost: &ChatBoostRemoved{Chat: Chat{ID: 0}}},
		{Message: &Message{Chat: Chat{ID: 0}}},
		{EditedMessage: &Message{Chat: Chat{ID: 0}}},
		{ChannelPost: &Message{Chat: Chat{ID: 0}}},
		{EditedChannelPost: &Message{Chat: Chat{ID: 0}}},
		{BusinessConnection: &BusinessConnection{User: User{ID: 0}}},
		{BusinessMessage: &Message{Chat: Chat{ID: 0}}},
		{EditedBusinessMessage: &Message{Chat: Chat{ID: 0}}},
		{DeletedBusinessMessages: &BusinessMessagesDeleted{Chat: Chat{ID: 0}}},
		{MessageReaction: &MessageReactionUpdated{Chat: Chat{ID: 0}}},
		{MessageReactionCount: &MessageReactionCountUpdated{Chat: Chat{ID: 0}}},
		{InlineQuery: &InlineQuery{From: &User{ID: 0}}},
		{ChosenInlineResult: &ChosenInlineResult{From: &User{ID: 0}}},
		{CallbackQuery: &CallbackQuery{Message: &Message{Chat: Chat{ID: 0}}}},
		{ShippingQuery: &ShippingQuery{From: User{ID: 0}}},
		{PreCheckoutQuery: &PreCheckoutQuery{From: User{ID: 0}}},
		{PollAnswer: &PollAnswer{User: &User{ID: 0}}},
		{MyChatMember: &ChatMemberUpdated{Chat: Chat{ID: 0}}},
		{ChatMember: &ChatMemberUpdated{Chat: Chat{ID: 0}}},
	}

	for _, update := range updates {
		dsp.updates <- &update
	}

	// Wait for updates to be processed
	time.Sleep(time.Second)

	// Stop polling
	err := dsp.Stop()
	if err != nil {
		t.Fatal("Failed to stop dispatcher: %v", err)
	}
}
