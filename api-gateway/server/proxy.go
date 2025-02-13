package server

import (
	loadbalancing "api-gateway/load-balancing"
	"log"
	"net/http/httputil"
)

func CreateProxy(lb loadbalancing.LoadBalancer) httputil.ReverseProxy {
	if lb == nil {
		log.Fatal("a load balancer must be provided for a proxy")
	}
	proxy := httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			target, err := lb.GetNextEndpoint(r.In)
			if err != nil {
				log.Fatal(err)
			}
			r.SetURL(target)
			r.Out.Host = r.In.Host
		},
	}

	return proxy
}
