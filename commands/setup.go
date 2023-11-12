package commands

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/info"
	"github.com/spf13/cobra"
)

func configurationAware(cmd *cobra.Command) {
	var configPath = ""
	cmd.Flags().StringP("configuration", "c", configPath, "path to your CaptainHook config")
}

func repositoryAware(cmd *cobra.Command) {
	var repoPath = ""
	cmd.Flags().StringP("repository", "r", repoPath, "path to your git repository")
}

func setUpConfig(cmd *cobra.Command) (*configuration.Configuration, error) {
	noColor, _ := cmd.Flags().GetBool("no-color")
	repoPath := ""
	confPath := info.CONFIG

	repoOption, _ := cmd.Flags().GetString("repository")
	if len(repoOption) > 0 {
		repoPath = repoOption
	}
	confOption, _ := cmd.Flags().GetString("config")
	if len(confOption) > 0 {
		confPath = confOption
	}
	settings := &configuration.AppSettings{
		AnsiColors:   !noColor,
		GitDirectory: repoPath,
		Verbosity:    getVerbosity(cmd),
	}

	conf, confErr := configuration.NewConfiguration(confPath, true, settings)
	if confErr != nil {
		return nil, confErr
	}
	return conf, nil
}

func getVerbosity(cmd *cobra.Command) string {
	verbosity := "normal"
	quiet, _ := cmd.Flags().GetBool("quiet")
	if quiet {
		verbosity = "quiet"
	}
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		verbosity = "verbose"
	}
	debug, _ := cmd.Flags().GetBool("debug")
	if debug {
		verbosity = "debug"
	}
	return verbosity
}

func mapArgs(names []string, args []string, cmd string) map[string]string {
	m := make(map[string]string)
	for index, name := range names {
		if len(args) > index {
			m[name] = args[index]
		}
	}
	m["command"] = cmd
	return m
}
