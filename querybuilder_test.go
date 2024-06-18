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
