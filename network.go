/*
 * Echotron
 * Copyright (C) 2018 The Echotron Contributors
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

// lclient is an HTTP client with built-in dual-level rate limiting.
// One instance is shared across all API objects for the same bot token.
type lclient struct {
	mu       sync.RWMutex
	cl       map[string]*rate.Limiter // per-chat rate limiter, keyed by chat_id
	gl       *rate.Limiter            // global rate limiter across all chats
	climiter func() *rate.Limiter     // factory for new per-chat limiters
	http     *http.Client
}

// clients caches one lclient per base URL (i.e. one per bot token).
var clients smap[string, *lclient]

func loadClient(url string) *lclient {
	// Fast path: client already exists for this base URL.
	if lc, ok := clients.load(url); ok {
		return lc
	}

	// Slow path: first access, create a new client and store it atomically.
	// If two goroutines race here, loadOrStore guarantees only one wins and
	// both get back the same *lclient.
	lc, _ := clients.loadOrStore(
		url,
		&lclient{
			http: new(http.Client),
			cl:   make(map[string]*rate.Limiter),
			gl:   rate.NewLimiter(rate.Every(time.Second/30), 30),
			climiter: func() *rate.Limiter {
				return rate.NewLimiter(rate.Every(time.Minute/20), 20)
			},
		},
	)
	return lc
}

// SetGlobalRequestLimit sets the global rate limit for requests to the Telegram API.
// An interval of 0 disables the rate limiter, allowing unlimited requests.
// By default the interval of this limiter is set to time.Second/30 and the
// burstSize is set to 30.
func (c *lclient) SetGlobalRequestLimit(interval time.Duration, burstSize int) {
	c.mu.Lock()
	c.gl = rate.NewLimiter(rate.Every(interval), burstSize)
	c.mu.Unlock()
}

// SetChatRequestLimit sets the per-chat rate limit for requests to the Telegram API.
// An interval of 0 disables the rate limiter, allowing unlimited requests.
// By default the interval of this limiter is set to time.Minute/20 and the
// burstSize is set to 20.
func (c *lclient) SetChatRequestLimit(interval time.Duration, burstSize int) {
	c.mu.Lock()
	// Reset the existing limiters so the new factory applies to all chats.
	c.cl = make(map[string]*rate.Limiter)
	c.climiter = func() *rate.Limiter {
		return rate.NewLimiter(rate.Every(interval), burstSize)
	}
	c.mu.Unlock()
}

// wait blocks until both the per-chat and global rate limiters allow the request.
func (c *lclient) wait(chatID string) error {
	ctx := context.Background()
	// If the chatID is empty, it's a general API call like GetUpdates, GetMe
	// and similar, so skip the per-chat request limit wait.
	if chatID != "" {
		c.mu.RLock()
		l, ok := c.cl[chatID]
		c.mu.RUnlock()

		if !ok {
			c.mu.Lock()
			// Re-check after acquiring the write lock to avoid overwriting
			// a limiter created by another goroutine in the meantime.
			if l, ok = c.cl[chatID]; !ok {
				l = c.climiter()
				c.cl[chatID] = l
			}
			c.mu.Unlock()
		}

		// Make sure to respect the single chat limit of requests.
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}

	// Make sure to respect the global limit of requests.
	return c.gl.Wait(ctx)
}

// dispatch is the common path for all API calls: rate-limit, send, decode, check.
func (c *lclient) dispatch(chatID string, send func() ([]byte, error), v APIResponse) error {
	if err := c.wait(chatID); err != nil {
		return err
	}
	cnt, err := send()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(cnt, v); err != nil {
		return err
	}
	return check(v)
}

// doGet performs a raw HTTP GET and returns the response body.
func (c *lclient) doGet(reqURL string) ([]byte, error) {
	resp, err := c.http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// doPost sends a multipart/form-data POST with the given files and returns the response body.
func (c *lclient) doPost(reqURL string, files ...content) ([]byte, error) {
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

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// doPostForm sends an application/x-www-form-urlencoded POST and returns the response body.
func (c *lclient) doPostForm(reqURL string, keyVals map[string]string) ([]byte, error) {
	var form = make(url.Values)

	for k, v := range keyVals {
		form.Add(k, v)
	}

	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// sendFile sends a single file to Telegram.
// If the file is identified by an ID or URL it is passed as a query parameter;
// otherwise it is uploaded via multipart. The thumbnail, if present, is always uploaded.
func (c *lclient) sendFile(file, thumbnail InputFile, url, fileType string) (res []byte, err error) {
	var cnt []content

	if file.id != "" {
		url = fmt.Sprintf("%s&%s=%s", url, fileType, file.id)
	} else if file.url != "" {
		url = fmt.Sprintf("%s&%s=%s", url, fileType, file.url)
	} else if f, e := toContent(fileType, file); e == nil {
		cnt = append(cnt, f)
	} else {
		err = e
	}

	// Thumbnail is optional; conversion errors are silently ignored.
	if t, e := toContent("thumbnail", thumbnail); e == nil {
		cnt = append(cnt, t)
	}

	if len(cnt) > 0 {
		res, err = c.doPost(url, cnt...)
	} else {
		res, err = c.doGet(url)
	}
	return
}

// get calls a Telegram API endpoint that requires no file upload.
func (c *lclient) get(base, endpoint string, vals url.Values, v APIResponse) error {
	u, err := url.JoinPath(base, endpoint)
	if err != nil {
		return err
	}

	if vals != nil {
		if q := vals.Encode(); q != "" {
			u = fmt.Sprintf("%s?%s", u, q)
		}
	}

	return c.dispatch(vals.Get("chat_id"), func() ([]byte, error) {
		return c.doGet(u)
	}, v)
}

// postFile calls a Telegram API endpoint that requires uploading a single file.
func (c *lclient) postFile(base, endpoint, fileType string, file, thumbnail InputFile, vals url.Values, v APIResponse) error {
	u, err := joinURL(base, endpoint, vals)
	if err != nil {
		return err
	}

	return c.dispatch(vals.Get("chat_id"), func() ([]byte, error) {
		return c.sendFile(file, thumbnail, u, fileType)
	}, v)
}

// postMedia calls a Telegram API endpoint that sends a media group or edits a single media item.
// editSingle serialises only the first element instead of the full array.
func (c *lclient) postMedia(base, endpoint string, editSingle bool, vals url.Values, v APIResponse, files ...InputMedia) error {
	u, err := joinURL(base, endpoint, vals)
	if err != nil {
		return err
	}

	return c.dispatch(vals.Get("chat_id"), func() ([]byte, error) {
		return c.sendMediaFiles(u, editSingle, files...)
	}, v)
}

// postStickers calls a Telegram API endpoint that sends one or more stickers.
func (c *lclient) postStickers(base, endpoint string, vals url.Values, v APIResponse, stickers ...InputSticker) error {
	u, err := joinURL(base, endpoint, vals)
	if err != nil {
		return err
	}

	return c.dispatch(vals.Get("chat_id"), func() ([]byte, error) {
		return c.sendStickers(u, stickers...)
	}, v)
}

// sendMediaFiles serialises the media group into JSON and uploads any local files via multipart.
func (c *lclient) sendMediaFiles(url string, editSingle bool, files ...InputMedia) (res []byte, err error) {
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
			return
		}

		im.InputMedia = file

		med = append(med, im)
		cnt = append(cnt, cntArr...)
	}

	// editSingle is set when editing a single media message; Telegram expects
	// the object directly rather than wrapped in an array.
	if editSingle {
		jsn, err = json.Marshal(med[0])
	} else {
		jsn, err = json.Marshal(med)
	}

	if err != nil {
		return
	}

	url = fmt.Sprintf("%s&media=%s", url, jsn)

	if len(cnt) > 0 {
		return c.doPost(url, cnt...)
	}
	return c.doGet(url)
}

// postProfilePhoto calls a Telegram API endpoint that sets a profile photo.
func (c *lclient) postProfilePhoto(base, endpoint, param string, photo InputProfilePhoto, vals url.Values, v APIResponse) error {
	u, err := joinURL(base, endpoint, vals)
	if err != nil {
		return err
	}

	return c.dispatch(vals.Get("chat_id"), func() ([]byte, error) {
		return c.sendProfilePhotoFile(u, param, photo)
	}, v)
}

// sendProfilePhotoFile serialises the profile photo and uploads it if it's a local file.
func (c *lclient) sendProfilePhotoFile(u, param string, photo InputProfilePhoto) ([]byte, error) {
	env, cnt, err := processProfilePhoto(photo)
	if err != nil {
		return nil, err
	}

	jsn, err := json.Marshal(env)
	if err != nil {
		return nil, err
	}

	u = fmt.Sprintf("%s&%s=%s", u, param, jsn)

	if len(cnt) > 0 {
		return c.doPost(u, cnt...)
	}
	return c.doGet(u)
}

// postStoryContent calls a Telegram API endpoint that posts or edits a story.
func (c *lclient) postStoryContent(base, endpoint string, sc InputStoryContent, vals url.Values, v APIResponse) error {
	u, err := joinURL(base, endpoint, vals)
	if err != nil {
		return err
	}

	return c.dispatch(vals.Get("chat_id"), func() ([]byte, error) {
		return c.sendStoryContentFile(u, sc)
	}, v)
}

// sendStoryContentFile serialises the story content and uploads it if it's a local file.
func (c *lclient) sendStoryContentFile(u string, sc InputStoryContent) ([]byte, error) {
	env, cnt, err := processStoryContent(sc)
	if err != nil {
		return nil, err
	}

	jsn, err := json.Marshal(env)
	if err != nil {
		return nil, err
	}

	u = fmt.Sprintf("%s&content=%s", u, jsn)

	if len(cnt) > 0 {
		return c.doPost(u, cnt...)
	}
	return c.doGet(u)
}

// sendStickers serialises the stickers and uploads any local files via multipart.
// A single sticker uses the "sticker" parameter; multiple use "stickers".
func (c *lclient) sendStickers(url string, stickers ...InputSticker) (res []byte, err error) {
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
			return
		}

		se.InputSticker = s

		sti = append(sti, se)
		cnt = append(cnt, cntArr...)
	}

	// Telegram uses different parameter names for one vs many stickers.
	if len(sti) == 1 {
		jsn, _ = json.Marshal(sti[0])
		url = fmt.Sprintf("%s&sticker=%s", url, jsn)
	} else {
		jsn, _ = json.Marshal(sti)
		url = fmt.Sprintf("%s&stickers=%s", url, jsn)
	}

	if len(cnt) > 0 {
		return c.doPost(url, cnt...)
	}
	return c.doGet(url)
}
