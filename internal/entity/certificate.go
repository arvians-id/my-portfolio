package entity

import "github.com/arvians-id/go-portfolio/internal/http/controller/model"

type Certificate struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Organization   string  `json:"organization"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	CredentialID   *string `json:"credential_id,omitempty"`
	ImageURL       *string `json:"image_url,omitempty"`
}

func (c *Certificate) ToModel() *model.Certificate {
	return &model.Certificate{
		ID:             c.ID,
		Name:           c.Name,
		Organization:   c.Organization,
		IssueDate:      c.IssueDate,
		ExpirationDate: c.ExpirationDate,
		CredentialID:   c.CredentialID,
		ImageURL:       c.ImageURL,
	}
}
