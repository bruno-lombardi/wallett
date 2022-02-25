package users

import (
	"net/http"
	"wallett/domain/usecases/users"
	"wallett/presentation/protocols"
)

type DeleteUserController struct {
	deleteUserByIDUsecase *users.DeleteUserByIDUsecase
}

func NewDeleteUserController(deleteUserByIDUsecase users.DeleteUserByIDUsecase) *DeleteUserController {
	return &DeleteUserController{
		deleteUserByIDUsecase: &deleteUserByIDUsecase,
	}
}

func (c *DeleteUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	id := req.PathParams["id"]

	if err := (*c.deleteUserByIDUsecase).Delete(id); err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusNoContent,
	}

	return response, nil
}
