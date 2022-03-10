package wallets

import (
	"net/http"
	"wallett/domain/models"
	"wallett/domain/usecases/wallets"
	"wallett/presentation/protocols"
)

type ListWalletsController struct {
	listWalletsUsecase *wallets.ListWalletsUsecase
}

func NewListWalletsController(listWalletsUsecase wallets.ListWalletsUsecase) *ListWalletsController {
	return &ListWalletsController{
		listWalletsUsecase: &listWalletsUsecase,
	}
}

func (c *ListWalletsController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	listWalletsDto := req.Body.(*models.ListWalletsDTO)

	wallet, err := (*c.listWalletsUsecase).List(listWalletsDto)
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
