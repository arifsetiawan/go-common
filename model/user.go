package model

// User is
type User struct {
	ID             string `json:"id,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	Username       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty"`
	Name           string `json:"name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

// Profile is
type Profile struct {
	ID             string         `json:"id,omitempty"`
	TenantID       string         `json:"tenant_id,omitempty"`
	Email          string         `json:"email,omitempty"`
	Name           string         `json:"name,omitempty"`
	Username       string         `json:"username,omitempty"`
	ProfilePicture string         `json:"profile_picture,omitempty"`
	UserMetadata   interface{}    `json:"user_metadata"`
	AppMetadata    interface{}    `json:"app_metadata"`
	Providers      []ProviderInfo `json:"providers"`
	GroupIDs       []string       `json:"group_ids,omitempty"`
	Suspended      bool           `json:"suspended"`
	LastLogin      *LoginInfo     `json:"last_login,omitempty"`
}
