package wallets

import (
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/presentation/protocols"
)

type GetWalletByIDFileSystemUsecase struct {
	data *data.WSD
}

func NewGetWalletByIDFileSystemUseCase(data *data.WSD) *GetWalletByIDFileSystemUsecase {
	u := &GetWalletByIDFileSystemUsecase{
		data: data,
	}
	return u
}

func (u *GetWalletByIDFileSystemUsecase) GetByID(id string) (*models.Wallet, error) {
	foundWallet := &models.Wallet{}
	for _, wallet := range *u.data.Wallets {
		if wallet.ID == id {
			foundWallet = &wallet
			break
		}
	}
	if foundWallet.ID == "" {
		return nil, protocols.NewHttpError("a wallet with that id was not found", http.StatusNotFound)
	}

	return foundWallet, nil
}
