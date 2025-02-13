package services

import (
	loadbalancing "api-gateway/load-balancing"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Service struct {
	Name         string `json:"name"`
	Targets      []ServiceTarget
	LoadBalancer loadbalancing.LoadBalancer `json:"loadBalancer,omitempty"`
	Routes       []Route
}

func (s *Service) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name         string          `json:"name"`
		Targets      json.RawMessage `json:"targets"`
		LoadBalancer string          `json:"loadBalancer,omitempty"`
		Routes       json.RawMessage `json:"routes"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.Name = temp.Name
	if err := json.Unmarshal(temp.Targets, &s.Targets); err != nil {
		return err
	}
	if err := json.Unmarshal(temp.Routes, &s.Routes); err != nil {
		return err
	}

	lbTargets := func() []loadbalancing.Target {
		targets := make([]loadbalancing.Target, len(s.Targets))
		for i, t := range s.Targets {
			targets[i] = &t
		}
		return targets
	}

	switch temp.LoadBalancer {
	default:
		s.LoadBalancer = loadbalancing.CreateDefaultLB(lbTargets())
	}

	return nil
}

func LoadServices() []Service {
	jsonFile, err := os.Open(os.Getenv("SERVICES_PATH"))
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var services []Service

	if err := json.Unmarshal(byteValue, &services); err != nil {
		log.Fatal(err)
	}

	return services
}
