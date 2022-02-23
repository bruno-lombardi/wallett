package models

type Wallet struct {
	ID           string        `json:"id"`
	UserID       string        `json:"user_id"`
	CurrencyCode string        `json:"currency_code"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type PaginatedWalletResultDTO struct {
	TotalPages int      `json:"total_pages"`
	Count      int      `json:"count"`
	PerPage    int      `json:"per_page"`
	Page       int      `json:"page"`
	Data       []Wallet `json:"data"`
}

func (w *Wallet) AddTransaction(t Transaction) {
	t.PreviousBalance = w.Balance
	w.Transactions = append(w.Transactions, t)
	w.Balance = w.calculateTotalBalance()
}

func (w *Wallet) DeleteTransaction(transactionID string) {
	filteredTransactions := w.Transactions[:0]

	for _, t := range w.Transactions {
		if t.ID != transactionID {
			filteredTransactions = append(filteredTransactions, t)
		}
	}

	w.Transactions = filteredTransactions
	w.Balance = w.calculateTotalBalance()
}

func (w *Wallet) calculateTotalBalance() float64 {
	var total float64
	for _, transaction := range w.Transactions {
		total += transaction.Amount
	}
	return total
}
