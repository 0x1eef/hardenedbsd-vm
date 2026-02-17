package input

import (
	"os"
)

var (
	Release    = get("INPUT_RELEASE", "16-CURRENT")
	Arch       = "x86_64"
	Filesystem = "ufs"
	Mem        = get("INPUT_MEM", "6144")
	Run        = get("INPUT_RUN", "uname -a")
)

func get(key, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	} else {
		return v
	}
}
