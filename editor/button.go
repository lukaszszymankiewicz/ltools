package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

const (
	BORDER_SIZE   = 2
	BUTTON_WIDTH  = 30
	BUTTON_HEIGHT = 30
)

type Button struct {
	name       string
	left_line  FilledRectElement
	up_line    FilledRectElement
	down_line  FilledRectElement
	right_line FilledRectElement
	Area       FilledRectElement
	content    Elementer
	state      *int
}

func NewButton(name string, x int, y int, img *ebiten.Image, state *int) *Button {
	var b Button

	b.name = name
	width, height := img.Size()
	b.content = NewImageElement(
		"Icon",
		x+int((BUTTON_WIDTH-width)/2),
		y+int((BUTTON_HEIGHT-height)/2),
		img,
	)

	b.left_line = NewFilledRectElement("LeftLine", x, y, BORDER_SIZE, BUTTON_HEIGHT, whiteColor)
	b.up_line = NewFilledRectElement("UpLine", x, y, BUTTON_WIDTH, BORDER_SIZE, whiteColor)
	b.down_line = NewFilledRectElement("DownLine", x, y+BUTTON_HEIGHT-BORDER_SIZE, BUTTON_WIDTH, BORDER_SIZE, darkGreyColor)
	b.right_line = NewFilledRectElement("RightLine", x+BUTTON_WIDTH-BORDER_SIZE, y, BORDER_SIZE, BUTTON_HEIGHT, darkGreyColor)
	b.Area = NewFilledRectElement("Background", x, y, BUTTON_WIDTH, BUTTON_HEIGHT, midGreyColor)
	b.state = state

	return &b
}

func NewButtonWithText(name string, x int, y int, text string, state *int) *Button {
	var b Button

	b.name = name
	b.content = NewTextElement(text, x, y, 80, 30, greyColor)

	width := b.content.Area().Max.X - b.content.Area().Min.X
	height := b.content.Area().Max.Y - b.content.Area().Min.Y

	b.left_line = NewFilledRectElement("LeftLine", x, y, BORDER_SIZE, height, whiteColor)
	b.up_line = NewFilledRectElement("UpLine", x, y, width, BORDER_SIZE, whiteColor)
	b.down_line = NewFilledRectElement("DownLine", x, y+height-BORDER_SIZE, width, BORDER_SIZE, darkGreyColor)
	b.right_line = NewFilledRectElement("RightLine", x+width-BORDER_SIZE, y, BORDER_SIZE, height, darkGreyColor)
	b.Area = NewFilledRectElement("Background", x, y, width, height, greyColor)

	b.state = state

	return &b
}

func (b *Button) Draw(screen *ebiten.Image) {
	b.Area.Draw(screen)
	b.content.Draw(screen)
	b.left_line.Draw(screen)
	b.up_line.Draw(screen)
	b.down_line.Draw(screen)
	b.right_line.Draw(screen)
}

func (b *Button) SetActive() {
	b.left_line.color = blackColor
	b.up_line.color = blackColor
	b.down_line.color = greyColor
	b.right_line.color = greyColor
	b.Area.color = whiteColor
}

func (b *Button) SetUnactive() {
	b.left_line.color = greyColor
	b.up_line.color = greyColor
	b.down_line.color = blackColor
	b.right_line.color = blackColor
	b.Area.color = greyColor
}

func (b *Button) Update(g *Game) {
	if b.state == nil || *b.state == 0 {
		b.SetUnactive()
	} else {
		b.SetActive()
	}
}

func (b *Button) Name() string {
	return b.name
}
