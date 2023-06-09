package entity

import "time"

type Skill struct {
	ID              int64          `json:"id"`
	CategorySkillID int64          `json:"category_skill_id"`
	CategorySkill   *CategorySkill `json:"category_skill"`
	Name            string         `json:"name"`
	Icon            *string        `json:"icon,omitempty"`
}
type SkillBelongsTo struct {
	ID               int64          `json:"id"`
	CategorySkillID  int64          `json:"category_skill_id"`
	CategorySkill    *CategorySkill `json:"category_skill"`
	Name             string         `json:"name"`
	Icon             *string        `json:"icon,omitempty"`
	ProjectID        int64          `json:"project_id"`
	WorkExperienceID int64          `json:"work_experience_id"`
}

type CategorySkill struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Skills    []*Skill  `json:"skills,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
