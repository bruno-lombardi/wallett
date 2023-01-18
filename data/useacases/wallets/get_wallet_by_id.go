package wallets

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbGetWalletByIDUsecase struct {
	walletRepository *repositories.WalletRepository
}

func NewDbGetWalletByIUseCase(walletRepository repositories.WalletRepository) *DbGetWalletByIDUsecase {
	u := &DbGetWalletByIDUsecase{
		walletRepository: &walletRepository,
	}
	return u
}

func (u *DbGetWalletByIDUsecase) GetByID(id string) (wallet *models.Wallet, err error) {
	wallet, err = (*u.walletRepository).Get(id)
	return wallet, err
}
