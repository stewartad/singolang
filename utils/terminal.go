package utils

import (
	//"unicode/utf8"
	"log"
	"os"
	"os/exec"
	"strings"
	"fmt"
	"io"
	"bytes"
	"sync"
)

// RunCommand runs a terminal command
/*	cmd is a slice of strings, 
	the first element of cmd must be the name of the command
	the remaining elements are arguments 
*/
func RunCommand(cmd []string, sudo bool, quiet bool) (string, string) {
	if sudo {
		cmd = append([]string{"sudo"}, cmd...)
	}
	name := cmd[0]
	// create command instance
	process := exec.Command(name, cmd[1:]...)

	var stdoutBuf, stderrBuf bytes.Buffer

	// get stdout and stderr
	stdout, err := process.StdoutPipe()
	if err != nil {
		fmt.Printf("Error getting stdout: %s\n", err)
	}
	stderr, err := process.StderrPipe()
	if err != nil {
		fmt.Printf("Error getting stderr: %s\n", err)
	}

	process.Start()

	var errStdout, errStderr error
	var outWriter, errWriter io.Writer
	if quiet {
		outWriter = io.Writer(&stdoutBuf)
		errWriter = io.Writer(&stderrBuf)
	} else {
		outWriter = io.MultiWriter(os.Stdout, &stdoutBuf)
		errWriter = io.MultiWriter(os.Stderr, &stderrBuf)
	}
	

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(outWriter, stdout)
		wg.Done()
	}()

	_, errStderr = io.Copy(errWriter, stderr)
	wg.Wait()

	err = process.Wait()
	if err != nil {
		log.Fatalf("Command failed with %s\n", err)
		// return nil, err
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalln("Failed to capture strout or stderr")
	}
	// return output
	return string(stdoutBuf.Bytes()), string(stderrBuf.Bytes());
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
	return parts[0], parts[1]
}

// RemoveURI strips the protocol and returns only the path
func RemoveURI(container string) string {
	_, path := SplitURI(container)
	return path
}