package wallets

import (
	"errors"
	"net/http"
	"wallett/data"
	"wallett/models"
	"wallett/presentation/protocols"
)

type GetWalletByIDController struct {
	data *data.WSD
}

func NewGetWalletByIDController(data *data.WSD) *GetWalletByIDController {
	return &GetWalletByIDController{
		data: data,
	}
}

func (c *GetWalletByIDController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	id := req.PathParams["id"]

	foundWallet := &models.Wallet{}
	for _, wallet := range *c.data.Wallets {
		if wallet.ID == id {
			foundWallet = &wallet
			break
		}
	}
	if foundWallet.ID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("a wallet with that id was not found")
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       foundWallet,
	}

	return response, nil
}
