package ext

import (
	"github.com/captainhook-go/captainhook/io"
	"os/exec"
	"strings"
)

func ExecuteCommand(aIO io.IO, command string) error {
	splits := strings.Split(command, " ")
	cmd := exec.Command(splits[0], splits[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		if len(out) > 0 {
			aIO.Write("<info>output:</info>\n"+string(out), true, io.NORMAL)
		}
		return err
	}

	if len(out) > 0 {
		aIO.Write("<info>output:</info>\n"+string(out), false, io.VERBOSE)
	}
	return nil
}
