package singolang

import (
	"fmt"
	"strings"
)

// Instance holds information about a currently running image instance
type Instance struct {
	Name     string
	ImageURI string
	Protocol string
	Image    string
	Cleanenv bool
	ImgLabels	 map[string]string
	ImgEnvVars	 map[string]string
	EnvOpts	 *EnvOptions
	Sudo	 bool
	cmd      []string
	running	 bool
}

var instanceOpts = runCommandOptions{
	sudo:     false,
	quietout: true,
	quieterr: true,
}

// TODO; Clean this up
func (i *Instance) String() string {
	if i.Protocol != "" {
		return fmt.Sprintf("%s:\\%s", i.Protocol, i.Image)
	}
	return i.Image
}

// GetInstance returns a new Instance with image information
func getInstance(image string, name string, sudo bool) *Instance {
	i := new(Instance)
	i.parseImageName(image)

	if name != "" {
		i.Name = name
	}
	i.Cleanenv = true
	i.ImgEnvVars = make(map[string]string)
	i.ImgLabels = make(map[string]string)
	i.EnvOpts = DefaultEnvOptions()
	i.Sudo = sudo
	i.running = false

	return i
}

// IsRunning returns true if the instance ins currently running
func (i *Instance) IsRunning() bool {
	return i.running
}

// parseImageName processes the image name and protocol
func (i *Instance) parseImageName(image string) {
	i.ImageURI = image
	i.Protocol, i.Image = SplitURI(image)
}

// Start starts an instance
func (i *Instance) Start() error {
	cmd := initCommand("instance", "start")

	cmd = append(cmd, i.ImageURI, i.Name)

	if i.Cleanenv {
		cmd = append(cmd, "--cleanenv")
	}

	var err error
	var status = 1
	if !i.running {
		_, _, status, err = runCommand(cmd, &instanceOpts)
		if status == 0 {
			i.running = true
		}
	}
	return err
}

// Stop stops an instance.
func (i *Instance) Stop() error {
	cmd := initCommand("instance", "stop")
	cmd = append(cmd, i.Name)

	var status = 1
	var err error

	if i.running {
		_, _, status, err = runCommand(cmd, &instanceOpts)
		if status == 0 {
			i.running = false
		}
	}
	
	return err
}

// RetrieveLabels gets the labels of the instance using `singularity inspect` and loads them into the map ImgLabels
func (i *Instance) RetrieveLabels() error {
	i.ImgLabels = make(map[string]string)
	cmd := []string{"singularity", "inspect", "--labels", i.Image}
	stdout, _, _, err := runCommand(cmd, defaultRunCommandOptions())

	for _, label := range strings.Split(string(stdout.Bytes()), "\n") {
		v := strings.Split(label, ":")
		if len(v) > 1 {
			i.ImgLabels[strings.TrimSpace(v[0])] = strings.TrimSpace(v[1])
		}
	}
	return err
}

// RetrieveEnv retrieves all env variables in the instance and stores them in the map ImgEnvVars
func (i *Instance) RetrieveEnv() error {
	i.ImgEnvVars = make(map[string]string)
	cmd := initCommand("exec")
	if i.Cleanenv {
		cmd = append(cmd, "--cleanenv")
	}
	cmd = append(cmd, i.Image, "env")

	stdout, _, _, err := runCommand(cmd, defaultRunCommandOptions())
	output := string(stdout.Bytes())
	if err != nil {
		return err
	}

	for _, env := range strings.Split(output, "\n") {
		v := strings.Split(env, "=")
		if len(v) > 1 {
			i.ImgEnvVars[strings.TrimSpace(v[0])] = strings.TrimSpace(v[1])
		}
	}

	return err
}

// SetEnv replaces the environment options of an instance and then processes the changes
func (i *Instance) SetEnv(opts *EnvOptions) {
	i.EnvOpts = opts
	i.RetrieveEnv()
	i.RetrieveLabels()
	i.EnvOpts.ProcessEnvVars()
}

// GetCmd returns a slice of strings that represent the full command created when i.Start() was called.
func (i *Instance) GetCmd() []string {
	return i.cmd
}
