package integration

import (
	"fmt"
	"github.com/arvians-id/go-portfolio/cmd/config"
	httpfiber "github.com/arvians-id/go-portfolio/internal/http"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Error struct {
	Message    string          `json:"message"`
	Path       []string        `json:"path"`
	Extensions ErrorExtensions `json:"extensions"`
}

type ErrorExtensions struct {
	Code string `json:"code"`
}

type ListCategorySkillResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		CategorySkill []*model.CategorySkill `json:"category_skills"`
	} `json:"data"`
}

type CategorySkillResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		CategorySkill *model.CategorySkill `json:"category_skill"`
	} `json:"data"`
}

type CategorySkillSuite struct {
	suite.Suite
	configuration config.Config
	server        *fiber.App
	db            *gorm.DB
}

func TestCategorySkillSuite(t *testing.T) {
	suite.Run(t, new(CategorySkillSuite))
}

func (suite *CategorySkillSuite) SetupSuite() {
	// Init Config
	configuration := config.New("../../.env")

	// Init DB
	db, err := config.NewPostgresSQLGormTest(configuration)
	suite.Require().NoError(err)

	// Init Server
	server, err := httpfiber.NewInitializedRoutes(configuration, nil, db)
	suite.Assertions.NoError(err)

	// Set Suite Variables
	suite.configuration = configuration
	suite.server = server
	suite.db = db
}

func (suite *CategorySkillSuite) TearDownTest() {
	suite.NoError(config.TearDownDBTest(suite.db))
}

func (suite *CategorySkillSuite) TestFindAllSuccess() {
	_ = suite.TestCreateSuccess()

	suite.Run("The category_skills table is not empty, should return all data", func() {
		query := `query FindAllCategorySkill {
			category_skills: FindAllCategorySkill {
				id
				name
				created_at
				updated_at
				skills {
					name
				}
			}
		}`
		query = util.CleanQuery(query)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListCategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.CategorySkill)
		suite.Assertions.Equal(responseBody.Data.CategorySkill[0].Name, "Frameworks")
	})
}

func (suite *CategorySkillSuite) TestFindAllError() {
	suite.Run("The category_skills table is empty, failed to get all data", func() {
		query := `query FindAllCategorySkill {
			category_skills: FindAllCategorySkill {
				id
				name
				created_at
				updated_at
				skills {
					name
				}
			}
		}`
		cleanQuery := util.CleanQuery(query)
		bodyRequest := strings.NewReader(cleanQuery)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListCategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Data.CategorySkill)
	})
}

func (suite *CategorySkillSuite) TestFindOneSuccess() string {
	id := suite.TestCreateSuccess()

	var name string
	suite.Run("The category_skills table is not empty, should return one data", func() {
		query := `query FindByIDCategorySkill ($id: ID!) {
					category_skill: FindByIDCategorySkill (id: $id) {
						id
						name
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

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.CategorySkill)
		suite.Assertions.Equal(responseBody.Data.CategorySkill.Name, "Frameworks")

		name = responseBody.Data.CategorySkill.Name
	})

	return name
}

func (suite *CategorySkillSuite) TestFindOneError() {
	suite.Run("The category_skills table is empty, failed to get one data", func() {
		query := `query FindByIDCategorySkill ($id: ID!) {
					category_skill: FindByIDCategorySkill (id: $id) {
						id
						name
						created_at
						updated_at
					}
				}`
		variables := `"id": 1`

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.CategorySkill)
	})
}

func (suite *CategorySkillSuite) TestCreateSuccess() int64 {
	var id int64
	suite.Run("Create category skill", func() {
		query := `mutation CreateCategorySkill ($input: CreateCategorySkillRequest!) {
						category_skill: CreateCategorySkill (input: $input) {
							id
							name
							created_at
							updated_at
						}
					}`
		variables := `"input": {
						"name": "Frameworks"
					  }`

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.CategorySkill)
		suite.Assertions.Equal(responseBody.Data.CategorySkill.Name, "Frameworks")

		id = responseBody.Data.CategorySkill.ID
	})

	return id
}

func (suite *CategorySkillSuite) TestCreateError() {
	suite.Run("The name field is empty, failed to create category skill", func() {
		query := `mutation CreateCategorySkill ($input: CreateCategorySkillRequest!) {
						category_skill: CreateCategorySkill (input: $input) {
							id
							name
							created_at
							updated_at
						}
					}`
		variables := `"input": {
						"name": ""
					  }`

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Contains(responseBody.Errors[0].Message, "validation error")
		suite.Assertions.Empty(responseBody.Data.CategorySkill)
	})
}

func (suite *CategorySkillSuite) TestUpdateSuccess() {
	id := suite.TestCreateSuccess()

	suite.Run("Update category skill", func() {
		query := `mutation UpdateCategorySkill ($input: UpdateCategorySkillRequest!) {
						category_skill: UpdateCategorySkill (input: $input) {
							id
							name
							created_at
							updated_at
						}
					}`
		variables := fmt.Sprintf(`"input": {
						"id": %d,
						"name": "Tools"
					  }`, id)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.CategorySkill)
		suite.Assertions.Equal(responseBody.Data.CategorySkill.Name, "Tools")
	})
}

func (suite *CategorySkillSuite) TestUpdateError() {
	suite.Run("The id field is empty, failed to update category skill", func() {
		query := `mutation UpdateCategorySkill ($input: UpdateCategorySkillRequest!) {
						category_skill: UpdateCategorySkill (input: $input) {
							id
							name
							created_at
							updated_at
						}
					}`
		variables := `"input": {
						"id": 0,
						"name": "Tools"
					  }`

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Contains(responseBody.Errors[0].Message, "validation error")
		suite.Assertions.Empty(responseBody.Data.CategorySkill)
	})
}

func (suite *CategorySkillSuite) TestDeleteSuccess() {
	id := suite.TestCreateSuccess()

	suite.Run("Delete category skill", func() {
		query := `mutation DeleteCategorySkill ($id: ID!) {
						deleted: DeleteCategorySkill (id: $id)
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

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("The category_skills table is empty, failed to get one data", func() {
		query := `query FindByIDCategorySkill ($id: ID!) {
					category_skill: FindByIDCategorySkill (id: $id) {
						id
						name
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

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.CategorySkill)
	})
}
