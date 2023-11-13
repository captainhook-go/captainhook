package exec

import (
	"encoding/json"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
)

type Initializer struct {
	appIO  io.IO
	config string
	force  bool
}

func NewInitializer(appIO io.IO) *Initializer {
	i := Initializer{appIO: appIO, config: "captainhook.json", force: false}
	return &i
}

func (i *Initializer) UseConfig(config string) {
	i.config = config
}

func (i *Initializer) Force(force bool) {
	i.force = force
}

func (i *Initializer) Run() error {
	i.appIO.Write("Initializing CaptainHook", true, io.NORMAL)

	i.appIO.Write("  writing config to '"+i.config+"'", true, io.VERBOSE)

	defaultGitDir := ".git"
	hookConfigs := i.createJsonHookConfigs()

	jsonConfig := &configuration.JsonConfiguration{
		Settings: &configuration.JsonAppSettings{GitDirectory: &defaultGitDir},
		Hooks:    hookConfigs,
	}

	res, jsonErr := json.MarshalIndent(jsonConfig, "", "  ")
	if jsonErr != nil {
		i.appIO.Write("<warning>json generation failed</warning>", true, io.NORMAL)
		return jsonErr
	}

	err := i.writeConfigFile(res)
	if err != nil {
		i.appIO.Write("<warning>initializing failed</warning>", true, io.NORMAL)
		return err
	}
	i.appIO.Write("<ok>successfully initialized</ok>", true, io.NORMAL)
	return nil
}

func (i *Initializer) createJsonHookConfigs() *map[string]*configuration.JsonHook {
	configs := map[string]*configuration.JsonHook{}

	for _, hook := range info.GetValidHooks() {
		configs[hook] = &configuration.JsonHook{
			Actions: make([]*configuration.JsonAction, 0),
		}
	}

	return &configs
}

func (i *Initializer) writeConfigFile(res []byte) error {
	doIt := true

	if i.needConfirmationToOverwrite() {
		answer := i.appIO.Ask("  <info>"+i.config+"</info> exists! Overwrite? <comment>[y,N]</comment> ", "n")
		doIt = io.AnswerToBool(answer)
	}

	if doIt {
		writeErr := os.WriteFile(i.config, res, 0644)
		if writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func (i *Initializer) needConfirmationToOverwrite() bool {
	return !i.force && io.FileExists(i.config)
}
