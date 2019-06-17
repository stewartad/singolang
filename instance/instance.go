package instance

import (
	"github.com/stewartad/singolang/utils"
	"fmt"
	"strings"
)

// Instance holds information about a currently running image instance
type Instance struct {
	name string
	imageURI string
	protocol string
	image string
	cmd []string
	options []string
	metadata []string // might go unused
}

func (i *Instance) String() string {
	if i.protocol != "" {
		return fmt.Sprintf("%s:\\%s", i.protocol, i.image)
	}
	return i.image
}

// GetInstance returns a new Instance with image information
func GetInstance(image string, name string, options ...string) *Instance {
	i := new(Instance)
	i.parseImageName(image)

	if name != "" {
		i.name = name
	}

	i.options = options
	return i
}

// parseImageName processes the image name and protocol
func (i *Instance) parseImageName(image string) {
	i.imageURI = image
	i.protocol, i.image = utils.SplitURI(image)
}

// TODO: make this do something
func (i *Instance) updateMetadata() {

}

// Start starts an instance
// Does not support startscript args
func (i *Instance) Start(sudo bool) error {
	cmd := utils.InitCommand("instance", "start")
	cmd = append(cmd, i.imageURI, i.name)

	_, err := utils.RunCommand(cmd, sudo, false)
	return err
}

// Stop stops an instance.
func (i *Instance) Stop(sudo bool) error {
	cmd := utils.InitCommand("instance", "stop")
	cmd = append(cmd, i.name)

	_, err := utils.RunCommand(cmd, sudo, false)
	return err
}

/*
 * Getters for Instance fields
 */

 // GetInfo returns the information about an Instance
func (i *Instance) GetInfo() map[string]string {
	m := make(map[string]string)
	m["name"] = i.name
	m["imageURI"] = i.imageURI
	m["protocol"] = i.protocol
	m["image"] = i.image
	m["cmd"] = strings.Join(i.cmd, " ")
	m["options"] = strings.Join(i.options, " ")
	return m
}

// GetCmd returns a slice of strings that represent the full command created when i.Start() was called.
// This slice can immediately be passed into RunCommand() to be ran again
func (i *Instance) GetCmd() []string {
	return i.cmd
}

