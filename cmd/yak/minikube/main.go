package minikube

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "minikube",
		Aliases: []string{"mk", "m"},
		Short:   "Manage minikube instance(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(stopCmd())
	cmd.AddCommand(startCmd())
	cmd.AddCommand(statusCmd())
	return cmd
}
