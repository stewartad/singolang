package main

import (
	"archive/tar"
	"fmt"
	"os"
	"path/filepath"
	"io"

	"github.com/stewartad/singolang/client"
)

func main() {
	fmt.Println("Singolang Version 0.0.1")
	fmt.Printf("Singularity Version: %s\n", client.GetSingularityVersion())

	execute := true
	archive := true

	// instantiate a new client and defer its teardown function
	cl, finish := client.NewClient()
	defer finish(cl)
	// img := client.Pull("docker://godlovedc/lolcow", "", "", "")

	// create a new instance
	instanceError := cl.NewInstance("lolcow_latest.sif", "lolcow3")
	if instanceError != nil {
		fmt.Printf("%s\n", instanceError)
	}

	opts := client.DefaultExecOptions()

	// Run some executes
	if execute {
		stdout, stderr, code, err := cl.Execute("lolcow3", []string{"which", "fortune"}, opts)
		fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

		// This one is designed to fail
		stdout, stderr, code, err = cl.Execute("lolcow3", []string{"which", "singularity"}, opts)
		fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

		stdout, stderr, code, err = cl.Execute("lolcow3", []string{"which", "lolcat"}, opts)
		fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

		stdout, stderr, code, err = cl.Execute("lolcow3", []string{"echo", "hello"}, opts)
		fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)
	}

	if archive {
		TarFile(cl, "lolcow3", "/usr/games/")
		TarFile(cl, "lolcow3", "/usr/games/cowsay")
	}
	
	// List client's stored images
	cl.ListInstances()

	// List all running singularity images
	client.ListAllInstances()
}

func TarFile(cl *client.Client, instance string, target string) {
	// need to make this work for both directories and files
	t, read, err := cl.CopyTarball(instance, target)
	if err != nil {
		panic(fmt.Sprintf("Error creating tar: %s", err))
	}
	defer os.RemoveAll(filepath.Dir(t))

	fmt.Println("-----------------")
	fmt.Println("DIRECTORY ARCHIVED")
	fmt.Println("-----------------")
	fmt.Println(t)
	file, err := os.Stat(t)
	if err != nil {
		panic("File not found")
	}
	fmt.Printf("%s - %d - %s\n", file.Name(), file.Size(), file.Mode())
	fmt.Println("-----------------")

	for {
		head, err := read.Next()
		if err == io.EOF {
			break
		}
		switch head.Typeflag {
		case tar.TypeDir, tar.TypeReg, tar.TypeLink, tar.TypeSymlink:
			fmt.Printf("%s -- %s\n", filepath.Clean(head.Name), filepath.Base(target))
		default:
			continue
		}
	}
	fmt.Println("-----------------")
}
