package github_client

import (
	"context"

	"github.com/selefra/selefra-provider-github/constants"

	"github.com/google/go-github/v45/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type Client struct {
	Github GithubServices
	Org    string
	Orgs   []string
}

func (c *Client) ID() string {
	return c.Org
}

func (c Client) withOrg(org string) *Client {
	return &Client{
		Github: c.Github,
		Org:    org,
		Orgs:   c.Orgs,
	}
}

func NewClients(configs Configs) ([]*Client, error) {
	var clients []*Client
	for _, provider := range configs.Providers {
		client, err := newClient(provider)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func newClient(config Config) (*Client, error) {
	if config.AccessToken == constants.Constants_0 {
		return nil, errors.New(constants.Missingpersonalaccesstokeninconfiguration)
	}
	if len(config.Orgs) == 0 {
		return nil, errors.New(constants.Noorganizationsdefinedinconfiguration)
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.AccessToken})
	tc := oauth2.NewClient(context.Background(), ts)
	c := github.NewClient(tc)
	return &Client{
		Github: GithubServices{
			Teams:         c.Teams,
			Billing:       c.Billing,
			Repositories:  c.Repositories,
			Organizations: c.Organizations,
			Issues:        c.Issues,
		},
		Orgs: config.Orgs,
	}, nil
}
