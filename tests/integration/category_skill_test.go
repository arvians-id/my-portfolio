package integration

import (
	"fmt"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
	"strings"
)

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

func (suite *E2ETestSuite) TestCategorySkill_FindAllSuccess() {
	_ = suite.TestCategorySkill_CreateSuccess()

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

func (suite *E2ETestSuite) TestCategorySkill_FindAllError() {
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

func (suite *E2ETestSuite) TestCategorySkill_FindByIDSuccess() model.CategorySkill {
	id := suite.TestCategorySkill_CreateSuccess()

	var categorySkill model.CategorySkill
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

		categorySkill = *responseBody.Data.CategorySkill
	})

	return categorySkill
}

func (suite *E2ETestSuite) TestCategorySkill_FindByIDError() {
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

func (suite *E2ETestSuite) TestCategorySkill_CreateSuccess() int64 {
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

func (suite *E2ETestSuite) TestCategorySkill_CreateError() {
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

func (suite *E2ETestSuite) TestCategorySkill_UpdateSuccess() {
	id := suite.TestCategorySkill_CreateSuccess()

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

func (suite *E2ETestSuite) TestCategorySkill_UpdateError() {
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

func (suite *E2ETestSuite) TestCategorySkill_DeleteSuccess() {
	id := suite.TestCategorySkill_CreateSuccess()

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
