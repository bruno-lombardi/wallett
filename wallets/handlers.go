package wallets

import (
	"net/http"
	"wallett/data"
	"wallett/models"

	"github.com/labstack/echo"
)

type WalletHandlers struct {
	data *data.WSD
}

type ListWalletsDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

func NewWalletHandlers(data *data.WSD) *WalletHandlers {
	h := &WalletHandlers{
		data: data,
	}
	return h
}

func (h *WalletHandlers) SetupRoutes(r *echo.Group) {
	r.GET("/wallets", h.ListWallets)
	r.GET("/wallets/:id", h.GetWalletByID)
	// r.POST("/wallets")
	// r.PUT("/wallets/:id")
	// r.DELETE("/wallets/:id")

	// r.GET("/wallets/:wallet_id/transactions")
	// r.GET("/wallets/:wallet_id/transactions/:trx_id")
	// r.POST("/wallets/:wallet_id/transactions")
	// r.PUT("/wallets/:wallet_id/transactions/:trx_id")
	// r.DELETE("/wallets/:wallet_id/transactions/:trx_id")
}

func (h *WalletHandlers) ListWallets(c echo.Context) (err error) {
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

	sliceSize := len(*h.data.Wallets) / listWalletsDto.Limit
	walletsSlices := make([][]models.Wallet, sliceSize)

	for i := 0; i < sliceSize; i++ {
		walletsSlices[i] = make([]models.Wallet, listWalletsDto.Limit)

		for j := i * listWalletsDto.Limit; j < (i*listWalletsDto.Limit)+listWalletsDto.Limit; j++ {
			innerSliceIdx := j - (i * listWalletsDto.Limit)
			walletsSlices[i][innerSliceIdx] = (*h.data.Wallets)[j]
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
		Count:      len(*h.data.Wallets),
		Data:       data,
	})
}

func (h *WalletHandlers) GetWalletByID(c echo.Context) (err error) {
	id := c.Param("id")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	foundWallet := &models.Wallet{}
	for _, wallet := range *h.data.Wallets {
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
