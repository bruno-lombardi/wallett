package wallets

import (
	"wallett/data"
	"wallett/domain/models"
)

type ListWalletsFileSystemUsecase struct {
	data *data.WSD
}

func NewListWalletsFileSystemUsecase(data *data.WSD) *ListWalletsFileSystemUsecase {
	u := &ListWalletsFileSystemUsecase{
		data: data,
	}
	return u
}

func (u *ListWalletsFileSystemUsecase) List(listWalletsDto *models.ListWalletsDTO) (*models.PaginatedWalletResultDTO, error) {
	sliceSize := len(*u.data.Wallets) / listWalletsDto.Limit
	walletsSlices := make([][]models.Wallet, sliceSize)

	for i := 0; i < sliceSize; i++ {
		walletsSlices[i] = make([]models.Wallet, listWalletsDto.Limit)

		for j := i * listWalletsDto.Limit; j < (i*listWalletsDto.Limit)+listWalletsDto.Limit; j++ {
			innerSliceIdx := j - (i * listWalletsDto.Limit)
			walletsSlices[i][innerSliceIdx] = (*u.data.Wallets)[j]
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
		Count:      len(*u.data.Wallets),
		Data:       data,
	}
	return paginateResult, nil
}
