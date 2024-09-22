package echotron

import "testing"

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
