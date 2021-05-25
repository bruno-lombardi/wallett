package users

import (
	"net/http"
	"wallett/data"
	"wallett/generators"
	"wallett/models"

	"github.com/labstack/echo"
)

type UserHandlers struct {
	data *data.WSD
}

type ListUsersDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

type CreateUserDTO struct {
	Email                string `json:"email" validate:"email,required"`
	Name                 string `json:"name" validate:"required,max=100,min=2"`
	Password             string `json:"password" validate:"required,max=64,min=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,max=64,min=6,eqcsfield=Password"`
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
	r.GET("/users", h.ListUsers)
	r.GET("/users/:id", h.GetUserByID)
	r.POST("/users", h.CreateUser)
	r.PUT("/users/:id", h.UpdateUser)
	r.DELETE("/users/:id", h.DeleteUserByID)
}

func (h *UserHandlers) GetUserByID(c echo.Context) (err error) {
	id := c.Param("id")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	foundUser := &models.User{}
	for _, user := range *h.data.Users {
		if user.ID == id {
			foundUser = &user
			break
		}
	}
	if foundUser.ID == "" {
		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
	}
	return c.JSON(http.StatusOK, foundUser)
}

func (h *UserHandlers) ListUsers(c echo.Context) (err error) {
	listUsersDto := &ListUsersDTO{
		Page:  1,
		Limit: 10,
	}
	if err = c.Bind(listUsersDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(listUsersDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sliceSize := len(*h.data.Users) / listUsersDto.Limit
	usersSlices := make([][]models.User, sliceSize)

	for i := 0; i < sliceSize; i++ {
		usersSlices[i] = make([]models.User, listUsersDto.Limit)
		// i = 0, j = 0...10
		// i = 9, j = 90...99
		for j := i * listUsersDto.Limit; j < (i*listUsersDto.Limit)+listUsersDto.Limit; j++ {
			innerSliceIdx := j - (i * listUsersDto.Limit)
			usersSlices[i][innerSliceIdx] = (*h.data.Users)[j]
		}
	}
	var data []models.User = []models.User{}
	if listUsersDto.Page <= len(usersSlices) {
		data = usersSlices[listUsersDto.Page-1]
	}

	return c.JSON(http.StatusOK, &models.PaginatedUserResultDTO{
		TotalPages: sliceSize,
		PerPage:    listUsersDto.Limit,
		Page:       listUsersDto.Page,
		Count:      len(*h.data.Users),
		Data:       data,
	})
}

func (h *UserHandlers) CreateUser(c echo.Context) (err error) {
	createUserDto := &CreateUserDTO{}
	if err = c.Bind(createUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(createUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := &models.User{
		ID:       generators.ID("u"),
		Email:    createUserDto.Email,
		Name:     createUserDto.Name,
		Password: createUserDto.Password,
	}
	*h.data.Users = append(*h.data.Users, *user)
	if err = h.data.PersistWSD(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "An error ocurred while saving this user.")
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandlers) UpdateUser(c echo.Context) (err error) {
	updateUserDto := &UpdateUserDTO{}
	updateUserDto.ID = c.Param("id")
	if err = c.Bind(updateUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(updateUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	foundUser := &models.User{}
	var idx int
	for i, user := range *h.data.Users {
		if user.ID == updateUserDto.ID {
			foundUser = &user
			idx = i
			break
		}
	}
	if foundUser.ID == "" {
		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
	}
	if foundUser.Password != updateUserDto.CurrentPassword {
		return echo.NewHTTPError(http.StatusUnauthorized, "The password provided for this user is not valid.")
	}

	foundUser.Email = updateUserDto.Email
	foundUser.Name = updateUserDto.Name
	foundUser.Password = updateUserDto.NewPassword

	(*h.data.Users)[idx] = *foundUser

	if err = h.data.PersistWSD(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "An error ocurred while saving this user.")
	}
	return c.JSON(http.StatusOK, foundUser)
}

func (h *UserHandlers) DeleteUserByID(c echo.Context) (err error) {
	id := c.Param("id")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	foundUser := &models.User{}
	var idx int
	for i, user := range *h.data.Users {
		if user.ID == id {
			foundUser = &user
			idx = i
			break
		}
	}
	if foundUser.ID == "" {
		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
	}

	var users = *h.data.Users
	*h.data.Users = append(users[:idx], users[idx+1:]...)

	if err = h.data.PersistWSD(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "An error ocurred while deleting this user.")
	}

	return c.JSON(http.StatusOK, foundUser)
}
