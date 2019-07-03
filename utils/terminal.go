package utils

import (
	//"unicode/utf8"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
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

type RunCommandOptions struct {
	Sudo 		bool
	Quietout	bool
	Quieterr	bool
}

func DefaultRunCommandOptions() *RunCommandOptions {
	return &RunCommandOptions{
		Sudo: false,
		Quietout: false,
		Quieterr: false,
	}
}

/*RunCommand runs a terminal command
	cmd - a slice of strings, of which the first element must be the command name
		all subsequent elements are the arguments
	sudo - set to True to run command as su
	quiet - set to True to not print stdout or stderr to the screen
*/
func RunCommand(cmd []string, opts *RunCommandOptions) (bytes.Buffer, bytes.Buffer, int, error) {
	// add sudo to front of command if requested
	if opts.Sudo {
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
	if opts.Quietout {
		// only write stdout to buffer, not to screen
		outWriter = io.Writer(&stdoutBuf)
	} else {
		// write stdout to both buffer and screen
		outWriter = io.MultiWriter(os.Stdout, &stdoutBuf)
	}
	if opts.Quieterr {
		// only write stderr to buffer, not to screen
		errWriter = io.Writer(&stderrBuf)
	} else {
		// write stderr to both buffer and screen
		errWriter = io.MultiWriter(os.Stderr, &stderrBuf)
	}
	

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

	var waitStatus syscall.WaitStatus

	// wait for command to exit
	err = process.Wait()
	// handle erros
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			// log.Printf("Command failed with %s\n", err)
			return stdoutBuf, stderrBuf, waitStatus.ExitStatus(), err
		} 
		// log.Printf("Unknown Error, Exit Status: %s", err)
		return stdoutBuf, stderrBuf, -1, err
		
	}

	
	waitStatus = process.ProcessState.Sys().(syscall.WaitStatus)
	if errStdout != nil || errStderr != nil {
		log.Printf("Failed to capture strout or stderr")
	}
	// return stdout and stderr as strings
	return stdoutBuf, stderrBuf, waitStatus.ExitStatus(), err
}

// GetSingularityVersion gets installed Singularity version
func GetSingularityVersion() string {
	version, _, status, _ := RunCommand([]string{"singularity", "--version"}, DefaultRunCommandOptions())
	// log.Println(status)
	if status == 0 {
		return strings.TrimSpace(string(version.Bytes()))
	}
	return ""
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