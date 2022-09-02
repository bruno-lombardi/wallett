package users

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbListUserUseCase struct {
	usersRepository *repositories.UserRepository
}

func NewDbListUserUseCase(usersRepository repositories.UserRepository) *DbListUserUseCase {
	u := &DbListUserUseCase{
		usersRepository: &usersRepository,
	}
	return u
}

func (u *DbListUserUseCase) List(listUsersDto *models.ListUsersDTO) (users *models.PaginatedUserResultDTO, err error) {
	// sliceSize := len(*u.data.Users) / listUsersDto.Limit
	// usersSlices := make([][]models.User, sliceSize)

	// for i := 0; i < sliceSize; i++ {
	// 	usersSlices[i] = make([]models.User, listUsersDto.Limit)
	// 	// i = 0, j = 0...10
	// 	// i = 9, j = 90...99
	// 	for j := i * listUsersDto.Limit; j < (i*listUsersDto.Limit)+listUsersDto.Limit; j++ {
	// 		innerSliceIdx := j - (i * listUsersDto.Limit)
	// 		usersSlices[i][innerSliceIdx] = (*u.data.Users)[j]
	// 	}
	// }
	// var data []models.User = []models.User{}
	// if listUsersDto.Page <= len(usersSlices) {
	// 	data = usersSlices[listUsersDto.Page-1]
	// }

	// paginatedResultDto := &models.PaginatedUserResultDTO{
	// 	TotalPages: sliceSize,
	// 	PerPage:    listUsersDto.Limit,
	// 	Page:       listUsersDto.Page,
	// 	Count:      int64(len(*u.data.Users)),
	// 	Data:       data,
	// }

	users, err = (*u.usersRepository).List(listUsersDto)
	return users, err
}
