package minikube

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/vdemeester/yak/pkgs/config"
	"github.com/vdemeester/yak/pkgs/minikube"
)

func stopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop a minikube profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := homedir.Dir()
			if err != nil {
				return err
			}
			configFile := filepath.Join(home, ".config", "yak", "minikube", args[0]+".hcl")
			cfg, err := config.LoadFile(configFile)
			if err != nil {
				return err
			}
			return minikube.Stop(cfg)
		},
	}

	return cmd
}
