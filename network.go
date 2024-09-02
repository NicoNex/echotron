/*
 * Echotron
 * Copyright (C) 2018-2022 The Echotron Devs
 *
 * Echotron is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Echotron is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type client struct {
	httpClient *http.Client
	globalRL   *rate.Limiter
	mu         sync.Mutex
	chats      map[string]*chatClient
	config     rateLimitConfig
}

type rateLimitConfig struct {
	enabled bool
	rps     rate.Limit
	burst   int
	cleanup time.Duration
}

type chatClient struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

/*
When initializing the Echotron client, the following default rate limit settings are applied:

Global Rate Limit:
30 requests per second
Burst of 10 requests

Per-Chat Rate Limit:
20 requests per minute
Burst of 5 requests
Cleanup interval of 1 minute

Cleanup Interval: 1 minute
Rate Limiting: Enabled by default
 */

func newClient() *client {
	c := &client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		globalRL:   rate.NewLimiter(rate.Every(time.Second/30), 10), // Default global rate limit
		chats:      make(map[string]*chatClient),
		config: rateLimitConfig{
			enabled: true,
			rps:     rate.Every(time.Minute / 20), // Default per-chat rate limit: 20 requests per minute
			burst:   5,
			cleanup: time.Minute,
		},
	}

	// Start the cleanup goroutine
	go c.cleanupChats()

	return c
}

// Helper function to create a new chat client with a rate limiter
func newChatClient(rps rate.Limit, burst int) *chatClient {
	return &chatClient{
		// Create a new rate limiter with the specified rate and burst
		limiter:  rate.NewLimiter(rps, burst),
		// Set the initial last seen time to now
		lastSeen: time.Now(),
	}
}

// Internal global client instance
var lclient = newClient()

// SetGlobalRequestLimit allows changing the global rate limit
func SetGlobalRequestLimit(rps rate.Limit, burst int) {
	lclient.mu.Lock()
	defer lclient.mu.Unlock()
	lclient.globalRL = rate.NewLimiter(rps, burst)
}

// SetChatRequestLimit function should be updated to reflect these changes
func SetChatRequestLimit(rps rate.Limit, burst int, cleanup time.Duration) {
	lclient.mu.Lock()
	defer lclient.mu.Unlock()

	lclient.config.rps = rps
	lclient.config.burst = burst
	lclient.config.cleanup = cleanup

	// Reset existing chat limiters to apply new settings
	for _, chat := range lclient.chats {
		chat.limiter = rate.NewLimiter(rps, burst)
	}
}

// SetRateLimiterEnabled allows enabling or disabling the rate limiter
func SetRateLimiterEnabled(enabled bool) {
	lclient.mu.Lock()
	defer lclient.mu.Unlock()
	lclient.config.enabled = enabled
}

// Cleanup goroutine (typically started in init() or similar)
func (c *client) cleanupChats() {
	for {
		time.Sleep(c.config.cleanup)
		c.mu.Lock()
		now := time.Now()
		for chatID, chat := range c.chats {
			if now.Sub(chat.lastSeen) > 3*c.config.cleanup {
				delete(c.chats, chatID)
			}
		}
		c.mu.Unlock()
	}
}

