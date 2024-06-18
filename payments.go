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

// RevenueWithdrawalState describes the state of a revenue withdrawal operation.
type RevenueWithdrawalState interface {
	ImplementsRevenueWithdrawalState()
}

// RevenueWithdrawalStatePending describes the state of a withdrawal in progress.
type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

// ImplementsRevenueWithdrawalState is used to implement the RevenueWithdrawalState interface.
func (r RevenueWithdrawalStatePending) ImplementsRevenueWithdrawalState() {}

// RevenueWithdrawalStateSucceeded describes the state of a succeeded withdrawal.
type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	URL  string `json:"url"`
	Date int    `json:"date"`
}

// ImplementsRevenueWithdrawalState is used to implement the RevenueWithdrawalState interface.
func (r RevenueWithdrawalStateSucceeded) ImplementsRevenueWithdrawalState() {}

// RevenueWithdrawalStateFailed describes the state of a failed withdrawal, in which the transaction was refunded.
type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

// ImplementsRevenueWithdrawalState is used to implement the RevenueWithdrawalState interface.
func (r RevenueWithdrawalStateFailed) ImplementsRevenueWithdrawalState() {}

// TransactionPartner describes the source of a transaction, or its recipient for outgoing transactions.
type TransactionPartner interface {
	ImplementsTransactionPartner()
}

// TransactionPartnerFragment describes a withdrawal transaction with Fragment.
type TransactionPartnerFragment struct {
	WithdrawalState RevenueWithdrawalState `json:"withdrawal_state"`
	Type            string                 `json:"type"`
}

// ImplementsTransactionPartner is used to implement the TransactionPartner interface.
func (t TransactionPartnerFragment) ImplementsTransactionPartner() {}

// TransactionPartnerUser describes a transaction with a user.
type TransactionPartnerUser struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// ImplementsTransactionPartner is used to implement the TransactionPartner interface.
func (t TransactionPartnerUser) ImplementsTransactionPartner() {}

// TransactionPartnerOther describes a transaction with an unknown source or recipient.
type TransactionPartnerOther struct {
	Type string `json:"type"`
}

// ImplementsTransactionPartner is used to implement the TransactionPartner interface.
func (t TransactionPartnerOther) ImplementsTransactionPartner() {}

// StarTransaction describes a Telegram Star transaction.
type StarTransaction struct {
	Source   TransactionPartner `json:"source"`
	Receiver TransactionPartner `json:"receiver"`
	ID       string             `json:"id"`
	Amount   int                `json:"amount"`
	Date     int                `json:"date"`
}

// StarTransactions contains a list of Telegram Star transactions.
type StarTransactions struct {
	Transaction []StarTransaction `json:"transaction"`
}

// StarTransactionsOptions contains the optional parameters used by the GetStarTransactions method.
type StarTransactionsOptions struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

// SendInvoice is used to send invoices.
func (a API) SendInvoice(chatID int64, title, description, payload, currency string, prices []LabeledPrice, opts *InvoiceOptions) (res APIResponseMessage, err error) {
	var vals = make(url.Values)

	p, err := json.Marshal(prices)
	if err != nil {
		return res, err
	}

	vals.Set("chat_id", itoa(chatID))
	vals.Set("title", title)
	vals.Set("description", description)
	vals.Set("payload", payload)
	vals.Set("currency", currency)
	vals.Set("prices", string(p))
	return res, a.client.get(a.base, "sendInvoice", addValues(vals, opts), &res)
}

// CreateInvoiceLink creates a link for an invoice.
func (a API) CreateInvoiceLink(title, description, payload, currency string, prices []LabeledPrice, opts *CreateInvoiceLinkOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	p, err := json.Marshal(prices)
	if err != nil {
		return res, err
	}

	vals.Set("title", title)
	vals.Set("description", description)
	vals.Set("payload", payload)
	vals.Set("currency", currency)
	vals.Set("prices", string(p))
	return res, a.client.get(a.base, "createInvoiceLink", addValues(vals, opts), &res)
}

// AnswerShippingQuery is used to reply to shipping queries.
// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified,
// the Bot API will send an Update with a shipping_query field to the bot.
func (a API) AnswerShippingQuery(shippingQueryID string, ok bool, opts *ShippingQueryOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("shipping_query_id", shippingQueryID)
	vals.Set("ok", btoa(ok))
	return res, a.client.get(a.base, "answerShippingQuery", addValues(vals, opts), &res)
}

// AnswerPreCheckoutQuery is used to respond to such pre-checkout queries.
// Once the user has confirmed their payment and shipping details,
// the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query.
// NOTE: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (a API) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, opts *PreCheckoutOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("pre_checkout_query_id", preCheckoutQueryID)
	vals.Set("ok", btoa(ok))
	return res, a.client.get(a.base, "answerPreCheckoutQuery", addValues(vals, opts), &res)
}

// GetStarTransactions returns the bot's Telegram Star transactions in chronological order.
func (a API) GetStarTransactions(opts *StarTransactionsOptions) (res APIResponseStarTransactions, err error) {
	return res, a.client.get(a.base, "getStarTransactions", urlValues(opts), &res)
}

// RefundStarPayment refunds a successful payment in Telegram Stars.
func (a API) RefundStarPayment(userID int64, telegramPaymentChargeID string) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", itoa(userID))
	vals.Set("telegram_payment_charge_id", telegramPaymentChargeID)
	return res, a.client.get(a.base, "refundStarPayment", vals, &res)
}
