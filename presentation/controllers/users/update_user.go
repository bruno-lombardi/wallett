package users

import (
	"net/http"
	"wallett/domain/models"
	"wallett/domain/usecases/users"
	"wallett/presentation/protocols"
)

type UpdateUserController struct {
	updateUserUsecase *users.UpdateUserUsecase
}

func NewUpdateUsersController(updateUserUsecase users.UpdateUserUsecase) *UpdateUserController {
	return &UpdateUserController{
		updateUserUsecase: &updateUserUsecase,
	}
}

func (c *UpdateUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	updateUserDto := req.Body.(*models.UpdateUserDTO)

	var err error
	user, err := (*c.updateUserUsecase).Update(updateUserDto)
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
