package services

import (
	"api-gateway/middleware/auth"
	"encoding/json"
)

type Route struct {
	Endpoint      string             `json:"endpoint"`
	Authenticator auth.Authenticator `json:"authenticator,omitempty"`
}

func (r *Route) UnmarshalJSON(data []byte) error {
	var temp struct {
		Endpoint      string `json:"endpoint"`
		Authenticator string `json:"authenticator,omitempty"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	r.Endpoint = temp.Endpoint

	switch temp.Authenticator {
	case "auth0":
		r.Authenticator = auth.GetAuth0Authenticator()
	default:
		r.Authenticator = nil
	}

	return nil
}
