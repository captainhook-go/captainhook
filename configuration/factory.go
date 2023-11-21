package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"path/filepath"
)

type Factory struct {
	includeLevel    int
	maxIncludeLevel int
}

// CreateConfig creates a default configuration in case the file exists it is loaded
func (f *Factory) CreateConfig(path string, cliSettings *JsonAppSettings) (*Configuration, error) {
	c, err := f.setupConfig(path)
	// load the local config "captainhook.config.json"
	decodeErr := f.loadSettingsFile(c)
	if decodeErr != nil {
		return c, decodeErr
	}
	// everything provided from the command line should overwrite any loaded configuration
	// this works even if there is an error because then you have a default configuration
	c.overwriteSettings(cliSettings)
	return c, err
}

// setupConfig creates a new configuration and loads the json file if it exists
func (f *Factory) setupConfig(path string) (*Configuration, error) {
	var err error
	c := NewConfiguration(path, io.FileExists(path))
	if c.fileExists {
		err = f.loadFromFile(c)
	}
	return c, err
}

func (f *Factory) loadFromFile(c *Configuration) error {
	jsonBytes, readError := f.readConfigFile(c.path)
	if readError != nil {
		return readError
	}
	configurationJson, decodeErr := f.decodeConfigJson(jsonBytes)
	if decodeErr != nil {
		return fmt.Errorf("unable to parse json: %s %s", c.path, decodeErr.Error())
	}
	c.overwriteSettings(configurationJson.Settings)

	if configurationJson.Hooks == nil {
		return errors.New("no hooks config found")
	}
	includeErr := f.appendIncludedConfiguration(c)
	if includeErr != nil {
		return includeErr
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

func (f *Factory) loadSettingsFile(c *Configuration) error {
	directory := filepath.Dir(c.path)
	filePath := directory + "/captainhook.config.json"

	// no local config file to load just exit
	if !io.FileExists(filePath) {
		return nil
	}

	jsonBytes, readError := f.readConfigFile(filePath)
	if readError != nil {
		return readError
	}
	appSettingJson, decodeErr := f.decodeSettingJson(jsonBytes)
	if decodeErr != nil {
		return fmt.Errorf("unable to parse json: %s %s", filePath, decodeErr.Error())
	}
	// overwrite current settings
	c.overwriteSettings(appSettingJson)
	return nil
}

func (f *Factory) appendIncludedConfiguration(c *Configuration) error {
	f.detectMaxIncludeLevel(c)
	if f.includeLevel < f.maxIncludeLevel {
		f.includeLevel++
		includes, err := f.loadIncludedConfigs(c.Includes(), c.path)
		if err != nil {
			return err
		}
		for _, configToInclude := range includes {
			f.mergeHookConfigs(configToInclude, c)
		}
		f.includeLevel--
	}
	return nil
}

func (f *Factory) mergeHookConfigs(from, to *Configuration) {
	for _, hook := range info.GetValidHooks() {
		// This `Enable` is solely to overwrite the main configuration in the special case that the hook
		// is not configured at all. In this case the empty config is disabled by default, and adding an
		// empty hook config just to enable the included actions feels a bit dull.
		// Since the main hook is processed last (if one is configured) the enabled flag will be overwritten
		// once again by the main config value. This is to make sure that if somebody disables a hook in its
		// main configuration no actions will get executed, even if we have enabled hooks in any include file.
		targetHookConfig := to.HookConfig(hook)
		targetHookConfig.Enable()
		f.copyActionsFromTo(from.HookConfig(hook), targetHookConfig)
	}
}

func (f *Factory) readConfigFile(path string) ([]byte, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("given configuration path is a directory: %s", path)
	}
	jsonData, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, fmt.Errorf("could not read configuration file at: %s", path)
	}
	return jsonData, nil
}

func (f *Factory) decodeConfigJson(jsonInBytes []byte) (JsonConfiguration, error) {
	var jConfig JsonConfiguration
	if !json.Valid(jsonInBytes) {
		return jConfig, fmt.Errorf("json configuration is invalid")
	}
	marshalError := json.Unmarshal(jsonInBytes, &jConfig)
	if marshalError != nil {
		return jConfig, fmt.Errorf("could not load json to struct: %s", marshalError.Error())
	}
	return jConfig, nil
}

func (f *Factory) decodeSettingJson(jsonInBytes []byte) (*JsonAppSettings, error) {
	var jSettings JsonAppSettings
	if !json.Valid(jsonInBytes) {
		return nil, fmt.Errorf("json configuration is invalid")
	}
	marshalError := json.Unmarshal(jsonInBytes, &jSettings)
	if marshalError != nil {
		return nil, fmt.Errorf("could not load json to struct: %s", marshalError.Error())
	}
	return &jSettings, nil
}

func (f *Factory) detectMaxIncludeLevel(c *Configuration) {
	// read the include-level setting only for the actual configuration not any included ones
	if f.includeLevel == 0 {
		f.maxIncludeLevel = c.MaxIncludeLevel()
	}
}

func (f *Factory) loadIncludedConfigs(includes []string, path string) ([]*Configuration, error) {
	var configs []*Configuration
	directory := filepath.Dir(path)

	for _, file := range includes {
		config, err := f.includeConfig(directory + "/" + file)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

func (f *Factory) includeConfig(path string) (*Configuration, error) {
	if !io.FileExists(path) {
		return nil, fmt.Errorf("config to include not found: %s", path)
	}
	return f.setupConfig(path)
}

func (f *Factory) copyActionsFromTo(from *Hook, to *Hook) {
	for _, action := range from.GetActions() {
		to.AddAction(action)
	}
}

func NewFactory() *Factory {
	return &Factory{includeLevel: 0}
}
