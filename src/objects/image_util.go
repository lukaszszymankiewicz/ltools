package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

// loads image from given path
func loadImage(path string) *ebiten.Image {
	var err error
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
