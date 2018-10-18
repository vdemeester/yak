package minishift

import (
	"os/exec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type minishiftOption struct {
	profile string
}

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "minishift",
		Aliases: []string{"ms"},
		Short:   "Manage minishift instance(s)",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := exec.LookPath("minishift")
			if err != nil {
				return errors.New("minishift command not found")
			}
			return nil
		},
	}
	cmd.AddCommand(stopCmd())
	cmd.AddCommand(startCmd())
	cmd.AddCommand(statusCmd())
	return cmd
}
