/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"fmt"
	"path/filepath"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/minishift/minishift/pkg/minikube/constants"
)

type kvmDriver struct {
	*drivers.BaseDriver

	Memory         int
	DiskSize       int
	CPU            int
	Network        string
	PrivateNetwork string
	ISO            string
	Boot2DockerURL string
	DiskPath       string
	CacheMode      string
	IOMode         string
}

func createKVMHost(config MachineConfig) *kvmDriver {
	return &kvmDriver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: constants.MachineName,
			StorePath:   constants.Minipath,
			SSHUser:     "docker",
		},
		Memory:         config.Memory,
		CPU:            config.CPUs,
		PrivateNetwork: "docker-machines",
		Network:        "default",
		Boot2DockerURL: config.GetISOFileURI(),
		DiskSize:       config.DiskSize,
		DiskPath:       filepath.Join(constants.Minipath, "machines", constants.MachineName, fmt.Sprintf("%s.img", constants.MachineName)),
		ISO:            filepath.Join(constants.Minipath, "machines", constants.MachineName, "boot2docker.iso"),
		CacheMode:      "default",
		IOMode:         "threads",
	}
}
