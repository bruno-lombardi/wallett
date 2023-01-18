package handlers

import (
	"wallett/data/useacases/wallets"
	"wallett/data/useacases/wallets/transactions"
	"wallett/domain/models"
	"wallett/infra/persistence/db/sqlite"
	"wallett/main/adapters"
	walletsControllers "wallett/presentation/controllers/wallets"
	transactionsControllers "wallett/presentation/controllers/wallets/transactions"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type WalletHandlers struct {
	db *gorm.DB
}

func NewWalletHandlers(db *gorm.DB) *WalletHandlers {
	h := &WalletHandlers{
		db: db,
	}
	return h
}

func (h *WalletHandlers) SetupHandlers(r *echo.Group) {
	walletsRepository := sqlite.NewSQLiteWalletRepository(h.db)

	dbCreateWalletUsecase := wallets.NewDbCreateWalletUseCase(walletsRepository)
	dbGetWalletByIDUsecase := wallets.NewDbGetWalletByIUseCase(walletsRepository)
	dbListWalletsUsecase := wallets.NewDbListWalletsUsecase(walletsRepository)
	dbAddTransactionToWalletUsecase := transactions.NewDbAddTransactionToWalletUsecase(walletsRepository)

	r.GET("/wallets",
		adapters.AdaptHandlerJSON(
			walletsControllers.NewListWalletsController(dbListWalletsUsecase),
			&models.ListWalletsDTO{}))
	r.GET("/wallets/:id",
		adapters.AdaptHandlerJSON(
			walletsControllers.NewGetWalletByIDController(dbGetWalletByIDUsecase), nil))
	r.POST("/wallets",
		adapters.AdaptHandlerJSON(
			walletsControllers.NewCreateWalletController(dbCreateWalletUsecase),
			&models.CreateWalletDTO{}))
	// r.PUT("/wallets/:id")
	// r.DELETE("/wallets/:id")

	r.POST("/wallets/:wallet_id/transactions",
		adapters.AdaptHandlerJSON(
			transactionsControllers.NewAddTransactionController(dbAddTransactionToWalletUsecase),
			&models.AddTransactionDTO{}))

	// r.GET("/wallets/:wallet_id/transactions")
	// r.GET("/wallets/:wallet_id/transactions/:trx_id")

	// r.PUT("/wallets/:wallet_id/transactions/:trx_id")
	// r.DELETE("/wallets/:wallet_id/transactions/:trx_id",
	// 	adapters.AdaptHandlerJSON(
	// 		transactionsControllers.NewDeleteTransactionController(h.data), nil))
}
