package curl

import (
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

const (
	target = "image.raw"
	image  = "https://github.com/0x1eef/hardenedbsd-builder/actions/runs/20037038355/artifacts/4800738241"
)

func Run() (string, error) {
	args := []string{"-L", "-o", target, url()}
	return target, cmd.Run(exec.Command("curl", args...))
}

func url() string {
	return image
}
