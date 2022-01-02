package whois

type WhoisClient struct{}

func (c WhoisClient) Lookup(domain string) (Response, error) {
	tld := getTldFromDomain(domain)
	adapter := getAdapterForTld(tld)
	response, err := adapter.request(domain)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func NewWhoisClient() *WhoisClient {
	return &WhoisClient{}
}
