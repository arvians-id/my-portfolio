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

type ListEducationResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Education []*model.Education `json:"educations"`
	} `json:"data"`
}

type EducationResponse struct {
	Errors []Error `json:"errors"`
	Data   struct {
		Education *model.Education `json:"education"`
	} `json:"data"`
}

func (suite *E2ETestSuite) TestEducation_FindAllSuccess() {
	_ = suite.TestEducation_CreateSuccess()

	suite.Run("Find all education", func() {
		query := `query FindAllEducation {
						educations: FindAllEducation {
							id
							institution
							degree
							field_of_study
							grade
							description
							start_date
							end_date
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

		var responseBody ListEducationResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Education)
		suite.Assertions.Equal(responseBody.Data.Education[0].Institution, "UMMI")
		suite.Assertions.Equal(responseBody.Data.Education[0].Degree, "Bachelor")
		suite.Assertions.Equal(responseBody.Data.Education[0].FieldOfStudy, "Computer Science")
		suite.Assertions.Equal(responseBody.Data.Education[0].Grade, 3.55)
		suite.Assertions.Equal(responseBody.Data.Education[0].Description, proto.String("Test"))
	})
}

func (suite *E2ETestSuite) TestEducation_FindByIDSuccess() model.Education {
	id := suite.TestEducation_CreateSuccess()

	var education model.Education
	suite.Run("Find education by ID", func() {
		query := `query FindByIDEducation ($id: ID!) {
						education: FindByIDEducation (id: $id) {
							id
							institution
							degree
							field_of_study
							grade
							description
							start_date
							end_date
						}
					}`
		variables := `"id": ` + fmt.Sprintf("%d", id)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody EducationResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Education)
		suite.Assertions.Equal(responseBody.Data.Education.Institution, "UMMI")
		suite.Assertions.Equal(responseBody.Data.Education.Degree, "Bachelor")
		suite.Assertions.Equal(responseBody.Data.Education.FieldOfStudy, "Computer Science")
		suite.Assertions.Equal(responseBody.Data.Education.Grade, 3.55)
		suite.Assertions.Equal(responseBody.Data.Education.Description, proto.String("Test"))

		education = *responseBody.Data.Education
	})

	return education
}

func (suite *E2ETestSuite) TestEducation_CreateSuccess() int64 {
	var id int64
	suite.Run("Create education skill", func() {
		query := `mutation CreateEducation ($input: CreateEducationRequest!) {
						education: CreateEducation (input: $input) {
							id
							institution
							degree
							field_of_study
							grade
							description
							start_date
							end_date
						}
					}`
		variables := `"input": {
						"institution": "UMMI",
						"degree": "Bachelor",
						"field_of_study": "Computer Science",
						"grade": "3.55",
						"description": "Test",
						"start_date": "2023-01-01",
						"end_date": "2018-07-06"
					  }`

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody EducationResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Education)
		suite.Assertions.Equal(responseBody.Data.Education.Institution, "UMMI")
		suite.Assertions.Equal(responseBody.Data.Education.Degree, "Bachelor")
		suite.Assertions.Equal(responseBody.Data.Education.FieldOfStudy, "Computer Science")
		suite.Assertions.Equal(responseBody.Data.Education.Grade, 3.55)
		suite.Assertions.Equal(responseBody.Data.Education.Description, proto.String("Test"))

		id = responseBody.Data.Education.ID
	})

	return id
}

func (suite *E2ETestSuite) TestEducation_UpdateSuccess() {
	id := suite.TestEducation_CreateSuccess()

	suite.Run("Update education skill", func() {
		query := `mutation UpdateEducation ($input: UpdateEducationRequest!) {
						education: UpdateEducation (input: $input) {
							id
							institution
							degree
							field_of_study
							grade
							description
							start_date
							end_date
						}
					}`
		variables := fmt.Sprintf(`"input": {
											"id": %d,
											"institution": "Universitas Muhammadiyah Sukabumi",
											"degree": "Sarjana",
											"field_of_study": "Teknik Informatika",
											"grade": "3.6",
											"description": "Test Bang",
											"start_date": "2023-01-02",
											"end_date": "2018-07-02"
										  }`, id)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody EducationResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
		suite.Assertions.NotEmpty(responseBody.Data.Education)
		suite.Assertions.Equal(responseBody.Data.Education.Institution, "Universitas Muhammadiyah Sukabumi")
		suite.Assertions.Equal(responseBody.Data.Education.Degree, "Sarjana")
		suite.Assertions.Equal(responseBody.Data.Education.FieldOfStudy, "Teknik Informatika")
		suite.Assertions.Equal(responseBody.Data.Education.Grade, 3.6)
		suite.Assertions.Equal(responseBody.Data.Education.Description, proto.String("Test Bang"))
	})
}

func (suite *E2ETestSuite) TestEducation_DeleteSuccess() {
	id := suite.TestEducation_CreateSuccess()

	suite.Run("Delete education skill", func() {
		query := `mutation DeleteEducation ($id: ID!) {
					deleted: DeleteEducation (id: $id)
				}`
		variables := `"id": ` + fmt.Sprintf("%d", id)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody EducationResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.Empty(responseBody.Errors)
	})

	suite.Run("Find education by ID", func() {
		query := `query FindByIDEducation ($id: ID!) {
						education: FindByIDEducation (id: $id) {
							id
							institution
							degree
							field_of_study
							grade
							description
							start_date
							end_date
						}
					}`
		variables := `"id": ` + fmt.Sprintf("%d", id)

		query = util.CleanQueryWithVariables(query, variables)
		bodyRequest := strings.NewReader(query)
		req := httptest.NewRequest(http.MethodPost, "/query", bodyRequest)
		req.Header.Add("X-API-KEY", suite.configuration.Get("X_API_KEY"))
		req.Header.Add("Content-Type", fiber.MIMEApplicationJSON)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", suite.configuration.Get("JWT_TEST")))

		res, err := suite.server.Test(req)
		suite.Assertions.NoError(err)

		var responseBody EducationResponse
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		suite.Assertions.NoError(err)

		suite.Assertions.NotEmpty(responseBody.Errors)
		suite.Assertions.Empty(responseBody.Data.Education)
	})
}
