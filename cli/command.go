package cli

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/git"
	"github.com/spf13/cobra"
)

func ConfigurationAware(cmd *cobra.Command) {
	var configPath = "captainhook.json"
	cmd.Flags().StringP("config", "c", configPath, "path to your CaptainHook config")
}

func RepositoryAware(cmd *cobra.Command) {
	var repoPath = ".git"
	cmd.Flags().StringP("repository", "r", repoPath, "path to your git repository")
}

func SetUpConfig(confPath string) (*config.Configuration, error) {
	settings := config.Settings{}
	conf, confErr := config.NewConfiguration(confPath, true, settings)
	if confErr != nil {
		return nil, confErr
	}
	return conf, nil
}

func SetUpRepo(repoPath string) (*git.Repository, error) {
	repo, repoErr := git.NewRepository(repoPath)
	if repoErr != nil {
		return nil, repoErr
	}
	return repo, nil
}

func SetUpConfigAndRepo(confPath string, repoPath string) (*config.Configuration, *git.Repository, error) {
	conf, confErr := SetUpConfig(confPath)
	if confErr != nil {
		return nil, nil, confErr
	}
	repo, repoErr := SetUpRepo(repoPath)
	if repoErr != nil {
		return nil, nil, repoErr
	}
	return conf, repo, nil
}

func MapArgs(names []string, args []string) map[string]string {
	m := make(map[string]string)
	for index, name := range names {
		if len(args) > index {
			m[name] = args[index]
		}
	}
	return m
}
