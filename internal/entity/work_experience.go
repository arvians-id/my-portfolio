package entity

import (
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"time"
)

type WorkExperience struct {
	ID          int64     `json:"id"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Description *string   `json:"description,omitempty"`
	StartDate   string    `json:"start_date"`
	EndDate     *string   `json:"end_date,omitempty"`
	JobType     string    `json:"job_type"`
	Skills      []*Skill  `json:"skills,omitempty" gorm:"many2many:work_experience_skill;"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (w *WorkExperience) ToModel() *model.WorkExperience {
	return &model.WorkExperience{
		ID:          w.ID,
		Role:        w.Role,
		Company:     w.Company,
		Description: w.Description,
		StartDate:   w.StartDate,
		EndDate:     w.EndDate,
		JobType:     w.JobType,
		CreatedAt:   w.CreatedAt.String(),
		UpdatedAt:   w.UpdatedAt.String(),
	}
}