// Internal function to get or create a limiter for a specific chat
func (c *client) getLimiter(chatID string) *rate.Limiter {
	// If rate limiting is disabled or no chat ID is provided, it returns the global rate limiter.
	if !c.config.enabled || chatID == "" {
		return c.globalRL
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	// Check if a limiter already exists for this chat ID
	if chat, found := c.chats[chatID]; found {
		// If found, update the last seen time to prevent premature cleanup
		chat.lastSeen = time.Now()
		// Return the existing limiter for this chat
		return chat.limiter
	}

	// If no limiter exists for this chat, create a new one
	// Use the current per-chat rate limit settings from the config
	newChat := newChatClient(c.config.rps, c.config.burst)
	// Store the new chat client in the chats map
	c.chats[chatID] = newChat
	// Return the limiter of the newly created chat client
	return newChat.limiter
}

func (c *client) doGet(reqURL string) ([]byte, error) {
	resp, err := c.httpClient.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *client) doPost(reqURL string, files ...content) ([]byte, error) {
	var (
		buf = new(bytes.Buffer)
		w   = multipart.NewWriter(buf)
	)

	for _, f := range files {
		part, err := w.CreateFormFile(f.ftype, filepath.Base(f.fname))
		if err != nil {
			return nil, err
		}
		part.Write(f.fdata)
	}
	w.Close()

	req, err := http.NewRequest(http.MethodPost, reqURL, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", w.FormDataContentType())

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func (c *client) doPostForm(reqURL string, keyVals map[string]string) ([]byte, error) {
	form := make(url.Values)
	for k, v := range keyVals {
		form.Add(k, v)
	}

	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// This method doesn't need rate limiting as it's called by postFile which already has rate limiting
func (c *client) sendFile(file, thumbnail InputFile, url, fileType string) (res []byte, err error) {
	var cnt []content

	if file.id != "" {
		url = fmt.Sprintf("%s&%s=%s", url, fileType, file.id)
	} else if file.url != "" {
		url = fmt.Sprintf("%s&%s=%s", url, fileType, file.url)
	} else if c, e := toContent(fileType, file); e == nil {
		cnt = append(cnt, c)
	} else {
		err = e
	}

	if c, e := toContent("thumbnail", thumbnail); e == nil {
		cnt = append(cnt, c)
	} else {
		err = e
	}

	if len(cnt) > 0 {
		res, err = c.doPost(url, cnt...)
	} else {
		res, err = c.doGet(url)
	}
	return
}

func (c *client) get(base, endpoint string, vals url.Values, v APIResponse) error {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}

	chatID := vals.Get("chat_id")
	if err := c.waitForLimiter(chatID); err != nil {
		return err
	}

	cnt, err := c.doGet(url)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(cnt, v); err != nil {
		return err
	}
	return check(v)
}


func (c *client) postFile(base, endpoint, fileType string, file, thumbnail InputFile, vals url.Values, v APIResponse) error {
	url, err := joinURL(base, endpoint, vals)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	chatID := vals.Get("chat_id")
	if err := c.waitForLimiter(chatID); err != nil {
		return err
	}

	cnt, err := c.sendFile(file, thumbnail, url, fileType)
	if err != nil {
		return fmt.Errorf("sending file: %w", err)
	}

	if err := json.Unmarshal(cnt, v); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}

	return check(v)
}

func (c *client) postMedia(base, endpoint string, editSingle bool, vals url.Values, v APIResponse, files ...InputMedia) error {
	url, err := joinURL(base, endpoint, vals)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	chatID := vals.Get("chat_id")
	if err := c.waitForLimiter(chatID); err != nil {
		return err
	}

	cnt, err := c.sendMediaFiles(url, editSingle, files...)
	if err != nil {
		return fmt.Errorf("sending media files: %w", err)
	}

	if err := json.Unmarshal(cnt, v); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}
	return check(v)
}

func (c *client) postStickers(base, endpoint string, vals url.Values, v APIResponse, stickers ...InputSticker) error {
	url, err := joinURL(base, endpoint, vals)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	chatID := vals.Get("chat_id")
	if err := c.waitForLimiter(chatID); err != nil {
		return err
	}

	cnt, err := c.sendStickers(url, stickers...)
	if err != nil {
		return fmt.Errorf("sending stickers: %w", err)
	}
	if err := json.Unmarshal(cnt, v); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}
	return check(v)
}

// This method doesn't need rate limiting as it's called by postMedia which already has rate limiting
func (c *client) sendMediaFiles(url string, editSingle bool, files ...InputMedia) (res []byte, err error) {
	var (
		med []mediaEnvelope
		cnt []content
		jsn []byte
	)

	for _, file := range files {
		var im mediaEnvelope
		var cntArr []content

		media := file.media()
		thumbnail := file.thumbnail()

		im, cntArr, err = processMedia(media, thumbnail)
		if err != nil {
			return nil, fmt.Errorf("processing media: %w", err)
		}

		im.InputMedia = file

		med = append(med, im)
		cnt = append(cnt, cntArr...)
	}

	if editSingle {
		jsn, err = json.Marshal(med[0])
	} else {
		jsn, err = json.Marshal(med)
	}

	if err != nil {
		return nil, fmt.Errorf("marshaling media: %w", err)
	}

	url = fmt.Sprintf("%s&media=%s", url, jsn)

	if len(cnt) > 0 {
		return c.doPost(url, cnt...)
	}
	return c.doGet(url)
}


// This method doesn't need rate limiting as it's called by postStickers which already has rate limiting
func (c *client) sendStickers(url string, stickers ...InputSticker) (res []byte, err error) {
	var (
		sti []stickerEnvelope
		cnt []content
		jsn []byte
	)

	for _, s := range stickers {
		var se stickerEnvelope
		var cntArr []content

		se, cntArr, err = processSticker(s.Sticker)
		if err != nil {
			return nil, fmt.Errorf("processing sticker: %w", err)
		}

		se.InputSticker = s

		sti = append(sti, se)
		cnt = append(cnt, cntArr...)
	}

	if len(sti) == 1 {
		jsn, err = json.Marshal(sti[0])
		url = fmt.Sprintf("%s&sticker=%s", url, jsn)
	} else {
		jsn, err = json.Marshal(sti)
		url = fmt.Sprintf("%s&stickers=%s", url, jsn)
	}

	if err != nil {
		return nil, fmt.Errorf("marshaling stickers: %w", err)
	}

	if len(cnt) > 0 {
		return c.doPost(url, cnt...)
	}
	return c.doGet(url)
}

// waitForLimiter is called before making any API request
func (c *client) waitForLimiter(chatID string) error {
	limiter := c.getLimiter(chatID)
	if err := limiter.Wait(context.Background()); err != nil {
		return fmt.Errorf("rate limit exceeded: %w", err)
	}
	return nil
}
