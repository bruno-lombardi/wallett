package transactions

import (
	"errors"
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/presentation/protocols"
)

type DeleteTransactionController struct {
	data *data.WSD
}

func NewDeleteTransactionController(data *data.WSD) *DeleteTransactionController {
	return &DeleteTransactionController{
		data: data,
	}
}

func (c *DeleteTransactionController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	walletID := req.PathParams["wallet_id"]
	transactionID := req.PathParams["trx_id"]

	if walletID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusBadRequest,
		}, errors.New("please provide a wallet id")
	}

	if transactionID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusBadRequest,
		}, errors.New("please provide a transaction id")
	}

	foundWallet := &models.Wallet{}
	var walletIdx int
	for i, wallet := range *c.data.Wallets {
		if wallet.ID == walletID {
			foundWallet = &wallet
			walletIdx = i
			break
		}
	}

	if foundWallet.ID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("a wallet with that id was not found")
	}

	foundWallet.DeleteTransaction(transactionID)
	(*c.data.Wallets)[walletIdx] = *foundWallet

	var err error
	if err = c.data.PersistWSD(); err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
		}, errors.New("an error ocurred while deleting this transaction")
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusNoContent,
	}

	return response, nil
}

// func (h *WalletHandlers) CreateWallet(c echo.Context) (err error) {
// 	createWalletDto := &CreateWalletDTO{}
// 	if err = c.Bind(createWalletDto); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	if err = c.Validate(createWalletDto); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	foundUser := &models.User{}
// 	for _, user := range *h.data.Users {
// 		if user.ID == createWalletDto.UserID {
// 			foundUser = &user
// 			break
// 		}
// 	}

// 	if foundUser.ID == "" {
// 		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
// 	}

// 	wallet := &models.Wallet{
// 		ID:           generators.ID("wa"),
// 		CurrencyCode: createWalletDto.CurrencyCode,
// 		UserID:       createWalletDto.UserID,
// 		Balance:      0.0,
// 		Transactions: []models.Transaction{},
// 	}
// 	*h.data.Wallets = append(*h.data.Wallets, *wallet)
// 	if err = h.data.PersistWSD(); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "An error ocurred while saving this user.")
// 	}
// 	return c.JSON(http.StatusCreated, wallet)
// }
