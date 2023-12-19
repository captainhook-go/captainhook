package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"sort"
	"strings"
	"text/template"
)

// Installer is responsible to write the hook files into your local .git/hooks directory
// Normally it makes sure you don't overwrite your existing git hooks.
// If you don't want to be bothered to acknowledge every hook you can use
// the Force function to activate the `force` mode
type Installer struct {
	appIO         io.IO
	config        *configuration.Configuration
	repo          git.Repo
	force         bool
	skipExisting  bool
	onlyEnabled   bool
	backupEnabled bool
}

// SkipExisting makes sure you don't overwrite existing hooks
func (i *Installer) SkipExisting(skip bool) {
	i.skipExisting = skip
}

// OnlyEnabled makes sure you only install hooks that are activated in the configuration
func (i *Installer) OnlyEnabled(enabled bool) {
	i.onlyEnabled = enabled
}

// Force makes sure to install all hooks without asking any questions
func (i *Installer) Force(force bool) {
	i.force = force
}

// EnableBackup makes sure existing hooks will be moved away to keep a backup
func (i *Installer) EnableBackup(backup bool) {
	i.backupEnabled = backup
}

func (i *Installer) Run() error {
	hooks := i.getHooksToInstall()

	// do some sort magic because go range random is weird
	var keys []string
	for key := range hooks {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, hook := range keys {
		err := i.installHook(hook, hooks[hook] && !i.force)
		if err != nil {
			return err
		}
	}
	i.appIO.Write("<ok>hooks installed successfully</ok>", true, io.NORMAL)
	return nil
}

func (i *Installer) installHook(hook string, ask bool) error {
	if i.shouldHookBeSkipped(hook) {
		hint := "  <info>" + hook + "</info>" + strings.Repeat(" ", 30-len(hook)) + "skipped"
		if i.appIO.IsDebug() {
			hint = ", remove the --skip-existing option to overwrite."
		}

		i.appIO.Write(hint, true, io.VERBOSE)
		return nil
	}
	doIt := true
	if ask {
		answer := i.appIO.Ask("  Install <info>"+hook+"</info> hook? <comment>[Y,n]</comment> ", "y")
		doIt = io.AnswerToBool(answer)
	}

	if doIt {
		if i.isBackupEnabled() {
			i.backupHook(hook)
		}
		return i.writeHookFile(hook)
	}
	return nil
}

func (i *Installer) writeHookFile(hook string) error {
	doIt := true
	// if hook is configured and no force option is set
	// ask the user if overwriting the hook is ok
	if i.needConfirmationToOverwrite(hook) {
		answer := i.appIO.Ask("  The <info>"+hook+"</info> hook exists! Overwrite? <comment>[y,N]</comment> ", "n")
		doIt = io.AnswerToBool(answer)
	}

	if doIt {
		vars := make(map[string]interface{})
		vars["HOOK_NAME"] = hook
		vars["RUN_PATH"] = i.config.RunPath()
		vars["INTERACTION"] = false
		vars["VERSION"] = info.Version
		vars["CONFIGURATION"] = i.config.Path()

		tpl, _ := template.New("hook").Parse(i.HookTemplate())

		file, _ := os.Create(i.repo.HooksDir() + "/" + hook)
		defer file.Close()

		i.appIO.Write("  <info>"+hook+"</info>"+strings.Repeat(" ", 30-len(hook))+"<ok>installed</ok>", true, io.VERBOSE)
		tplErr := tpl.Execute(file, vars)
		if tplErr != nil {
			return tplErr
		}
		return os.Chmod(file.Name(), 0700)
	}
	return nil
}

func (i *Installer) shouldHookBeSkipped(hook string) bool {
	return i.skipExisting && i.repo.HookExists(hook)
}

func (i *Installer) needConfirmationToOverwrite(hook string) bool {
	return !i.force && i.repo.HookExists(hook)
}

func (i *Installer) getHooksToInstall() map[string]bool {
	hooks := i.hooksToHandle()
	// if only enabled hooks should be installed remove disabled ones from hooks map
	if i.onlyEnabled {
		filter := func(ss map[string]bool, test func(string) bool) map[string]bool {
			ret := map[string]bool{}
			for hook, ask := range ss {
				if test(hook) {
					ret[hook] = ask
				}
			}
			return ret
		}
		test := func(hook string) bool {
			return i.config.IsHookEnabled(hook)
		}
		hooks = filter(hooks, test)
	}
	return hooks
}

func (i *Installer) hooksToHandle() map[string]bool {
	hooks := map[string]bool{}
	for _, hook := range info.GetNativeHooks() {
		hooks[hook] = true
	}
	return hooks
}

func (i *Installer) isBackupEnabled() bool {
	return i.backupEnabled
}

func (i *Installer) backupHook(hook string) {
	original := i.repo.HooksDir() + "/" + hook

	if !io.FileExists(original) {
		return
	}
	backup := original + ".old"
	data, _ := os.ReadFile(original)
	err := os.WriteFile(backup, data, 0644)
	if err == nil {
		i.appIO.Write("backup '"+hook+"' to '"+backup+"'", true, io.VERBOSE)
	}
	return
}

func (i *Installer) HookTemplate() string {
	return "#!/bin/sh\n" +
		"\n" +
		"# installed by CaptainHook {{ .VERSION }}\n" +
		"\n" +
		"INTERACTIVE=\"{{ if .INTERACTION }}--no-interaction {{ end }}\"\n" +
		"\n" +
		"{{ .RUN_PATH }}captainhook $INTERACTIVE--configuration={{ .CONFIGURATION }} hook {{ .HOOK_NAME }} \"$@\" <&0\n\n"
}

func NewInstaller(appIO io.IO, config *configuration.Configuration, repo git.Repo) *Installer {
	return &Installer{
		appIO:         appIO,
		config:        config,
		repo:          repo,
		force:         false,
		skipExisting:  false,
		onlyEnabled:   false,
		backupEnabled: false,
	}
}
