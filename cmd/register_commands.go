package cmd

func (runner *Runner) registerCommands() {
	runner.RootCmd.AddCommand(runner.versionCommand())
	runner.RootCmd.AddCommand(runner.isLatestCommitCommand())
}
