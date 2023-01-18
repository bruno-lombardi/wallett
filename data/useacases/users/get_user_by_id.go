package users

import (
	"wallett/data/protocols/repositories"
	"wallett/domain/models"
)

type DbGetUserByIDUsecase struct {
	usersRepository *repositories.UserRepository
}

func NewDbGetUserByIDUsecase(usersRepository repositories.UserRepository) *DbGetUserByIDUsecase {
	u := &DbGetUserByIDUsecase{
		usersRepository: &usersRepository,
	}
	return u
}

func (u *DbGetUserByIDUsecase) Get(ID string) (*models.User, error) {
	user, err := (*u.usersRepository).Get(ID)
	return user, err
}
