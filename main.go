package main

import (
	"fmt"
	"log"
	"path/filepath"
	"github.com/stewartad/singolang/utils"
	"github.com/stewartad/singolang/sing"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())
	
	utils.RunCommand([]string{"rm", "-r", "./testdir"}, false)
	client := sing.GetClient()
	img := client.Pull("docker://godlovedc/lolcow", "", "", "~")
	out, err := utils.RunCommand([]string{"ls", "-l", filepath.Dir(img)}, false)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", out)
}