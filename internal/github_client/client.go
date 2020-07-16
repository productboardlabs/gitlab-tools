package githubclient

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Client contains the set up github client and allowed queries to communicate with Github.
type Client struct {
	ghClient *githubv4.Client
}

// New bootstraps the Client struct
func New(token string) *Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	return &Client{
		ghClient: client,
	}
}
