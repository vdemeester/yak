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
	"github.com/docker/machine/drivers/vmwarefusion"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/minishift/minishift/pkg/minikube/constants"
	"github.com/pborman/uuid"
)

func createVMwareFusionHost(config MachineConfig) drivers.Driver {
	d := vmwarefusion.NewDriver(constants.MachineName, constants.Minipath).(*vmwarefusion.Driver)
	d.Boot2DockerURL = config.GetISOFileURI()
	d.Memory = config.Memory
	d.CPU = config.CPUs
	d.DiskSize = config.DiskSize

	// TODO(philips): push these defaults upstream to fixup this driver
	d.SSHPort = 22
	d.ISO = d.ResolveStorePath("boot2docker.iso")
	return d
}

type xhyveDriver struct {
	*drivers.BaseDriver
	Boot2DockerURL string
	BootCmd        string
	CPU            int
	CaCertPath     string
	DiskSize       int64
	MacAddr        string
	Memory         int
	PrivateKeyPath string
	UUID           string
	NFSShare       bool
	DiskNumber     int
	Virtio9p       bool
	Virtio9pFolder string
	QCow2          bool
	RawDisk        bool
}

type hyperkitDriver struct {
	*drivers.BaseDriver
	Boot2DockerURL string
	DiskSize       int
	CPU            int
	Memory         int
	Cmdline        string
	NFSShares      []string
	NFSSharesRoot  string
	UUID           string
	BootKernel     string
	BootInitrd     string
	Initrd         string
	Vmlinuz        string
}

func createXhyveHost(config MachineConfig) *xhyveDriver {
	return &xhyveDriver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: constants.MachineName,
			StorePath:   constants.Minipath,
		},
		Memory:         config.Memory,
		CPU:            config.CPUs,
		Boot2DockerURL: config.GetISOFileURI(),
		DiskSize:       int64(config.DiskSize),
		QCow2:          false,
		RawDisk:        true,
	}
}

func createHyperkitHost(config MachineConfig) *hyperkitDriver {
	return &hyperkitDriver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: constants.MachineName,
			StorePath:   constants.Minipath,
			SSHUser:     "docker",
		},
		Boot2DockerURL: config.GetISOFileURI(),
		DiskSize:       config.DiskSize,
		Memory:         config.Memory,
		CPU:            config.CPUs,
		UUID:           uuid.NewUUID().String(),
	}
}
