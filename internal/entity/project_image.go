package entity

type ProjectImage struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"project_id"`
	Image     string `json:"image"`
}
