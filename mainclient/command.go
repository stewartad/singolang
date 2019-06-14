package mainclient

import "github.com/stewartad/singolang/utils"

func InitCommand(args ...string) []string {
	cmd := []string{"singularity"}
	cmd = append(cmd, args...)
	// append quiet or debug if flags are set in client 
	return cmd
}

func generateBindList(bindlist []string) {

}

func sendCommand() {

}

func runCommand(cmd []string, sudo bool) {
	utils.RunCommand(cmd, sudo, false)
}