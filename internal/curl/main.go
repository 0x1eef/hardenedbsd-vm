package curl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
)

var (
	Destination = "hardenedbsd-vm.raw.xz"
	URLMap      = map[string]string{
		"16-CURRENT": "https://github.com/0x1eef/hardenedbsd-builder/releases/download/hardenedbsd-16-latest/hardenedbsd-16.0-ufs-amd64.raw.xz",
		"15-STABLE":  "FIXME",
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
		destNoSuffix string   = strings.TrimSuffix(Destination, ".xz")
		targets      []string = []string{Destination, destNoSuffix}
		url          string
		err          error
	)
	if url, err = Source(); err != nil {
		return "", err
	}
	for _, target := range targets {
		if _, err = os.Stat(target); err == nil {
			return Destination, nil
		}
	}
	args := []string{"-L", "-o", Destination, url}
	return Destination, cmd.Run(exec.Command("curl", args...))
}
