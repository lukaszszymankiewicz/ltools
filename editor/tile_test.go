package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestNewTile(t *testing.T) {
	image := ebiten.NewImage(100, 200)
	tile := NewTile(image, "name", false, 0, 0, 0)

	if tile.unique != false {
		t.Errorf("Tile creation does not succeeded")
	}

	if tile.tileset != "name" {
		t.Errorf("Tile creation does not succeeded")
	}
}
