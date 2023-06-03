package entity

import "time"

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
