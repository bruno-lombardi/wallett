package users

import (
	"errors"
	"wallett/data"
	"wallett/domain/models"
)

type GetUserByIDFileSystemUseCase struct {
	data *data.WSD
}

func NewGetUserByIDFileSystemUseCase(data *data.WSD) *GetUserByIDFileSystemUseCase {
	u := &GetUserByIDFileSystemUseCase{
		data: data,
	}
	return u
}

func (u *GetUserByIDFileSystemUseCase) Get(ID string) (*models.User, error) {
	foundUser := &models.User{}
	for _, user := range *u.data.Users {
		if user.ID == ID {
			foundUser = &user
			break
		}
	}
	if foundUser.ID == "" {
		return nil, errors.New("an user with that ID was not found")
	}
	return foundUser, nil
}
