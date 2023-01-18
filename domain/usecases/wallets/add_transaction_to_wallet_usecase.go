package wallets

import "wallett/domain/models"

type AddTransactionToWalletUsecase interface {
	AddTransaction(walletID string, addTransactionDTO models.AddTransactionDTO) (*models.Wallet, error)
}
