package wallets

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbCreateWalletUsecase struct {
	walletRepository *repositories.WalletRepository
}

func NewDbCreateWalletUseCase(walletRepository repositories.WalletRepository) *DbCreateWalletUsecase {
	u := &DbCreateWalletUsecase{
		walletRepository: &walletRepository,
	}
	return u
}

func (u *DbCreateWalletUsecase) Create(createWalletDto *models.CreateWalletDTO) (wallet *models.Wallet, err error) {
	wallet, err = (*u.walletRepository).Create(createWalletDto)
	return wallet, err
}
