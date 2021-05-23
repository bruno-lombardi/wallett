package controllers

import (
	"net/http"
	"wallett/data"
	"wallett/models"

	"github.com/labstack/echo"
)

type GetUserDTO struct {
	ID string `param:"id"`
}

type ListUsersDTO struct {
	Page  int `query:"page" validate:"gte=1"`
	Limit int `query:"limit" validate:"gte=1,lte=20"`
}

func HandleGetUserByID(c echo.Context) (err error) {
	getUserDto := &GetUserDTO{}
	if err = c.Bind(getUserDto); err != nil {
		return err
	}
	wsd := data.GetWSD()
	foundUser := &models.User{}
	for _, user := range *wsd.Users {
		if user.ID == getUserDto.ID {
			foundUser = &user
			break
		}
	}
	if foundUser.ID == "" {
		return echo.NewHTTPError(http.StatusNotFound, "An user with that id was not found.")
	}
	return c.JSON(http.StatusOK, foundUser)
}

func HandleListUsers(c echo.Context) (err error) {
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

	wsd := data.GetWSD()
	sliceSize := len(*wsd.Users) / listUsersDto.Limit
	usersSlices := make([][]models.User, sliceSize)

	for i := 0; i < sliceSize; i++ {
		usersSlices[i] = make([]models.User, listUsersDto.Limit)
		// i = 0, j = 0...10
		// i = 9, j = 90...99
		for j := i * listUsersDto.Limit; j < (i*listUsersDto.Limit)+listUsersDto.Limit; j++ {
			innerSliceIdx := j - (i * listUsersDto.Limit)
			usersSlices[i][innerSliceIdx] = (*wsd.Users)[j]
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
		Data:       data,
	})
}
