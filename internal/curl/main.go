package curl

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
)

var (
	dest = "hardenedbsd-vm.raw.xz"
	base = "https://github.com/0x1eef/hardenedbsd-builder/releases/download/"
	urls = map[string]string{
		"16-CURRENT": fmt.Sprintf("%s/16CURRENT_UFS_AMD64_MODIFIED/hardenedbsd-vm.raw.xz", base),
		"15-STABLE":  fmt.Sprintf("%s/15STABLE_UFS_AMD64_MODIFIED/hardenedbsd-vm.raw.xz", base),
	}
)

func Run(release string) (string, error) {
	var (
		destNoSuffix string   = strings.TrimSuffix(dest, ".xz")
		targets      []string = []string{dest, destNoSuffix}
		url          string
		err          error
		ok           bool
	)
	for _, target := range targets {
		if _, err = os.Stat(target); err == nil {
			return dest, nil
		}
	}
	if url, ok = urls[release]; !ok {
		return "", fmt.Errorf("unknown release: %s", release)
	}
	args := []string{"-L", "-o", dest, url}
	return dest, cmd.Run(exec.Command("curl", args...))
}
