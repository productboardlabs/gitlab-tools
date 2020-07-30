package cmd

import (
	"errors"
	"fmt"
	"strings"

	githubclient "github.com/productboardlabs/gitlab-tools/internal/github_client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*setStatus sets a given job
name: jobName that will show up in Github
status: queued, in_progress, or completed
*/
func (runner *Runner) setStatus() *cobra.Command {

	var command = &cobra.Command{
		Use:   "set-check-status",
		Short: "Sets status to provided in Github UI",
		Long:  `In order to show more granuality in Github UI we want to provide status breakdown per job.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			token := viper.GetString("GITHUB_TOKEN")

			if token == "" {
				return errors.New("missing github token")
			}

			github := githubclient.New(token)

			repository := strings.Split(viper.GetString("GITHUB_REPOSITORY"), "/")

			commitSha := viper.GetString("CI_COMMIT_SHA")

			status := viper.GetString("STATUS")

			description := viper.GetString("DESCRIPTION")

			jobURL := viper.GetString("CI_JOB_URL")

			exitCode, err := cmd.Flags().GetInt("exit-code")

			if err == nil {
				if exitCode == 0 {
					status = "success"
				}
				if exitCode > 0 {
					status = "failure"
				}
			}

			if repository[0] == "" {
				return errors.New("missing repository")
			}

			if len(repository) < 2 {
				return errors.New("repository is missing repo name")
			}

			if commitSha == "" {
				return errors.New("commitSha missing")
			}

			if status == "" {
				return errors.New("missing status, available are error, failure, pending, or success")
			}

			jobName := cmd.Flag("status-name").Value.String()

			if jobName == "" {
				jobName = fmt.Sprintf("%s/%s", viper.GetString("CI_JOB_STAGE"), viper.GetString("CI_JOB_NAME"))
			}

			err = github.SetCheckStatus(
				repository[0],
				repository[1],
				status,
				jobName,
				description,
				jobURL,
				commitSha,
			)

			return err
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

	command.Flags().String("status", "", "status to set in github")
	err = viper.BindPFlag("STATUS", command.Flags().Lookup("status"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag Status %s", err)
	}

	command.Flags().String("commit", "", "commit hash of current commit")
	err = viper.BindPFlag("CI_COMMIT_SHA", command.Flags().Lookup("commit"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag commit %s", err)
	}

	command.Flags().String("status-name", "", "name of status that show up in Github")

	command.Flags().String("url", "", "commit hash of current commit")
	err = viper.BindPFlag("CI_JOB_URL", command.Flags().Lookup("url"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag url %s", err)
	}

	command.Flags().String("description", "", "commit hash of current commit")
	err = viper.BindPFlag("DESCRIPTION", command.Flags().Lookup("description"))

	if err != nil {
		runner.debugLogger.Printf("failed to bindPflag description %s", err)
	}

	command.Flags().Int("exit-code", 0, "exit code of previous process")

	return command
}
