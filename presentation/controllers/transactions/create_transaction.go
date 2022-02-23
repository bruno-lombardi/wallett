package transactions

import (
	"errors"
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/infra/generators"
	"wallett/presentation/protocols"
)

type CreateTransactionDTO struct {
	WalletID     string
	CurrencyCode string  `json:"currency_code" validate:"required,max=5,min=2"`
	Amount       float64 `json:"amount" validate:"required"`
}

type CreateTransactionController struct {
	data *data.WSD
}

func NewCreateWalletController(data *data.WSD) *CreateTransactionController {
	return &CreateTransactionController{
		data: data,
	}
}

func (c *CreateTransactionController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	createWalletDto := req.Body.(*CreateTransactionDTO)
	createWalletDto.WalletID = req.PathParams["wallet_id"]

	if createWalletDto.WalletID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusBadRequest,
		}, errors.New("please provide a wallet id")
	}

	foundWallet := &models.Wallet{}
	var idx int
	for i, wallet := range *c.data.Wallets {
		if wallet.ID == createWalletDto.WalletID {
			foundWallet = &wallet
			idx = i
			break
		}
	}

	if foundWallet.ID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("a wallet with that id was not found")
	}

	transaction := &models.Transaction{
		ID:              generators.ID("trx"),
		WalletID:        createWalletDto.WalletID,
		Amount:          createWalletDto.Amount,
		CurrencyCode:    createWalletDto.CurrencyCode,
		PreviousBalance: foundWallet.Balance,
	}
	foundWallet.AddTransaction(*transaction)
	(*c.data.Wallets)[idx] = *foundWallet

	var err error
	if err = c.data.PersistWSD(); err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("an error ocurred while creating this transaction")
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       transaction,
	}

	return response, nil
}
