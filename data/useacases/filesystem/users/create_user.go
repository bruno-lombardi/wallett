package users

import (
	"fmt"
	"wallett/data"
	"wallett/domain/models"
	"wallett/infra/generators"
)

type CreateUserFileSystemUseCase struct {
	data *data.WSD
}

func NewCreateUserFileSystemUseCase(data *data.WSD) *CreateUserFileSystemUseCase {
	u := &CreateUserFileSystemUseCase{
		data: data,
	}
	return u
}

func (u *CreateUserFileSystemUseCase) Create(createUserDto *models.CreateUserDTO) (*models.User, error) {

	user := &models.User{
		ID:       generators.ID("u"),
		Email:    createUserDto.Email,
		Name:     createUserDto.Name,
		Password: createUserDto.Password,
	}

	*u.data.Users = append(*u.data.Users, *user)
	var err error
	if err = u.data.PersistWSD(); err != nil {
		return nil, fmt.Errorf("could not save user data to file system: %v", err)
	}

	return user, nil
}
