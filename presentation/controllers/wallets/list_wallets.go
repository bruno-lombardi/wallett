package wallets

import (
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/presentation/protocols"
)

type ListWalletsDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

type ListWalletsController struct {
	data *data.WSD
}

func NewListWalletsController(data *data.WSD) *ListWalletsController {
	return &ListWalletsController{
		data: data,
	}
}

func (c *ListWalletsController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	listWalletsDto := req.Body.(*ListWalletsDTO)

	sliceSize := len(*c.data.Wallets) / listWalletsDto.Limit
	walletsSlices := make([][]models.Wallet, sliceSize)

	for i := 0; i < sliceSize; i++ {
		walletsSlices[i] = make([]models.Wallet, listWalletsDto.Limit)

		for j := i * listWalletsDto.Limit; j < (i*listWalletsDto.Limit)+listWalletsDto.Limit; j++ {
			innerSliceIdx := j - (i * listWalletsDto.Limit)
			walletsSlices[i][innerSliceIdx] = (*c.data.Wallets)[j]
		}
	}
	var data []models.Wallet = []models.Wallet{}
	if listWalletsDto.Page <= len(walletsSlices) {
		data = walletsSlices[listWalletsDto.Page-1]
	}

	paginateResult := &models.PaginatedWalletResultDTO{
		TotalPages: sliceSize,
		PerPage:    listWalletsDto.Limit,
		Page:       listWalletsDto.Page,
		Count:      len(*c.data.Wallets),
		Data:       data,
	}
	res := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       paginateResult,
		Headers:    map[string][]string{},
	}
	return res, nil
}
