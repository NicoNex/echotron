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
	"net/url"
)

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label string `json:"label"`
	// Price of the product in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code.
	Currency string `json:"currency"`
	// Total amount in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	Name            string          `json:"name,omitempty"`
	PhoneNumber     string          `json:"phone_number,omitempty"`
	Email           string          `json:"email,omitempty"`
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	OrderInfo               OrderInfo `json:"order_info"`
	Currency                string    `json:"currency"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionID        string    `json:"shipping_option_id"`
	TelegramPaymentChargeID string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string    `json:"provider_payment_charge_id"`
	TotalAmount             int       `json:"total_amount"`
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	ShippingAddress ShippingAddress `json:"shipping_address"`
	ID              string          `json:"id"`
	InvoicePayload  string          `json:"invoice_payload"`
	From            User            `json:"from"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	OrderInfo        OrderInfo `json:"order_info,omitempty"`
	Currency         string    `json:"currency"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionID string    `json:"shipping_option_id,omitempty"`
	ID               string    `json:"id"`
	From             User      `json:"from"`
	TotalAmount      int       `json:"total_amount"`
}

// SendInvoice is used to send invoices.
func (a API) SendInvoice(chatID int64, title, description, payload, providerToken, currency string, prices []LabeledPrice, opts *InvoiceOptions) (res APIResponseMessage, err error) {
	var vals = make(url.Values)

	p, err := json.Marshal(prices)
	if err != nil {
		return res, err
	}

	vals.Set("chat_id", itoa(chatID))
	vals.Set("title", title)
	vals.Set("description", description)
	vals.Set("payload", payload)
	vals.Set("provider_token", providerToken)
	vals.Set("currency", currency)
	vals.Set("prices", string(p))
	return get[APIResponseMessage](a.base, "sendInvoice", addValues(vals, opts))
}

// AnswerShippingQuery is used to reply to shipping queries.
// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified,
// the Bot API will send an Update with a shipping_query field to the bot.
func (a API) AnswerShippingQuery(shippingQueryID string, ok bool, opts *ShippingQueryOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("shipping_query_id", shippingQueryID)
	vals.Set("ok", btoa(ok))
	return get[APIResponseBase](a.base, "answerShippingQuery", addValues(vals, opts))
}

// AnswerPreCheckoutQuery is used to respond to such pre-checkout queries.
// Once the user has confirmed their payment and shipping details,
// the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query.
// NOTE: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (a API) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, opts *PreCheckoutOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("pre_checkout_query_id", preCheckoutQueryID)
	vals.Set("ok", btoa(ok))
	return get[APIResponseBase](a.base, "answerPreCheckoutQuery", addValues(vals, opts))
}

// CreateInvoiceLink creates a link for an invoice.
func (a API) CreateInvoiceLink(title, description, payload, providerToken, currency string, prices []LabeledPrice, opts *CreateInvoiceLinkOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	p, err := json.Marshal(prices)
	if err != nil {
		return res, err
	}

	vals.Set("title", title)
	vals.Set("description", description)
	vals.Set("payload", payload)
	vals.Set("provider_token", providerToken)
	vals.Set("currency", currency)
	vals.Set("prices", string(p))
	return get[APIResponseBase](a.base, "createInvoiceLink", addValues(vals, opts))
}
