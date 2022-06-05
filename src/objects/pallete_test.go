package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestFillPallete(t *testing.T) {
	FIRST_LAYER := 0

	// This image contains (5) * (5) tiles = 25 Tiles
	img := ebiten.NewImage(32*5, 32*5)

	// this pallete can handle 10*10 Tiles (so it will be filled by 25%)
	p := NewPallete(0, 0, 5*32, 5*32, 32, 10, 10, 1)

	// we fill only first layer
	tileset := Tileset{0, "dupaa", img}
	p.FillPallete(tileset, FIRST_LAYER)

	// in this case all pallete is filled and all row and col should have some *Tile in it
	for n := 0; n < 25; n++ {
		row := int(n / 10)
		col := n % 10
		tile := p.GetTileOnDrawingArea(row, col, FIRST_LAYER)

		if tile == nil {
			t.Errorf("Tile creation does not succeeded on tile: %d, row:%d, col:%d", n, row, col)
		}
	}
}
