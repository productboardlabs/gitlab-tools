package githubclient

import (
	"context"

	"github.com/shurcooL/githubv4"
)

var latestCommitQuery struct {
	Repository struct {
		Ref struct {
			Target struct {
				OID string
			}
		} `graphql:"ref(qualifiedName: $reference)"`
	} `graphql:"repository(owner: $owner, name: $repo)"`
}

func (client *Client) LatestCommitOnBranch(owner, repo, reference, commit string) (bool, error) {
	variables := map[string]interface{}{
		"owner":     githubv4.String(owner),
		"repo":      githubv4.String(repo),
		"reference": githubv4.String(reference),
	}

	err := client.ghClient.Query(context.Background(), &latestCommitQuery, variables)

	if err != nil {
		return false, err
	}

	return latestCommitQuery.Repository.Ref.Target.OID == commit, nil
}
