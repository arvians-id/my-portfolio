package integration

import (
	"bytes"
	"fmt"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/proto"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type ListUserResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		User []*model.User `json:"users"`
	} `json:"data"`
}

type UserResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		User *model.User `json:"user"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestUser_CreateSuccess() int64 {
	var id int64

	suite.Run("Create User", func() {
		query := `mutation CreateUser ($input: CreateUserRequest!) {
						user: CreateUser (input: $input) {
							id
							name
							email
							password
							bio
							pronouns
							country
							job_title
							image
							created_at
							updated_at
						}
					}`
		variables := `"input": {
						"name": "Widdy",
						"email": "widdy@gmail.com",
						"password": "Rahasia123",
						"pronouns": "he/him",
						"country": "Jawa Barat, Indonesia",
						"job_title": "Software Engineer"
					  }`
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody UserResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.User)
		suite.Assertions.Equal(responseBody.Data.User.Name, "Widdy")
		suite.Assertions.Equal(responseBody.Data.User.Email, "widdy@gmail.com")
		suite.Assertions.Equal(responseBody.Data.User.Pronouns, "he/him")
		suite.Assertions.Equal(responseBody.Data.User.Country, "Jawa Barat, Indonesia")
		suite.Assertions.Equal(responseBody.Data.User.JobTitle, "Software Engineer")

		id = responseBody.Data.User.ID
	})

	return id
}

func (suite *E2ETestSuite) TestUser_FindAllSuccess() {
	suite.TestUser_CreateSuccess()

	suite.Run("Find All User", func() {
		query := `query FindAllUser {
						users: FindAllUser {
							id
							name
							email
							password
							bio
							pronouns
							country
							job_title
							image
							created_at
							updated_at
						}
					}`
		query = util.CleanQuery(query)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListUserResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.User)
		suite.Assertions.Equal(responseBody.Data.User[0].Name, "Widdy")
		suite.Assertions.Equal(responseBody.Data.User[0].Email, "widdy@gmail.com")
		suite.Assertions.Equal(responseBody.Data.User[0].Pronouns, "he/him")
		suite.Assertions.Equal(responseBody.Data.User[0].Country, "Jawa Barat, Indonesia")
		suite.Assertions.Equal(responseBody.Data.User[0].JobTitle, "Software Engineer")
	})
}

func (suite *E2ETestSuite) TestUser_FindByIDSuccess() {
	id := suite.TestUser_CreateSuccess()

	suite.Run("Find User By ID", func() {
		query := `query FindByIDUser ($id: ID!) {
						user: FindByIDUser (id: $id) {
							id
							name
							email
							password
							bio
							pronouns
							country
							job_title
							image
							created_at
							updated_at
						}
					}`
		variables := fmt.Sprintf(`"id": %d`, id)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody UserResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.User)
		suite.Assertions.Equal(responseBody.Data.User.Name, "Widdy")
		suite.Assertions.Equal(responseBody.Data.User.Email, "widdy@gmail.com")
		suite.Assertions.Equal(responseBody.Data.User.Pronouns, "he/him")
		suite.Assertions.Equal(responseBody.Data.User.Country, "Jawa Barat, Indonesia")
		suite.Assertions.Equal(responseBody.Data.User.JobTitle, "Software Engineer")
	})
}

func (suite *E2ETestSuite) TestUser_UpdateSuccess() {
	id := suite.TestUser_CreateSuccess()

	suite.Run("Update User", func() {
		query := `mutation UpdateUser ($input: UpdateUserRequest!) {
						user: UpdateUser (input: $input) {
							id
							name
							email
							password
							bio
							pronouns
							country
							job_title
							image
							created_at
							updated_at
						}
					}`
		variables := fmt.Sprintf(`"input": {
							"id": %d,
							"name": "Arfiansyah",
							"password": "Rahasia123",
							"bio": "Test Bang Hehe",
							"pronouns": "he/him",
							"country": "Jakarta, Indonesia",
							"job_title": "Software Engineer - Backend",
							"image": null
						}`, id)

		maps := `{ "image": ["variables.input.image"] }`
		query = util.CleanQueryWithVariables(query, variables)
		pathImages := "./database/images/test.png"
		var b bytes.Buffer
		wr := multipart.NewWriter(&b)

		// Operations
		operations, err := wr.CreateFormField("operations")
		suite.Assertions.NoError(err)
		_, err = operations.Write([]byte(query))
		suite.Assertions.NoError(err)

		// Map
		mapping, err := wr.CreateFormField("map")
		suite.Assertions.NoError(err)
		_, err = mapping.Write([]byte(maps))
		suite.Assertions.NoError(err)

		// File
		destination, err := wr.CreateFormFile("image", pathImages)
		suite.Assertions.NoError(err)

		// Open file source
		source, err := os.OpenFile(pathImages, os.O_RDONLY|os.O_CREATE, 0644)
		suite.Assertions.NoError(err)
		defer source.Close()

		_, err = io.Copy(destination, source)
		suite.Assertions.NoError(err)
		wr.Close()

		req := httptest.NewRequest(http.MethodPost, "/query", &b)
		req.Header.Add("Content-Type", wr.FormDataContentType())
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody UserResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Equal(responseBody.Data.User.Name, "Arfiansyah")
		suite.Assertions.Equal(responseBody.Data.User.Email, "widdy@gmail.com")
		suite.Assertions.Equal(responseBody.Data.User.Bio, proto.String("Test Bang Hehe"))
		suite.Assertions.Equal(responseBody.Data.User.Pronouns, "he/him")
		suite.Assertions.Equal(responseBody.Data.User.Country, "Jakarta, Indonesia")
		suite.Assertions.Equal(responseBody.Data.User.JobTitle, "Software Engineer - Backend")
		suite.Assertions.NotEmpty(responseBody.Data.User.Image)
	})
}

func (suite *E2ETestSuite) TestUser_DeleteSuccess() {
	id := suite.TestUser_CreateSuccess()

	suite.Run("Delete User", func() {
		query := `mutation DeleteUser ($id: ID!) {
					DeleteUser (id: $id)
				}`
		variables := fmt.Sprintf(`"id": %d`, id)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody UserResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find User By ID", func() {
		query := `query FindByIDUser ($id: ID!) {
						user: FindByIDUser (id: $id) {
							id
							name
							email
							password
							bio
							pronouns
							country
							job_title
							image
							created_at
							updated_at
						}
					}`
		variables := fmt.Sprintf(`"id": %d`, id)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody UserResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.User)
	})
}
