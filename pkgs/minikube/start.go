package minikube

import (
	"fmt"

	"github.com/vdemeester/yak/pkgs/config"
)

func Start(cfg config.Config) error {
	commands := []func() error{
		minikube("profile", cfg.Name),
	}
	for k, v := range cfg.Configs {
		commands = append(commands, minikube("config", "set", k, fmt.Sprintf("%v", v)))
	}
	for _, c := range commands {
		if err := c(); err != nil {
			return err
		}
	}
	/*
		api, err := machine.NewAPIClient()
		if err != nil {
			return err
		}
		defer api.Close()
		status, err := cluster.GetHostStatus(api)
		if err != nil {
			return err
		}
		if status == state.Running.String() {
			// FIXME(vdemeester) restart if needed
			fmt.Println(cfg.Name, "is already running")
			return nil
		}
	*/
	return minikube("start")()
}
