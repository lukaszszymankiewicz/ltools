package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text"
	"image"
	"image/color"
    _ "image/png"
	"golang.org/x/image/font"
    "ltools/src/drawer"
)

const (
    test_string ="abcdefghijklmnoprstuwyzABCEFGHIJKLMNOPRSTUWYZ1234567890"
)

// STRUCTS
type Element struct {
	rect image.Rectangle
}

type FilledRectElement struct {
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

type ImageBasedElement2 struct {
	FilledRectElement
	img        *ebiten.Image
}

type TextElement struct {
	FilledRectElement
	content  string
    face font.Face
    text_pos_x int
    text_pos_y int
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

func (e FloatingEmptyRectElement) Draw(screen *ebiten.Image, x1 int, y1 int, x2 int, y2 int) {
	rect := image.Rect(x1, y1, x2, y2)
	drawer.EmptyBorder(screen, rect, e.color)
}

func (e ImageBasedElement2) Draw(screen *ebiten.Image) {
    x := e.FilledRectElement.rect.Min.X
    y := e.FilledRectElement.rect.Min.Y

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.Scale(1.0, 1.0, 1.0, 1.0)

	screen.DrawImage(e.img, op)
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

func NewImageBasedElement2(
    image *ebiten.Image,
    x int,
    y int,
    bg_color *color.RGBA,
) ImageBasedElement2 {
	var e ImageBasedElement2 

	e.img = image
    width, height := e.img.Size()
    e.FilledRectElement = NewFilledRectElement(x, y, width, height, bg_color)

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

    e.text_pos_x = (border_width + text_width) / 2 + e.FilledRectElement.rect.Min.X - text_rect.Max.X
    e.text_pos_y = (border_height - text_height) / 2 + e.FilledRectElement.rect.Min.Y - teoretical_rect.Min.Y

	return e
}

func (e TextElement)Draw(screen *ebiten.Image) {
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

func (e Element) Area() image.Rectangle {
	return e.rect
}
