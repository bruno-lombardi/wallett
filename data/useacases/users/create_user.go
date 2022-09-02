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
	// user := &models.User{
	// 	ID:       generators.ID("u"),
	// 	Email:    createUserDto.Email,
	// 	Name:     createUserDto.Name,
	// 	Password: createUserDto.Password,
	// }

	// *u.data.Users = append(*u.data.Users, *user)
	// var err error
	// if err = u.data.PersistWSD(); err != nil {
	// 	return nil, protocols.NewHttpError(
	// 		fmt.Sprintf("could not save user data to file system: %v", err),
	// 		http.StatusInternalServerError)
	// }

	// return user, nil
}
