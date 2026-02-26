package vm

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/input"
)

func Run(image string) (string, error) {
	var (
		vm  string = "testvm"
		ip  string
		err error
	)
	if err := create(vm, image); err != nil {
		return "", err
	}
	if ip, err = waitForIP(vm, 25); err != nil {
		return "", err
	}
	return ip, nil
}

func create(vm, image string) error {
	args := []string{
		"--name", vm,
		"--memory", input.Mem,
		"--vcpus", input.Cpu,
		"--arch", input.Arch,
		"--disk", "path=" + image + ",format=raw,bus=virtio",
		"--os-variant", "freebsd13.1",
		"--network", "network=default,model=e1000",
		"--graphics", "vnc,listen=0.0.0.0",
		"--noautoconsole",
		"--import",
	}
	return cmd.Run(exec.Command("sudo", append([]string{"virt-install"}, args...)...))
}

func waitForIP(vmName string, maxAttempts int) (string, error) {
	re := regexp.MustCompile(`ipv4\s+([0-9.]+)\/`)
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Discover VM (attempt %d/%d)\n", attempt, maxAttempts)
		cmd := exec.Command("sudo", "virsh", "domifaddr", vmName)
		out, err := cmd.Output()
		if err == nil {
			matches := re.FindSubmatch(bytes.TrimSpace(out))
			if len(matches) == 2 {
				fmt.Printf("The VM has an IP (%s)\n", matches[1])
				return string(matches[1]), nil
			} else {
				fmt.Printf("The VM doesn't have an IP yet\n%s\n", string(out))
			}
		}
		time.Sleep(5 * time.Second)
		fmt.Println()
	}
	return "", exec.ErrNotFound
}
