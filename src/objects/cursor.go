package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
    "image/color"
)

type Cursor struct {
	ImageElement
	RectElement
	size int
    borders []*color.RGBA
}

func NewCursor(size int) Cursor {
	var crs Cursor

	dummy := ebiten.NewImage(32, 32)
	crs.ImageElement = NewImageElement(0, 0, dummy)
	crs.RectElement = NewRectElement(0, 0, size, size, blackColor)
	crs.size = size
    crs.borders = []*color.RGBA{blackColor, blackColor, whiteColor, blackColor, blackColor}

	return crs
}

// cursor on pallete is created from one thick rect, to draw it, simply few rects, each one smaller
// than another is drawn.
func (crs Cursor) DrawBorder(screen *ebiten.Image, x int, y int) {
    for i, clr := range(crs.borders) {
        crs.RectElement.rect.Min.X = x + i
        crs.RectElement.rect.Min.Y = y + i
        crs.RectElement.rect.Max.X = x + crs.size - i
        crs.RectElement.rect.Max.Y = y + crs.size - i
        crs.RectElement.color = clr
        crs.RectElement.Draw(screen)
    }
}

func (crs Cursor) DrawIcon(screen *ebiten.Image, image *ebiten.Image, x int, y int) {
	_, height := image.Size()

	crs.ImageElement.rect.Min.X = x
	crs.ImageElement.rect.Min.Y = y - height
	crs.ImageElement.image = image

	crs.ImageElement.Draw(screen)
}
