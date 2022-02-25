package handlers

import (
	"wallett/data"
	"wallett/data/useacases/filesystem/users"
	"wallett/domain/models"
	"wallett/main/adapters"
	usersControllers "wallett/presentation/controllers/users"

	"github.com/labstack/echo"
)

type UserHandlers struct {
	data *data.WSD
}

func NewUserHandlers(data *data.WSD) *UserHandlers {
	h := &UserHandlers{
		data: data,
	}
	return h
}

func (h *UserHandlers) SetupRoutes(r *echo.Group) {
	createUserFileSystemUseCase := users.NewCreateUserFileSystemUseCase(h.data)
	listUsersFileSystemUseCase := users.NewListUserFileSystemUseCase(h.data)
	updateUserFileSystemUseCase := users.NewUpdateUserFileSystemUseCase(h.data)
	getUserByIDUseCase := users.NewGetUserByIDFileSystemUseCase(h.data)
	deleteUserByIDUseCase := users.NewDeleteUserByIDFileSystemUseCase(h.data)

	r.GET("/users", adapters.AdaptHandlerJSON(
		usersControllers.NewListUsersController(listUsersFileSystemUseCase),
		&models.ListUsersDTO{}))
	r.GET("/users/:id", adapters.AdaptHandlerJSON(
		usersControllers.NewGetUserController(getUserByIDUseCase),
		nil))
	r.POST("/users", adapters.AdaptHandlerJSON(
		usersControllers.NewCreateUserController(createUserFileSystemUseCase),
		&models.CreateUserDTO{}))

	r.PUT("/users/:id", adapters.AdaptHandlerJSON(
		usersControllers.NewUpdateUsersController(updateUserFileSystemUseCase),
		&models.UpdateUserDTO{}))
	r.DELETE("/users/:id", adapters.AdaptHandlerJSON(
		usersControllers.NewDeleteUserController(deleteUserByIDUseCase),
		&models.UpdateUserDTO{}))
}
