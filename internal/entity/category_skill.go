package entity

import (
	"time"
)

type CategorySkill struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Skills    []*Skill  `json:"skills,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
