package models

import "fmt"

type Transaction struct {
	ID              string  `json:"id"`
	WalletID        string  `json:"wallet_id"`
	Amount          float64 `json:"amount"`
	PreviousBalance float64 `json:"previous_balance"`
	CurrencyCode    string  `json:"currency_code"`
}

func (t Transaction) String() string {
	return fmt.Sprintf("TRX_ID: %v, AMOUNT: %.2f,", t.ID, t.Amount)
}
