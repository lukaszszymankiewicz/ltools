package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	ImageElement
	n       int    // number of tile is used on Canvas
	tileset string // tilseset from which Tile is taken
	unique  bool   // indication if Tile is unique
	layer   int    // layer of tile
	row     int    // position of tile on tileset
	col     int    // position of tile on tileset
	place_x int    // unique tile has its unique placement
	place_y int    // unique tile has its unique placement
}

// creates new Tile
func NewTile(
	image *ebiten.Image,
	tileset string,
	unique bool,
	row int,
	col int,
	layer int,
) Tile {
	var t Tile

	t.ImageElement = NewImageElement(0, 0, image)
	t.tileset = tileset
	t.unique = unique
	t.row = row
	t.col = col
	t.layer = layer
	t.place_x = -1
	t.place_y = -1

	return t
}

func (t *Tile) GetTileTileset() string {
	return t.tileset
}

func (t *Tile) GetImage() *ebiten.Image {
	return t.image
}

func (t *Tile) GetLayer() int {
	return t.layer
}

func (t *Tile) IsUnique() bool {
	return t.unique
}

func (t *Tile) IsSet() bool {
	if t.place_x == -1 && t.place_y == -1 {
		return false
	} else {
		return true
	}
}

func (t *Tile) Loc() (int, int) {
	return t.place_x, t.place_y
}

func (t *Tile) Put(x int, y int) {
	t.place_x = x
	t.place_y = y
}

func (t *Tile) PosOnTileset() (int, int) {
	return t.row, t.col
}
