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

type ListProjectResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Project []*model.Project `json:"projects"`
	} `json:"data"`
}

type ProjectResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Project *model.Project `json:"project"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestProject_CreateWithFileSuccess() int64 {
	var id int64
	idSkill := suite.TestSkill_CreateSuccess()

	suite.Run("Create project", func() {
		query := `mutation CreateProject ($input: CreateProjectRequest!) {
						project: CreateProject (input: $input) {
							id 
							category 
							title 
							description 
							url 
							is_featured 
							date 
							working_type 
							created_at 
							updated_at
						}
					}`
		variables := fmt.Sprintf(`"input": {
											"category": "Tutorial",
											"title": "Learn Go RabbitMQ",
											"description": "Belajar Message Broker",
											"url": "github.com/arvians-id/mq",
											"is_featured": true,
											"date": "1994-10-27",
											"working_type": "self",
											"skills": [%d],
											"images": [null]
										}`, idSkill)
		maps := `{ "0": ["variables.input.images.0"] }`
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
		destination, err := wr.CreateFormFile("0", pathImages)
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

		var responseBody ProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Project)
		suite.Assertions.Equal(responseBody.Data.Project.Category, "Tutorial")
		suite.Assertions.Equal(responseBody.Data.Project.Title, "Learn Go RabbitMQ")
		suite.Assertions.Equal(responseBody.Data.Project.Description, proto.String("Belajar Message Broker"))
		suite.Assertions.Equal(responseBody.Data.Project.URL, proto.String("github.com/arvians-id/mq"))
		suite.Assertions.Equal(responseBody.Data.Project.IsFeatured, proto.Bool(true))
		suite.Assertions.Equal(responseBody.Data.Project.WorkingType, "self")

		id = responseBody.Data.Project.ID
	})

	return id
}

func (suite *E2ETestSuite) TestProject_CreateSuccess() int64 {
	var id int64
	idSkill := suite.TestSkill_CreateSuccess()

	suite.Run("Create project", func() {
		query := `mutation CreateProject ($input: CreateProjectRequest!) {
						project: CreateProject (input: $input) {
							id 
							category 
							title 
							description 
							url 
							is_featured 
							date 
							working_type 
							created_at 
							updated_at
						}
					}`
		variables := fmt.Sprintf(`"input": {
											"category": "Tutorial",
											"title": "Learn Go RabbitMQ",
											"description": "Belajar Message Broker",
											"url": "github.com/arvians-id/mq",
											"is_featured": true,
											"date": "1994-10-27",
											"working_type": "self",
											"skills": [%d]
										}`, idSkill)
		query = util.CleanQueryWithVariables(query, variables)

		req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewBufferString(query))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Project)
		suite.Assertions.Equal(responseBody.Data.Project.Category, "Tutorial")
		suite.Assertions.Equal(responseBody.Data.Project.Title, "Learn Go RabbitMQ")
		suite.Assertions.Equal(responseBody.Data.Project.Description, proto.String("Belajar Message Broker"))
		suite.Assertions.Equal(responseBody.Data.Project.URL, proto.String("github.com/arvians-id/mq"))
		suite.Assertions.Equal(responseBody.Data.Project.IsFeatured, proto.Bool(true))
		suite.Assertions.Equal(responseBody.Data.Project.WorkingType, "self")

		id = responseBody.Data.Project.ID
	})

	return id
}

func (suite *E2ETestSuite) TestProject_FindAllSuccess() {
	suite.TestProject_CreateSuccess()

	suite.Run("Find All Project", func() {
		query := `query FindAllProject($search: String) {
						projects: FindAllProject(name: $search) {
							id
							category
							title
							description
							url
							is_featured
							date
							working_type
							created_at
							updated_at
							skills {
								name
								category_skill {
									name
								}
							}
							images {
								id
								project_id
								image
							}
						}
					}`

		variables := ""
		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Project)
		suite.Assertions.Equal(responseBody.Data.Project[0].Category, "Tutorial")
		suite.Assertions.Equal(responseBody.Data.Project[0].Title, "Learn Go RabbitMQ")
		suite.Assertions.Equal(responseBody.Data.Project[0].Description, proto.String("Belajar Message Broker"))
		suite.Assertions.Equal(responseBody.Data.Project[0].URL, proto.String("github.com/arvians-id/mq"))
		suite.Assertions.Equal(responseBody.Data.Project[0].IsFeatured, proto.Bool(true))
		suite.Assertions.Equal(responseBody.Data.Project[0].WorkingType, "self")
		suite.Assertions.NotEmpty(responseBody.Data.Project[0].Skills)
		suite.Assertions.Empty(responseBody.Data.Project[0].Images)
	})
}

func (suite *E2ETestSuite) TestProject_FindByIDSuccess() {
	id := suite.TestProject_CreateWithFileSuccess()

	suite.Run("Find By ID Project", func() {
		query := `query FindByIDProject ($id: ID!) {
					project: FindByIDProject (id: $id) {
						id
						category
						title
						description
						url
						is_featured
						date
						working_type
						skills {
							name
							icon
						}
						images {
							image
						}
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

		var responseBody ProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Project)
		suite.Assertions.Equal(responseBody.Data.Project.Category, "Tutorial")
		suite.Assertions.Equal(responseBody.Data.Project.Title, "Learn Go RabbitMQ")
		suite.Assertions.Equal(responseBody.Data.Project.Description, proto.String("Belajar Message Broker"))
		suite.Assertions.Equal(responseBody.Data.Project.URL, proto.String("github.com/arvians-id/mq"))
		suite.Assertions.Equal(responseBody.Data.Project.IsFeatured, proto.Bool(true))
		suite.Assertions.Equal(responseBody.Data.Project.WorkingType, "self")
		suite.Assertions.NotEmpty(responseBody.Data.Project.Skills)
		suite.Assertions.NotEmpty(responseBody.Data.Project.Images)
	})
}

