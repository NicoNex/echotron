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
	"encoding/json"
	"fmt"
)

func (a API) SendInvoice(chatID int64, title, description, payload, providerToken, currency string, prices []LabeledPrice, opts *InvoiceOptions) (res APIResponseMessage, err error) {
	p, err := json.Marshal(prices)
	if err != nil {
		return
	}

	var url = fmt.Sprintf(
		"%ssendInvoice?chat_id=%d&title=%s&description=%s&payload=%s&provider_token=%s&currency=%s&prices=%s&%s",
		a.base,
		chatID,
		title,
		description,
		payload,
		providerToken,
		currency,
		string(p),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

func (a API) AnswerShippingQuery(shippingQueryID string, ok bool, opts *ShippingQueryOptions) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%sanswerShippingQuery?shipping_query_id=%s&ok=%T&%s",
		a.base,
		shippingQueryID,
		ok,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

func (a API) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, opts *PreCheckoutOptions) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%sanswerPreCheckoutQuery?pre_checkout_query_id=%s&ok=%T&error_message=%s",
		a.base,
		preCheckoutQueryID,
		ok,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}
