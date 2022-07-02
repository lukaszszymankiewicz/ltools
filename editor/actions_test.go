package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestHasTile(t *testing.T) {
	tileset_name := "name"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := NewTile(i3, tileset_name, false, 0, 2, 0)

	ts := NewTileStack()

	ts.AddNew(&t1)
	ts.AddNew(&t2)

	if ts.hasTile(&t1) == false {
		t.Errorf("Adding tiles to stack failed")
	}

	if ts.hasTile(&t2) == false {
		t.Errorf("Adding tiles to stack failed")
	}

	if ts.hasTile(&t3) == true {
		t.Errorf("Adding tiles to stack failed")
	}

	ts.AddNew(&t3)

	if ts.hasTile(&t3) == false {
		t.Errorf("Adding tiles to stack failed")
	}
}

func TestPrepareTiles(t *testing.T) {
	var g Game

	g.Canvas = NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	tileset_name := "assets/basic_tileset.png"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := NewTile(i3, tileset_name, false, 0, 2, 1)

	g.Canvas.SetTileOnDrawingArea(10, 10, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(11, 11, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(12, 12, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(13, 13, 0, &t1)

	g.Canvas.SetTileOnDrawingArea(14, 14, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(15, 15, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(16, 16, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(17, 17, 0, &t2)

	g.Canvas.SetTileOnDrawingArea(5, 5, 1, &t3)
	g.Canvas.SetTileOnDrawingArea(6, 6, 1, &t3)

	stack := g.FillStack()
	g.PrepareTileset("sample_name", stack)
}
