package users

import (
	"net/http"
	"wallett/domain/usecases/users"
	"wallett/presentation/protocols"
)

type GetUserController struct {
	getUserByIDUsecase *users.GetUserByIDUsecase
}

func NewGetUserController(getUserByIDUsecase users.GetUserByIDUsecase) *GetUserController {
	return &GetUserController{
		getUserByIDUsecase: &getUserByIDUsecase,
	}
}

func (c *GetUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	id := req.PathParams["id"]

	var err error
	user, err := (*c.getUserByIDUsecase).Get(id)
	if err != nil {
		if err.Error() == "record not found" {
			return &protocols.HttpResponse{
				StatusCode: http.StatusNotFound,
				Body: &map[string]string{
					"message": "User not found",
				},
			}, err
		}
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if user.ID == "" {
		return &protocols.HttpResponse{
			StatusCode: http.StatusNotFound,
			Body: &map[string]string{
				"message": "User not found",
			},
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       user,
	}

	return response, nil
}
