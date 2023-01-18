package users

import (
	"wallett/data/protocols/repositories"
)

type DbDeleteUserByIDUseCase struct {
	usersRepository *repositories.UserRepository
}

func NewDbDeleteUserByIDUseCase(usersRepository repositories.UserRepository) *DbDeleteUserByIDUseCase {
	u := &DbDeleteUserByIDUseCase{
		usersRepository: &usersRepository,
	}
	return u
}

func (u *DbDeleteUserByIDUseCase) Delete(ID string) error {
	err := (*u.usersRepository).Delete(ID)
	return err
}
