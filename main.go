package main

import (
	"fmt"
	"path/filepath"
	"github.com/stewartad/singolang/utils"
	"github.com/stewartad/singolang/sing"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())
	
	client := sing.GetClient()
	img := client.Pull("docker://godlovedc/lolcow", "", "", "")
	utils.RunCommand([]string{"ls", "-l", filepath.Dir(img)}, false, false)
}