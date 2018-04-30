package model

// Client is
type Client struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"client_name,omitempty"`
	Secret        string   `json:"client_secret,omitempty"`
	RedirectUris  []string `json:"redirect_uris,omitempty"`
	GrantTypes    []string `json:"grant_types,omitempty"`
	ResponseTypes []string `json:"response_types,omitempty"`
	Scope         string   `json:"scope,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	ClientURI     string   `json:"client_uri,omitempty"`
	LogoURI       string   `json:"logo_uri,omitempty"`
}