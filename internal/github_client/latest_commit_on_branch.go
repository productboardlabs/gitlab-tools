package githubclient

import (
	"context"
	"fmt"
)

func (client *Client) LatestCommitOnBranch(owner, repo, reference, commit string) (bool, error) {
	ctx := context.Background()

	ref, _, err := client.ghClient.Git.GetRef(ctx, owner, repo, fmt.Sprintf("heads/%s", reference))

	if err != nil {
		return false, err
	}

	return *ref.Object.SHA == commit, nil
}
