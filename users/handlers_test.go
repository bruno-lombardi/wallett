package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wallett/data"
	"wallett/middlewares"
	"wallett/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockData       = data.NewWSD("test.dat")
	createUserJSON = `{
		"name": "John Doe",
		"email": "john.doe@gmail.com",
		"password": "654321",
		"password_confirmation": "654321"
	}`
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = middlewares.NewCustomValidator()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Println(c.Path(), c.QueryParams(), err)
		e.DefaultHTTPErrorHandler(err, c)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	api := e.Group("/api/v1")
	h := NewUserHandlers(mockData)
	h.SetupRoutes(api)

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		decodedUser := &models.User{}
		err := json.Unmarshal(rec.Body.Bytes(), decodedUser)
		assert.Nil(t, err)

		assert.Equal(t, decodedUser.Name, "John Doe")
		assert.Equal(t, decodedUser.Email, "john.doe@gmail.com")
		assert.NotEmpty(t, decodedUser.ID)
	}
}

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/users/:email")
// 	c.SetParamNames("email")
// 	c.SetParamValues("jon@labstack.com")
// 	h := &handler{mockDB}

// 	// Assertions
// 	if assert.NoError(t, h.getUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
