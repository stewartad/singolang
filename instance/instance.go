package instance

import (
	"github.com/stewartad/singolang/utils"
	"github.com/stewartad/singolang/mainclient"
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

// ParseImageName sets elements of the Instance i to values gathered from image. 
// It is automatically called by GetInstance(), so it only needs to be manually called on a manually defined struct
func (i *Instance) parseImageName(image string) {
	i.imageURI = image
	i.protocol, i.image = utils.SplitURI(image)
}

func (i *Instance) updateMetadata() {

}

// Start starts an instance
// Does not support startscript args
func (i *Instance) Start(sudo bool) {
	cmd := mainclient.InitCommand("instance", "start")
	cmd = append(cmd, i.imageURI, i.name)

	_, _ = utils.RunCommand(cmd, sudo, false)
}

// Stop stops an instance.
func (i *Instance) Stop(sudo bool) {
	cmd := mainclient.InitCommand("instance", "stop")
	cmd = append(cmd, i.name)

	_, _ = utils.RunCommand(cmd, sudo, false)
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

