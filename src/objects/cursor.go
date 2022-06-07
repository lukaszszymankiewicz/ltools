package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Cursor struct {
	ImageElement
	RectElement
	size int
}

func NewCursor(size int) Cursor {
	var crs Cursor

	dummy := ebiten.NewImage(32, 32)
	crs.ImageElement = NewImageElement(0, 0, dummy)
	crs.RectElement = NewRectElement(0, 0, size, size, blackColor)

	crs.size = size

	return crs
}

func (crs Cursor) DrawOnPallete(screen *ebiten.Image, x int, y int) {
	crs.RectElement.rect.Min.X = x
	crs.RectElement.rect.Min.Y = y
	crs.RectElement.rect.Max.X = x + crs.size
	crs.RectElement.rect.Max.Y = y + crs.size

	crs.RectElement.Draw(screen)
}

func (crs Cursor) DrawIconOnCanvas(screen *ebiten.Image, image *ebiten.Image, x int, y int) {
	_, height := image.Size()

	crs.ImageElement.rect.Min.X = x
	crs.ImageElement.rect.Min.Y = y - height
	crs.ImageElement.image = image

	crs.ImageElement.Draw(screen)
}

func (crs Cursor) DrawPlaceOnCanvas(screen *ebiten.Image, x int, y int) {
	crs.RectElement.rect.Min.X = x
	crs.RectElement.rect.Min.Y = y
	crs.RectElement.rect.Max.X = x + crs.size
	crs.RectElement.rect.Max.Y = y + crs.size

	crs.RectElement.Draw(screen)
}
