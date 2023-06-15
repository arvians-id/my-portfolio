package integration

import (
	"bytes"
	"fmt"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type ListSkillResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Skill []*model.Skill `json:"skills"`
	} `json:"data"`
}

type SkillResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Skill *model.Skill `json:"skill"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestSkill_CreateSuccess() int64 {
	var id int64
	idCategorySkill := suite.TestCategorySkill_CreateSuccess()

	suite.Run("Create Skill", func() {
		query := `mutation CreateSkill ($input: CreateSkillRequest!) {
						skill: CreateSkill (input: $input) {
							id 
							category_skill_id 
							name 
							icon
						}
					}`
		variables := fmt.Sprintf(`"input": { 
											"category_skill_id": %d, 
											"name": "Golang", 
											"icon": null 
										}`, idCategorySkill)
		query = util.CleanQueryWithVariables(query, variables)

		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody SkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Skill)
		suite.Assertions.Equal(responseBody.Data.Skill.Name, "Golang")
		suite.Assertions.Equal(responseBody.Data.Skill.CategorySkillID, idCategorySkill)
		suite.Assertions.Empty(responseBody.Data.Skill.Icon)

		id = responseBody.Data.Skill.ID
	})

	return id
}

func (suite *E2ETestSuite) TestSkill_CreateWithFileSuccess() int64 {
	var id int64
	idCategorySkill := suite.TestCategorySkill_CreateSuccess()

	suite.Run("Create Skill", func() {
		query := `mutation CreateSkill ($input: CreateSkillRequest!) {
						skill: CreateSkill (input: $input) {
							id 
							category_skill_id 
							name 
							icon
						}
					}`
		variables := fmt.Sprintf(`"input": { 
											"category_skill_id": %d, 
											"name": "Golang", 
											"icon": null 
										}`, idCategorySkill)
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

		var responseBody SkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Skill)
		suite.Assertions.Equal(responseBody.Data.Skill.Name, "Golang")
		suite.Assertions.Equal(responseBody.Data.Skill.CategorySkillID, idCategorySkill)

		id = responseBody.Data.Skill.ID
	})

	return id
}

func (suite *E2ETestSuite) TestSkill_FindAllSuccess() {
	id := suite.TestSkill_CreateSuccess()

	suite.Run("Find All Skill", func() {
		query := `query FindAllSkill {
						skills: FindAllSkill {
							id
							category_skill_id
							name
							icon
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

		var responseBody ListSkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Skill)
		suite.Assertions.Equal(responseBody.Data.Skill[0].ID, id)
		suite.Assertions.Equal(responseBody.Data.Skill[0].Name, "Golang")
		suite.Assertions.NotEmpty(responseBody.Data.Skill[0].CategorySkillID)
	})
}

func (suite *E2ETestSuite) TestSkill_FindByIDSuccess() model.Skill {
	id := suite.TestSkill_CreateSuccess()

	var skill model.Skill
	suite.Run("Find By ID Skill", func() {
		query := `query FindByIDSkill ($id: ID!) {
						skill: FindByIDSkill (id: $id) {
							id
							category_skill_id
							name
							icon
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

		var responseBody SkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Skill)
		suite.Assertions.Equal(responseBody.Data.Skill.ID, id)
		suite.Assertions.Equal(responseBody.Data.Skill.Name, "Golang")
		suite.Assertions.NotEmpty(responseBody.Data.Skill.CategorySkillID)

		skill = *responseBody.Data.Skill
	})

	return skill
}

func (suite *E2ETestSuite) TestSkill_UpdateWithFileSuccess() {
	id := suite.TestSkill_CreateSuccess()
	idCategorySkill := suite.TestCategorySkill_CreateSuccess()

	suite.Run("Update Skill", func() {
		query := `mutation UpdateSkill ($input: UpdateSkillRequest!) {
						skill: UpdateSkill (input: $input) {
							id
							category_skill_id
							name
							icon
						}
					}`
		variables := fmt.Sprintf(`"input": {
											"id": %d,
											"category_skill_id": %d,
											"name": "Golang Fiber",
											"icon": null
										}`, id, idCategorySkill)
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

		var responseBody SkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Skill)
		suite.Assertions.Equal(responseBody.Data.Skill.ID, id)
		suite.Assertions.Equal(responseBody.Data.Skill.Name, "Golang Fiber")
		suite.Assertions.Equal(responseBody.Data.Skill.CategorySkillID, idCategorySkill)
	})
}

func (suite *E2ETestSuite) TestSkill_DeleteSuccess() {
	id := suite.TestSkill_CreateSuccess()

	suite.Run("Delete Skill", func() {
		query := `mutation DeleteSkill ($id: ID!) {
						DeleteSkill (id: $id)
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

		var responseBody SkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find By ID Skill", func() {
		query := `query FindByIDSkill ($id: ID!) {
						skill: FindByIDSkill (id: $id) {
							id
							category_skill_id
							name
							icon
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

		var responseBody SkillResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.Skill)
	})
}
