package main

import (
	"archive/tar"
	"fmt"
	"log"
	"os"
	"bytes"
	_"path"
	"path/filepath"
	"io"
	"io/ioutil"
	"compress/gzip"
	"github.com/stewartad/singolang/client"
	"github.com/stewartad/singolang/utils"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println("hello world")
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

	opts := client.DefaultExecOptions()

	// Run some executes
	stdout, stderr, code, err := cl.Execute("lolcow3", []string{"which", "fortune"}, opts)
	fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

	// This one is designed to fail
	stdout, stderr, code, err = cl.Execute("lolcow3", []string{"which", "singularity"}, opts)
	fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

	stdout, stderr, code, err = cl.Execute("lolcow3", []string{"which", "lolcat"}, opts)
	fmt.Printf("%s\n%s\n%d\t%s\n", stdout, stderr, code, err)

	tarDir := "/usr/games/cowsay"
	t, dir, err := cl.CopyTarball("lolcow3", tarDir)
	if err != nil {
		panic(fmt.Sprintf("Error creating tar: %s", err))
	}

	fmt.Println(t)
	file, err := os.Stat(t)
	if err != nil {
		panic("File not found")
	}
	fmt.Printf("%s - %d - %s\n\n", file.Name(), file.Size(), file.Mode())

	b, err := ioutil.ReadFile(t)
	if err != nil {
		panic(fmt.Sprintf("Could not read file %s", err))
	}
	gzr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		log.Fatalf("bad")
	}
	read := tar.NewReader(gzr)
	target := tarDir
	for {
		head, err := read.Next()
		if err == io.EOF {
			break
		}
		switch head.Typeflag {
		case tar.TypeDir, tar.TypeReg, tar.TypeLink, tar.TypeSymlink:
			fmt.Printf("%s -- %s\n", filepath.Clean(head.Name), filepath.Join(dir, filepath.Base(target)))
		default:
			continue
		}
		

	}

	// List client's stored images
	cl.ListInstances()

	// List all running singularity images
	client.ListAllInstances()
}
