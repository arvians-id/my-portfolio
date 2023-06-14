package entity

import "gorm.io/gorm"

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

func (e *Education) BeforeCreate(tx *gorm.DB) error {
	if e.Description == nil || *e.Description == "" {
		e.Description = nil
	}

	if e.EndDate == nil || *e.EndDate == "" {
		e.EndDate = nil
	}

	return nil
}
