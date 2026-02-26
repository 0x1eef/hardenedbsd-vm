package input

import (
	"os"
	"runtime"
	"strconv"
)

var (
	Arch       = "x86_64"
	Filesystem = "ufs"
	Release    = getString("INPUT_RELEASE", "16-CURRENT")
	Mem        = getString("INPUT_MEM", "6144")
	Run        = getString("INPUT_RUN", "uname -a")
	Cpu        = getString("INPUT_CPU", strconv.Itoa(runtime.NumCPU()))
	CopyBack   = getBool("INPUT_COPYBACK", true)
)

func getString(key string, def string) string {
	if s, ok := get(key, def).(string); ok {
		return s
	} else {
		// This can't happen
		return ""
	}
}

func getBool(key string, def bool) bool {
	if b, ok  := get(key, def).(bool); ok {
		return b
	} else {
		// This can't happen
		return false
	}
}

func get(key string, def any) any {
	v, ok := os.LookupEnv(key)
	if v == "" || !ok {
		return def
	} else {
		return v
	}
}
