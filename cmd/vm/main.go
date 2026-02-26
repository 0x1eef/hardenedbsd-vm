package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/apt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/keys"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/rsync"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/ssh"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/vm"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/xz"
)

func main() {
	var (
		ip, archive, image, dir, script string
		payload                         []byte
		ok                              bool
		session                         *ssh.Session
		err                             error
	)
	step("Save payload", func() {
		if dir, ok = os.LookupEnv("GITHUB_WORKSPACE"); !ok {
			u, err := user.Current()
			if err != nil {
				abort("error: cannot get current user (%s)\n", err)
			}
			dir = u.HomeDir
		}
		script = path.Join(dir, "hardenedbsd-vm.sh")
		payload = fmt.Appendf(payload, "#!/bin/sh\nset -ex\ncd %s\n%s\n", dir, input.Run)
		if err = os.WriteFile(script, payload, 0755); err != nil {
			abort("error: %s\n", err)
		} else {
			fmt.Printf("OK: %s", path.Base(script))
		}
	})
	step("Install tools", func() {
		if err := apt.Run(); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("Tools installed")
	})
	step("Download VM", func() {
		if archive, err = curl.Run(input.Release); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("VM downloaded:", archive)
	})
	step("Extract VM", func() {
		if image, err = xz.Run(archive); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("VM extracted:", image)
	})
	step("Remove VM archive", func() {
		if _, err := os.Stat(archive); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "warn: archive does not exist")
		} else {
			if err := os.Remove(archive); err != nil {
				abort("error: %s", err)
			}
			fmt.Println("Removed VM archive: ", archive)
		}
	})
	step("Boot VM", func() {
		if ip, err = vm.Run(image); err != nil {
			abort("error: %s\n", err)
		}
	})
	step("Install SSH keys", func() {
		if err := keys.Install(); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("SSH keys installed")
	})
	step("Establish SSH session", func() {
		if session, err = ssh.Run(ip); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("SSH session established")
	})
	step("Sync files to VM", func() {
		if err := rsync.CopyToVM(ip, dir); err != nil {
			abort("error: %s\n", err)
		}
		fmt.Println("Files synced")
	})
	step("Run payload", func() {
		defer session.Close()
		fmt.Printf("payload: %s\n", script)
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		if err := session.Run(script); err != nil {
			abort("error: \n%s\n\n", err)
		}
	})
	step("Sync files from VM", func() {
		if input.CopyBack {
			if err := rsync.CopyFromVM(ip, dir); err != nil {
				abort("error: %s\n", err)
			}
			fmt.Println("Files synced")
		} else {
			fmt.Println("Nothing to do")
		}
	})
}

func step(label string, fn func()) {
	fmt.Printf("::group::%s\n", label)
	fn()
	fmt.Println("::endgroup::")
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}
