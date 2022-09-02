package handlers

import (
	"wallett/data/useacases/users"
	"wallett/domain/models"
	"wallett/infra/persistence/db/sqlite"
	"wallett/main/adapters"
	usersControllers "wallett/presentation/controllers/users"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type UserHandlers struct {
	db *gorm.DB
}

func NewUserHandlers(db *gorm.DB) *UserHandlers {
	h := &UserHandlers{
		db: db,
	}
	return h
}

func (h *UserHandlers) SetupRoutes(r *echo.Group) {
	usersRepository := sqlite.NewSQLiteUserRepository(h.db)
	dbCreateUserUsecase := users.NewDbCreateUserUsecase(usersRepository)
	listUsersFileSystemUseCase := users.NewDbListUserUseCase(usersRepository)
	updateUserFileSystemUseCase := users.NewDbUpdateUserUseCase(usersRepository)
	getUserByIDUseCase := users.NewDbGetUserByIDUsecase(usersRepository)
	deleteUserByIDUseCase := users.NewDbDeleteUserByIDUseCase(usersRepository)

	r.GET("/users", adapters.AdaptHandlerJSON(
		usersControllers.NewListUsersController(listUsersFileSystemUseCase),
		&models.ListUsersDTO{}))
	r.GET("/users/:id", adapters.AdaptHandlerJSON(
		usersControllers.NewGetUserController(getUserByIDUseCase),
		nil))
	r.POST("/users", adapters.AdaptHandlerJSON(
		usersControllers.NewCreateUserController(dbCreateUserUsecase),
		&models.CreateUserDTO{}))

	r.PUT("/users/:id", adapters.AdaptHandlerJSON(
		usersControllers.NewUpdateUsersController(updateUserFileSystemUseCase),
		&models.UpdateUserDTO{}))
	r.DELETE("/users/:id", adapters.AdaptHandlerJSON(
		usersControllers.NewDeleteUserController(deleteUserByIDUseCase),
		&models.UpdateUserDTO{}))
}
