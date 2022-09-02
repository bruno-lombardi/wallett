package wallets

import (
	"fmt"
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/infra/generators"
	"wallett/presentation/protocols"
)

type CreateWalletFileSystemUsecase struct {
	data *data.WSD
}

func NewCreateWalletFileSystemUseCase(data *data.WSD) *CreateWalletFileSystemUsecase {
	u := &CreateWalletFileSystemUsecase{
		data: data,
	}
	return u
}

func (u *CreateWalletFileSystemUsecase) Create(createWalletDto *models.CreateWalletDTO) (*models.Wallet, error) {

	foundUser := &models.User{}
	for _, user := range *u.data.Users {
		if user.ID == createWalletDto.UserID {
			foundUser = &user
			break
		}
	}

	if foundUser.ID == "" {
		return nil, protocols.NewHttpError("an user with that id was not found", http.StatusNotFound)
	}

	wallet := &models.Wallet{
		ID:           generators.ID("wa"),
		CurrencyCode: createWalletDto.CurrencyCode,
		UserID:       createWalletDto.UserID,
		Balance:      0.0,
		Transactions: []models.Transaction{},
	}
	*u.data.Wallets = append(*u.data.Wallets, *wallet)
	var err error
	if err = u.data.PersistWSD(); err != nil {
		return nil, protocols.NewHttpError(fmt.Sprintf("an error ocurred while saving this wallet: %v", err), http.StatusInternalServerError)
	}

	return wallet, nil
}
