package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	stream = func(t io.ReadCloser, o io.Writer) {
		defer t.Close()
		chunk := make([]byte, 1024)
		for {
			n, err := t.Read(chunk)
			if n > 0 {
				fmt.Fprintf(o, "%s", chunk[:n])
			}
			if err != nil {
				break
			}
		}
	}
)

func Run(c *exec.Cmd) error {
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	err = c.Start()
	if err != nil {
		return err
	}
	go stream(stdout, os.Stdout)
	go stream(stderr, os.Stderr)
	debug(c.Args)
	err = c.Wait()
	if err != nil {
		return err
	}
	return nil
}

func debug(args []string) {
	for _, v := range args {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}
