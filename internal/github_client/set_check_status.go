package githubclient

import (
	"context"

	"github.com/google/go-github/v32/github"
)

func (client *Client) SetCheckStatus(owner, repo, status, jobName, description, jobURL, commit string) error {
	ctx := context.Background()

	options := github.RepoStatus{
		State:       &status,
		Description: &description,
		Context:     &jobName,
		TargetURL:   &jobURL,
	}

	_, _, err := client.ghClient.Repositories.CreateStatus(ctx, owner, repo, commit, &options)

	return err
}
