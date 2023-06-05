package entity

import (
	"github.com/arvians-id/go-portfolio/internal/http/controller/model"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       *string   `json:"bio,omitempty"`
	Pronouns  string    `json:"pronouns"`
	Country   string    `json:"country"`
	JobTitle  string    `json:"job_title"`
	Image     *string   `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) ToModel() *model.User {
	return &model.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Bio:       u.Bio,
		Pronouns:  u.Pronouns,
		Country:   u.Country,
		JobTitle:  u.JobTitle,
		Image:     u.Image,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}
