package singolang

import (
	"fmt"
	"strings"
)

// instance holds information about a currently running image instance
type instance struct {
	name     string
	imageURI string
	protocol string
	image    string
	cleanenv bool
	env		 map[string]string
	cmd      []string
	options  []string
	metadata []string // might go unused
}

var instanceOpts = runCommandOptions{
	sudo:     false,
	quietout: true,
	quieterr: true,
}

func (i *instance) String() string {
	if i.protocol != "" {
		return fmt.Sprintf("%s:\\%s", i.protocol, i.image)
	}
	return i.image
}

// GetInstance returns a new Instance with image information
func getInstance(image string, name string, options ...string) *instance {
	i := new(instance)
	i.parseImageName(image)

	if name != "" {
		i.name = name
	}
	i.cleanenv = true
	i.env = make(map[string]string)

	i.options = options
	return i
}

// parseImageName processes the image name and protocol
func (i *instance) parseImageName(image string) {
	i.imageURI = image
	i.protocol, i.image = SplitURI(image)
}

// TODO: make this do something
func (i *instance) updateMetadata() {

}

// Start starts an instance
// Does not support startscript args
func (i *instance) start(sudo bool) error {
	cmd := initCommand("instance", "start")

	cmd = append(cmd, i.imageURI, i.name)

	if i.cleanenv {
		cmd = append(cmd, "--cleanenv")
	}

	stdout, stderr, status, err := runCommand(cmd, &instanceOpts)
	// TODO: use these
	_, _, _ = stdout, stderr, status

	err = i.processEnv()

	return err
}

// Stop stops an instance.
func (i *instance) stop(sudo bool) error {
	cmd := initCommand("instance", "stop")
	cmd = append(cmd, i.name)

	stdout, stderr, status, err := runCommand(cmd, &instanceOpts)
	// TODO: use these
	_, _, _ = stdout, stderr, status
	return err
}

// GetEnvVar retrieves an environment variable
func (i *instance) getEnvVar(varname string) (string, string) {
	return varname, i.env[varname]
}

// SetEnvVar sets an environment variable
func (i *instance) setEnvVar() {

}

// ProcessEnv retrieves all env variables in the instance and stores them in a map
func (i *instance) processEnv() error {
	cmd := []string{"singularity", "exec", "--cleanenv", fmt.Sprintf("instance://%s", i.name), "env"}
	stdout, _, _, err := runCommand(cmd, &runCommandOptions{
		quietout: true,
		quieterr: true,
		sudo: false,
	})
	if err != nil {
		return err
	}

	for _, env := range strings.Split(string(stdout.Bytes()), "\n") {
		v := strings.Split(env, "=")
		if len(v) > 1 {
			i.env[v[0]] = v[1]
		}
	}

	return err
}

/*
 * Getters for Instance fields
 */

// GetInfo returns the information about an Instance
func (i *instance) GetInfo() map[string]string {
	m := make(map[string]string)
	m["name"] = i.name
	m["imageURI"] = i.imageURI
	m["protocol"] = i.protocol
	m["image"] = i.image
	m["cmd"] = strings.Join(i.cmd, " ")
	m["options"] = strings.Join(i.options, " ")
	return m
}

func (i *instance) getEnv() map[string]string {
	return i.env
}

// GetCmd returns a slice of strings that represent the full command created when i.Start() was called.
// This slice can immediately be passed into RunCommand() to be ran again
func (i *instance) GetCmd() []string {
	return i.cmd
}
