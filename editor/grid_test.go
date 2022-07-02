package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(10, 20, 320, 320, 32, 10, 10, 3, nil)

	if g.rect.Min.X != 10 {
		t.Errorf("Grid created inproperly: rect.Min.X = %d, want = %d", g.rect.Min.X, 10)
	}

	if g.rect.Min.Y != 20 {
		t.Errorf("Grid created inproperly: rect.Min.Y = %d, want = %d", g.rect.Min.Y, 20)
	}

	if g.rect.Max.X != 330 {
		t.Errorf("Grid created inproperly: rect.Max.X = %d, want = %d", g.rect.Max.X, 110)
	}

	if g.rect.Max.Y != 340 {
		t.Errorf("Grid created inproperly: rect.Max.Y = %d, want = %d", g.rect.Max.Y, 220)
	}

	if g.rows != 10 {
		t.Errorf("Grid created inproperly: rows = %d, want = %d", g.rows, 10)
	}

	if g.cols != 10 {
		t.Errorf("Grid created inproperly: cols = %d, want = %d", g.cols, 10)
	}

	if g.viewportCols != 10 {
		t.Errorf("Grid created inproperly: ViewportCols = %d, want = %d", g.viewportCols, 10)
	}

	if g.viewportRows != 10 {
		t.Errorf("Grid created inproperly: ViewportRows = %d, want = %d", g.viewportRows, 10)
	}

	if g.n_layers != 3 {
		t.Errorf("Grid created inproperly: n_layers = %d, want = %d", g.n_layers, 3)
	}

}

func TestGetTileOnDrawingArea(t *testing.T) {
	g := NewGrid(0, 0, 32*5, 32*5, 32, 5, 5, 3, nil)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, "name", false, 0, 0, 0)
	t2 := NewTile(i2, "name", false, 0, 1, 0)

	x1 := 3
	y1 := 3

	x2 := 2
	y2 := 2

	// adding some tiles by hand!
	g.drawingAreas[0][x1*g.rows+y1] = &t1
	g.drawingAreas[1][x2*g.rows+y2] = &t2

	if g.GetTileOnDrawingArea(x1, y1, 0) == nil {
		t.Errorf("getting tile on grid does not succeeded on (%d, %d, %d)", x1, y1, 0)
	}

	if g.GetTileOnDrawingArea(x2, y2, 1) == nil {
		t.Errorf("getting tile on grid does not succeeded on (%d, %d, %d)", x2, y2, 1)
	}
}

func TestMoveViewport(t *testing.T) {
	g := NewGrid(0, 0, 32*5, 32*5, 32, 5, 5, 3, nil)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, "name", false, 0, 0, 0)
	t2 := NewTile(i2, "name", false, 0, 1, 0)

	x1 := 3
	y1 := 3

	x2 := 2
	y2 := 2

	g.SetTileOnDrawingArea(x1, y1, 0, &t1)
	g.SetTileOnDrawingArea(x2, y2, 1, &t2)

	// reading tiles by hand!
	if g.drawingAreas[0][x1*g.rows+y1] == nil {
		t.Errorf("setting tile on grid does not succeeded on (%d, %d, %d)", x1, y1, 0)
	}

	if g.drawingAreas[1][x2*g.rows+y2] == nil {
		t.Errorf("setting tile on grid does not succeeded on (%d, %d, %d)", x2, y2, 1)
	}
}

func TestSetTileOnDrawingArea(t *testing.T) {
	// for such grid viewport is eqaul to whole grid - no movement is available and viewport_x and
	// viewport_y will always be 0
	g := NewGrid(0, 0, 32*5, 32*5, 32, 5, 5, 3, nil)

	for x := -10; x < 10; x++ {
		for y := -10; y < 10; y++ {
			g.MoveViewport(x, y)

			if g.viewport_x != 0 || g.viewport_y != 0 {
				t.Errorf("moving viewport does not succeeded want: (0, 0) have (%d, %d)", g.viewport_x, g.viewport_y)
			}
		}
	}

	// such grid can be moved at max 5 tile at each direction
	g = NewGrid(0, 0, 32*5, 32*5, 32, 10, 10, 3, nil)

	g.MoveViewport(10, 0)
	if g.viewport_x != 5 || g.viewport_y != 0 {
		t.Errorf("moving viewport does not succeeded want: (5, 0) have (%d, %d)", g.viewport_x, g.viewport_y)
	}

	g.MoveViewport(0, 10)
	if g.viewport_x != 5 || g.viewport_y != 5 {
		t.Errorf("moving viewport does not succeeded want: (5, 5) have (%d, %d)", g.viewport_x, g.viewport_y)
	}

	g.MoveViewport(-10, 0)
	if g.viewport_x != 0 || g.viewport_y != 5 {
		t.Errorf("moving viewport does not succeeded want: (0, 5) have (%d, %d)", g.viewport_x, g.viewport_y)
	}

	g.MoveViewport(0, -10)
	if g.viewport_x != 0 || g.viewport_y != 0 {
		t.Errorf("moving viewport does not succeeded want: (0, 0) have (%d, %d)", g.viewport_x, g.viewport_y)
	}

	for x := 1; x < 6; x++ {
		g.MoveViewport(1, 0)
		if g.viewport_x != x || g.viewport_y != 0 {
			t.Errorf("moving viewport does not succeeded want: (%d, 0) have (%d, %d)", x, g.viewport_x, g.viewport_y)
		}
	}
}

func TestMousePosToTile(t *testing.T) {
	g := NewGrid(0, 0, 32*5, 32*5, 32, 5, 5, 3, nil)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, "name", false, 0, 0, 0)
	t2 := NewTile(i2, "name", false, 0, 1, 0)

	g.SetTileOnDrawingArea(1, 1, 0, &t1)
	g.SetTileOnDrawingArea(2, 2, 0, &t2)
	g.SetTileOnDrawingArea(3, 3, 0, &t2)
	g.SetTileOnDrawingArea(4, 4, 1, &t1)

	// all mouse selection is in pixel!
	if g.MousePosToTile(10, 10, 0) != nil {
		t.Errorf("Tile selected by mouse is incorrect!")
	}

	if g.MousePosToTile(g.grid_size+10, g.grid_size+10, 0) != &t1 {
		t.Errorf("Tile selected by mouse is incorrect!")
	}

	if g.MousePosToTile(2*g.grid_size+10, 2*g.grid_size+10, 0) != &t2 {
		t.Errorf("Tile selected by mouse is incorrect!")
	}

	if g.MousePosToTile(3*g.grid_size+10, 3*g.grid_size+10, 0) != &t2 {
		t.Errorf("Tile selected by mouse is incorrect!")
	}

	if g.MousePosToTile(4*g.grid_size+10, 4*g.grid_size+10, 1) != &t1 {
		t.Errorf("Tile selected by mouse is incorrect!")
	}
}
