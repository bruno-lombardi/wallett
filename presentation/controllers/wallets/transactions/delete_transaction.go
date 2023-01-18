package transactions

import (
	"fmt"
	"wallett/data"
	"wallett/presentation/protocols"
)

type DeleteTransactionController struct {
	data *data.WSD
}

func NewDeleteTransactionController(data *data.WSD) *DeleteTransactionController {
	return &DeleteTransactionController{
		data: data,
	}
}

func (c *DeleteTransactionController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
