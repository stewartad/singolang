package sing

import "github.com/stewartad/singolang/utils"

func initCommand(action string) []string {
	cmd := []string{"singularity", action}
	// append quiet or debug if flags are set in client 
	return cmd
}

func generateBindList(bindlist []string) {

}

func sendCommand() {

}

func runCommand(cmd []string, sudo bool) {
	utils.RunCommand(cmd, sudo)
}