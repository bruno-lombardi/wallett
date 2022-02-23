package users

import (
	"wallett/data"
	"wallett/domain/models"
)

type ListUserFileSystemUseCase struct {
	data *data.WSD
}

func NewListUserFileSystemUseCase(data *data.WSD) *ListUserFileSystemUseCase {
	u := &ListUserFileSystemUseCase{
		data: data,
	}
	return u
}

func (u *ListUserFileSystemUseCase) List(listUsersDto *models.ListUsersDTO) (*models.PaginatedUserResultDTO, error) {
	sliceSize := len(*u.data.Users) / listUsersDto.Limit
	usersSlices := make([][]models.User, sliceSize)

	for i := 0; i < sliceSize; i++ {
		usersSlices[i] = make([]models.User, listUsersDto.Limit)
		// i = 0, j = 0...10
		// i = 9, j = 90...99
		for j := i * listUsersDto.Limit; j < (i*listUsersDto.Limit)+listUsersDto.Limit; j++ {
			innerSliceIdx := j - (i * listUsersDto.Limit)
			usersSlices[i][innerSliceIdx] = (*u.data.Users)[j]
		}
	}
	var data []models.User = []models.User{}
	if listUsersDto.Page <= len(usersSlices) {
		data = usersSlices[listUsersDto.Page-1]
	}

	paginatedResultDto := &models.PaginatedUserResultDTO{
		TotalPages: sliceSize,
		PerPage:    listUsersDto.Limit,
		Page:       listUsersDto.Page,
		Count:      len(*u.data.Users),
		Data:       data,
	}

	return paginatedResultDto, nil
}
