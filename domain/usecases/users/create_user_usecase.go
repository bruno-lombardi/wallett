package users

import "wallett/domain/models"

type CreateUserUsecase interface {
	Create(*models.CreateUserDTO) (*models.User, error)
}
