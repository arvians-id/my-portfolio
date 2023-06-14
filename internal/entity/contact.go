package entity

import "gorm.io/gorm"

type Contact struct {
	ID       int64   `json:"id"`
	Platform string  `json:"platform"`
	URL      string  `json:"url"`
	Icon     *string `json:"icon,omitempty"`
}

func (c *Contact) BeforeCreate(tx *gorm.DB) error {
	if c.Icon == nil || *c.Icon == "" {
		c.Icon = nil
	}

	return nil
}
