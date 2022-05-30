package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	_ "image/png"
)

const (
    border_size = 2
)

type Fill struct {
	SingleImageBasedElement       // image itselt
	tile                    *Tile // pointer to whole Tile
}

type Tool struct {
    left_line FilledRectElement 
    up_line FilledRectElement 
    down_line FilledRectElement 
    right_line FilledRectElement 
    icon ImageBasedElement2 
    state int
}

type Toolbox struct {
	x int
	y int
	Fill
    Tools []Tool
}

func NewTool(x int, y int, img *ebiten.Image) Tool {
    var t Tool

    t.icon = NewImageBasedElement2(img, x, y, greyColor) 
    width, height := t.icon.img.Size()

    t.left_line = NewFilledRectElement(x, y, border_size , height, whiteColor)
    t.up_line = NewFilledRectElement(x, y, width, border_size, whiteColor)
    t.down_line = NewFilledRectElement(x, y+height-border_size, width, border_size, darkGreyColor) 
    t.right_line = NewFilledRectElement(x+width-border_size, y, border_size, height, darkGreyColor) 

    return t
}

func NewToolbox(x int, y int) Toolbox {
	var tb Toolbox
	tb.x = x
	tb.y = y

	tb.Tools = make([]Tool, 0)

    pen := NewTool(x, y, LoadImage("src/objects/assets/buttons/pen15.png"))
    rubber := NewTool(x+32, y, LoadImage("src/objects/assets/buttons/rubber.png"))

    tb.Tools = append(tb.Tools, pen)
    tb.Tools = append(tb.Tools, rubber)

	tb.Fill.SingleImageBasedElement = NewSingleImageBasedElement(ebiten.NewImage(32, 32))
	dummy_image := ebiten.NewImage(32, 32)
	dummy_image.Fill(color.RGBA{0xff, 0, 0, 0xff})
	tb.SetFill(dummy_image)


	return tb
}

func (tb *Toolbox) Draw(screen *ebiten.Image) {
	tb.Fill.Draw(screen, tb.x+32*5, tb.y, 1.0)

    for _, tool := range tb.Tools {
        tool.Draw(screen)
    }
}

func (t *Tool)Draw(screen *ebiten.Image) {
    t.icon.FilledRectElement.Draw(screen)
    t.icon.Draw(screen)
    t.left_line.Draw(screen)
    t.up_line.Draw(screen)
    t.down_line.Draw(screen)
    t.right_line.Draw(screen)
}

func (tb *Toolbox) SetFill(image *ebiten.Image) {
	tb.Fill.SingleImageBasedElement.image = image
}

func (tb *Toolbox) GetFill() *ebiten.Image {
	return tb.Fill.SingleImageBasedElement.image
}

func (tb *Toolbox) GetFillTile() *Tile {
	return tb.Fill.tile
}

func (tb *Toolbox) SetFillTile(tile *Tile) {
	tb.Fill.tile = tile
	tb.Fill.image = tile.image
}
