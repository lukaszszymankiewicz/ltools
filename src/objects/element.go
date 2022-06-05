package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

const (
	test_string = "abcdefghijklmnoprstuwyzABCEFGHIJKLMNOPRSTUWYZ1234567890"
)

// STRUCTS
type Element struct {
	rect image.Rectangle
}

type RectElement struct {
	Element
	color *color.RGBA // border
}

type FilledRectElement struct {
	Element
	color *color.RGBA //fill color
}

type ImageElement struct {
	Element
	image *ebiten.Image
	alpha float64
}

type TextElement struct {
	FilledRectElement
	content    string
	face       font.Face
	text_pos_x int
	text_pos_y int
}

// DRAWS
func (e RectElement) Draw(screen *ebiten.Image) {
	drawer.EmptyRect(screen, e.rect, e.color)
}

func (e FilledRectElement) Draw(screen *ebiten.Image) {
	drawer.FilledRect(screen, e.rect, e.color)
}

func (e ImageElement) Draw(screen *ebiten.Image) {
	x := e.rect.Min.X
	y := e.rect.Min.Y

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.Scale(1.0, 1.0, 1.0, e.alpha)

	screen.DrawImage(e.image, op)
}

func (e TextElement) Draw(screen *ebiten.Image) {
	e.FilledRectElement.Draw(screen)

	text.Draw(
		screen,
		e.content,
		e.face,
		e.text_pos_x,
		e.text_pos_y,
		blackColor,
	)
}

// CONSTRUCTORS
func NewRectElement(
	x int,
	y int,
	width int,
	height int,
	color *color.RGBA,
) RectElement {
	var e RectElement

	e.rect = image.Rect(x, y, x+width, y+height)
	e.color = color

	return e
}

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

func NewImageElement(
	x int,
	y int,
	img *ebiten.Image,
) ImageElement {
	var e ImageElement

	e.image = img
	width, height := e.image.Size()
	e.rect = image.Rect(x, y, x+width, y+height)
	e.alpha = 1.0

	return e
}

func NewTextElement(
	content string,
	x int,
	y int,
	width int,
	height int,
	bg_color *color.RGBA,
) TextElement {
	var e TextElement

	e.content = content
	e.face = normalFont

	text_rect := text.BoundString(e.face, content)
	teoretical_rect := text.BoundString(e.face, test_string)

	text_width := text_rect.Max.X - text_rect.Min.X
	text_height := teoretical_rect.Max.Y - teoretical_rect.Min.Y

	e.FilledRectElement = NewFilledRectElement(x, y, width, height, bg_color)

	// calculating text position on element (is centered horizontally and vertically)
	border_width := e.FilledRectElement.rect.Max.X - e.FilledRectElement.rect.Min.X
	border_height := e.FilledRectElement.rect.Max.Y - e.FilledRectElement.rect.Min.Y

	e.text_pos_x = (border_width+text_width)/2 + e.FilledRectElement.rect.Min.X - text_rect.Max.X
	e.text_pos_y = (border_height-text_height)/2 + e.FilledRectElement.rect.Min.Y - teoretical_rect.Min.Y

	return e
}

// UTIL
func (e Element) Area() image.Rectangle {
	return e.rect
}

func (e Element) Set(x int, y int) {
	e.rect.Min.X = x
	e.rect.Min.Y = y
}

func (e ImageElement) SetAlpha(alpha float64) {
	e.alpha = alpha
}
