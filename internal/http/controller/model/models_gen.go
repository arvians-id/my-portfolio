// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CategorySkill struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Skills    []*Skill `json:"skills,omitempty"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Certificate struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Organization   string  `json:"organization"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	CredentialID   *string `json:"credential_id,omitempty"`
	ImageURL       *string `json:"image_url,omitempty"`
}

type Contact struct {
	ID       int64   `json:"id"`
	Platform string  `json:"platform"`
	URL      string  `json:"url"`
	Icon     *string `json:"icon,omitempty"`
}

type CreateCategorySkillRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateCertificateRequest struct {
	Name           string  `json:"name"`
	Organization   string  `json:"organization"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	CredentialID   *string `json:"credential_id,omitempty"`
	ImageURL       *string `json:"image_url,omitempty"`
}

type CreateContactRequest struct {
	Platform string  `json:"platform"`
	URL      string  `json:"url"`
	Icon     *string `json:"icon,omitempty"`
}

type CreateEducationRequest struct {
	Institution  string  `json:"institution"`
	Degree       string  `json:"degree"`
	FieldOfStudy string  `json:"field_of_study"`
	Grade        float64 `json:"grade"`
	Description  *string `json:"description,omitempty"`
	StartDate    string  `json:"start_date"`
	EndDate      *string `json:"end_date,omitempty"`
}

type CreateProjectRequest struct {
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
	URL         *string `json:"url,omitempty"`
	IsFeatured  *bool   `json:"is_featured,omitempty"`
	Date        string  `json:"date"`
	WorkingType string  `json:"working_type"`
	Skills      []int64 `json:"skills"`
}

type CreateSkillRequest struct {
	CategorySkillID int64   `json:"category_skill_id"`
	Name            string  `json:"name"`
	Icon            *string `json:"icon,omitempty"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Pronouns string `json:"pronouns"`
	Country  string `json:"country"`
	JobTitle string `json:"job_title"`
}

type CreateWorkExperienceRequest struct {
	Role        string  `json:"role"`
	Company     string  `json:"company"`
	Description *string `json:"description,omitempty"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
	JobType     string  `json:"job_type"`
	Skills      []int64 `json:"skills"`
}

type Education struct {
	ID           int64   `json:"id"`
	Institution  string  `json:"institution"`
	Degree       string  `json:"degree"`
	FieldOfStudy string  `json:"field_of_study"`
	Grade        float64 `json:"grade"`
	Description  *string `json:"description,omitempty"`
	StartDate    string  `json:"start_date"`
	EndDate      *string `json:"end_date,omitempty"`
}

type Project struct {
	ID          int64    `json:"id"`
	Category    string   `json:"category"`
	Title       string   `json:"title"`
	Description *string  `json:"description,omitempty"`
	Image       *string  `json:"image,omitempty"`
	URL         *string  `json:"url,omitempty"`
	IsFeatured  *bool    `json:"is_featured,omitempty"`
	Date        string   `json:"date"`
	WorkingType string   `json:"working_type"`
	Skills      []*Skill `json:"skills,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type Skill struct {
	ID              int64          `json:"id"`
	CategorySkillID int64          `json:"category_skill_id"`
	CategorySkill   *CategorySkill `json:"category_skill"`
	Name            string         `json:"name"`
	Icon            *string        `json:"icon,omitempty"`
}

type UpdateCategorySkillRequest struct {
	ID   int64   `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdateCertificateRequest struct {
	ID             int64   `json:"id"`
	Name           *string `json:"name,omitempty"`
	Organization   *string `json:"organization,omitempty"`
	IssueDate      *string `json:"issue_date,omitempty"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	CredentialID   *string `json:"credential_id,omitempty"`
	ImageURL       *string `json:"image_url,omitempty"`
}

type UpdateContactRequest struct {
	ID       int64   `json:"id"`
	Platform *string `json:"platform,omitempty"`
	URL      *string `json:"url,omitempty"`
	Icon     *string `json:"icon,omitempty"`
}

type UpdateEducationRequest struct {
	ID           int64    `json:"id"`
	Institution  *string  `json:"institution,omitempty"`
	Degree       *string  `json:"degree,omitempty"`
	FieldOfStudy *string  `json:"field_of_study,omitempty"`
	Grade        *float64 `json:"grade,omitempty"`
	Description  *string  `json:"description,omitempty"`
	StartDate    *string  `json:"start_date,omitempty"`
	EndDate      *string  `json:"end_date,omitempty"`
}

type UpdateProjectRequest struct {
	ID          int64   `json:"id"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
	URL         *string `json:"url,omitempty"`
	IsFeatured  *bool   `json:"is_featured,omitempty"`
	Date        string  `json:"date"`
	WorkingType string  `json:"working_type"`
}

type UpdateSkillRequest struct {
	ID              int64   `json:"id"`
	CategorySkillID int64   `json:"category_skill_id"`
	Name            string  `json:"name"`
	Icon            *string `json:"icon,omitempty"`
}

type UpdateUserRequest struct {
	ID       int64   `json:"id"`
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	Bio      *string `json:"bio,omitempty"`
	Pronouns *string `json:"pronouns,omitempty"`
	Country  *string `json:"country,omitempty"`
	JobTitle *string `json:"job_title,omitempty"`
	Image    *string `json:"image,omitempty"`
}

type UpdateWorkExperienceRequest struct {
	ID          int64   `json:"id"`
	Role        *string `json:"role,omitempty"`
	Company     *string `json:"company,omitempty"`
	Description *string `json:"description,omitempty"`
	StartDate   *string `json:"start_date,omitempty"`
	EndDate     *string `json:"end_date,omitempty"`
	JobType     *string `json:"job_type,omitempty"`
}

type User struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Bio       *string `json:"bio,omitempty"`
	Pronouns  string  `json:"pronouns"`
	Country   string  `json:"country"`
	JobTitle  string  `json:"job_title"`
	Image     *string `json:"image,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type WorkExperience struct {
	ID          int64    `json:"id"`
	Role        string   `json:"role"`
	Company     string   `json:"company"`
	Description *string  `json:"description,omitempty"`
	StartDate   string   `json:"start_date"`
	EndDate     *string  `json:"end_date,omitempty"`
	JobType     string   `json:"job_type"`
	Skills      []*Skill `json:"skills,omitempty"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
