package loadbalancing

import (
	"net/http"
	"net/url"
)

type LoadBalancer interface {
	RegisterTarget(Target)
	GetNextEndpoint(*http.Request) (*url.URL, error)
}

// had to create this to avoid cycle import with services package
type Target interface {
	GetUrl() *url.URL
	GetWeight() int
	GetIsWorking() bool
	SetIsWorking(bool)
}
