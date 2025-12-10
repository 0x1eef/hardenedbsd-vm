package curl

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
)

var (
	Destination = "hardenedbsd-vm.raw.xz"
	URLMap      = map[string]string{
		"16-CURRENT": "FIXME",
		"15-STABLE":  "https://github.com/0x1eef/hardenedbsd-builder/releases/download/hardenedbsd-15-latest/hardenedbsd-15.0-ufs-amd64.raw.xz",
		"14-STABLE":  "FIXME",
	}
)

func Source() (string, error) {
	url, ok := URLMap[input.Release]
	if !ok {
		return "", fmt.Errorf("unknown release: %s", input.Release)
	}
	return url, nil
}

func Run() (string, error) {
	var (
		url string
		err error
	)
	if url, err = Source(); err != nil {
		return "", err
	}
	if _, err = os.Stat(Destination); errors.Is(err, os.ErrNotExist) {
		args := []string{"-L", "-o", Destination, url}
		return Destination, cmd.Run(exec.Command("curl", args...))
	} else {
		return Destination, nil
	}
}
