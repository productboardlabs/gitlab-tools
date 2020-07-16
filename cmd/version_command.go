package cmd

import (
	"github.com/spf13/cobra"
)

// version is a global variable passed during build time
var version string

// commit is a global variable passed during build time. Should be used if version is not available.
var commit string

// date is a global variable passed during build time
var date string

func (runner *Runner) versionCommand() *cobra.Command {
	// Version returns undefined if not on a tag. This needs to reset it.
	if version == "undefined" {
		version = ""
	}

	if version == "" && commit != "" {
		version = commit
	}
	if version == "" && commit == "" {
		version = "development"
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gitlab-tools",
		Long:  `All software has versions. This is of gitlab-tools`,
		RunE: func(cmd *cobra.Command, args []string) error {
			runner.logger.Printf("gitlab-tools version: %s\t Built on: %s", version, date)
			return nil
		},
	}

	return versionCmd
}
