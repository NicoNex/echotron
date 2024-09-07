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
	"errors"
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

var (
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
)

type client struct {
	httpClient  *http.Client
	globalRL    *rate.Limiter
	mu          sync.Mutex
	chats       map[string]*chatClient
	config      rateLimitConfig
	stopCleanup context.CancelFunc
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

func newClient() *client {
	ctx, cancel := context.WithCancel(context.Background())
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
		stopCleanup: cancel,
	}

	// Start the cleanup goroutine
	go c.cleanupChats(ctx)

	return c
}

// Helper function to create a new chat client with a rate limiter
func newChatClient(rps rate.Limit, burst int) *chatClient {
	return &chatClient{
		limiter:  rate.NewLimiter(rps, burst),
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

// Cleanup goroutine
func (c *client) cleanupChats(ctx context.Context) {
	ticker := time.NewTicker(c.config.cleanup)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
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
}

// StopClient stops the cleanup goroutine and releases resources
func StopClient() {
	lclient.stopCleanup()
}

// Internal function to get or create a limiter for a specific chat
func (c *client) getLimiter(chatID string) *rate.Limiter {
	if !c.config.enabled || chatID == "" {
		return c.globalRL
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if chat, found := c.chats[chatID]; found {
		chat.lastSeen = time.Now()
		return chat.limiter
	}

	newChat := newChatClient(c.config.rps, c.config.burst)
	c.chats[chatID] = newChat
	return newChat.limiter
}

func (c *client) doGet(ctx context.Context, reqURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *client) doPost(ctx context.Context, reqURL string, files ...content) ([]byte, error) {
	var (
		buf = new(bytes.Buffer)
		w   = multipart.NewWriter(buf)
	)

	for _, f := range files {
		part, err := w.CreateFormFile(f.ftype, filepath.Base(f.fname))
		if err != nil {
			return nil, err
		}
		_, err = part.Write(f.fdata)
		if err != nil {
			return nil, err
		}
	}
	err := w.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, buf)
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

func (c *client) doPostForm(ctx context.Context, reqURL string, keyVals map[string]string) ([]byte, error) {
	form := make(url.Values)
	for k, v := range keyVals {
		form.Add(k, v)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, strings.NewReader(form.Encode()))
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
func (c *client) sendFile(ctx context.Context, file, thumbnail InputFile, url, fileType string) (res []byte, err error) {
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
		res, err = c.doPost(ctx, url, cnt...)
	} else {
		res, err = c.doGet(ctx, url)
	}
	return
}

func (c *client) get(ctx context.Context, base, endpoint string, vals url.Values, v APIResponse) error {
	fullURL, err := url.JoinPath(base, endpoint)
	if err != nil {
		return err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			fullURL = fmt.Sprintf("%s?%s", fullURL, queries)
		}
	}

	if err := c.waitForLimiter(ctx, vals.Get("chat_id")); err != nil {
		return err
	}

	cnt, err := c.doGet(ctx, fullURL)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(cnt, v); err != nil {
		return err
	}
	return check(v)
}

func (c *client) postFile(ctx context.Context, base, endpoint, fileType string, file, thumbnail InputFile, vals url.Values, v APIResponse) error {
	fullURL, err := joinURL(base, endpoint, vals)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	if err := c.waitForLimiter(ctx, vals.Get("chat_id")); err != nil {
		return err
	}

	cnt, err := c.sendFile(ctx, file, thumbnail, fullURL, fileType)
	if err != nil {
		return fmt.Errorf("sending file: %w", err)
	}

	if err := json.Unmarshal(cnt, v); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}

	return check(v)
}

func (c *client) postMedia(ctx context.Context, base, endpoint string, editSingle bool, vals url.Values, v APIResponse, files ...InputMedia) error {
	fullURL, err := joinURL(base, endpoint, vals)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	if err := c.waitForLimiter(ctx, vals.Get("chat_id")); err != nil {
		return err
	}

	cnt, err := c.sendMediaFiles(ctx, fullURL, editSingle, files...)
	if err != nil {
		return fmt.Errorf("sending media files: %w", err)
	}

	if err := json.Unmarshal(cnt, v); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}
	return check(v)
}

func (c *client) postStickers(ctx context.Context, base, endpoint string, vals url.Values, v APIResponse, stickers ...InputSticker) error {
	fullURL, err := joinURL(base, endpoint, vals)
	if err != nil {
		return fmt.Errorf("joining URL: %w", err)
	}

	if err := c.waitForLimiter(ctx, vals.Get("chat_id")); err != nil {
		return err
	}

	cnt, err := c.sendStickers(ctx, fullURL, stickers...)
	if err != nil {
		return fmt.Errorf("sending stickers: %w", err)
	}
	if err := json.Unmarshal(cnt, v); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}
	return check(v)
}

// This method doesn't need rate limiting as it's called by postMedia which already has rate limiting
func (c *client) sendMediaFiles(ctx context.Context, url string, editSingle bool, files ...InputMedia) (res []byte, err error) {
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
		return c.doPost(ctx, url, cnt...)
	}
	return c.doGet(ctx, url)
}

// This method doesn't need rate limiting as it's called by postStickers which already has rate limiting
// This method doesn't need rate limiting as it's called by postStickers which already has rate limiting
func (c *client) sendStickers(ctx context.Context, url string, stickers ...InputSticker) (res []byte, err error) {
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
		return c.doPost(ctx, url, cnt...)
	}
	return c.doGet(ctx, url)
}

// waitForLimiter is called before making any API request
func (c *client) waitForLimiter(ctx context.Context, chatID string) error {
	limiter := c.getLimiter(chatID)
	if err := limiter.Wait(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("%w: %v", ErrRateLimitExceeded, err)
		}
		return err
	}
	return nil
}
