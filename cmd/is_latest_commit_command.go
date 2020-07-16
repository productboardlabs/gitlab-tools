package cmd

import (
	"errors"
	"fmt"
	"strings"

	githubclient "github.com/productboardlabs/gitlab-tools/internal/github_client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*isLatestCommitCommand compares the current commit hash to the one provided by Github API and throws an error if the commit is not latest.
Requires:
- Github token
- Github repo in format owner/repo
- Commit hash
- Branch
*/
func (runner *Runner) isLatestCommitCommand() *cobra.Command {

	var command = &cobra.Command{
		Use:   "is-latest-commit",
		Short: "Finds if the commit is latest on the github branch",
		Long:  `In order to find whether we can run expensive parts of our CI we would like to know whether the current commit is the latest on the branch`,
		RunE: func(cmd *cobra.Command, args []string) error {
			token := viper.GetString("GITHUB_TOKEN")

			if token == "" {
				return errors.New("missing github token")
			}

			github := githubclient.New(token)

			repository := strings.Split(viper.GetString("GITHUB_REPOSITORY"), "/")

			commitSha := viper.GetString("CI_COMMIT_SHA")

			ref := viper.GetString("CI_COMMIT_REF_NAME")

			if repository[0] == "" {
				return errors.New("missing repository")
			}

			if len(repository) < 2 {
				return errors.New("repository is missing repo name")
			}

			if commitSha == "" {
				return errors.New("commitSha missing")
			}

			if ref == "" {
				return errors.New("reference is missing")
			}

			ok, err := github.LatestCommitOnBranch(
				repository[0],
				repository[1],
				ref,
				commitSha,
			)

			if err != nil {
				return err
			}

			if ok {
				runner.logger.Print("Success! This the latest commit on branch")
				return nil
			}

			return fmt.Errorf("[FAIL] Commit: %s is not the latest on the branch %s", commitSha, ref)
		},
	}

	command.Flags().String("github-token", "", "the API token for github https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token")
	err := viper.BindPFlag("GITHUB_TOKEN", command.Flags().Lookup("github-token"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag Github token %s", err)
	}

	command.Flags().String("repository", "", "repository in format owner/repo")
	err = viper.BindPFlag("GITHUB_REPOSITORY", command.Flags().Lookup("repository"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag Github token %s", err)
	}

	command.Flags().String("reference", "", "branch name to check")
	err = viper.BindPFlag("CI_COMMIT_REF_NAME", command.Flags().Lookup("reference"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag Github token %s", err)
	}

	command.Flags().String("commit", "", "commit hash of current commit")
	err = viper.BindPFlag("CI_COMMIT_SHA", command.Flags().Lookup("commit"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag Github token %s", err)
	}

	return command
}
