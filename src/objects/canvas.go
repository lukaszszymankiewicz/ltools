package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Canvas struct {
	canvasRows   int
	canvasCols   int
	drawingArea  []int
	Rect         image.Rectangle
	viewportCols int
	viewportRows int
	viewport_x   int
	viewport_y   int
	scroller_x   Scroller
	scroller_y   Scroller
}

// clears whole Canvas from any Tile
func (c *Canvas) ClearDrawingArea() {
	for x := 0; x < c.canvasCols; x++ {
		for y := 0; y < c.canvasRows; y++ {
			c.drawingArea[x*c.canvasRows+y] = -1
		}
	}
}

// cleares only visible part of Canvas from any Tile
func (c *Canvas) ClearViewport() {
	for x := c.viewport_x; x < c.viewportRows+c.viewport_x; x++ {
		for y := c.viewport_y; y < c.viewportCols+c.viewport_y; y++ {
			c.drawingArea[x*c.canvasRows+y] = -1
		}
	}
}

// draws all Canvas parts:
// - Canvas border
// - actual Canvas (or rather its part visible by viewport)
// - x and y Scrollers
func (c *Canvas) DrawCanvas(screen *ebiten.Image, tiles []Tile) {
	// drawing the border around Canvas
	drawer.EmptyBorder(screen, c.Rect, color.White)

	// drawing Canvas and tiles on it
	for x := c.viewport_x; x < c.viewportCols+c.viewport_x; x++ {
		for y := c.viewport_y; y < c.viewportRows+c.viewport_y; y++ {
			if tile_index := c.GetTileOnCanvas(x, y); tile_index != -1 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64((x-c.viewport_x)*TileWidth+c.Rect.Min.X),
					float64((y-c.viewport_y)*TileWidth+c.Rect.Min.Y),
				)
				screen.DrawImage(tiles[tile_index].Image, op)
			}
		}
	}

	// drawing the scrollers
	c.DrawScrollers(screen)
}

// translates mouse position to Tile position where it will be drawn on Canvas (Canvas x and y)
func (c *Canvas) PosToTileCoordsOnCanvas(x int, y int) (int, int) {
	return int((x - c.Rect.Min.X) / TileWidth), int((y - c.Rect.Min.Y) / TileHeight)
}

// translates mouse position to Tile position where it will be drawn on Canvas (screen x and y)
func (c *Canvas) PosToTileHoveredOnCanvas(x int, y int) (int, int) {
	tileX, tileY := c.PosToTileCoordsOnCanvas(x, y)
	return tileX*TileWidth + c.Rect.Min.X, tileY*TileHeight + c.Rect.Min.Y
}

// returns Tile from Canvas by given position (Canvas x and y). Returned number is an
// index on stack.
func (c *Canvas) GetTileOnCanvas(x int, y int) int {
	return c.drawingArea[x*c.canvasRows+y]
}

// sets current Tile to draw (brush) on Canvas. Canvas x and y are used, and Tile is
// selected by its position on stack
func (c *Canvas) SetTileOnCanvas(x int, y int, value int) {
	c.drawingArea[(x+c.viewport_x)*c.canvasRows+(y+c.viewport_y)] = value
}

// moves viewport
func (c *Canvas) MoveCanvas(x int, y int) {
	new_x_value := c.viewport_x + x
	new_y_value := c.viewport_y + y

	c.viewport_x = MinVal(MaxVal(0, new_x_value), (c.canvasCols - c.viewportCols))
	c.viewport_y = MinVal(MaxVal(0, new_y_value), (c.canvasRows - c.viewportRows))

	// c.UpdateScrollers()
}

// creates new canvas struct
func NewCanvas(
	viewportRows int, viewportCols int,
	canvasRows int, canvasCols int,
	x_pos int, y_pos int,
) Canvas {
	var canvas Canvas

	canvas.viewportRows = viewportRows
	canvas.viewportCols = viewportCols
	canvas.canvasRows = canvasRows
	canvas.canvasCols = canvasCols
	canvas.drawingArea = make([]int, canvasRows*canvasCols)

	canvas.Rect = image.Rect(
		x_pos,
		y_pos,
		x_pos+TileWidth*viewportCols,
		y_pos+TileHeight*viewportRows,
	)

	canvas.scroller_x = NewScroller(
		x_pos,
		y_pos+TileHeight*viewportRows+2,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
		color.RGBA{96, 96, 96, 255},
		color.RGBA{192, 192, 192, 255},
	)

	canvas.scroller_y = NewScroller(
		x_pos+TileWidth*viewportCols+2,
		y_pos,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
		color.RGBA{96, 96, 96, 255},
		color.RGBA{192, 192, 192, 255},
	)

    // canvas.UpdateScrollers()

	return canvas
}
