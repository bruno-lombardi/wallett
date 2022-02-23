package users

import (
	"net/http"
	"wallett/domain/models"
	"wallett/domain/usecases/users"
	"wallett/presentation/protocols"
)

type ListUserController struct {
	listUsersUsecase *users.ListUsersUsecase
}

func NewListUsersController(listUsersUsecase users.ListUsersUsecase) *ListUserController {
	return &ListUserController{
		listUsersUsecase: &listUsersUsecase,
	}
}

func (c *ListUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	listUsersDto := req.Body.(*models.ListUsersDTO)

	if listUsersDto == nil {
		listUsersDto = &models.ListUsersDTO{
			Page:  1,
			Limit: 10,
		}
	}

	var err error
	user, err := (*c.listUsersUsecase).List(listUsersDto)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       user,
	}

	return response, nil
}
