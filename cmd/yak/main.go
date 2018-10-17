package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yak",
	Short: "Kubernetes pot-pourri command-line",
	RunE: func(cml *cobra.Command, args []string) error {
		fmt.Println("Hello yak !")
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
