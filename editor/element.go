package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"image/color"
	_ "image/png"
)

const (
	test_string = "abcdefghijklmnoprstuwyzABCEFGHIJKLMNOPRSTUWYZ1234567890"
)

const (
	POPUP_POS_CENTER = iota
)

type Elementer interface {
	Update(*Game)
	Draw(*ebiten.Image)
	Name() string
	Area() image.Rectangle
}

// STRUCTS
type Element struct {
	Rect image.Rectangle
	name string
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

type PopupElement struct {
	FilledRectElement
	content    string
	name       string
	face       font.Face
	text_pos_x int
	text_pos_y int
}

// DRAWS
func (e RectElement) Draw(screen *ebiten.Image) {
	EmptyRect(screen, e.Rect, e.color)
}

func (e FilledRectElement) Draw(screen *ebiten.Image) {
	FilledRect(screen, e.Rect, e.color)
}

func (e ImageElement) Draw(screen *ebiten.Image) {
	x := e.Rect.Min.X
	y := e.Rect.Min.Y

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorM.Scale(1.0, 1.0, 1.0, e.alpha)

	screen.DrawImage(e.image, op)
}

func (e PopupElement) Draw(screen *ebiten.Image) {
	e.FilledRectElement.Draw(screen)

	text.Draw(
		screen,
		e.content,
		e.face,
		e.text_pos_x,
		e.text_pos_y,
		blackColor,
	)

	EmptyBorder(screen, e.FilledRectElement.Rect, blackColor)
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

	e.Rect = image.Rect(x, y, x+width, y+height)
	e.color = color

	return e
}

func NewFilledRectElement(
	name string,
	x int,
	y int,
	width int,
	height int,
	color *color.RGBA,
) FilledRectElement {
	var e FilledRectElement

	e.name = name
	e.Rect = image.Rect(x, y, x+width, y+height)
	e.color = color

	return e
}

func NewImageElement(
	name string,
	x int,
	y int,
	img *ebiten.Image,
) ImageElement {
	var e ImageElement

	e.name = name
	e.image = img
	width, height := e.image.Size()
	e.Rect = image.Rect(x, y, x+width, y+height)
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

	e.FilledRectElement = NewFilledRectElement("Background", x, y, width, height, bg_color)

	// calculating text position on element (is centered horizontally and vertically)
	border_width := e.FilledRectElement.Rect.Max.X - e.FilledRectElement.Rect.Min.X
	border_height := e.FilledRectElement.Rect.Max.Y - e.FilledRectElement.Rect.Min.Y

	e.text_pos_x = (border_width+text_width)/2 + e.FilledRectElement.Rect.Min.X - text_rect.Max.X
	e.text_pos_y = (border_height-text_height)/2 + e.FilledRectElement.Rect.Min.Y - teoretical_rect.Min.Y

	return e
}

func NewPopupElement(
	name string,
	content string,
	pos int,
	pad_width int,
	pad_height int,
	bg_color *color.RGBA,
) PopupElement {
	var e PopupElement
	var x int
	var y int

	e.content = content
	e.name = name
	e.face = normalFont

	text_rect := text.BoundString(e.face, content)
	teoretical_rect := text.BoundString(e.face, test_string)

	text_width := text_rect.Max.X - text_rect.Min.X
	text_height := teoretical_rect.Max.Y - teoretical_rect.Min.Y

	// placing popup at dummy position
	if pos == POPUP_POS_CENTER {
		x = (ScreenWidth - text_width - pad_width) / 2
		y = (ScreenHeight - text_height - pad_height) / 2
	}
	e.FilledRectElement = NewFilledRectElement(
		"Background",
		x,
		y,
		text_width+pad_width,
		text_height+pad_height,
		bg_color,
	)

	// calculating text position on element (is centered horizontally and vertically)
	border_width := e.FilledRectElement.Rect.Max.X - e.FilledRectElement.Rect.Min.X
	border_height := e.FilledRectElement.Rect.Max.Y - e.FilledRectElement.Rect.Min.Y

	e.text_pos_x = (border_width+text_width)/2 + e.FilledRectElement.Rect.Min.X - text_rect.Max.X
	e.text_pos_y = (border_height-text_height)/2 + e.FilledRectElement.Rect.Min.Y - teoretical_rect.Min.Y

	return e
}

// UTIL
func (e Element) Area() image.Rectangle {
	return e.Rect
}

func (e Element) Name() string {
	return e.name
}

func (e Element) Update(g *Game) {}

func (e Element) Set(x int, y int) {
	e.Rect.Min.X = x
	e.Rect.Min.Y = y
}

func (e ImageElement) SetAlpha(alpha float64) {
	e.alpha = alpha
}
