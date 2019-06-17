package mainclient

import (
	"strings"
	"github.com/stewartad/singolang/utils"
)

// Execute runs a command inside a container
func (c *Client) Execute(image string, command string) string {
	// TODO: check install

	cmd := InitCommand("exec")

	// --nv for graphics card drivers

	// use client's loaded image by default

	// if instance, use its uri

	// TODO: bind paths
	
	// TODO: run an app

	// TODO: sudo/writable

	splitCommand := strings.Split(command, " ")
	cmd = append(cmd, image)
	cmd = append(cmd, splitCommand...)

	out, _ := utils.RunCommand(cmd, false, false)
	return out
	
}