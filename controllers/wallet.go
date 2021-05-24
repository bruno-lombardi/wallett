package controllers

import (
	"net/http"
	"wallett/data"
	"wallett/models"

	"github.com/labstack/echo"
)

type ListWalletsDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

func HandleListWallets(c echo.Context) (err error) {
	listWalletsDto := &ListWalletsDTO{
		Page:  1,
		Limit: 10,
	}
	if err = c.Bind(listWalletsDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(listWalletsDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wsd := data.GetWSD()
	sliceSize := len(*wsd.Wallets) / listWalletsDto.Limit
	walletsSlices := make([][]models.Wallet, sliceSize)

	for i := 0; i < sliceSize; i++ {
		walletsSlices[i] = make([]models.Wallet, listWalletsDto.Limit)

		for j := i * listWalletsDto.Limit; j < (i*listWalletsDto.Limit)+listWalletsDto.Limit; j++ {
			innerSliceIdx := j - (i * listWalletsDto.Limit)
			walletsSlices[i][innerSliceIdx] = (*wsd.Wallets)[j]
		}
	}
	var data []models.Wallet = []models.Wallet{}
	if listWalletsDto.Page <= len(walletsSlices) {
		data = walletsSlices[listWalletsDto.Page-1]
	}

	return c.JSON(http.StatusOK, &models.PaginatedWalletResultDTO{
		TotalPages: sliceSize,
		PerPage:    listWalletsDto.Limit,
		Page:       listWalletsDto.Page,
		Count:      len(*wsd.Wallets),
		Data:       data,
	})
}

func HandleGetWalletByID(c echo.Context) (err error) {
	id := c.Param("id")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wsd := data.GetWSD()
	foundWallet := &models.Wallet{}
	for _, wallet := range *wsd.Wallets {
		if wallet.ID == id {
			foundWallet = &wallet
			break
		}
	}
	if foundWallet.ID == "" {
		return echo.NewHTTPError(http.StatusNotFound, "An wallet with that id was not found.")
	}
	return c.JSON(http.StatusOK, foundWallet)
}
