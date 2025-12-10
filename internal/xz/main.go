package xz

import (
	"os/exec"
	"strings"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Run(archive string) (string, error) {
	args := []string{"-d", archive}
	image, _ := strings.CutSuffix(archive, ".xz")
	return image, cmd.Run(exec.Command("xz", args...))
}
