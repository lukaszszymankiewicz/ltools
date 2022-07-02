package editor

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

type TilesCollection struct {
    tiles  [][]*Tile
    n_layers int
}

func NewTilesCollection(n_layers int) TilesCollection {
	var at TilesCollection

	at.tiles = make([][]*Tile, n_layers)
    at.n_layers = n_layers

	return at 
}

func (tc *TilesCollection) Fill(tileset Tileset, layer int) {
	width, height := tileset.img.Size()
	rows := int(width / GridSize)
	cols := int(height / GridSize)

	n := 0

	for y := 0; y < cols; y++ {
		for x := 0; x < rows; x++ {
			r := image.Rect(x*GridSize, y*GridSize, x*GridSize+GridSize, y*GridSize+GridSize)

			// despite ebiten.Image.SubImage returning another ebiten.Image, type assertion is still
			// needed.
			i := tileset.img.SubImage(r)
			i2, err := i.(*ebiten.Image)

			if err != false {
				fmt.Errorf("something strange happen while dividing Tileset")
			}

			t := NewTile(i2, tileset.name, false, x, y, layer)

            tc.tiles[layer] = append(tc.tiles[layer], &t)

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

        for i:=0; i<n; i++ {
            tc.tiles[layer][i] = nil
        }
    }
}

// func (tc *TilesCollection) PosHasTile(x int, y int, l int) bool {
// 	if at.GetTileOnDrawingArea(x, y, l) == nil {
// 		return false
// 	} else {
// 		return true
// 	}
// }
