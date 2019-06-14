package instance

import (
	"github.com/stewartad/singolang/utils"
	"github.com/stewartad/singolang/mainclient"
)

// Stop stops an instance
func (i *Instance) Stop(name string, sudo bool) {
	cmd := mainclient.InitCommand("instance", "stop")

	instanceName := i.name
	if name != "" {
		instanceName = name
	}

	cmd = append(cmd, instanceName)

	_, _ = utils.RunCommand(cmd, sudo, false)
}