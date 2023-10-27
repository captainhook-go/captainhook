package cli

import "github.com/spf13/cobra"

func ConfigurationAware(cmd *cobra.Command) {
	var configPath = "captainhook.json"
	cmd.Flags().StringP("config", "c", configPath, "path to your CaptainHook config")
}

func RepositoryAware(cmd *cobra.Command) {
	var repoPath = ".git"
	cmd.Flags().StringP("repository", "r", repoPath, "path to your git repository")
}
