package xz

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Run(archive string) (string, error) {
	args := []string{"-k", "-T", "0", "-d", archive}
	image := strings.TrimSuffix(archive, ".xz")
	if _, err := os.Stat(image); errors.Is(err, os.ErrNotExist) {
		return image, cmd.Run(exec.Command("xz", args...))
	} else {
		return image, nil
	}
}
