package mainclient

import (
	"strings"
	"fmt"
	"os"
	"github.com/stewartad/singolang/utils"
)

// Execute runs a command inside a container
func (c *Client) Execute(image string, command string) string {
	// TODO: check install

	cmd := utils.InitCommand("exec")

	// --nv for graphics card drivers

	// use client's loaded image by default

	// if instance, use its uri

	// TODO: bind paths
	
	// TODO: run an app

	// TODO: sudo/writable

	splitCommand := strings.Split(command, " ")
	cmd = append(cmd, image)
	cmd = append(cmd, splitCommand...)

	out, err := utils.RunCommand(cmd, false, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %s\n", strings.Join(cmd, " "))
	}
	return out
	
}