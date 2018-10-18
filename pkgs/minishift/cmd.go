package minishift

import (
	"os"
	"os/exec"
)

func minishift(args ...string) func() error {
	return func() error {
		cmd := exec.Command("minishift", args...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stderr
		return cmd.Run()
	}
}
