package users

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbCreateUserUsecase struct {
	usersRepository *repositories.UserRepository
}

func NewDbCreateUserUsecase(usersRepository repositories.UserRepository) *DbCreateUserUsecase {
	u := &DbCreateUserUsecase{
		usersRepository: &usersRepository,
	}
	return u
}

func (u *DbCreateUserUsecase) Create(createUserDto *models.CreateUserDTO) (user *models.User, err error) {
	user, err = (*u.usersRepository).Create(createUserDto)
	return user, err
}
