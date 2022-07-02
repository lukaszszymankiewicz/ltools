package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
)

const (
	border_size   = 2
	button_width  = 30
	button_height = 30
)

type Fill struct {
	ImageElement
	tile *Tile
}

type Tool struct {
	Brush
	// graphic elements
	left_line  FilledRectElement
	up_line    FilledRectElement
	down_line  FilledRectElement
	right_line FilledRectElement
	bg         FilledRectElement
	icon       ImageElement
}

type Toolbox struct {
	x int
	y int
	Fill
	Tools  []Tool          // functional tools
    pipette *ebiten.Image  // pipette icon - just image, non functional by now
    active int
}

func NewTool(x int, y int, img *ebiten.Image, brush Brush) Tool {
	var t Tool
	t.Brush = brush

	width, height := img.Size()
	t.icon = NewImageElement(
		x+int((button_width-width)/2),
		y+int((button_height-height)/2),
		img,
	)

	t.left_line = NewFilledRectElement(x, y, border_size, button_height, whiteColor)
	t.up_line = NewFilledRectElement(x, y, button_width, border_size, whiteColor)
	t.down_line = NewFilledRectElement(x, y+button_height-border_size, button_width, border_size, darkGreyColor)
	t.right_line = NewFilledRectElement(x+button_width-border_size, y, border_size, button_height, darkGreyColor)
	t.bg = NewFilledRectElement(x, y, button_width, button_height, midGreyColor)

	return t
}

func NewToolbox(x int, y int) Toolbox {
	var tb Toolbox
	tb.x = x
	tb.y = y

	tb.Tools = make([]Tool, 0)

	pen := NewTool(x, y, LoadImage("editor/assets/pen15.png"), PenBrush)
	rubber := NewTool(x+32, y, LoadImage("editor/assets/rubber.png"), RubberBrush)

	tb.Tools = append(tb.Tools, pen)
	tb.Tools = append(tb.Tools, rubber)

	tb.Fill.ImageElement = NewImageElement(0, 0, ebiten.NewImage(32, 32))
	dummy_image := ebiten.NewImage(32, 32)
	dummy_image.Fill(color.RGBA{0xff, 0, 0, 0xff})
	tb.SetFill(dummy_image)
    
    tb.pipette = LoadImage("editor/assets/pippete.png")

	tb.Activate(0)

	return tb
}

func (tb *Toolbox) Draw(screen *ebiten.Image) {
	if tb.active != 1 {
		tb.Fill.ImageElement.rect.Min.X = tb.x + 32*5
		tb.Fill.ImageElement.rect.Min.Y = tb.y
		tb.Fill.Draw(screen)
	}

	for _, tool := range tb.Tools {
		tool.Draw(screen)
	}
}

func (t *Tool) Draw(screen *ebiten.Image) {
	t.bg.Draw(screen)
	t.icon.Draw(screen)
	t.left_line.Draw(screen)
	t.up_line.Draw(screen)
	t.down_line.Draw(screen)
	t.right_line.Draw(screen)
}

func (tb *Toolbox) SetFill(image *ebiten.Image) {
	tb.Fill.ImageElement.image = image
}

func (tb *Toolbox) GetFill() *ebiten.Image {
	return tb.Fill.ImageElement.image
}

func (tb *Toolbox) GetCursor() *ebiten.Image {
	return tb.Tools[tb.active].icon.image
}

func (tb *Toolbox) GetFillTile() *Tile {
	return tb.Fill.tile
}

func (tb *Toolbox) SetFillTile(tile *Tile) {
	tb.Fill.tile = tile
	tb.Fill.image = tile.image
}

func (tb *Toolbox) Activate(i int) {
	tb.Tools[tb.active].SetUnactive()
	tb.active = i
	tb.Tools[tb.active].SetActive()
}

func (t *Tool) SetActive() {
	t.left_line.color = blackColor
	t.up_line.color = blackColor
	t.down_line.color = greyColor
	t.right_line.color = greyColor
	t.bg.color = whiteColor
}

func (t *Tool) SetUnactive() {
	t.left_line.color = whiteColor
	t.up_line.color = whiteColor
	t.down_line.color = darkGreyColor
	t.right_line.color = darkGreyColor
	t.bg.color = midGreyColor
}

func (tb *Toolbox) Area(i int) image.Rectangle {
	return tb.Tools[i].icon.rect
}

func (tb *Toolbox) GetTool(i int) Tool {
	return tb.Tools[i]
}

func (tb *Toolbox) GetPipette() *ebiten.Image {
	return tb.pipette
}
