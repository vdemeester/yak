package minishift

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/vdemeester/yak/pkgs/config"
	"github.com/vdemeester/yak/pkgs/minishift"
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start a minishift profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := homedir.Dir()
			if err != nil {
				return err
			}
			configFile := filepath.Join(home, ".config", "yak", "minishift", args[0]+".hcl")
			cfg, err := config.LoadFile(configFile)
			if err != nil {
				return err
			}
			return minishift.Start(cfg)
		},
	}

	return cmd
}
