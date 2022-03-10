package wallets

import "wallett/domain/models"

type GetWalletByIDUsecase interface {
	GetByID(id string) (*models.Wallet, error)
}
