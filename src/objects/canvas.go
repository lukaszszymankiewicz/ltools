package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Canvas struct {
	Grid
}

// this show how to set alpha channel for every layer if given layer is on
var layers_alpha = [][]float64{
	{1.0, 1.0, 1.0}, // LAYER_DRAW is on
	{0.5, 1.0, 0.5}, // LAYER_LIGHT in on
	{0.5, 0.0, 1.0}, // LAYER_ENTITY in on
}

func (c *Canvas) Draw(screen *ebiten.Image, layer int) {
	// drawing the border around Canvas
	c.FilledRectElement.Draw(screen)

	// drawing Tiles and tiles on it
	for x := 0; x < c.viewportCols; x++ {
		for y := 0; y < c.viewportRows; y++ {
			for l := 0; l < c.n_layers; l++ {

				tile := c.GetTileOnDrawingArea(x+c.viewport_x, y+c.viewport_y, l)

				pos_x := c.rect.Min.X + (x * c.grid_size)
				pos_y := c.rect.Min.Y + (y * c.grid_size)

				if tile != nil {
					tile.SingleImageBasedElement.Draw(screen, pos_x, pos_y, layers_alpha[layer][l])
				}
			}
		}
	}

	// drawing the scrollers
	c.scroller_x.Draw(screen)
	c.scroller_y.Draw(screen)
}

// creates new canvas struct
func NewCanvas(
	x int,
	y int,
	width int,
	height int,
	grid_size int,
	rows int,
	cols int,
	n_layers int,
) Canvas {
	var c Canvas

	c.Grid = NewGrid(x, y, width, height, grid_size, rows, cols, n_layers, greyColor)

	c.UpdateXScroller()
	c.UpdateYScroller()

	return c
}

func (c *Canvas) EraseOneRecord(record Record) {
	for i := len(record.x_coords) - 1; i > -1; i-- {
		c.SetTileOnDrawingArea(
			record.x_coords[i],
			record.y_coords[i],
			record.layers[i],
			record.old_tiles[i],
		)
	}
}

func (c *Canvas) MoveCanvasDown(screen *ebiten.Image) {
	c.MoveViewport(0, 1)
}

func (c *Canvas) MoveCanvasUp(screen *ebiten.Image) {
	c.MoveViewport(0, -1)
}

func (c *Canvas) MoveCanvasRight(screen *ebiten.Image) {
	c.MoveViewport(1, 0)
}

func (c *Canvas) MoveCanvasLeft(screen *ebiten.Image) {
	c.MoveViewport(-1, 0)
}
