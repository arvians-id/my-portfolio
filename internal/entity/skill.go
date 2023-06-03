package entity

type Skill struct {
	ID              int64          `json:"id"`
	CategorySkillID int64          `json:"category_skill_id"`
	CategorySkill   *CategorySkill `json:"category_skill"`
	Name            string         `json:"name"`
	Icon            *string        `json:"icon,omitempty"`
}
