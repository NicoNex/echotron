package echotron

import "testing"

func TestRevenueWithdrawalStatePending(t *testing.T) {
	r := RevenueWithdrawalStatePending{}
	r.ImplementsRevenueWithdrawalState()
}

func TestRevenueWithdrawalStateSucceeded(t *testing.T) {
	r := RevenueWithdrawalStateSucceeded{}
	r.ImplementsRevenueWithdrawalState()
}

func TestRevenueWithdrawalStateFailed(t *testing.T) {
	r := RevenueWithdrawalStateFailed{}
	r.ImplementsRevenueWithdrawalState()
}

func TestTransactionPartnerFragment(t *testing.T) {
	r := TransactionPartnerFragment{}
	r.ImplementsTransactionPartner()
}

func TestTransactionPartnerUser(t *testing.T) {
	r := TransactionPartnerUser{}
	r.ImplementsTransactionPartner()
}
func TestTransactionPartnerTelegramAds(t *testing.T) {
	r := TransactionPartnerTelegramAds{}
	r.ImplementsTransactionPartner()
}
func TestTransactionPartnerOther(t *testing.T) {
	r := TransactionPartnerOther{}
	r.ImplementsTransactionPartner()
}

func TestSendInvoice(t *testing.T) {
	_, err := api.SendInvoice(
		chatID,
		"TestSendInvoice",
		"TestSendInvoiceDesc",
		"echotron_test",
		"XTR",
		[]LabeledPrice{
			{
				Label:  "Test",
				Amount: 1,
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateInvoiceLink(t *testing.T) {
	_, err := api.CreateInvoiceLink(
		"TestCreateInvoiceLink",
		"TestCreateInvoiceLinkDesc",
		"echotron_test",
		"XTR",
		[]LabeledPrice{
			{
				Label:  "Test",
				Amount: 1,
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStarTransactions(t *testing.T) {
	_, err := api.GetStarTransactions(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}
