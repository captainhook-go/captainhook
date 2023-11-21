package commands

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/info"
	"github.com/spf13/cobra"
)

func configurationAware(cmd *cobra.Command) {
	var configPath = info.CONFIG
	cmd.Flags().StringP("configuration", "c", configPath, "path to your CaptainHook config")
}

func repositoryAware(cmd *cobra.Command) {
	var repoPath = ""
	cmd.Flags().StringP("repository", "r", repoPath, "path to your git repository")
}

func setUpConfig(cmd *cobra.Command) (*configuration.Configuration, error) {
	nullableSettings := &configuration.JsonAppSettings{}

	detectColor(cmd, nullableSettings)
	detectGitDir(cmd, nullableSettings)
	detectVerbosity(cmd, nullableSettings)

	confPath := info.CONFIG
	confOption, _ := cmd.Flags().GetString("configuration")
	if len(confOption) > 0 {
		confPath = confOption
	}

	factory := configuration.NewFactory()
	conf, confErr := factory.CreateConfig(confPath, nullableSettings)
	if confErr != nil {
		return nil, confErr
	}
	return conf, nil
}

func detectColor(cmd *cobra.Command, settings *configuration.JsonAppSettings) {
	noColor := getNoColor(cmd)
	if noColor {
		falsePointer := false
		settings.AnsiColors = &falsePointer
	}
}

func getNoColor(cmd *cobra.Command) bool {
	noColor, _ := cmd.Flags().GetBool("no-color")
	return noColor
}

func detectGitDir(cmd *cobra.Command, settings *configuration.JsonAppSettings) {
	repoOption := getGitDit(cmd)
	if len(repoOption) > 0 {
		settings.GitDirectory = &repoOption
	}
}

func getGitDit(cmd *cobra.Command) string {
	repoPath := ""
	repoOption, _ := cmd.Flags().GetString("repository")
	if len(repoOption) > 0 {
		repoPath = repoOption
	}
	return repoPath
}

func detectVerbosity(cmd *cobra.Command, settings *configuration.JsonAppSettings) {
	v := getVerbosity(cmd)
	if v != "normal" {
		settings.Verbosity = &v
	}
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
	m[info.ArgCommand] = cmd
	return m
}
