package minishift

import (
	"fmt"

	"github.com/minishift/minishift/pkg/minishift/profile"
	"github.com/pkg/errors"
	"github.com/vdemeester/yak/pkgs/config"
)

func Stop(cfg config.Config) error {
	profiles := profile.GetProfileList()
	if !profileExists(cfg.Name, profiles) {
		return nil
	}
	fmt.Println("🐂 Stoping profile", cfg.Name)
	if err := minishift("profile", "set", cfg.Name)(); err != nil {
		return errors.Wrapf(err, "couldn't activate profile %s", cfg.Name)
	}
	return minishift("stop")()
}
