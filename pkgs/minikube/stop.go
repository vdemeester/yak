package minikube

import (
	"fmt"

	"github.com/vdemeester/yak/pkgs/config"
)

func Stop(cfg config.Config) error {
	fmt.Println("🐂 Stoping profile", cfg.Name)
	if err := minikube("profile", cfg.Name)(); err != nil {
		return err
	}
	return minikube("stop")()
}
