package rsync

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func CopyToVM(ip, src string) error {
	dest := fmt.Sprintf("runner@%s:%s/", ip, filepath.Dir(src))
	args := []string{"-rvah", "--mkpath", "-e", "ssh -o StrictHostKeyChecking=no", src, dest}
	return cmd.Run(exec.Command("rsync", args...))
}

func CopyFromVM(ip, dir string) error {
	src := fmt.Sprintf("runner@%s:%s", ip, dir)
	dest := filepath.Dir(dir)
	args := []string{"-rvah", "--mkpath", "-e", "ssh -o StrictHostKeyChecking=no", src, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
