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
	// foundUser := &models.User{}
	// for _, user := range *u.data.Users {
	// 	if user.ID == ID {
	// 		foundUser = &user
	// 		break
	// 	}
	// }
	// if foundUser.ID == "" {
	// 	return nil, protocols.NewHttpError(
	// 		"an user with that ID was not found",
	// 		http.StatusNotFound)
	// }
	user, err := (*u.usersRepository).Get(ID)
	return user, err
}
