package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	ImageElement
	file             string // file from which Tile is taken
	layer            int    // layer of tile
	row              int    // position of tile on tileset
	col              int    // position of tile on tileset
	unique_condition int    // index of unique condition of tile to be set
	set              bool
}

func NewTile(
	image *ebiten.Image,
	file string,
	row int,
	col int,
	layer int,
) Tile {
	var t Tile

	t.ImageElement = NewImageElement("Img", 0, 0, image)
	t.file = file
	t.row = row
	t.col = col
	t.layer = layer

	return t
}

func (t *Tile) GetImage() *ebiten.Image {
	return t.image
}

func (t *Tile) GetLayer() int {
	return t.layer
}

func (t *Tile) IsSet() bool {
	return t.set
}

func (t *Tile) PosOnTileset() (int, int) {
	return t.row, t.col
}
