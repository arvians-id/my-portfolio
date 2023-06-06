package entity

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
