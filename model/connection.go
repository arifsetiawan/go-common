package model

import "time"

// Connection is
type Connection struct {
	ID             string          `json:"id,omitempty"`
	TenantID       string          `json:"tenant_id,omitempty"`
	Name           string          `json:"name,omitempty"`
	Group          string          `json:"group,omitempty"`
	Form           string          `json:"form,omitempty"`
	Login          string          `json:"login,omitempty"`
	CollectionID   string          `json:"collection_id,omitempty"`
	ConnectionData *ConnectionData `json:"connection,omitempty"`
}

// ConnectionData is
type ConnectionData struct {
	ProviderOAuth2
}

// ProviderOAuth2 is
type ProviderOAuth2 struct {
	ClientID     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
	AuthURL      string   `json:"auth_url,omitempty"`
	TokenURL     string   `json:"token_url,omitempty"`
	RedirectURL  string   `json:"redirect_url,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
}

// ProviderInfo is extra info for user idp
type ProviderInfo struct {
	Name         string     `json:"name"`
	ConnectionID string     `json:"connection_id"`
	UserID       string     `json:"user_id"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	AccessToken  string     `json:"access_token,omitempty"`
}

// LoginInfo is
type LoginInfo struct {
	IP           string     `json:"ip,omitempty"`
	Time         *time.Time `json:"last_login,omitempty"`
	ProviderType string     `json:"provider_type,omitempty"`
}

// UserProfile is
type UserProfile struct {
	User         *User         `json:"user,omitempty"`
	ProviderInfo *ProviderInfo `json:"provider_info,omitempty"`
	LoginInfo    *LoginInfo    `json:"login_info,omitempty"`
}
