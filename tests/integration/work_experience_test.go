package integration

import (
	"fmt"
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"github.com/arvians-id/go-portfolio/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/proto"
	"net/http"
	"net/http/httptest"
	"strings"
)

type ListWorkExperienceResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		WorkExperience []*model.WorkExperience `json:"work_experiences"`
	} `json:"data"`
}

type WorkExperienceResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		WorkExperience *model.WorkExperience `json:"work_experience"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestWorkExperience_CreateSuccess() int64 {
	var id int64
	suite.Run("Create work experience", func() {
		query := `mutation CreateWorkExperience ($input: CreateWorkExperienceRequest!) {
						work_experience: CreateWorkExperience (input: $input) {
							id
							role
							company
							description
							start_date
							end_date
							job_type
							created_at
							updated_at
						}
					}`
		variables := `"input": {
						"role": "Software Engineer - Backend",
						"company": "Tokopedia",
						"description": "Hello",
						"start_date": "2023-06-15",
						"job_type": "Full Time",
						"skills": []
					  }`

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody WorkExperienceResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.WorkExperience)
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Role, "Software Engineer - Backend")
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Company, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Description, proto.String("Hello"))
		suite.Assertions.Equal(responseBody.Data.WorkExperience.JobType, "Full Time")

		id = responseBody.Data.WorkExperience.ID
	})

	return id
}

func (suite *E2ETestSuite) TestWorkExperience_FindAllSuccess() {
	_ = suite.TestWorkExperience_CreateSuccess()

	suite.Run("Find all work experience", func() {
		query := `query FindAllWorkExperience {
						work_experiences: FindAllWorkExperience {
							id
							role
							company
							description
							start_date
							end_date
							job_type
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

		var responseBody ListWorkExperienceResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.WorkExperience)
		suite.Assertions.Equal(responseBody.Data.WorkExperience[0].Role, "Software Engineer - Backend")
		suite.Assertions.Equal(responseBody.Data.WorkExperience[0].Company, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.WorkExperience[0].Description, proto.String("Hello"))
		suite.Assertions.Equal(responseBody.Data.WorkExperience[0].JobType, "Full Time")
	})
}

func (suite *E2ETestSuite) TestWorkExperience_FindByIDSuccess() model.WorkExperience {
	id := suite.TestWorkExperience_CreateSuccess()

	var workExperience model.WorkExperience
	suite.Run("Find by ID work experience", func() {
		query := `query FindByIDWorkExperience ($id: ID!) {
						work_experience: FindByIDWorkExperience (id: $id) {
							id
							role
							company
							description
							start_date
							end_date
							job_type
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

		var responseBody WorkExperienceResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.WorkExperience)
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Role, "Software Engineer - Backend")
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Company, "Tokopedia")
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Description, proto.String("Hello"))
		suite.Assertions.Equal(responseBody.Data.WorkExperience.JobType, "Full Time")

		workExperience = *responseBody.Data.WorkExperience
	})

	return workExperience
}

func (suite *E2ETestSuite) TestWorkExperience_UpdateSuccess() {
	id := suite.TestWorkExperience_CreateSuccess()

	suite.Run("Update work experience", func() {
		query := `mutation UpdateWorkExperience ($input: UpdateWorkExperienceRequest!) {
						work_experience: UpdateWorkExperience (input: $input) {
							id
							role
							company
							description
							start_date
							end_date
							job_type
						}
					}`
		variables := fmt.Sprintf(`"input": {
											"id": %d,
											"role": "Software Engineer",
											"company": "Goto Finance",
											"description": "Hi, I'm a backend engineer",
											"start_date": "2023-06-15",
											"job_type": "Permanent Employee",
											"skills": []
										  }`, id)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody WorkExperienceResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.WorkExperience)
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Role, "Software Engineer")
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Company, "Goto Finance")
		suite.Assertions.Equal(responseBody.Data.WorkExperience.Description, proto.String("Hi, I'm a backend engineer"))
		suite.Assertions.Equal(responseBody.Data.WorkExperience.JobType, "Permanent Employee")
	})
}

func (suite *E2ETestSuite) TestWorkExperience_DeleteSuccess() {
	id := suite.TestWorkExperience_CreateSuccess()

	suite.Run("Delete work experience", func() {
		query := `mutation DeleteWorkExperience ($id: ID!) {
					DeleteWorkExperience (id: $id)
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

		var responseBody WorkExperienceResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find by id work experience", func() {
		query := `query FindByIDWorkExperience ($id: ID!) {
						work_experience: FindByIDWorkExperience (id: $id) {
							id
							role
							company
							description
							start_date
							end_date
							job_type
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

		var responseBody WorkExperienceResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.WorkExperience)
	})
}
