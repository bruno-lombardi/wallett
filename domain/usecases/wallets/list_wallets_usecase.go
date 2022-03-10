package wallets

import "wallett/domain/models"

type ListWalletsUsecase interface {
	List(listWalletsDto *models.ListWalletsDTO) (*models.PaginatedWalletResultDTO, error)
}
