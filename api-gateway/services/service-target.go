package services

import (
	"encoding/json"
	"net/url"
)

type ServiceTarget struct {
	Url       *url.URL
	Weight    int
	IsWorking bool
}

func (t *ServiceTarget) GetUrl() *url.URL {
	return t.Url
}

func (t *ServiceTarget) GetWeight() int {
	return t.Weight
}

func (t *ServiceTarget) GetIsWorking() bool {
	return t.IsWorking
}

func (t *ServiceTarget) SetIsWorking(v bool) {
	t.IsWorking = v
}

func (t *ServiceTarget) UnmarshalJSON(data []byte) error {
	var temp struct {
		Url    string `json:"url"`
		Weight int    `json:"weight"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	var err error
	t.Url, err = url.Parse(temp.Url)
	if err != nil {
		return err
	}
	t.Weight = temp.Weight
	t.IsWorking = true

	return nil
}
