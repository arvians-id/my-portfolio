package integration

import (
	"github.com/arvians-id/go-portfolio/cmd/config"
	httpfiber "github.com/arvians-id/go-portfolio/internal/http"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/tests/setup"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

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
		CategorySkill *model.CategorySkill `json:"category_skills"`
	} `json:"data"`
}

type CategorySkillSuite struct {
	suite.Suite
	server        *fiber.App
	configuration config.Config
	db            *gorm.DB
}

func TestCategorySkillSuite(t *testing.T) {
	suite.Run(t, new(CategorySkillSuite))
}

func (suite *CategorySkillSuite) SetupTest() {
	// Init Log File
	file, err := os.OpenFile("../../logs/main_test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	suite.Assertions.NoError(err)

	// Init Config
	configuration := config.New("../../.env")

	// Init DB
	db, err := setup.NewPostgresSQLGormTest(configuration)
	suite.Assertions.NoError(err)

	// Init Server
	server, err := httpfiber.NewInitializedRoutes(configuration, file, db)
	suite.Assertions.NoError(err)

	suite.server = server
	suite.configuration = configuration
	suite.db = db
}

func (suite *CategorySkillSuite) TearDownTest() {
	err := setup.TearDownDBTest(suite.db)
	suite.Assertions.NoError(err)
}

func (suite *CategorySkillSuite) TestFindAllSuccess() {
	suite.Run("Create category skill", func() {
		// Insert data
		query := `mutation CreateCategorySkill ($input: CreateCategorySkillRequest!) {
						category_skills: CreateCategorySkill (input: $input) {
							id
							name
							created_at
							updated_at
						}
					}`
		variables := `"input": {
						"name": "Test Bang"
					  }`

		cleanMutation := util.CleanMutation(query, variables)
		bodyRequestCreate := strings.NewReader(cleanMutation)

		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequestCreate)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFyZmlhbkBnbWFpbC5jb20iLCJleHAiOjE2ODY2MjQ1MTMsImlkIjoxfQ.8LQ3kHzUnPsU7L0xmT89JuqCB5U8IQ-7Y-DqZDxoS5o")
		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBodyCreate CategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBodyCreate)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBodyCreate.Data.CategorySkill)
		suite.Assertions.Equal(responseBodyCreate.Data.CategorySkill.Name, "Test Bang")
	})

	suite.Run("The category_skills table is not empty", func() {
		// Find all data
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
		req.Header.Add("Content-Type", "application/json")
		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListCategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Data.CategorySkill)
		suite.Assertions.Equal(responseBody.Data.CategorySkill[0].Name, "Test Bang")
	})

	err := setup.TearDownDBTest(suite.db)
	suite.Assertions.NoError(err)

	suite.Run("The category_skills table is empty", func() {
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
		req.Header.Add("Content-Type", "application/json")
		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody ListCategorySkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Data.CategorySkill)
	})
}
