package repositories

import "wallett/domain/models"

type UserRepository interface {
	Create(createUserDto *models.CreateUserDTO) (*models.User, error)
	Get(ID string) (*models.User, error)
	List(listUsersDto *models.ListUsersDTO) (*models.PaginatedUserResultDTO, error)
	Update(updateUserDto *models.UpdateUserDTO) (*models.User, error)
	Delete(ID string) error
}
