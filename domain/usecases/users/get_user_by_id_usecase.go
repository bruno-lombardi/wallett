package users

import "wallett/domain/models"

type GetUserByIDUsecase interface {
	Get(ID string) (*models.User, error)
}
