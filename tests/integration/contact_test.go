package integration

import (
	"bytes"
	"fmt"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type ListContactResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Contact []*model.Contact `json:"contacts"`
	} `json:"data"`
}

type ContactResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Contact *model.Contact `json:"contact"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestContact_CreateWithFileSuccess() int64 {
	var id int64

	suite.Run("Create Contact", func() {
		query := `mutation CreateContact ($input: CreateContactRequest!) {
						contact: CreateContact (input: $input) {
							id 
							platform 
							url
							icon
						}
					}`
		variables := `"input": {
						"platform": "Facebook",
						"url": "www.facebook.com",
						"icon": null
					}`
		maps := `{ "icon": ["variables.input.icon"] }`
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
		destination, err := wr.CreateFormFile("icon", pathImages)
		suite.Assertions.NoError(err)

		// Open file source
		source, err := os.OpenFile(pathImages, os.O_RDONLY|os.O_CREATE, 0644)
		suite.Assertions.NoError(err)

		_, err = io.Copy(destination, source)
		suite.Assertions.NoError(err)
		wr.Close()

		req := httptest.NewRequest(http.MethodPost, "/query", &b)
		req.Header.Add("Content-Type", wr.FormDataContentType())
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Contact)
		suite.Assertions.Equal(responseBody.Data.Contact.Platform, "Facebook")
		suite.Assertions.Equal(responseBody.Data.Contact.URL, "www.facebook.com")
		suite.Assertions.NotEmpty(responseBody.Data.Contact.Icon)

		id = responseBody.Data.Contact.ID
	})

	return id
}

func (suite *E2ETestSuite) TestContact_CreateSuccess() int64 {
	var id int64

	suite.Run("Create Contact", func() {
		query := `mutation CreateContact ($input: CreateContactRequest!) {
						contact: CreateContact (input: $input) {
							id 
							platform 
							url
							icon
						}
					}`
		variables := `"input": {
						"platform": "Facebook",
						"url": "www.facebook.com",
						"icon": null
					}`
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Contact)
		suite.Assertions.Equal(responseBody.Data.Contact.Platform, "Facebook")
		suite.Assertions.Equal(responseBody.Data.Contact.URL, "www.facebook.com")
		suite.Assertions.Empty(responseBody.Data.Contact.Icon)

		id = responseBody.Data.Contact.ID
	})

	return id
}

func (suite *E2ETestSuite) TestContact_FindAllSuccess() {
	suite.TestContact_CreateSuccess()

	suite.Run("Find All Contact", func() {
		query := `query FindAllContact {
						contacts: FindAllContact {
							id 
							platform 
							url
							icon
						}
					}`

		query = util.CleanQuery(query)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Contact)
		suite.Assertions.Equal(responseBody.Data.Contact[0].Platform, "Facebook")
		suite.Assertions.Equal(responseBody.Data.Contact[0].URL, "www.facebook.com")
		suite.Assertions.Empty(responseBody.Data.Contact[0].Icon)
	})
}

func (suite *E2ETestSuite) TestContact_FindByIDSuccess() {
	id := suite.TestContact_CreateSuccess()

	suite.Run("Find By ID Contact", func() {
		query := `query FindByIDContact ($id: ID!) {
						contact: FindByIDContact (id: $id) {
							id 
							platform 
							url
							icon
						}
					}`
		variables := fmt.Sprintf(`"id": %d`, id)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Contact)
		suite.Assertions.Equal(responseBody.Data.Contact.Platform, "Facebook")
		suite.Assertions.Equal(responseBody.Data.Contact.URL, "www.facebook.com")
		suite.Assertions.Empty(responseBody.Data.Contact.Icon)
	})
}

func (suite *E2ETestSuite) TestContact_UpdateWithFileSuccess() {
	id := suite.TestContact_CreateSuccess()

	suite.Run("Update Contact", func() {
		query := `mutation UpdateContact ($input: UpdateContactRequest!) {
						contact: UpdateContact (input: $input) {
							id 
							platform 
							url
							icon
						}
					}`
		variables := fmt.Sprintf(` "input": {
						"id": %d,
						"platform": "Google",
						"url": "www.google.com",
						"icon": null
					}`, id)
		maps := `{ "icon": ["variables.input.icon"] }`
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
		destination, err := wr.CreateFormFile("icon", pathImages)
		suite.Assertions.NoError(err)

		// Open file source
		source, err := os.OpenFile(pathImages, os.O_RDONLY|os.O_CREATE, 0644)
		suite.Assertions.NoError(err)

		_, err = io.Copy(destination, source)
		suite.Assertions.NoError(err)
		wr.Close()

		req := httptest.NewRequest(http.MethodPost, "/query", &b)
		req.Header.Add("Content-Type", wr.FormDataContentType())
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Contact)
		suite.Assertions.Equal(responseBody.Data.Contact.Platform, "Google")
		suite.Assertions.Equal(responseBody.Data.Contact.URL, "www.google.com")
		suite.Assertions.NotEmpty(responseBody.Data.Contact.Icon)
	})
}

func (suite *E2ETestSuite) TestContact_DeleteSuccess() {
	id := suite.TestContact_CreateSuccess()

	suite.Run("Delete Contact", func() {
		query := `mutation DeleteContact ($id: ID!) {
					DeleteContact (id: $id)
				}`
		variables := fmt.Sprintf(`"id": %d`, id)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find By ID Contact", func() {
		query := `query FindByIDContact ($id: ID!) {
						contact: FindByIDContact (id: $id) {
							id 
							platform 
							url
							icon
						}
					}`
		variables := fmt.Sprintf(`"id": %d`, id)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ContactResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.Contact)
	})
}
