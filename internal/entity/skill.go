package entity

import "github.com/arvians-id/go-portfolio/internal/http/controller/model"

type Skill struct {
	ID              int64          `json:"id"`
	CategorySkillID int64          `json:"category_skill_id"`
	CategorySkill   *CategorySkill `json:"category_skill"`
	Name            string         `json:"name"`
	Icon            *string        `json:"icon,omitempty"`
}

func (s *Skill) ToModel() *model.Skill {
	return &model.Skill{
		ID:              s.ID,
		CategorySkillID: s.CategorySkillID,
		Name:            s.Name,
		Icon:            s.Icon,
	}
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
