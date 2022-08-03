package editor

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"os"
	"strings"
)

const (
	ASSETS    = "assets"
	EXTENSION = ".png"
)

var resources map[string]*ebiten.Image
var resources_prefix string

func LoadResources(prefix string) {
	resources = make(map[string]*ebiten.Image)

	if prefix != "" {
		resources_prefix = prefix + "/"
	} else {
		resources_prefix = ""
	}

	files, _ := os.ReadDir(resources_prefix + ASSETS)

	if len(files) == 0 {
		fmt.Printf("desired path does not contain any files! %s\n", resources_prefix+ASSETS)
	}

	for _, file := range files {
		name := file.Name()
		filename := strings.Split(file.Name(), ".")[0]
		path := resources_prefix + ASSETS + "/" + name
		img := LoadImage(path)

		if img == nil {
			fmt.Printf("something went wrong while reading resource: %s \n", path)
		} else {
			resources[filename] = img
		}
	}
}

func GetResourceFullName(name string) string {
	return resources_prefix + ASSETS + "/" + name + EXTENSION
}
