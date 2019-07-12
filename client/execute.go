package client

import (
	"fmt"
	"strings"
)

const senv string = "SINGULARITYENV_"

// ExecOptions provide flags simulating options int he singularity command line
type ExecOptions struct {
	Pwd   string
	Quiet bool
	Cleanenv bool
	EnvVars map[string]string
	PrependPath []string
	AppendPath []string
	ReplacePath	string
}

// DefaultExecOptions provides a default options struct
func DefaultExecOptions() *ExecOptions {
	return &ExecOptions{
		Pwd:   "",
		Quiet: true,
		Cleanenv: true,
		EnvVars: make(map[string]string),
		PrependPath: []string{},
		AppendPath: []string{},
	}
}

type existError struct {
	instance string
}

func (e *existError) Error() string {
	return fmt.Sprintf("%s is not a loaded instance", e.instance)
}

// Execute runs a command inside a container
func (c *Client) Execute(instance string, command []string, opts *ExecOptions) (string, string, int, error) {
	// TODO: check install

	cmd := initCommand("exec")

	_, exists := c.instances[instance]
	if !exists {
		return "", "", -1, &existError{instance}
	}

	// TODO: bind paths

	if opts.Cleanenv {
		cmd = append(cmd, "--cleanenv")
	}
	if opts.Pwd != "" {
		cmd = append(cmd, "--pwd", opts.Pwd)
	}

	var image string
	if !strings.HasPrefix(instance, "instance://") {
		image = strings.Join([]string{"instance://", instance}, "")
	} else {
		image = instance
	}

	cmd = append(cmd, image)
	cmd = append(cmd, command...)

	finalcmd := []string{"bash", "-c"}
	// adding these is causing SEGFAULTS
	finalcmd = append(finalcmd, opts.processEnvVars()...)
	finalcmd = append(finalcmd, opts.processPathMod()...)
	finalcmd = append(finalcmd, cmd...)

	fmt.Println(finalcmd)

	stdout, stderr, status, err := runCommand(finalcmd, &runCommandOptions{
		sudo:     c.Sudo,
		quieterr: opts.Quiet,
		quietout: opts.Quiet,
	})
	// TODO: use status
	_ = status
	if err != nil {
		return string(stdout.Bytes()), string(stderr.Bytes()), -1, err
	}

	return string(stdout.Bytes()), string(stderr.Bytes()), 0, nil

}

func (opts *ExecOptions) processEnvVars() []string {
	if len(opts.EnvVars) < 1 {
		return []string{}
	}
	envVarStrings := []string{}
	for k, v := range opts.EnvVars {
		envVarStrings = append(envVarStrings, fmt.Sprintf("%s%s=%s", senv, k, v))
	}
	return envVarStrings
}

func (opts *ExecOptions) processPathMod() []string {
	if len(opts.PrependPath) < 1 || len(opts.AppendPath) < 1 {
		return []string{}
	}
	pathVars := []string{}
	for _, prepath := range opts.PrependPath {
		pathVars = append(pathVars, fmt.Sprintf("%sPREPEND_PATH=%s", senv, prepath))
	}
	for _, apppath := range opts.AppendPath {
		pathVars = append(pathVars, fmt.Sprintf("%sAPPEND_PATH=%s", senv, apppath))
	} 
	return pathVars
}