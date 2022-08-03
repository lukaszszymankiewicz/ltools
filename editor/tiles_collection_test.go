package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestFill(t *testing.T) {
	LoadResources("")
	FIRST_LAYER := 0

	// This image contains (5) * (5) tiles = 25 Tiles
	resources["dummy"] = ebiten.NewImage(32*5, 32*5)

	// this pallete can handle 10*10 Tiles (so it will be filled by 25%)
	v := NewViewport(10, 10, 5, 5)
	tc := NewTilesCollection(3)
	tc.Alloc(10 * 10)

	tc.Fill("dummy", FIRST_LAYER)

	// in this case all pallete is filled and all row and col should have some *Tile in it
	for n := 0; n < 25; n++ {
		col := int(n / 10)
		row := n % 10
		idx := v.GetTileIdx(row, col)
		tile := tc.Get(idx, FIRST_LAYER)

		if tile == nil {
			t.Errorf("Tile creation does not succeeded on tile: %d, row:%d, col:%d", n, row, col)
		}
	}
}
