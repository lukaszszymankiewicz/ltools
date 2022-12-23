package editor

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

type TilesCollection struct {
	tiles    [][]*Tile
	n_layers int
}

func NewTilesCollection(n_layers int) TilesCollection {
	var tc TilesCollection

	tc.tiles = make([][]*Tile, n_layers)

	for layer := 0; layer < tc.n_layers; layer++ {
		tc.tiles[layer] = make([]*Tile, 0)
	}

	tc.n_layers = n_layers

	return tc
}

func (tc *TilesCollection) Fill(name string, layer int) {
	img := resources[name]

	filename := GetResourceFullName(name)

	if img == nil {
		fmt.Printf("something went wrong while reading the resources\n")
	}

	width, height := img.Size()
	rows := int(width / GridSize)
	cols := int(height / GridSize)

	n := 0

	for y := 0; y < cols; y++ {
		for x := 0; x < rows; x++ {
			r := image.Rect(x*GridSize, y*GridSize, x*GridSize+GridSize, y*GridSize+GridSize)

			// despite ebiten.Image.SubImage returning another ebiten.Image, type assertion is still
			// needed.
			i := img.SubImage(r)
			i2, err := i.(*ebiten.Image)
			if err != false {
				fmt.Errorf("something strange happen while dividing Tileset")
			}

			t := NewTile(i2, filename, x, y, layer, y*rows + x)

			tc.tiles[layer][y*rows+x] = &t

			n++
		}
	}
}

func (tc *TilesCollection) TilesPerLayer(i int) int {
	return len(tc.tiles[i])
}

func (tc *TilesCollection) Alloc(n int) {
	for layer := 0; layer < tc.n_layers; layer++ {
		tc.tiles[layer] = make([]*Tile, n)

		for i := 0; i < n; i++ {
			tc.tiles[layer][i] = nil
		}
	}
}

func (tc *TilesCollection) Get(idx int, layer int) *Tile {
	return tc.tiles[layer][idx]
}

func (tc *TilesCollection) Set(idx int, layer int, tile *Tile) {
	tc.tiles[layer][idx] = tile
}

func (tc *TilesCollection) Empty(idx int) bool {
	empty := true

	for layer := 0; layer < tc.n_layers; layer++ {
		if tc.Get(idx, layer) != nil {
			empty = false
		}
	}
	return empty
}

func (tc *TilesCollection) HasTile(tile *Tile) int {
	for idx := 0; idx < len(tc.tiles[tile.layer]); idx++ {
		if tc.Get(idx, tile.layer) != nil {
			return idx
		}
	}

	return -1
}
