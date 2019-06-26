package main

import (
	"fmt"
	// "log"
	// "os"
	// "path/filepath"
	"github.com/stewartad/singolang/client"
	"github.com/stewartad/singolang/utils"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())

	// instantiate a new client and defer its teardown function
	cl, finish := client.NewClient()
	defer finish(cl)
	// img := client.Pull("docker://godlovedc/lolcow", "", "", "")

	// create a new instance
	instanceError := cl.NewInstance("lolcow_latest.sif", "lolcow3")
	if instanceError != nil {
		fmt.Printf("%s\n", instanceError)
	}

	// Run some executes
	stdout, stderr, code, err := cl.Execute("lolcow3", []string{"which", "fortune"}, false)
	fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

	// This one is designed to fail
	stdout, stderr, code, err = cl.Execute("lolcow3", []string{"which", "singularity"}, false)
	fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

	stdout, stderr, code, err = cl.Execute("lolcow3", []string{"which", "lolcat"}, false)
	fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

	// List client's stored images
	cl.ListInstances()

	// List all running singularity images
	client.ListAllInstances()
}
