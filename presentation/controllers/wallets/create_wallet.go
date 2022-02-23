package wallets

import (
	"errors"
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/infra/generators"
	"wallett/presentation/protocols"
)

type CreateWalletDTO struct {
	UserID       string `json:"user_id" validate:"required,max=32"`
	CurrencyCode string `json:"currency_code" validate:"required,max=5,min=2"`
}

type CreateWalletController struct {
	data *data.WSD
}

func NewCreateWalletController(data *data.WSD) *CreateWalletController {
	return &CreateWalletController{
		data: data,
	}
}

func (c *CreateWalletController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	createWalletDto := req.Body.(*CreateWalletDTO)

	foundUser := &models.User{}
	for _, user := range *c.data.Users {
		if user.ID == createWalletDto.UserID {
			foundUser = &user
			break
		}
	}

	if foundUser.ID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("an user with that id was not found")
	}

	wallet := &models.Wallet{
		ID:           generators.ID("wa"),
		CurrencyCode: createWalletDto.CurrencyCode,
		UserID:       createWalletDto.UserID,
		Balance:      0.0,
		Transactions: []models.Transaction{},
	}
	*c.data.Wallets = append(*c.data.Wallets, *wallet)
	var err error
	if err = c.data.PersistWSD(); err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("an error ocurred while saving this user")
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       wallet,
	}

	return response, nil
}
