package model

// User is
type User struct {
	ID             string `json:"id,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	Username       string `json:"username,omitempty"`
	Email          string `json:"email,omitempty" validate:"email"`
	Name           string `json:"name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

// Group is
type Group struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// SignupUser is
type SignupUser struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"password_confirm,omitempty"`
}

// Profile is
type Profile struct {
	ID             string         `json:"id,omitempty"`
	TenantID       string         `json:"tenant_id,omitempty"`
	Email          string         `json:"email,omitempty"`
	Name           string         `json:"name,omitempty"`
	Username       string         `json:"username,omitempty"`
	Organization   string         `json:"organization,omitempty"`
	ProfilePicture string         `json:"profile_picture,omitempty"`
	UserMetadata   interface{}    `json:"user_metadata"`
	AppMetadata    interface{}    `json:"app_metadata"`
	Providers      []ProviderInfo `json:"providers"`
	Groups         []Group        `json:"groups,omitempty"`
	Applications   []Client       `json:"applications,omitempty"`
	Suspended      bool           `json:"suspended"`
	LastLogin      *LoginInfo     `json:"last_login,omitempty"`
}

// EmailVerification is
type EmailVerification struct {
	ID string `json:"id,omitempty"`
}
