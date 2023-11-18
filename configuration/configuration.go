package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
)

type Configuration struct {
	size       int64
	path       string
	fileExists bool
	settings   *AppSettings
	hooks      map[string]*Hook
}

func NewConfiguration(path string, fileExists bool, settings *AppSettings) (*Configuration, error) {
	var err error
	c := Configuration{path: path, fileExists: fileExists, settings: NewDefaultAppSettings()}
	c.path = path
	c.fileExists = fileExists
	c.init()
	conf := &c

	if c.fileExists {
		err = c.load()
		if err != nil {
			println(err.Error())
		}
	}
	c.mergeSettings(settings, false)
	return conf, err
}

func (c *Configuration) init() {
	c.hooks = map[string]*Hook{}
	for _, hook := range info.GetValidHooks() {
		c.hooks[hook] = NewHook(hook, false)
	}
}

func (c *Configuration) IsLoadedFromFile() bool {
	return c.fileExists
}

func (c *Configuration) IsHookEnabled(hook string) bool {
	return c.HookConfig(hook).IsEnabled()
}

func (c *Configuration) Path() string {
	return c.path
}

func (c *Configuration) RunPath() string {
	return c.settings.RunPath
}

func (c *Configuration) CustomSettings() map[string]string {
	return c.settings.Custom
}

func (c *Configuration) GitDirectory() string {
	gitDir := c.settings.GitDirectory
	if len(gitDir) < 1 {
		gitDir = ".git"
	}
	return gitDir
}

func (c *Configuration) AnsiColors() bool {
	return c.settings.AnsiColors
}

func (c *Configuration) Verbosity() int {
	return MapVerbosity(c.settings.Verbosity)
}

func (c *Configuration) IsFailureAllowed() bool {
	return c.settings.AllowFailure
}

func (c *Configuration) FailOnFirstError() bool {
	return c.settings.FailOnFirstError
}

func (c *Configuration) RunAsync() bool {
	return c.settings.RunAsync
}

func (c *Configuration) HookConfig(hook string) *Hook {
	return c.hooks[hook]
}

func (c *Configuration) load() error {
	jsonBytes, readError := c.readConfigFile()
	if readError != nil {
		return readError
	}
	configurationJson, decodeErr := c.decodeConfigJson(jsonBytes)
	if decodeErr != nil {
		return fmt.Errorf("unable to parse json: %s %s", c.path, decodeErr.Error())
	}

	c.settings = createAppSettingsFromJson(configurationJson.Settings)

	if configurationJson.Hooks == nil {
		return errors.New("no hooks config found")
	}

	for hookName, hookConfigJson := range *configurationJson.Hooks {
		hookConfig := c.HookConfig(hookName)
		hookConfig.isEnabled = true
		for _, actionJson := range hookConfigJson.Actions {
			hookConfig.AddAction(CreateActionFromJson(actionJson))
		}
	}

	return nil
}

func (c *Configuration) readConfigFile() ([]byte, error) {
	fileInfo, err := os.Stat(c.path)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("given configuration path is a directory: %s", c.path)
	}
	c.size = fileInfo.Size()

	jsonData, readErr := os.ReadFile(c.path)
	if readErr != nil {
		return nil, fmt.Errorf("could not read configuration file at: %s", c.path)
	}
	return jsonData, nil
}

func (c *Configuration) decodeConfigJson(jsonInBytes []byte) (JsonConfiguration, error) {
	var config JsonConfiguration
	if !json.Valid(jsonInBytes) {
		return config, fmt.Errorf("json configuration is invalid: %s", c.path)
	}
	marshalError := json.Unmarshal(jsonInBytes, &config)
	if marshalError != nil {
		return config, fmt.Errorf("could not load json to struct: %s %s", c.path, marshalError.Error())
	}
	return config, nil
}

func (c *Configuration) mergeSettings(settings *AppSettings, complete bool) {
	if complete {
		c.settings.AllowFailure = settings.AllowFailure
		c.settings.FailOnFirstError = settings.FailOnFirstError
		for key, value := range settings.Custom {
			c.settings.Custom[key] = value
		}
	}
	if settings.AnsiColors == false {
		c.settings.AnsiColors = settings.AnsiColors
	}
	if MapVerbosity(settings.Verbosity) > c.Verbosity() {
		c.settings.Verbosity = settings.Verbosity
	}
	if len(settings.GitDirectory) > 0 {
		c.settings.GitDirectory = settings.GitDirectory
	}
}

func MapVerbosity(verbosity string) int {
	verbosityMap := map[string]int{
		"normal":  io.NORMAL,
		"verbose": io.VERBOSE,
		"debug":   io.DEBUG,
	}
	verbosityIO, ok := verbosityMap[verbosity]
	if !ok {
		verbosityIO = io.NORMAL
	}
	return verbosityIO
}

func UnMapVerbosity(verbosity int) string {
	verbosityMap := map[int]string{
		io.NORMAL:  "normal",
		io.VERBOSE: "verbose",
		io.DEBUG:   "debug",
	}
	verbosityConfig, ok := verbosityMap[verbosity]
	if !ok {
		verbosityConfig = "normal"
	}
	return verbosityConfig
}
