package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Cursor struct {
	SingleImageBasedElement
	size int
}

func NewCursor(size int) Cursor {
	var crs Cursor

	crs.SingleImageBasedElement = NewSingleImageBasedElement(nil)
	crs.SingleImageBasedElement.FloatingEmptyRectElement.color = whiteColor

	crs.size = size

	return crs
}

func (crs Cursor) DrawOnPallete(screen *ebiten.Image, x int, y int) {
	crs.FloatingEmptyRectElement.Draw(screen, x, y, x+crs.size, y+crs.size)
}

func (crs Cursor) DrawOnCanvas(screen *ebiten.Image, image *ebiten.Image, x int, y int) {
	crs.SingleImageBasedElement.image = image
	crs.SingleImageBasedElement.Draw(screen, x, y, 1.0)
}
