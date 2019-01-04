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
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SignupUser is
type SignupUser struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Password        string `json:"password,omitempty"`
	PasswordConfirm string `json:"password_confirm,omitempty"`
}

// Profile is
type Profile struct {
	ID             string         `json:"id"`
	TenantID       string         `json:"tenant_id"`
	Email          string         `json:"email"`
	Name           string         `json:"name"`
	Username       string         `json:"username"`
	Organization   string         `json:"organization"`
	ProfilePicture string         `json:"profile_picture"`
	UserMetadata   interface{}    `json:"user_metadata"`
	AppMetadata    interface{}    `json:"app_metadata"`
	Providers      []ProviderInfo `json:"providers"`
	Groups         []Group        `json:"groups"`
	Applications   []Client       `json:"applications"`
	Suspended      bool           `json:"suspended"`
	LastLogin      *LoginInfo     `json:"last_login"`
}

// EmailVerification is
type EmailVerification struct {
	ID string `json:"id,omitempty"`
}
