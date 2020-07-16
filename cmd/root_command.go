package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Runner houses the logging and rootcommand for all
type Runner struct {
	RootCmd     *cobra.Command
	logger      *log.Logger
	debugLogger *log.Logger
}

// New creates a instance of Runner with defaults for logger and defaultLogger in case they are not provided.
func New(logger *log.Logger, debugLogger *log.Logger) *Runner {
	if logger == nil {
		logger = log.New(os.Stdout, "", 0)
	}

	if debugLogger == nil {
		// If no debug logger is supplied then just silence output
		debugLogger = log.New(ioutil.Discard, "", 0)
	}

	rootCmd := rootCommand()

	runner := &Runner{
		RootCmd:     rootCmd,
		logger:      logger,
		debugLogger: debugLogger,
	}

	runner.registerCommands()

	return runner
}

// rootCommand returns the Cobra command all other commands register to
func rootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "gitlab-tools ",
		Short: "Tooling for gitlab",
		Long:  "Tooling for gitlab",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("There is no root command. Please check gitlab-tools --help.")
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	return rootCmd
}
