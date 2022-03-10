package wallets

import "wallett/domain/models"

type CreateWalletUsecase interface {
	Create(*models.CreateWalletDTO) (*models.Wallet, error)
}
