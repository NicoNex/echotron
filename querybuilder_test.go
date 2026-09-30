package echotron

import (
	"net/url"
	"reflect"
	"testing"
)

type scanTest struct {
	i          any
	predefined url.Values
	expected   url.Values
}

func TestScan(t *testing.T) {
	tests := []scanTest{
		{
			i: CommandOptions{
				LanguageCode: "it",
				Scope:        BotCommandScope{Type: BCSTChat, ChatID: 33288},
			},
			predefined: url.Values{"foo": {"bar"}},
			expected: url.Values{
				"foo":           {"bar"},
				"language_code": {"it"},
				"scope":         {`{"type":"chat","chat_id":33288,"user_id":0}`},
			},
		},
	}

	for i, tt := range tests {
		result := scan(tt.i, tt.predefined)
		if !reflect.DeepEqual(tt.expected, result) {
			t.Fatalf("test #%d: result differs from expected value\n", i)
		}
	}
}

func TestToStringDefault(t *testing.T) {
	ret := toString(reflect.ValueOf(nil))

	if ret != "" {
		t.Fatalf("expected empty string, got %+v", ret)
	}
}

func TestUrlValues(t *testing.T) {
	ret := urlValues(nil)

	if ret != nil {
		t.Fatalf("expected nil, got %+v", ret)
	}
}

func TestAddValues(t *testing.T) {
	vals := url.Values{}
	ret := addValues(vals, nil)

	if !reflect.DeepEqual(vals, ret) {
		t.Fatalf("expected nil, got %+v", ret)
	}
}

func TestAddValuesNil(t *testing.T) {
	opts := MessageOptions{ParseMode: MarkdownV2}
	vals := urlValues(opts)
	ret := addValues(nil, opts)

	if !reflect.DeepEqual(vals, ret) {
		t.Fatalf("expected %+v, got %+v", vals, ret)
	}
}

func TestScanPointer(t *testing.T) {
	vals := urlValues(MessageOptions{
		EphemeralMessageParameters: &EphemeralMessageParameters{ReceiverUserID: 42},
	})

	if exp := `{"receiver_user_id":42}`; vals.Get("ephemeral_message_parameters") != exp {
		t.Fatalf("expected %s, got %s", exp, vals.Get("ephemeral_message_parameters"))
	}
}

func TestScanEphemeralMessage(t *testing.T) {
	vals := urlValues(ephemeralMessage{1, 2, 3})
	exp := url.Values{"chat_id": {"1"}, "receiver_user_id": {"2"}, "ephemeral_message_id": {"3"}}

	if !reflect.DeepEqual(vals, exp) {
		t.Fatalf("expected %v, got %v", exp, vals)
	}
}

func TestScanPromoteOptions(t *testing.T) {
	vals := urlValues(PromoteOptions{CanManageTags: true})

	if vals.Get("can_manage_tags") != "true" {
		t.Fatalf("unexpected values: %v", vals)
	}
}
