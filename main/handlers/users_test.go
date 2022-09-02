package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"wallett/domain/models"
	"wallett/infra/generators"
	"wallett/infra/persistence/db/sqlite"
	"wallett/presentation/helpers"
	"wallett/test"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockUsers *[]models.User
)

type HttpTestCase = test.HttpTestCase
type User = models.User

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	sqlite.Connect("file::memory:?cache=shared")
	sqlite.GetDB().AutoMigrate(&sqlite.SQLiteUser{})
	mockUsers = &[]User{
		{
			ID:       generators.ID("u"),
			Email:    fmt.Sprintf("%v@%v.com", generators.RandomString(30), generators.RandomString(10)),
			Name:     "Bruno",
			Password: generators.RandomString(12),
		},
	}

	code := m.Run()

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
			BeforeTest:   func() error { return nil },
			AfterTest:    func() error { return nil },
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
			BeforeTest:   func() error { return nil },
			AfterTest:    func() error { return nil },
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
			BeforeTest:   func() error { return nil },
			AfterTest:    func() error { return nil },
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
			BeforeTest:   func() error { return nil },
			AfterTest: func() error {
				sqlite.GetDB().Exec("DELETE FROM sq_lite_users").Commit()
				return nil
			},
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
			WhenURL:      fmt.Sprintf("/api/v1/users/%s", (*mockUsers)[0].ID),
			WhenMethod:   http.MethodGet,
			WhenBody:     "",
			ExpectStatus: http.StatusOK,
			BeforeTest: func() error {
				user := (*mockUsers)[0]
				result := sqlite.GetDB().Model(&sqlite.SQLiteUser{}).Create(&sqlite.SQLiteUser{
					ID:       user.ID,
					Email:    user.Email,
					Name:     user.Name,
					Password: user.Password,
				})
				return result.Error
			},
			ExpectBody: func(t *testing.T, body *bytes.Buffer) (err error) {
				decodedUser := &models.User{}
				err = json.Unmarshal(body.Bytes(), decodedUser)
				assert.Nil(t, err)

				assert.Equal(t, decodedUser.ID, (*mockUsers)[0].ID)
				assert.NotEmpty(t, decodedUser.Email)
				assert.NotEmpty(t, decodedUser.Name)
				assert.Empty(t, decodedUser.Password)
				return err
			},
			AfterTest: func() error {
				user := (*mockUsers)[0]
				result := sqlite.GetDB().Delete(&sqlite.SQLiteUser{ID: user.ID})
				return result.Error
			},
		},
		{
			Name:         "Should return 404 not found if invalid ID is given",
			WhenURL:      "/api/v1/users/invalid_id",
			WhenMethod:   http.MethodGet,
			WhenBody:     "",
			ExpectStatus: http.StatusNotFound,
			BeforeTest:   func() error { return nil },
			AfterTest:    func() error { return nil },
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
			e.Validator = helpers.NewCustomValidator()
			api := e.Group("/api/v1")
			h := NewUserHandlers(sqlite.GetDB())
			h.SetupRoutes(api)

			assert.NoError(t, testCase.BeforeTest())

			req := httptest.NewRequest(testCase.WhenMethod, testCase.WhenURL, strings.NewReader(testCase.WhenBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			body := rec.Body

			assert.Equal(t, testCase.ExpectStatus, rec.Code)
			assert.NoError(t, testCase.ExpectBody(t, body))
			assert.NoError(t, testCase.AfterTest())
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
