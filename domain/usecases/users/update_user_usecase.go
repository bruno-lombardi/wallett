package users

import "wallett/domain/models"

type UpdateUserUsecase interface {
	Update(*models.UpdateUserDTO) (*models.User, error)
}
