package repositories

import "wallett/domain/models"

type WalletRepository interface {
	Create(createWalletDTO *models.CreateWalletDTO) (*models.Wallet, error)
	Get(ID string) (*models.Wallet, error)
	List(listWalletsDTO *models.ListWalletsDTO) (*models.PaginatedWalletResultDTO, error)
	AddTransaction(WalletID string, t models.AddTransactionDTO) (*models.Wallet, error)
	Delete(ID string) error
}
