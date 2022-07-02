package editor

import (
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20, 320, 320, 32, 100, 100, 3)

	for l := 0; l < c.n_layers; l++ {
		for x := 0; x < c.cols; x++ {
			for y := 0; y < c.rows; y++ {
				if c.drawingAreas[l][x*c.rows+y] != nil {
					t.Errorf("Cleaning whole canvas as start does not succeeded on pos: (%d, %d, %d)", x, y, l)
				}
			}
		}
	}
}
