package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
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

func loadImage(path string) *ebiten.Image {
	var err error
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
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
