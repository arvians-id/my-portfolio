package entity

import "github.com/arvians-id/go-portfolio/internal/http/controller/model"

type Education struct {
	ID           int64   `json:"id"`
	Institution  string  `json:"institution"`
	Degree       string  `json:"degree"`
	FieldOfStudy string  `json:"field_of_study"`
	Grade        float64 `json:"grade"`
	Description  *string `json:"description,omitempty"`
	StartDate    string  `json:"start_date"`
	EndDate      *string `json:"end_date,omitempty"`
}

func (e *Education) ToModel() *model.Education {
	return &model.Education{
		ID:           e.ID,
		Institution:  e.Institution,
		Degree:       e.Degree,
		FieldOfStudy: e.FieldOfStudy,
		Grade:        e.Grade,
		Description:  e.Description,
		StartDate:    e.StartDate,
		EndDate:      e.EndDate,
	}
}
