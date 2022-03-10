package wallets

import (
	"net/http"
	"wallett/domain/usecases/wallets"
	"wallett/presentation/protocols"
)

type GetWalletByIDController struct {
	getWalletByIdUsecase *wallets.GetWalletByIDUsecase
}

func NewGetWalletByIDController(getWalletByIdUsecase wallets.GetWalletByIDUsecase) *GetWalletByIDController {
	return &GetWalletByIDController{
		getWalletByIdUsecase: &getWalletByIdUsecase,
	}
}

func (c *GetWalletByIDController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	id := req.PathParams["id"]

	var err error
	wallet, err := (*c.getWalletByIdUsecase).GetByID(id)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       wallet,
	}

	return response, nil
}
