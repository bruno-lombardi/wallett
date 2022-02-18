package handlers

import (
	"wallett/data"
	"wallett/main/adapters"
	transactionsControllers "wallett/presentation/controllers/transactions"
	walletsControllers "wallett/presentation/controllers/wallets"

	"github.com/labstack/echo"
)

type WalletHandlers struct {
	data *data.WSD
}

func NewWalletHandlers(data *data.WSD) *WalletHandlers {
	h := &WalletHandlers{
		data: data,
	}
	return h
}

func (h *WalletHandlers) SetupHandlers(r *echo.Group) {
	r.GET("/wallets", adapters.AdaptHandlerJSON(walletsControllers.NewListWalletsController(h.data), &walletsControllers.ListWalletsDTO{}))
	r.GET("/wallets/:id", adapters.AdaptHandlerJSON(walletsControllers.NewGetWalletByIDController(h.data), nil))
	r.POST("/wallets", adapters.AdaptHandlerJSON(walletsControllers.NewCreateWalletController(h.data), &walletsControllers.CreateWalletDTO{}))
	// r.PUT("/wallets/:id")
	// r.DELETE("/wallets/:id")

	// r.GET("/wallets/:wallet_id/transactions")
	// r.GET("/wallets/:wallet_id/transactions/:trx_id")
	r.POST("/wallets/:wallet_id/transactions",
		adapters.AdaptHandlerJSON(
			transactionsControllers.NewCreateWalletController(h.data),
			&transactionsControllers.CreateTransactionDTO{}))
	// r.PUT("/wallets/:wallet_id/transactions/:trx_id")
	r.DELETE("/wallets/:wallet_id/transactions/:trx_id",
		adapters.AdaptHandlerJSON(
			transactionsControllers.NewDeleteTransactionController(h.data), nil))
}
