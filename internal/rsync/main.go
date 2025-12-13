package rsync

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func CopyToVM(ip string) error {
	var (
		wrkdir string
		err    error
	)
	if wrkdir, err = os.Getwd(); err != nil {
		return err
	}
	src := wrkdir
	dest := fmt.Sprintf("runner@%s:~/", ip)
	args := []string{"-rvah", "-e", "ssh -o StrictHostKeyChecking=no", src, dest}
	return cmd.Run(exec.Command("rsync", args...))
}
