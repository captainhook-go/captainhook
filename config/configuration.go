package config

import (
	"encoding/json"
	"fmt"
	json2 "github.com/captainhook-go/captainhook/config/json"
	"github.com/captainhook-go/captainhook/hooks"
	"os"
)

const (
	SETTING_ALLOW_FAILURE       = "allow-failure"
	SETTING_BOOTSTRAP           = "bootstrap"
	SETTING_COLORS              = "ansi-colors"
	SETTING_CUSTOM              = "custom"
	SETTING_GIT_DIR             = "git-directory"
	SETTING_INCLUDES            = "includes"
	SETTING_INCLUDES_LEVEL      = "includes-level"
	SETTING_RUN_PATH            = "run-path"
	SETTING_VERBOSITY           = "verbosity"
	SETTING_FAIL_ON_FIRST_ERROR = "fail-on-first-error"
)

type Configuration struct {
	size       int64
	path       string
	fileExists bool
	settings   Settings
	custom     map[string]string
	hooks      map[string]*Hook
}

func NewConfiguration(path string, fileExists bool, settings Settings) (*Configuration, error) {
	c := Configuration{path: path, fileExists: fileExists, settings: settings}
	c.path = path
	c.fileExists = fileExists
	c.init()
	if c.fileExists {
		err := c.load()
		if err != nil {
			println(err.Error())
		}
	}
	return &c, nil
}
func (c *Configuration) init() {
	c.hooks = map[string]*Hook{}
	for _, hook := range hooks.GetValidHooks() {
		c.hooks[hook] = NewHook(false)
	}
}
func (c *Configuration) IsLoadedFromFile() bool {
	return c.fileExists
}
func (c *Configuration) Path() string {
	return c.path
}
func (c *Configuration) CustomSettings() map[string]string {
	return c.custom
}
func (c *Configuration) IsFailureAllowed() bool {
	return c.settings.AllowFailure
}
func (c *Configuration) FailOnFirstError() bool {
	return c.settings.FailOnFirstError
}
func (c *Configuration) getHookConfig(hook string) *Hook {
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

	for hookName, hookConfigJson := range configurationJson.Hooks {
		hookConfig := c.getHookConfig(hookName)
		for _, actionJson := range hookConfigJson.Actions {
			hookConfig.AddAction(CreateActionFromJson(*actionJson))
		}
	}

	//fmt.Printf("config file: %s", configurationJson)

	return nil
}

func (c *Configuration) readConfigFile() ([]byte, error) {
	file, err := os.Open(c.path)
	if err != nil {
		return nil, err
	}
	stats, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file at: %s", c.path)
	}
	if stats.IsDir() {
		return nil, fmt.Errorf("given configuration path is a directory: %s", c.path)
	}

	c.size = stats.Size()
	file.Close()

	jsonData, err := os.ReadFile(c.path)
	if err != nil {
		return nil, fmt.Errorf("could not read configuration file at: %s", c.path)
	}
	return jsonData, nil
}

func (c *Configuration) decodeConfigJson(jsonInBytes []byte) (json2.ConfigurationJson, error) {
	var config json2.ConfigurationJson
	if !json.Valid(jsonInBytes) {
		return config, fmt.Errorf("json configuration is invalid: %s", c.path)
	}
	marshalError := json.Unmarshal(jsonInBytes, &config)
	if marshalError != nil {
		return config, fmt.Errorf("could not load json to struct: %s %s", c.path, marshalError.Error())
	}
	return config, nil
}
