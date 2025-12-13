package keys

import (
	"os/exec"
	"os/user"
	"path"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

func Install() error {
	if me, err := user.Current(); err != nil {
		return err
	} else {
		return install(path.Join(me.HomeDir, ".ssh"))
	}
}

func install(target string) error {
	payload := [][]string{
		{"mkdir", "-p", target},
		{"cp", "config/keys/id_ed25519", target},
		{"cp", "config/keys/id_ed25519.pub", target},
		{"chmod", "u=rw,go=", path.Join(target, "id_ed25519")},
		{"chmod", "u=rw,go=", path.Join(target, "id_ed25519.pub")},
	}
	for _, command := range payload {
		if err := cmd.Run(exec.Command(command[0], command[1:]...)); err != nil {
			return err
		}
	}
	return nil
}
