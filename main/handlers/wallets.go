package handlers

import (
	"wallett/data"
	"wallett/data/useacases/wallets"
	"wallett/domain/models"
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
	createWalletFileSystemUsecase := wallets.NewCreateWalletFileSystemUseCase(h.data)
	getWalletByIDFileSystemUsecase := wallets.NewGetWalletByIDFileSystemUseCase(h.data)
	listWalletsFileSystemUsecase := wallets.NewListWalletsFileSystemUsecase(h.data)

	r.GET("/wallets",
		adapters.AdaptHandlerJSON(
			walletsControllers.NewListWalletsController(listWalletsFileSystemUsecase),
			&models.ListWalletsDTO{}))
	r.GET("/wallets/:id",
		adapters.AdaptHandlerJSON(
			walletsControllers.NewGetWalletByIDController(getWalletByIDFileSystemUsecase), nil))
	r.POST("/wallets",
		adapters.AdaptHandlerJSON(
			walletsControllers.NewCreateWalletController(createWalletFileSystemUsecase),
			&models.CreateWalletDTO{}))
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
