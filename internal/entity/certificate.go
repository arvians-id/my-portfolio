package entity

import "gorm.io/gorm"

type Certificate struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Organization   string  `json:"organization"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	CredentialID   *string `json:"credential_id,omitempty"`
	Image          *string `json:"image,omitempty"`
}

func (c *Certificate) BeforeCreate(tx *gorm.DB) error {
	if c.ExpirationDate == nil || *c.ExpirationDate == "" {
		c.ExpirationDate = nil
	}

	if c.CredentialID == nil || *c.CredentialID == "" {
		c.CredentialID = nil
	}

	if c.Image == nil || *c.Image == "" {
		c.Image = nil
	}

	return nil
}
