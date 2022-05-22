package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

// STRUCTS
type Element struct {
	rect image.Rectangle
}

type FilledRectElement struct {
	Element
	color *color.RGBA
}

type EmptyRectElement struct {
	Element
	color *color.RGBA
}

type FloatingEmptyRectElement struct {
	Element
	color *color.RGBA
}

type SingleImageBasedElement struct {
	FloatingEmptyRectElement
	image *ebiten.Image
}

type ImageBasedElement struct {
	FilledRectElement
	images        []*ebiten.Image
	current_image int
}

// DRAWS
func DrawImageBasedElement(e ImageBasedElement, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.rect.Min.X), float64(e.rect.Min.Y))
	screen.DrawImage(e.images[e.current_image], op)
}

func (e ImageBasedElement) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.rect.Min.X), float64(e.rect.Min.Y))
	screen.DrawImage(e.images[e.current_image], op)
}

func (e FilledRectElement) Draw(screen *ebiten.Image) {
	drawer.FilledRect(screen, e.rect, e.color)
}

func (e EmptyRectElement) Draw(screen *ebiten.Image) {
	drawer.EmptyBorder(screen, e.rect, e.color)
}

func (e FloatingEmptyRectElement) Draw(screen *ebiten.Image, x1 int, y1 int, x2 int, y2 int) {
	rect := image.Rect(x1, y1, x2, y2)
	drawer.EmptyBorder(screen, rect, e.color)
}

func (e SingleImageBasedElement) Draw(screen *ebiten.Image, x int, y int, alpha float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	screen.DrawImage(e.image, op)
}

// CONSTRUCTORS
func NewFilledRectElement(
	x int,
	y int,
	width int,
	height int,
	color *color.RGBA,
) FilledRectElement {
	var e FilledRectElement

	e.rect = image.Rect(x, y, x+width, y+height)
	e.color = color

	return e
}

func NewEmptyRectElement(
	x int,
	y int,
	width int,
	height int,
	color *color.RGBA,
) EmptyRectElement {
	var e EmptyRectElement

	e.rect = image.Rect(x, y, x+width, y+height)
	e.color = color

	return e
}

func NewFloatingEmptyRectElement(color *color.RGBA) FloatingEmptyRectElement {
	var e FloatingEmptyRectElement
	e.color = color
	return e
}

func NewImageBasedElement(
	x int,
	y int,
	images []*ebiten.Image,
) ImageBasedElement {
	var e ImageBasedElement

	e.images = images
	width, height := e.images[e.current_image].Size()
	e.FilledRectElement.rect = image.Rect(x, y, x+width, y+height)

	return e
}

func NewSingleImageBasedElement(image *ebiten.Image) SingleImageBasedElement {
	var e SingleImageBasedElement

	e.image = image

	return e
}

func (e Element) Area() image.Rectangle {
	return e.rect
}
