package transactions

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbAddTransactionToWalletUsecase struct {
	walletRepository *repositories.WalletRepository
}

func NewDbAddTransactionToWalletUsecase(walletRepository repositories.WalletRepository) *DbAddTransactionToWalletUsecase {
	u := &DbAddTransactionToWalletUsecase{
		walletRepository: &walletRepository,
	}
	return u
}

func (u *DbAddTransactionToWalletUsecase) AddTransaction(walletID string, addTransactionDTO models.AddTransactionDTO) (wallet *models.Wallet, err error) {
	wallet, err = (*u.walletRepository).AddTransaction(walletID, addTransactionDTO)
	return wallet, err
}
