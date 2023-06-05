package entity

import "github.com/arvians-id/go-portfolio/internal/http/controller/model"

type Contact struct {
	ID       int64   `json:"id"`
	Platform string  `json:"platform"`
	URL      string  `json:"url"`
	Icon     *string `json:"icon,omitempty"`
}

func (c *Contact) ToModel() *model.Contact {
	return &model.Contact{
		ID:       c.ID,
		Platform: c.Platform,
		URL:      c.URL,
		Icon:     c.Icon,
	}
}
