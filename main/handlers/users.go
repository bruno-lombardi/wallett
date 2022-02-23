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

type UpdateUserDTO struct {
	ID                      string `param:"id" validate:"required"`
	Email                   string `json:"email" validate:"email,required"`
	Name                    string `json:"name" validate:"required,max=100,min=2"`
	CurrentPassword         string `json:"current_password" validate:"required,max=64,min=6"`
	NewPassword             string `json:"new_password" validate:"required,max=64,min=6"`
	NewPasswordConfirmation string `json:"new_password_confirmation" validate:"required,max=64,min=6,eqcsfield=NewPassword"`
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

	r.GET("/users", adapters.AdaptHandlerJSON(
		usersControllers.NewListUsersController(listUsersFileSystemUseCase),
		&models.ListUsersDTO{}))
	// r.GET("/users/:id", h.GetUserByID)
	r.POST("/users", adapters.AdaptHandlerJSON(
		usersControllers.NewCreateUserController(createUserFileSystemUseCase),
		&models.CreateUserDTO{}))

	// r.PUT("/users/:id", h.UpdateUser)
	// r.DELETE("/users/:id", h.DeleteUserByID)
}

// func (h *UserHandlers) GetUserByID(c echo.Context) (err error) {
// 	id := c.Param("id")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	foundUser := &models.User{}
// 	for _, user := range *h.data.Users {
// 		if user.ID == id {
// 			foundUser = &user
// 			break
// 		}
// 	}
// 	if foundUser.ID == "" {
// 		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
// 	}
// 	return c.JSON(http.StatusOK, foundUser)
// }

// func (h *UserHandlers) UpdateUser(c echo.Context) (err error) {
// 	updateUserDto := &UpdateUserDTO{}
// 	updateUserDto.ID = c.Param("id")
// 	if err = c.Bind(updateUserDto); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	if err = c.Validate(updateUserDto); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	foundUser := &models.User{}
// 	var idx int
// 	for i, user := range *h.data.Users {
// 		if user.ID == updateUserDto.ID {
// 			foundUser = &user
// 			idx = i
// 			break
// 		}
// 	}
// 	if foundUser.ID == "" {
// 		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
// 	}
// 	if foundUser.Password != updateUserDto.CurrentPassword {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "The password provided for this user is not valid.")
// 	}

// 	foundUser.Email = updateUserDto.Email
// 	foundUser.Name = updateUserDto.Name
// 	foundUser.Password = updateUserDto.NewPassword

// 	(*h.data.Users)[idx] = *foundUser

// 	if err = h.data.PersistWSD(); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "An error ocurred while saving this user.")
// 	}
// 	return c.JSON(http.StatusOK, foundUser)
// }

// func (h *UserHandlers) DeleteUserByID(c echo.Context) (err error) {
// 	id := c.Param("id")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	foundUser := &models.User{}
// 	var idx int
// 	for i, user := range *h.data.Users {
// 		if user.ID == id {
// 			foundUser = &user
// 			idx = i
// 			break
// 		}
// 	}
// 	if foundUser.ID == "" {
// 		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
// 	}

// 	var users = *h.data.Users
// 	*h.data.Users = append(users[:idx], users[idx+1:]...)

// 	if err = h.data.PersistWSD(); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "An error ocurred while deleting this user.")
// 	}

// 	return c.JSON(http.StatusOK, foundUser)
// }
