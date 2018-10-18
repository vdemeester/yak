package minishift

import (
	"fmt"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/mcnerror"
	"github.com/docker/machine/libmachine/state"
	cmdState "github.com/minishift/minishift/cmd/minishift/state"
	"github.com/minishift/minishift/pkg/minikube/cluster"
	"github.com/minishift/minishift/pkg/minikube/constants"
	"github.com/minishift/minishift/pkg/minishift/profile"
	"github.com/pkg/errors"
)

func init() {
	// Initialize the instance directory structure
	cmdState.InstanceDirs = cmdState.NewMinishiftDirs(constants.Minipath)
}

func Start(cfg Config) error {
	profiles := profile.GetProfileList()
	if !profileExists(cfg.Name, profiles) {
		if err := createProfile(cfg); err != nil {
			return errors.Wrapf(err, "couldn't create profile %s", cfg.Name)
		}
	}
	if err := minishift("profile", "set", cfg.Name)(); err != nil {
		return errors.Wrapf(err, "couldn't activate profile %s", cfg.Name)
	}
	machineClient := libmachine.NewClient(cmdState.InstanceDirs.Home, cmdState.InstanceDirs.Certs)
	defer machineClient.Close()
	host, err := machineClient.Load(cfg.Name)
	if err != nil {
		if _, ok := err.(mcnerror.ErrHostDoesNotExist); !ok {
			return err
		}
		// You're good to go, create it !
		fmt.Println("🐂 Starting profile", cfg.Name)
		if err := minishift("start")(); err != nil {
			return err
		}
	}
	fmt.Println("host", host)
	status, err := cluster.GetHostStatus(machineClient, constants.MachineName)
	if err != nil {
		return err
	}
	if status == state.Running.String() {
		// FIXME(vdemeester) restart if needed
		fmt.Println(cfg.Name, "is already running")
		return nil
	}
	return nil
}

func createProfile(cfg Config) error {
	fmt.Println("🐂 Creating profile", cfg.Name)
	commands := []func() error{
		minishift("profile", "set", cfg.Name),
	}
	for k, v := range cfg.Configs {
		commands = append(commands, minishift("config", "set", k, fmt.Sprintf("%v", v)))
	}
	for _, a := range cfg.Addons {
		commands = append(commands, minishift("addons", "enable", a))
	}
	for _, c := range commands {
		if err := c(); err != nil {
			return err
		}
	}
	return nil
}

func profileExists(name string, profiles []string) bool {
	for _, p := range profiles {
		if name == p {
			return true
		}
	}
	return false
}
