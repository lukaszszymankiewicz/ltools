package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestPenBrushFunc(t *testing.T) {
    image := ebiten.NewImage(100, 200)
	tile := NewTile(image, "name", false, 0, 0, 0)

    brush_result := PenBrushFunc(10, 13, 1, 1, &tile)

    if brush_result.n != 1 {
        t.Errorf("Brush Result of Pen Brush is invalid, expected %d results, got %d", 1, brush_result.n)
    }

    if brush_result.coords[0][0] != 10 || brush_result.coords[0][1] != 13 {
        t.Errorf(
            "Brush Result of Pen Brush is invalid, expected (10, 13) coords, got (%d %d)", brush_result.coords[0][0], brush_result.coords[0][1])
    }

    if brush_result.tiles[0] != &tile {
        t.Errorf("Brush Result of Pen Brush is invalid, tile is wrong")
    }
}


func TestRubberBrushFunc(t *testing.T) {
    image := ebiten.NewImage(100, 200)
	tile := NewTile(image, "name", false, 0, 0, 0)

    brush_result := RubberBrushFunc(10, 13, 1, 1, &tile)

    if brush_result.n != 1 {
        t.Errorf("Brush Result of Pen Brush is invalid, expected %d results, got %d", 1, brush_result.n)
    }

    if brush_result.coords[0][0] != 10 || brush_result.coords[0][1] != 13 {
        t.Errorf(
            "Brush Result of Pen Brush is invalid, expected (10, 13) coords, got (%d %d)", brush_result.coords[0][0], brush_result.coords[0][1])
    }

    if brush_result.tiles[0] != nil {
        t.Errorf("Brush Result of Pen Brush is invalid, tile is wrong")
    }
}
