package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

const (
	DUMMY_IMAGE_W = 10
	DUMMY_IMAGE_H = 10
)

// loads image from given path
func LoadImage(path string) *ebiten.Image {
	var err error
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		// if image cannot be read, dummy image is returned instead
		// log.Fatal(err)
		dummy_image := ebiten.NewImage(DUMMY_IMAGE_W, DUMMY_IMAGE_H)

		return dummy_image
	}

	return img
}
