package instance

import (
	"github.com/stewartad/singolang/utils"
	"github.com/stewartad/singolang/mainclient"
)

// Start starts an instance
func (i *Instance) Start(imgURI string, name string, sudo bool) {
	if name != "" {
		i.name = name
	}

	if imgURI == "" {
		imgURI = i.imageURI
	}

	cmd := mainclient.InitCommand("instance", "start")
	cmd = append(cmd, imgURI, i.name)

	_, _ = utils.RunCommand(cmd, sudo, false)
}