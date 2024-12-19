package exec

import (
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"sort"
	"strings"
)

// ConfigInfo
// blah blah
type ConfigInfo struct {
	appIO      io.IO
	config     *configuration.Configuration
	repo       git.Repo
	show       map[string]bool
	isExtended bool
	hook       string
}

// Display allows to change the displayed details
func (c *ConfigInfo) Display(what string, show bool) {
	if show {
		c.show[what] = show
	}
}

// Extended gives you an installation status
func (c *ConfigInfo) Extended(ext bool) {
	c.isExtended = ext
}

// Hook restricts the info to show only information of given hooks
func (c *ConfigInfo) Hook(hook string) {
	c.hook = hook
}

func (c *ConfigInfo) Run() error {
	hooks := info.GetValidHooks()
	sort.Strings(hooks)
	for _, hook := range hooks {
		c.displayHook(c.config.HookConfig(hook))
	}
	return nil
}

func (c *ConfigInfo) displayHook(config *configuration.Hook) {
	if c.shouldHookBeDisplayed(config.Name()) {
		c.appIO.Write("<ok>"+config.Name()+"</ok>", !c.isExtended, io.NORMAL)
		c.displayExtended(config)
		c.displayActions(config)
	}
}

func (c *ConfigInfo) displayExtended(config *configuration.Hook) {
	if c.isExtended {
		c.appIO.Write(
			" "+strings.Repeat("-", 50-len(config.Name()))+
				"--[enabled: "+c.yesOrNo(config.IsEnabled())+
				", installed: "+c.yesOrNo(c.repo.HookExists(config.Name()))+"]",
			true,
			io.NORMAL,
		)
	}
}

func (c *ConfigInfo) displayActions(config *configuration.Hook) {
	for _, action := range config.GetActions() {
		c.displayaction(action)
	}
}

func (c *ConfigInfo) displayaction(action *configuration.Action) {
	c.appIO.Write(" - <info>"+action.Run()+"</info>", true, io.NORMAL)
	c.displayOptions(action.Options())
	c.displayConditions(action.Conditions(), "")
}

func (c *ConfigInfo) displayOptions(opts *configuration.Options) {
	if len(opts.All()) == 0 {
		return
	}
	if c.shouldDisplay("options") {
		c.appIO.Write("   <comment>Options:</comment>", true, io.NORMAL)
		for key, value := range opts.All() {
			c.displayOption(key, value, "")
		}
	}
}

func (c *ConfigInfo) displayOption(key string, value interface{}, prefix string) {
	c.appIO.Write(prefix+"    - "+key+": "+fmt.Sprintf("%v", value), true, io.NORMAL)
}

func (c *ConfigInfo) displayConditions(conditions []*configuration.Condition, prefix string) {
	if len(conditions) == 0 {
		return
	}
	if c.shouldDisplay("conditions") {
		if prefix == "" {
			c.appIO.Write("   <comment>Conditions:</comment>", true, io.NORMAL)
		}
		for _, condition := range conditions {
			c.displayCondition(condition, prefix)
		}
	}
}

func (c *ConfigInfo) displayCondition(condition *configuration.Condition, prefix string) {
	c.appIO.Write(prefix+"    - "+condition.Run(), true, io.NORMAL)

	for _, subC := range condition.Conditions() {
		c.displayCondition(subC, prefix+"  ")
	}

	if c.shouldDisplay("options") {
		if len(condition.Options().All()) == 0 {
			return
		}
		c.appIO.Write(prefix+"      <comment>Args:</comment>", true, io.NORMAL)
		for key, value := range condition.Options().All() {
			c.displayOption(key, value, prefix+"   ")
		}
	}
}

func (c *ConfigInfo) shouldHookBeDisplayed(hook string) bool {
	return c.config.IsHookEnabled(hook)
}

func (c *ConfigInfo) shouldDisplay(configPart string) bool {
	if len(c.show) == 0 {
		return true
	}
	_, ok := c.show[configPart]
	return ok
}

func (c *ConfigInfo) yesOrNo(val bool) string {
	if val {
		return "✅ "
	}
	return "❌ "
}

func NewConfigInfo(appIO io.IO, config *configuration.Configuration, repo git.Repo) *ConfigInfo {
	return &ConfigInfo{
		appIO:      appIO,
		config:     config,
		repo:       repo,
		show:       map[string]bool{},
		isExtended: false,
		hook:       "",
	}
}
