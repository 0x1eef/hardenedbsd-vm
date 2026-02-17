package input

import (
	"os"
	"runtime"
	"strconv"
)

var (
	Arch       = "x86_64"
	Filesystem = "ufs"
	Release    = get("INPUT_RELEASE", "16-CURRENT")
	Mem        = get("INPUT_MEM", "6144")
	Run        = get("INPUT_RUN", "uname -a")
	Cpu        = get("INPUT_CPU", strconv.Itoa(runtime.NumCPU()))
)

func get(key, def string) string {
	v, ok := os.LookupEnv(key)
	if v == "" || !ok {
		return def
	} else {
		return v
	}
}