func (suite *E2ETestSuite) TestProject_UpdateSuccess() {
	id := suite.TestProject_CreateSuccess()
	idSkill := suite.TestSkill_CreateSuccess()

	suite.Run("Update Project", func() {
		query := `mutation UpdateProject ($input: UpdateProjectRequest!) { 
					project: UpdateProject (input: $input) { 
							id 
							category 
							title 
							description 
							url 
							is_featured 
							date 
							working_type 
							created_at 
							updated_at 
						} 
					}`

		variables := fmt.Sprintf(`"input": {
											"id": %d,
											"category": "Paid Project",
											"title": "Surat Kuasa",
											"description": "Pengadilan Negeri Sukabumi",
											"url": "github.com/arvians-id/surat-kuasa",
											"is_featured": true,
											"date": "1994-10-27",
											"working_type": "self",
											"skills": [%d]
										}`, id, idSkill)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Project)
		suite.Assertions.Equal(responseBody.Data.Project.Category, "Paid Project")
		suite.Assertions.Equal(responseBody.Data.Project.Title, "Surat Kuasa")
		suite.Assertions.Equal(responseBody.Data.Project.Description, proto.String("Pengadilan Negeri Sukabumi"))
		suite.Assertions.Equal(responseBody.Data.Project.URL, proto.String("github.com/arvians-id/surat-kuasa"))
		suite.Assertions.Equal(responseBody.Data.Project.IsFeatured, proto.Bool(true))
		suite.Assertions.Equal(responseBody.Data.Project.WorkingType, "self")
	})
}

func (suite *E2ETestSuite) TestProject_DeleteSuccess() {
	id := suite.TestProject_CreateSuccess()

	suite.Run("Update Project", func() {
		query := `mutation DeleteProject ($id: ID!) {
						DeleteProject (id: $id)
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

		var responseBody ProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find By ID Project", func() {
		query := `query FindByIDProject ($id: ID!) {
					project: FindByIDProject (id: $id) {
						id
						category
						title
						description
						url
						is_featured
						date
						working_type
						skills {
							name
							icon
						}
						images {
							image
						}
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

		var responseBody ProjectResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.Project)
	})
}
