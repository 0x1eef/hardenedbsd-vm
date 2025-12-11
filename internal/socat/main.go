package socat

import (
	"fmt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/cmd"
	"os/exec"
)

func Run() error {
	fmt.Println("VM doesn't have IP yet, please wait up to 100 seconds")
	if out, err := exec.Command("bash", "socat.sh").Output(); err != nil {
		return err
	} else {
		go func() {
			dest := fmt.Sprintf("TCP:%s:22", string(out))
			exec.Command("socat", "TCP-LISTEN:2222,fork,reuseaddr", dest).Run()
		}()
		return nil
	}
}
