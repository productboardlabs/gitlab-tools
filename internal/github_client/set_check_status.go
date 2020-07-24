package githubclient

import (
	"context"

	"github.com/google/go-github/v32/github"
)

func (client *Client) SetCheckStatus(owner, repo, status, jobName, commit string) error {
	ctx := context.Background()

	options := github.CreateCheckRunOptions{
		Name:    jobName,
		HeadSHA: commit,
		Status:  &status,
	}

	_, _, err := client.ghClient.Checks.CreateCheckRun(ctx, owner, repo, options)

	return err
}
