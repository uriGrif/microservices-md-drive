package loadbalancing

import (
	"net/http"
	"net/url"

	"github.com/jmcvetta/randutil"
)

type Default struct {
	targets []Target
}

func (lb *Default) RegisterTarget(target Target) {
	lb.targets = append(lb.targets, target)
}

func (lb *Default) GetNextEndpoint(req *http.Request) (*url.URL, error) {

	choices := make([]randutil.Choice, 0, len(lb.targets))
	for _, t := range lb.targets {
		if !t.GetIsWorking() {
			continue
		}
		choices = append(choices, randutil.Choice{Item: t.GetUrl(), Weight: t.GetWeight()})
	}
	c, err := randutil.WeightedChoice(choices)
	if err != nil {
		return &url.URL{}, err
	}

	return c.Item.(*url.URL), nil
}

func CreateDefaultLB(targets []Target) LoadBalancer {
	lb := Default{}
	for _, t := range targets {
		lb.RegisterTarget(t)
	}
	return &lb
}
