package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Tileset struct {
	id   int
	name string
	img  *ebiten.Image
}

type Tilesets struct {
	tilesets []Tileset
}

func NewTileset(id int, name string, img *ebiten.Image) Tileset {
	var t Tileset

	t.id = id
	t.name = name
	t.img = img

	return t
}

func NewTilesets(images []string) Tilesets {
	var tl Tilesets

	tl.tilesets = make([]Tileset, 0)

	for i, name := range images {
		tileset := NewTileset(i, name, LoadImage(name))
		tl.tilesets = append(tl.tilesets, tileset)
	}
	return tl
}

func (tl *Tilesets) AvailableTilesets() int {
	return len(tl.tilesets)
}

func (tl *Tilesets) GetById(i int) Tileset {
	return tl.tilesets[i]
}
