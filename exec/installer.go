package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"sort"
	"text/template"
)

type Installer struct {
	appIO        io.IO
	config       *configuration.Configuration
	repo         *git.Repository
	force        bool
	skipExisting bool
	onlyEnabled  bool
}

func NewInstaller(appIO io.IO, config *configuration.Configuration, repo *git.Repository) *Installer {
	i := Installer{appIO: appIO, config: config, repo: repo, force: false, skipExisting: false}
	return &i
}

func (i *Installer) SkipExisting(skip bool) {
	i.skipExisting = skip
}

func (i *Installer) OnlyEnabled(enabled bool) {
	i.onlyEnabled = enabled
}

func (i *Installer) Force(force bool) {
	i.force = force
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
	return nil
}

func (i *Installer) installHook(hook string, ask bool) error {
	if i.shouldHookBeSkipped(hook) {
		hint := "  <info>" + hook + "</info> is already installed"
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
		if i.shouldHookBeMoved() {
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
		vars["VERSION"] = info.VERSION
		vars["CONFIGURATION"] = i.config.Path()

		tpl, _ := template.New("hook").Parse(i.HookTemplate())

		file, _ := os.Create(i.repo.HooksDir() + "/" + hook)
		defer file.Close()

		i.appIO.Write("  installing <info>"+hook+"</info> to "+i.repo.HooksDir()+"/"+hook, true, io.VERBOSE)
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

func (i *Installer) shouldHookBeMoved() bool {
	return false
}

func (i *Installer) backupHook(hook string) {
	// TODO: add backup functionality
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
