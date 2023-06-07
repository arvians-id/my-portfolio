package entity

import (
	"time"
)

type ProjectImages struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"project_id"`
	Image     string `json:"image"`
}

type Project struct {
	ID          int64            `json:"id"`
	Category    string           `json:"category"`
	Title       string           `json:"title"`
	Description *string          `json:"description,omitempty"`
	URL         *string          `json:"url,omitempty"`
	IsFeatured  *bool            `json:"is_featured,omitempty"`
	Date        string           `json:"date"`
	WorkingType string           `json:"working_type"`
	Skills      []*Skill         `json:"skills,omitempty" gorm:"many2many:project_skill;"`
	Images      []*ProjectImages `json:"images,omitempty"`
	CreatedAt   time.Time        `json:"created_at,omitempty"`
	UpdatedAt   time.Time        `json:"updated_at,omitempty"`
}
