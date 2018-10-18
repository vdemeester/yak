package minikube

import (
	"os"
	"os/exec"
)

func minikube(args ...string) func() error {
	return func() error {
		cmd := exec.Command("minikube", args...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stderr
		return cmd.Run()
	}
}
