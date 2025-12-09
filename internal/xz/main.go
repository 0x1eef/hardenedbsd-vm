package xz

import (
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Run() error {
	args := []string{
		"-d", "image.qcow2.xz",
	}
	return cmd.Run(exec.Command("xz", args...))
}
