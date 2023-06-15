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

type ListCertificateResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Certificate []*model.Certificate `json:"certificates"`
	} `json:"data"`
}

type CertificateResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Certificate *model.Certificate `json:"certificate"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestCertificate_CreateWithFileSuccess() int64 {
	var id int64

	suite.Run("Create Certificate", func() {
		query := `mutation CreateCertificate ($input: CreateCertificateRequest!) { 
					certificate: CreateCertificate (input: $input) { 
						id 
						name 
						organization 
						issue_date 
						expiration_date 
						credential_id 
						image 
					}
				}`
		variables := `"input": {
							"name": "Tokopedia Devcamp Participant",
							"organization": "Tokopedia",
							"issue_date": "2022-10-17",
							"expiration_date": "2025-10-17",
							"credential_id": "CFFSSSNDWE",
							"image": null
						}`
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

		_, err = io.Copy(destination, source)
		suite.Assertions.NoError(err)
		wr.Close()

		req := httptest.NewRequest(http.MethodPost, "/query", &b)
		req.Header.Add("Content-Type", wr.FormDataContentType())
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Certificate)
		suite.Assertions.Equal(responseBody.Data.Certificate.Name, "Tokopedia Devcamp Participant")
		suite.Assertions.Equal(responseBody.Data.Certificate.Organization, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.Certificate.CredentialID, proto.String("CFFSSSNDWE"))
		suite.Assertions.NotEmpty(responseBody.Data.Certificate.Image)

		id = responseBody.Data.Certificate.ID
	})

	return id
}

func (suite *E2ETestSuite) TestCertificate_CreateSuccess() int64 {
	var id int64

	suite.Run("Create Certificate", func() {
		query := `mutation CreateCertificate ($input: CreateCertificateRequest!) { 
					certificate: CreateCertificate (input: $input) { 
						id 
						name 
						organization 
						issue_date 
						expiration_date 
						credential_id 
						image 
					}
				}`
		variables := `"input": {
							"name": "Tokopedia Devcamp Participant",
							"organization": "Tokopedia",
							"issue_date": "2022-10-17",
							"expiration_date": "2025-10-17",
							"credential_id": "CFFSSSNDWE",
							"image": null
						}`
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Certificate)
		suite.Assertions.Equal(responseBody.Data.Certificate.Name, "Tokopedia Devcamp Participant")
		suite.Assertions.Equal(responseBody.Data.Certificate.Organization, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.Certificate.CredentialID, proto.String("CFFSSSNDWE"))
		suite.Assertions.Empty(responseBody.Data.Certificate.Image)

		id = responseBody.Data.Certificate.ID
	})

	return id
}

func (suite *E2ETestSuite) TestCertificate_FindAllSuccess() {
	suite.TestCertificate_CreateSuccess()

	suite.Run("Find All Certificate", func() {
		query := `query FindAllCertificate { 
					certificates: FindAllCertificate { 
						id 
						name 
						organization 
						issue_date 
						expiration_date 
						credential_id 
						image 
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

		var responseBody ListCertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Certificate)
		suite.Assertions.Equal(responseBody.Data.Certificate[0].Name, "Tokopedia Devcamp Participant")
		suite.Assertions.Equal(responseBody.Data.Certificate[0].Organization, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.Certificate[0].CredentialID, proto.String("CFFSSSNDWE"))
		suite.Assertions.Empty(responseBody.Data.Certificate[0].Image)
	})
}

func (suite *E2ETestSuite) TestCertificate_FindByIDSuccess() {
	id := suite.TestCertificate_CreateSuccess()

	suite.Run("Find By ID Certificate", func() {
		query := `query FindByIDCertificate ($id: ID!) { 
					certificate: FindByIDCertificate (id: $id) { 
						id 
						name 
						organization 
						issue_date 
						expiration_date 
						credential_id 
						image 
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

		var responseBody CertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Certificate)
		suite.Assertions.Equal(responseBody.Data.Certificate.Name, "Tokopedia Devcamp Participant")
		suite.Assertions.Equal(responseBody.Data.Certificate.Organization, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.Certificate.CredentialID, proto.String("CFFSSSNDWE"))
		suite.Assertions.Empty(responseBody.Data.Certificate.Image)
		suite.Assertions.Equal(responseBody.Data.Certificate.ID, id)
	})
}

func (suite *E2ETestSuite) TestCertificate_UpdateWithFileSuccess() {
	id := suite.TestCertificate_CreateSuccess()

	suite.Run("Update Certificate With File", func() {
		query := `mutation UpdateCertificate ($input: UpdateCertificateRequest!) { 
					certificate: UpdateCertificate (input: $input) { 
						id 
						name 
						organization 
						issue_date 
						expiration_date 
						credential_id 
						image 
					}
				}`
		variables := fmt.Sprintf(`"input": {
							"id": %d,
							"name": "Tokopedia Academy",
							"organization": "Tokopedia",
							"issue_date": "2022-10-17",
							"expiration_date": "2025-10-17",
							"credential_id": "KKSJJDBBAH",
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

		_, err = io.Copy(destination, source)
		suite.Assertions.NoError(err)
		wr.Close()

		req := httptest.NewRequest(http.MethodPost, "/query", &b)
		req.Header.Add("Content-Type", wr.FormDataContentType())
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Certificate)
		suite.Assertions.Equal(responseBody.Data.Certificate.Name, "Tokopedia Academy")
		suite.Assertions.Equal(responseBody.Data.Certificate.Organization, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.Certificate.CredentialID, proto.String("KKSJJDBBAH"))
		suite.Assertions.NotEmpty(responseBody.Data.Certificate.Image)
	})
}

func (suite *E2ETestSuite) TestCertificate_DeleteSuccess() {
	id := suite.TestCertificate_CreateSuccess()

	suite.Run("Delete Certificate", func() {
		query := `mutation DeleteCertificate ($id: ID!) {
					DeleteCertificate (id: $id)
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

		var responseBody CertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find By ID Certificate", func() {
		query := `query FindByIDCertificate ($id: ID!) { 
					certificate: FindByIDCertificate (id: $id) { 
						id 
						name 
						organization 
						issue_date 
						expiration_date 
						credential_id 
						image 
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

		var responseBody CertificateResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.Certificate)
	})
}
