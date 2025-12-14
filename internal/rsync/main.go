package rsync

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func CopyToVM(ip string) error {
	var (
		dir string
		ok  bool
	)
	if dir, ok = os.LookupEnv("GITHUB_WORKSPACE"); !ok {
		return fmt.Errorf("GITHUB_WORKSPACE not set\nEnvironment: %v", os.Environ())
	}
	dest := fmt.Sprintf("runner@%s:%s/", ip, filepath.Dir(dir))
	args := []string{"-rvah", "--mkpath", "-e", "ssh -o StrictHostKeyChecking=no", dir, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
