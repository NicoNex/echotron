package echotron

import (
	"fmt"
	"net/url"
)

// webhookListenAddr returns the TCP address for the local webhook listener.
// When the webhook URL omits an explicit port, https defaults to :443 and
// http (and anything else) defaults to :80, matching the documented examples
// that use scheme-default ports without spelling them out in the URL.
func webhookListenAddr(u *url.URL) (string, error) {
	if u == nil {
		return "", fmt.Errorf("webhook URL is nil")
	}
	if u.Hostname() == "" {
		return "", fmt.Errorf("webhook URL must include a host")
	}
	port := u.Port()
	if port == "" {
		switch u.Scheme {
		case "https":
			port = "443"
		default:
			port = "80"
		}
	}
	return ":" + port, nil
}
