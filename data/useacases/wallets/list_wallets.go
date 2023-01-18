package wallets

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbListWalletsUsecase struct {
	walletRepository *repositories.WalletRepository
}

func NewDbListWalletsUsecase(walletRepository repositories.WalletRepository) *DbListWalletsUsecase {
	u := &DbListWalletsUsecase{
		walletRepository: &walletRepository,
	}
	return u
}

func (u *DbListWalletsUsecase) List(listWalletsDTO *models.ListWalletsDTO) (wallets *models.PaginatedWalletResultDTO, err error) {
	wallets, err = (*u.walletRepository).List(listWalletsDTO)
	return wallets, err
}
