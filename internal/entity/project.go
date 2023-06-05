package entity

import (
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"time"
)

type Project struct {
	ID          int64     `json:"id"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	Image       *string   `json:"image,omitempty"`
	URL         *string   `json:"url,omitempty"`
	IsFeatured  *bool     `json:"is_featured,omitempty"`
	Date        string    `json:"date"`
	WorkingType string    `json:"working_type"`
	Skills      []*Skill  `json:"skills,omitempty" gorm:"many2many:project_skill;"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (p *Project) ToModel() *model.Project {
	return &model.Project{
		ID:          p.ID,
		Category:    p.Category,
		Title:       p.Title,
		Description: p.Description,
		Image:       p.Image,
		URL:         p.URL,
		IsFeatured:  p.IsFeatured,
		Date:        p.Date,
		WorkingType: p.WorkingType,
		CreatedAt:   p.CreatedAt.String(),
		UpdatedAt:   p.UpdatedAt.String(),
	}
}
