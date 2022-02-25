package users

import (
	"fmt"
	"net/http"
	"wallett/data"
	"wallett/domain/models"
	"wallett/presentation/protocols"
)

type UpdateUserFileSystemUseCase struct {
	data *data.WSD
}

func NewUpdateUserFileSystemUseCase(data *data.WSD) *UpdateUserFileSystemUseCase {
	u := &UpdateUserFileSystemUseCase{
		data: data,
	}
	return u
}

func (u *UpdateUserFileSystemUseCase) Update(updateUserDto *models.UpdateUserDTO) (*models.User, error) {
	foundUser := &models.User{}
	var idx int
	for i, user := range *u.data.Users {
		if user.ID == updateUserDto.ID {
			foundUser = &user
			idx = i
			break
		}
	}

	if foundUser.ID == "" {
		return nil, protocols.NewHttpError(
			"an user with that ID was not found",
			http.StatusNotFound)
	}

	if foundUser.Password != updateUserDto.CurrentPassword {
		return nil, protocols.NewHttpError(
			"the password provided is not valid",
			http.StatusBadRequest)
	}

	foundUser.Email = updateUserDto.Email
	foundUser.Name = updateUserDto.Name
	foundUser.Password = updateUserDto.NewPassword

	(*u.data.Users)[idx] = *foundUser

	var err error
	if err = u.data.PersistWSD(); err != nil {
		return nil, protocols.NewHttpError(
			fmt.Sprintf("an error ocurred while saving this user: %v", err),
			http.StatusBadRequest)
	}

	return foundUser, nil
}
