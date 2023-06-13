package integration

import "github.com/arvians-id/go-portfolio/internal/http/controller/model"

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
