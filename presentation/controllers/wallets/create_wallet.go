package wallets

import (
	"net/http"
	"wallett/domain/models"
	"wallett/domain/usecases/wallets"
	"wallett/presentation/protocols"
)

type CreateWalletController struct {
	createWalletUsecase *wallets.CreateWalletUsecase
}

func NewCreateWalletController(createWalletUsecase wallets.CreateWalletUsecase) *CreateWalletController {
	return &CreateWalletController{
		createWalletUsecase: &createWalletUsecase,
	}
}

func (c *CreateWalletController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	createWalletDto := req.Body.(*models.CreateWalletDTO)

	var err error
	user, err := (*c.createWalletUsecase).Create(createWalletDto)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       user,
	}

	return response, nil
}
