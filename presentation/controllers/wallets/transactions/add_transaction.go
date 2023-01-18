package transactions

import (
	"net/http"
	"wallett/domain/models"
	"wallett/domain/usecases/wallets"
	"wallett/presentation/protocols"
)

type AddTransactionController struct {
	addTransactionUsecase *wallets.AddTransactionToWalletUsecase
}

func NewAddTransactionController(addTransactionUsecase wallets.AddTransactionToWalletUsecase) *AddTransactionController {
	return &AddTransactionController{
		addTransactionUsecase: &addTransactionUsecase,
	}
}

func (c *AddTransactionController) Handle(req *protocols.HttpRequest) (response *protocols.HttpResponse, err error) {
	id := req.PathParams["wallet_id"]
	addTransactionDTO := req.Body.(*models.AddTransactionDTO)

	wallet, err := (*c.addTransactionUsecase).AddTransaction(id, *addTransactionDTO)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response = &protocols.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       wallet,
	}

	return response, nil
}
