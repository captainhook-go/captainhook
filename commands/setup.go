package commands

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/io"
	"github.com/spf13/cobra"
)

func configurationAware(cmd *cobra.Command) {
	var configPath = "captainhook.json"
	cmd.Flags().StringP("config", "c", configPath, "path to your CaptainHook config")
}

func repositoryAware(cmd *cobra.Command) {
	var repoPath = ".git"
	cmd.Flags().StringP("repository", "r", repoPath, "path to your git repository")
}

func setUpConfig(cmd *cobra.Command) (*config.Configuration, error) {
	confPath, _ := cmd.Flags().GetString("config")
	settings := config.Settings{}
	conf, confErr := config.NewConfiguration(confPath, true, settings)
	if confErr != nil {
		return nil, confErr
	}
	return conf, nil
}

func setUpRepo(cmd *cobra.Command) (*git.Repository, error) {
	repoPath, _ := cmd.Flags().GetString("repository")
	repo, repoErr := git.NewRepository(repoPath)
	if repoErr != nil {
		return nil, repoErr
	}
	return repo, nil
}

func setUpConfigAndRepo(cmd *cobra.Command) (*config.Configuration, *git.Repository, error) {
	conf, confErr := setUpConfig(cmd)
	if confErr != nil {
		return nil, nil, confErr
	}
	repo, repoErr := setUpRepo(cmd)
	if repoErr != nil {
		return nil, nil, repoErr
	}
	return conf, repo, nil
}

func setupGlobalFlags(cmd *cobra.Command) int {
	verbosity := io.NORMAL
	noColor, _ := cmd.Flags().GetBool("no-color")
	if noColor {
		io.DeactivateColors()
	}
	quiet, _ := cmd.Flags().GetBool("quiet")
	if quiet {
		verbosity = io.QUIET
	}
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		verbosity = io.VERBOSE
	}
	debug, _ := cmd.Flags().GetBool("debug")
	if debug {
		verbosity = io.DEBUG
	}

	return verbosity
}

func mapArgs(names []string, args []string) map[string]string {
	m := make(map[string]string)
	for index, name := range names {
		if len(args) > index {
			m[name] = args[index]
		}
	}
	return m
}
