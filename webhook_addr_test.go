package echotron

import (
	"net/url"
	"testing"
)

func TestWebhookListenAddr(t *testing.T) {
	tests := []struct {
		raw     string
		want    string
		wantErr bool
	}{
		{raw: "https://example.com:8443/bot", want: ":8443"},
		{raw: "https://example.com/bot", want: ":443"},
		{raw: "http://example.com/bot", want: ":80"},
		{raw: "http://example.com:8080/bot", want: ":8080"},
		{raw: "https:///bot", wantErr: true},
		{raw: "/bot", wantErr: true},
	}
	for _, tc := range tests {
		u, err := url.Parse(tc.raw)
		if err != nil {
			t.Fatalf("parse %q: %v", tc.raw, err)
		}
		got, err := webhookListenAddr(u)
		if tc.wantErr {
			if err == nil {
				t.Fatalf("%q: expected error", tc.raw)
			}
			continue
		}
		if err != nil {
			t.Fatalf("%q: unexpected error: %v", tc.raw, err)
		}
		if got != tc.want {
			t.Fatalf("%q: got %q want %q", tc.raw, got, tc.want)
		}
	}
}
