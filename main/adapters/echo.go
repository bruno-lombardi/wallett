package adapters

import (
	"net/http"
	"strings"
	"wallett/presentation/protocols"

	"github.com/labstack/echo"
)

func AdaptHandlerJSON(controller protocols.Controller, body interface{}) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		httpRequest := &protocols.HttpRequest{
			Body: body,
		}
		// Set Headers
		httpRequest.Headers = c.Request().Header

		// Set Params and Query Params
		params := map[string]string{}
		for _, p := range c.ParamNames() {
			for _, v := range c.ParamValues() {
				params[p] = v
			}
		}
		httpRequest.PathParams = params
		httpRequest.QueryParams = c.QueryParams()

		if body != nil {
			if err = c.Bind(httpRequest.Body); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			if err = c.Validate(httpRequest.Body); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}

		response, err := controller.Handle(httpRequest)
		if err != nil {
			switch err := err.(type) {
			case *protocols.HttpError:
				return echo.NewHTTPError(err.StatusCode, map[string]string{"message": err.Error()})
			default:
				return echo.NewHTTPError(response.StatusCode, map[string]string{"message": err.Error()})
			}
		}

		for k, v := range response.Headers {
			c.Response().Header().Set(k, strings.Join(v[:], ";"))
		}
		return c.JSON(response.StatusCode, response.Body)
	}
}
