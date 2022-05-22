package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Tileset struct {
	tilesets map[string]*ebiten.Image
}

func NewTileset(images []string) Tileset {
	var tl Tileset

	tl.tilesets = make(map[string]*ebiten.Image)

	for _, image := range images {
		tl.tilesets[image] = LoadImage(image)
	}

	return tl
}

func (tl *Tileset) AvailableTilesetsImage() map[string]*ebiten.Image {
	return tl.tilesets
}
