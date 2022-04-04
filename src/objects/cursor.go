package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

// cursor struct
type Cursor struct {
	size  int
	color color.Color
}

// creates new cursor object
func NewCursor(size int) Cursor {
	var crs Cursor

	crs.color = cursorColor
	crs.size = size

	return crs
}

// draws cursor on screen
func (crs Cursor) DrawCursor(screen *ebiten.Image, x int, y int) {
	rect := image.Rect(x, y, x+crs.size, y+crs.size)
	drawer.EmptyRect(screen, rect, crs.color)
}
