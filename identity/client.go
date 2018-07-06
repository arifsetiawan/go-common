package identity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arifsetiawan/go-common/model"
	"github.com/parnurzeal/gorequest"
)

// SDK is
type SDK interface {
	GetCollectionClientsWithID(tenant string, collectionID string) (*model.Client, error)
	CheckWarden(token string) error
	Me(token string) error
}

// SDKClient is
type SDKClient struct {
	IdentityAPIURL string
	IdentityAPIKey string
}

// ClientsData is
type ClientsData struct {
	Clients []model.Client `json:"data,omitempty"`
}

// ClientData is
type ClientData struct {
	Client model.Client `json:"data"`
}

// GetCollectionClientsWithID is
func (s *SDKClient) GetCollectionClientsWithID(tenant string, collectionID string) (*model.Client, error) {

	url := s.IdentityAPIURL + "/" + tenant + "/collections/clients/" + collectionID
	resp, body, errs := gorequest.New().Get(url).
		Set("API-Key", s.IdentityAPIKey).
		End()
	if errs != nil {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get connection status is %d with message %s", resp.StatusCode, body)
	}

	clientData := ClientData{}
	buff := bytes.NewBufferString(body)
	decoder := json.NewDecoder(buff)

	if err := decoder.Decode(&clientData); err != nil {
		if terr, ok := err.(*json.UnmarshalTypeError); ok {
			return nil, fmt.Errorf("failed to unmarshal field %s", terr.Field)
		}

		return nil, err
	}

	return &clientData.Client, nil
}

// CheckWarden is
func (s *SDKClient) CheckWarden(token string) error {
	return error("Not implemented")
}

// Me is
func (s *SDKClient) Me(token string) error {
	return error("Not implemented")
}
