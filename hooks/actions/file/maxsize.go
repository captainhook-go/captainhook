package file

import (
	"errors"
	"fmt"
	"github.com/captainhook-go/captainhook/configuration"
	"github.com/captainhook-go/captainhook/git"
	"github.com/captainhook-go/captainhook/hooks"
	"github.com/captainhook-go/captainhook/io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MaxSize struct {
	hookBundle *hooks.HookBundle
}

func (a *MaxSize) IsApplicableFor(hook string) bool {
	return a.hookBundle.Restriction.IsApplicableFor(hook)
}

func (a *MaxSize) Run(action *configuration.Action) error {
	a.hookBundle.AppIO.Write("checking max file size", true, io.VERBOSE)

	size := action.Options().StringValueOf("max-size", "0")
	sizeInBytes := a.toBytes(size)
	if sizeInBytes == 0 {
		return errors.New("the 'size' option is missing or wrong")
	}
	files, err := a.hookBundle.Repo.StagedFiles()
	if err != nil {
		return err
	}
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		stats, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to read file at: %s", path)
		}

		if sizeInBytes < stats.Size() {
			_ = file.Close()
			return fmt.Errorf("file '%s' is bigger than the limit of %s", path, size)
		}
		_ = file.Close()
	}
	return nil
}

func (a *MaxSize) toBytes(value string) int64 {
	matched, err := regexp.MatchString("^[0-9]*[BKMGTP]$", value)
	if err != nil {
		return 0
	}
	if !matched {
		return 0
	}
	units := map[string]float64{"B": 0, "K": 1, "M": 2, "G": 3, "T": 4, "P": 5}
	unit := strings.ToUpper(io.SubString(value, -1, 1))
	number, err := strconv.ParseInt(io.SubString(value, 0, -1), 10, 64)
	if err != nil {
		return 0
	}
	unitBytes := math.Pow(1024, units[unit])
	return number * int64(math.Round(unitBytes))
}

func NewMaxSize(appIO io.IO, conf *configuration.Configuration, repo *git.Repository) hooks.Action {
	a := MaxSize{
		hookBundle: hooks.NewHookBundle(appIO, conf, repo, []string{}),
	}
	return &a
}
