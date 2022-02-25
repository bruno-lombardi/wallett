package users

import (
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/presentation/protocols"
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
		return nil, protocols.NewHttpError(
			"an user with that ID was not found",
			http.StatusNotFound)
	}
	return foundUser, nil
}
