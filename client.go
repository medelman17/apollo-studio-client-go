package apollostudio

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

const StudioURL string = "https://graphql.api.apollographql.com/api/graphql"
const ClientName string = "apollo-studio-client-go"

type StudioClient struct {
	host    string
	fetcher *http.Client
}

type StudioTransport struct {
	apikey string
	rt     http.RoundTripper
}

func (t *StudioTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("x-api-key", t.apikey)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("apollographql-client-name", ClientName)
	return t.rt.RoundTrip(r)
}

func NewClient(host, apikey *string) (*StudioClient, error) {

	fetcher := http.DefaultClient

	transporter := StudioTransport{
		apikey: *apikey,
		rt:     fetcher.Transport,
	}

	fetcher.Transport = &transporter

	client := StudioClient{
		fetcher: fetcher,
		host:    StudioURL,
	}
	if host != nil {
		client.host = *host
	}

	return &client, nil
}

func (c *StudioClient) memberships() (*GetCallerMembershipsResponse, error) {
	ctx := context.Background()
	client := graphql.NewClient(c.host, c.fetcher)
	resp, err := GetCallerMemberships(ctx, client)
	return resp, err
}

type CreateSupergraphInput struct {
	AccountId              string                 `json:"accountId"`
	NewServiceId           string                 `json:"newServiceId"`
	Name                   string                 `json:"name"`
	OnboardingArchitecture OnboardingArchitecture `json:"onboardingArchitecture"`
}

func (c *StudioClient) createSupergraph(input CreateSupergraphInput) (*CreateServiceResponse, error) {
	ctx := context.Background()
	client := graphql.NewClient(c.host, c.fetcher)
	resp, err := CreateService(ctx, client, input.AccountId, input.NewServiceId, input.Name, input.OnboardingArchitecture)
	return resp, err
}

func (c *StudioClient) deleteSupergraph(id string) (*DeleteServiceResponse, error) {
	ctx := context.Background()
	client := graphql.NewClient(c.host, c.fetcher)
	resp, err := DeleteService(ctx, client, id)
	return resp, err
}

func (c *StudioClient) createSupergraphKey(serviceId string, name string, role UserPermission) (*NewKeyResponse, error) {
	ctx := context.Background()
	client := graphql.NewClient(c.host, c.fetcher)
	resp, err := NewKey(ctx, client, serviceId, name, role)
	return resp, err
}

func (c *StudioClient) deleteSupergraphKey(serviceId string, keyId string) (*RemoveKeyResponse, error) {
	ctx := context.Background()
	client := graphql.NewClient(c.host, c.fetcher)
	resp, err := RemoveKey(ctx, client, serviceId, keyId)
	return resp, err
}
