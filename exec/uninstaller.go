package exec

import (
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/info"
	"github.com/captainhook-go/captainhook/io"
	"os"
	"sort"
	"strings"
)

// Uninstaller is responsible to delete the hook files from your local .git/hooks directory
// Normally it makes sure to ask for confirmation.
// If you don't want to be bothered to acknowledge every hook you can use
// the Force function to activate the `force` mode.
// If you are unsure use EnableBackup to create a {HOOK-FILE}.bak copy before deletion.
type Uninstaller struct {
	appIO         io.IO
	config        *configuration.Configuration
	repo          git.Repo
	force         bool
	backupEnabled bool
}

// Force makes sure to remove all hooks without asking any questions
func (u *Uninstaller) Force(force bool) {
	u.force = force
}

// EnableBackup makes sure existing hooks will be moved away to keep a backup
func (u *Uninstaller) EnableBackup(backup bool) {
	u.backupEnabled = backup
}

// Run execute the Uninstaller
func (u *Uninstaller) Run() error {
	hooks := u.hooksToHandle()

	// do some sort magic because go range random is weird
	var hookNames []string
	for key := range hooks {
		hookNames = append(hookNames, key)
	}
	sort.Strings(hookNames)
	for _, hook := range hookNames {
		// hook should not be handled or does not exist
		if !hooks[hook] || !u.repo.HookExists(hook) {
			continue
		}
		err := u.uninstallHook(hook, !u.force)
		if err != nil {
			u.appIO.Write("<warn>failed to uninstall hooks</warn>", true, io.NORMAL)
			return err
		}
	}
	u.appIO.Write("<ok>hooks uninstalled successfully</ok>", true, io.NORMAL)
	return nil
}

func (u *Uninstaller) hooksToHandle() map[string]bool {
	hooks := map[string]bool{}
	for _, hook := range info.GetNativeHooks() {
		hooks[hook] = true
	}
	return hooks
}

// uninstallHook asks for confirmation if needed creates a backup if requested and deletes the hook
func (u *Uninstaller) uninstallHook(hook string, ask bool) error {
	doIt := true
	if ask {
		answer := u.appIO.Ask("  Remove <info>"+hook+"</info> hook? <comment>[n,Y]</comment> ", "n")
		doIt = io.AnswerToBool(answer)
	}

	if doIt {
		return u.handleFiles(hook)
	}
	return nil
}

// handleFiles is copying the existing hook for backup purposes and deletes the existing one
func (u *Uninstaller) handleFiles(hook string) error {
	u.appIO.Write("  <info>"+hook+"</info>", true, io.VERBOSE)

	// handle backup
	backupErr := u.handleBackup(hook)
	if backupErr != nil {
		return backupErr
	}
	// handle hook deletion
	return u.handleDeletion(hook)
}

func (u *Uninstaller) handleBackup(hook string) error {
	u.appIO.Write("    backup"+strings.Repeat(" ", 24)+": ", false, io.VERBOSE)
	if u.backupEnabled {
		err := u.backup(hook)
		if err != nil {
			u.appIO.Write("<warning>failed</warning>", true, io.VERBOSE)
			return err
		}
		u.appIO.Write("<ok>done</ok>", true, io.VERBOSE)
		return nil
	}
	u.appIO.Write("<comment>skipped</comment>", true, io.VERBOSE)
	return nil
}

func (u *Uninstaller) handleDeletion(hook string) error {
	u.appIO.Write("    deletion"+strings.Repeat(" ", 22)+": ", false, io.VERBOSE)
	err := u.delete(hook)
	if err != nil {
		u.appIO.Write("<warning>failed</warning>", true, io.VERBOSE)
		return err
	}
	u.appIO.Write("<ok>done</ok>", true, io.VERBOSE)
	return nil
}

func (u *Uninstaller) backup(hook string) error {
	original := u.repo.HooksDir() + "/" + hook
	backup := original + ".bak"
	data, _ := os.ReadFile(original)
	return os.WriteFile(backup, data, 0644)
}

func (u *Uninstaller) delete(hook string) error {
	return os.Remove(u.repo.HooksDir() + "/" + hook)
}

func NewUninstaller(appIO io.IO, config *configuration.Configuration, repo git.Repo) *Uninstaller {
	return &Uninstaller{
		appIO:         appIO,
		config:        config,
		repo:          repo,
		force:         false,
		backupEnabled: false,
	}
}
