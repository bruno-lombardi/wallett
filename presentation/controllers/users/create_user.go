package users

import (
	"net/http"
	"wallett/domain/models"
	"wallett/domain/usecases/users"
	"wallett/presentation/protocols"
)

type CreateUserController struct {
	createUserUsecase *users.CreateUserUsecase
}

func NewCreateUserController(createUserUsecase users.CreateUserUsecase) *CreateUserController {
	return &CreateUserController{
		createUserUsecase: &createUserUsecase,
	}
}

func (c *CreateUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	createUserModel := req.Body.(*models.CreateUserDTO)

	var err error
	user, err := (*c.createUserUsecase).Create(createUserModel)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       user,
	}

	return response, nil
}
