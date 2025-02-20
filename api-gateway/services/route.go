package services

type Route struct {
	Endpoint      string  `json:"endpoint"`
	Authenticator *string `json:"authenticator,omitempty"`
}
