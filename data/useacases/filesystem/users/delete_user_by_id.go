package users

import (
	"fmt"
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/presentation/protocols"
)

type DeleteUserByIDFileSystemUseCase struct {
	data *data.WSD
}

func NewDeleteUserByIDFileSystemUseCase(data *data.WSD) *DeleteUserByIDFileSystemUseCase {
	u := &DeleteUserByIDFileSystemUseCase{
		data: data,
	}
	return u
}

func (u *DeleteUserByIDFileSystemUseCase) Delete(ID string) error {
	foundUser := &models.User{}
	var idx int
	for i, user := range *u.data.Users {
		if user.ID == ID {
			foundUser = &user
			idx = i
			break
		}
	}
	if foundUser.ID == "" {
		return protocols.NewHttpError(
			"an user with that ID was not found",
			http.StatusNotFound)
	}

	var users = *u.data.Users
	*u.data.Users = append(users[:idx], users[idx+1:]...)

	var err error
	if err = u.data.PersistWSD(); err != nil {
		return protocols.NewHttpError(
			fmt.Sprintf("an error ocurred while deleting this user: %v", err),
			http.StatusInternalServerError)
	}

	return nil
}
