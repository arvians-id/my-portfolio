package entity

import "time"

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
