package echotron

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name        string
		i           any
		val         url.Values
		except      url.Values
		description string
	}{
		{
			name: "#1",
			i: CommandOptions{
				LanguageCode: "ru",
				Scope:        BotCommandScope{Type: BCSTChat, ChatID: 33288},
			},
			val: url.Values{
				"foo": {"bar"},
			},

			except: url.Values{
				"foo":           {"bar"},
				"language_code": {"ru"},
				"scope":         {"{\"type\":\"chat\",\"chat_id\":33288,\"user_id\":0}"},
			},
			description: "Scopes doesn't serialized",
		},
		{
			// TODO: i don't know what ist...
			name: "#2",
			i: CommandOptions{
				Scope: BotCommandScope{Type: BCSTChat, ChatID: 0101010},
			},
			val: url.Values{},
			except: url.Values{
				"scope": {"{\"type\":\"chat\",\"chat_id\":0101010,\"user_id\":0}"},
			},
			description: "Invalid ChatID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := scan(tt.i, tt.val)
			assert.Equal(t, tt.except, result, tt.description)
		})
	}
}
