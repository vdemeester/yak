package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/vdemeester/yak/cmd/yak/minikube"
	"github.com/vdemeester/yak/cmd/yak/minishift"
)

var rootCmd = &cobra.Command{
	Use:   "yak",
	Short: "Kubernetes pot-pourri command-line",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := args[0]
		a := args[1:]
		_, err := exec.LookPath(c)
		if err != nil {
			return errors.Errorf("%s command not found", c)
		}
		return runCommand(c, a)
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// FIXME(vdemeester) handle config, env, …
		return nil
	},
}

func runCommand(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Env = os.Environ()
	return c.Run()
}

func main() {
	rootCmd.AddCommand(
		minishift.Cmd(),
		minikube.Cmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
