package users

import "wallett/domain/models"

type ListUsersUsecase interface {
	List(*models.ListUsersDTO) (*models.PaginatedUserResultDTO, error)
}
