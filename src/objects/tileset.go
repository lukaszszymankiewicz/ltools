package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

type Tileset struct {
	image     *ebiten.Image
	width     int
	height    int
	rows      int
	cols      int
	Num       int
	tilesUsed int
}

func NewTileset(path string) Tileset {
	var tileset Tileset

	tileset.image = loadImage(path)
	width, height := tileset.image.Size()
	tileset.width = width
	tileset.height = height
	tileset.rows = tileset.width / TileWidth
	tileset.cols = tileset.height / TileHeight
	tileset.Num = tileset.rows * tileset.cols

	return tileset
}

func (t *Tileset) TileNrToSubImageOnTileset(tileNr int) *ebiten.Image {
	RectX, RectY := tileNr%t.cols*TileWidth, tileNr/t.rows*TileHeight
	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

	return t.image.SubImage(tileRect).(*ebiten.Image)
}
