package entity

type Certificate struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Organization   string  `json:"organization"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	CredentialID   *string `json:"credential_id,omitempty"`
	Image          *string `json:"image,omitempty"`
}
