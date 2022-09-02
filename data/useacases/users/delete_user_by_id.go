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
	// foundUser := &models.User{}
	// var idx int
	// for i, user := range *u.data.Users {
	// 	if user.ID == ID {
	// 		foundUser = &user
	// 		idx = i
	// 		break
	// 	}
	// }
	// if foundUser.ID == "" {
	// 	return protocols.NewHttpError(
	// 		"an user with that ID was not found",
	// 		http.StatusNotFound)
	// }

	// var users = *u.data.Users
	// *u.data.Users = append(users[:idx], users[idx+1:]...)

	// var err error
	// if err = u.data.PersistWSD(); err != nil {
	// 	return protocols.NewHttpError(
	// 		fmt.Sprintf("an error ocurred while deleting this user: %v", err),
	// 		http.StatusInternalServerError)
	// }

	err := (*u.usersRepository).Delete(ID)
	return err
}
