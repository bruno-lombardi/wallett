package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"wallett/data"
	"wallett/middlewares"
	"wallett/models"
	"wallett/test"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockData *data.WSD
)

type HttpTestCase = test.HttpTestCase

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	mockData = data.NewWSD("test.dat")
	code := m.Run()
	mockData.ClearWSD()

	os.Exit(code)
}

func TestUserHandlers(t *testing.T) {
	httpTestCases := []HttpTestCase{
		{
			Name:       "Should create user with valid data",
			WhenURL:    "/api/v1/users",
			WhenMethod: http.MethodPost,
			WhenBody: `{
				"name": "John Doe",
				"email": "john.doe@gmail.com",
				"password": "654321",
				"password_confirmation": "654321"
			}`,
			ExpectStatus: http.StatusCreated,
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				decodedUser := &models.User{}
				err = json.Unmarshal(body.Bytes(), decodedUser)
				assert.Nil(t, err)

				assert.Equal(t, decodedUser.Name, "John Doe")
				assert.Equal(t, decodedUser.Email, "john.doe@gmail.com")
				assert.NotEmpty(t, decodedUser.ID)
				return err
			},
		},
		{
			Name:       "Should not create user if invalid email",
			WhenURL:    "/api/v1/users",
			WhenMethod: http.MethodPost,
			WhenBody: `{
				"name": "John Doe",
				"email": "invalid_email",
				"password": "654321",
				"password_confirmation": "654321"
			}`,
			ExpectStatus: http.StatusBadRequest,
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				var decodedResponse map[string]string
				err = json.Unmarshal(body.Bytes(), &decodedResponse)
				assert.Nil(t, err)

				assert.NotEmpty(t, decodedResponse["message"])
				return err
			},
		},
		{
			Name:       "Should not create user if invalid password",
			WhenURL:    "/api/v1/users",
			WhenMethod: http.MethodPost,
			WhenBody: `{
				"name": "John Doe",
				"email": "john.doe@gmail.com",
				"password": "inval",
				"password_confirmation": "inval"
			}`,
			ExpectStatus: http.StatusBadRequest,
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				var decodedResponse map[string]string
				err = json.Unmarshal(body.Bytes(), &decodedResponse)
				assert.Nil(t, err)

				assert.NotEmpty(t, decodedResponse["message"])
				return err
			},
		},
		{
			Name:       "Should not create user if invalid password confirmation",
			WhenURL:    "/api/v1/users",
			WhenMethod: http.MethodPost,
			WhenBody: `{
				"name": "John Doe",
				"email": "john.doe@gmail.com",
				"password": "123456",
				"password_confirmation": "123455"
			}`,
			ExpectStatus: http.StatusBadRequest,
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				var decodedResponse map[string]string
				err = json.Unmarshal(body.Bytes(), &decodedResponse)
				assert.Nil(t, err)

				assert.NotEmpty(t, decodedResponse["message"])
				return err
			},
		},
		{
			Name:         "Should return user if valid ID is given",
			WhenURL:      fmt.Sprintf("/api/v1/users/%s", (*mockData.Users)[0].ID),
			WhenMethod:   http.MethodGet,
			WhenBody:     "",
			ExpectStatus: http.StatusOK,
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				decodedUser := &models.User{}
				err = json.Unmarshal(body.Bytes(), decodedUser)
				assert.Nil(t, err)

				assert.Equal(t, decodedUser.ID, (*mockData.Users)[0].ID)
				assert.NotEmpty(t, decodedUser.Email)
				assert.NotEmpty(t, decodedUser.Name)
				assert.Empty(t, decodedUser.Password)
				return err
			},
		},
		{
			Name:         "Should return 404 not found if invalid ID is given",
			WhenURL:      "/api/v1/users/invalid_id",
			WhenMethod:   http.MethodGet,
			WhenBody:     "",
			ExpectStatus: http.StatusNotFound,
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				var decodedResponse map[string]string
				err = json.Unmarshal(body.Bytes(), &decodedResponse)
				assert.Nil(t, err)

				assert.NotEmpty(t, decodedResponse["message"])
				return err
			},
		},
	}

	for _, testCase := range httpTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			// Setup
			e := echo.New()
			e.Validator = middlewares.NewCustomValidator()
			api := e.Group("/api/v1")
			h := NewUserHandlers(mockData)
			h.SetupRoutes(api)

			req := httptest.NewRequest(testCase.WhenMethod, testCase.WhenURL, strings.NewReader(testCase.WhenBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			body := rec.Body
			assert.Equal(t, testCase.ExpectStatus, rec.Code)
			assert.NoError(t, testCase.ExpectBody(t, body))

			// Assertions
			// if assert.NoError(t, h.CreateUser(c)) {
			// 	assert.Equal(t, http.StatusCreated, rec.Code)

			// }
		})
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
