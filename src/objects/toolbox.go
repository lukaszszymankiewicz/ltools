package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	_ "image/png"
)

type Fill struct {
	SingleImageBasedElement       // image itselt
	tile                    *Tile // pointer to whole Tile
}

type Toolbox struct {
	x int
	y int
	Fill
}

func NewToolbox(x int, y int) Toolbox {
	var tb Toolbox
	tb.x = x
	tb.y = y

	tb.Fill.SingleImageBasedElement = NewSingleImageBasedElement(ebiten.NewImage(32, 32))

	// setting dummy image as starting fill
	dummy_image := ebiten.NewImage(32, 32)
	dummy_image.Fill(color.RGBA{0xff, 0, 0, 0xff})
	tb.SetFill(dummy_image)

	return tb
}

func (tb *Toolbox) Draw(screen *ebiten.Image) {
	tb.Fill.Draw(screen, tb.x, tb.y, 1.0)
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
