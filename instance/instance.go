package instance

import (
	"github.com/stewartad/singolang/utils"
	"fmt"
)

type Instance struct {
	name string
	imageURI string
	protocol string
	image string
	metadata []string
}

func (i *Instance) String() string {
	if i.protocol != "" {
		return fmt.Sprintf("%s:\\%s", i.protocol, i.image)
	}
	return i.image
}

func GetInstance(image string) *Instance {
	i := new(Instance)
	i.ParseImageName(image)

	return i
}

func (i *Instance) ParseImageName(image string) {
	i.imageURI = image
	i.protocol, i.image = utils.SplitURI(image)
}

func (i *Instance) updateMetadata() {

}