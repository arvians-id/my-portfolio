package entity

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID          int64           `json:"id"`
	Category    string          `json:"category"`
	Title       string          `json:"title"`
	Description *string         `json:"description,omitempty"`
	URL         *string         `json:"url,omitempty"`
	IsFeatured  *bool           `json:"is_featured,omitempty"`
	Date        string          `json:"date"`
	WorkingType string          `json:"working_type"`
	Skills      []*Skill        `json:"skills,omitempty" gorm:"many2many:project_skill;"`
	Images      []*ProjectImage `json:"images,omitempty"`
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	UpdatedAt   time.Time       `json:"updated_at,omitempty"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) error {
	if p.Description == nil || *p.Description == "" {
		p.Description = nil
	}

	if p.URL == nil || *p.URL == "" {
		p.URL = nil
	}

	if p.IsFeatured == nil {
		p.IsFeatured = nil
	}

	return nil
}
