package users

import (
	"errors"
	"fmt"
	"wallett/data"
	"wallett/domain/models"
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
		return errors.New("an user with that ID was not found")
	}

	var users = *u.data.Users
	*u.data.Users = append(users[:idx], users[idx+1:]...)

	var err error
	if err = u.data.PersistWSD(); err != nil {
		return fmt.Errorf("an error ocurred while deleting this user: %v", err)
	}

	return nil
}
