package models

type Wallet struct {
	ID           string        `json:"id"`
	UserID       string        `json:"user_id"`
	CurrencyCode string        `json:"currency_code"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}
type AddTransactionDTO struct {
	WalletID     string  `json:"wallet_id"`
	Amount       float32 `json:"amount"`
	CurrencyCode string  `json:"currency_code"`
}
type CreateWalletDTO struct {
	UserID       string `json:"user_id" validate:"required,max=32"`
	CurrencyCode string `json:"currency_code" validate:"required,max=5,min=2"`
}
type ListWalletsDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}
type PaginatedWalletResultDTO struct {
	TotalPages int      `json:"total_pages"`
	Count      int64    `json:"count"`
	PerPage    int      `json:"per_page"`
	Page       int      `json:"page"`
	Data       []Wallet `json:"data"`
}
