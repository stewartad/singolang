package utils

import (
	//"unicode/utf8"
	"log"
	"os"
	"os/exec"
	"strings"
	// "fmt"
	"io"
	"bytes"
	"sync"
)

// InitCommand creates a slice of strings of which the first element is "singularity" followed by all args
func InitCommand(args ...string) []string {
	cmd := []string{"singularity"}
	cmd = append(cmd, args...)
	// append quiet or debug if flags are set in client 
	return cmd
}

/*RunCommand runs a terminal command
	cmd - a slice of strings, of which the first element must be the command name
		all subsequent elements are the arguments
	sudo - set to True to run command as su
	quiet - set to True to not print stdout to the screen
		note that stderr will always print to screen
*/
func RunCommand(cmd []string, sudo bool, quiet bool) (string, error) {
	// add sudo to front of command if requested
	if sudo {
		cmd = append([]string{"sudo"}, cmd...)
	}
	name := cmd[0]

	// create command instance
	process := exec.Command(name, cmd[1:]...)

	// buffers to store stdout and stderr to be returned
	var stdoutBuf, stderrBuf bytes.Buffer

	// get stdout and stderr pipes
	stdout, err := process.StdoutPipe()
	if err != nil {
		log.Printf("Error getting stdout: %s\n", err)
	}
	stderr, err := process.StderrPipe()
	if err != nil {
		log.Printf("Error getting stderr: %s\n", err)
	}

	process.Start()

	var errStdout, errStderr error
	var outWriter, errWriter io.Writer
	if quiet {
		// only write stdout to buffer, bot to screen
		outWriter = io.Writer(&stdoutBuf)
	} else {
		// write stdout to both buffer and screen
		outWriter = io.MultiWriter(os.Stdout, &stdoutBuf)
	}
	// write stderr to both buffer and screen
	errWriter = io.MultiWriter(os.Stderr, &stderrBuf)

	// capture stdout concurrently
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(outWriter, stdout)
		wg.Done()
	}()

	// wait for stdout to be captured, then capture stderr
	_, errStderr = io.Copy(errWriter, stderr)
	wg.Wait()

	// wait for command to exit
	err = process.Wait()
	// handle erros
	if err != nil {
		log.Printf("Command failed with %s\n", err)
		return string(stderrBuf.Bytes()), err
	}
	if errStdout != nil || errStderr != nil {
		log.Printf("Failed to capture strout or stderr")
	}
	// return stdout and stderr as strings
	return string(stdoutBuf.Bytes()), err
}

// GetSingularityVersion gets installed Singularity version
func GetSingularityVersion() string {
	version, _ := RunCommand([]string{"singularity", "--version"}, false, true)
	return strings.TrimSpace(version)
}

// SplitURI splits the URI into protocol and path"
func SplitURI(container string) (string, string) {
	// Splits
	parts := strings.Split(container, "://")
	if len(parts) == 1 {
		return "", parts[0]
	}
	return parts[0], parts[1]
}

// RemoveURI strips the protocol and returns only the path
func RemoveURI(container string) string {
	_, path := SplitURI(container)
	return path
}