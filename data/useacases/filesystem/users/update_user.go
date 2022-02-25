package users

import (
	"errors"
	"wallett/data"
	"wallett/domain/models"
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
		return nil, errors.New("an user with that id was not found")
	}

	if foundUser.Password != updateUserDto.CurrentPassword {
		return nil, errors.New("the password provided for this user is not valid")
	}

	foundUser.Email = updateUserDto.Email
	foundUser.Name = updateUserDto.Name
	foundUser.Password = updateUserDto.NewPassword

	(*u.data.Users)[idx] = *foundUser

	var err error
	if err = u.data.PersistWSD(); err != nil {
		return nil, errors.New("an error ocurred while saving this user")
	}

	return foundUser, nil
}
