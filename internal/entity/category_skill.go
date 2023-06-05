package entity

import (
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"time"
)

type CategorySkill struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Skills    []*Skill  `json:"skills,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (c *CategorySkill) ToModel() *model.CategorySkill {
	return &model.CategorySkill{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}
