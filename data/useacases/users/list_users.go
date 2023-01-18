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
	users, err = (*u.usersRepository).List(listUsersDto)
	return users, err
}
