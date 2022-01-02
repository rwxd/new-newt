package whois

import (
	"strings"
	"time"
)

type Adapter interface {
	formatResponse(response string) (Response, error)
	request(domain string) (Response, error)
}

type defaultAdapter struct {
}

func (a defaultAdapter) request(domain string) (Response, error) {
	request := NewRequest(domain, "whois.iana.com", "43", time.Second*5)
	body, err := request.Query(domain)
	if err != nil {
		return Response{}, err
	}

	response, err := a.formatResponse(body)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func (a defaultAdapter) formatResponse(response string) (Response, error) {
	return Response{
		Body:      response,
		Available: false,
	}, nil
}

var DefaultAdapter = &defaultAdapter{}

type deAdapter struct {
	defaultAdapter
}

func (a deAdapter) formatResponse(response string) (Response, error) {
	available := false
	if strings.Contains(response, "Status: free") {
		available = true
	} else if strings.Contains(response, "Status: connect") {
		available = false
	}

	return Response{
		Body:      response,
		Available: available,
	}, nil
}

func (a deAdapter) request(domain string) (Response, error) {
	request := NewRequest(domain, "whois.denic.de", "43", time.Second*5)
	body, err := request.Query(domain)
	if err != nil {
		return Response{}, err
	}

	response, err := a.formatResponse(body)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

var adapters = map[string]Adapter{}

func bindAdapters(s Adapter, names ...string) {
	for _, name := range names {
		adapters[name] = s
	}
}

func getAdapterForTld(tld string) Adapter {
	if a, ok := adapters[tld]; ok {
		return a
	}
	return DefaultAdapter
}

func init() {
	bindAdapters(deAdapter{}, "de")
}
