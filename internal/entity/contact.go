package entity

type Contact struct {
	ID       int64   `json:"id"`
	Platform string  `json:"platform"`
	URL      string  `json:"url"`
	Icon     *string `json:"icon,omitempty"`
}
