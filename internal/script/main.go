package script

import (
	"fmt"
	"os"
)

func Save(str string) (string, error) {
	payload := fmt.Sprintf("#!bin/sh\nset -x\n%s\n", str)
	return "script.sh", os.WriteFile("script.sh", []byte(payload), 0644)
}
