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
	clickableAreas map[image.Rectangle]func()
	hoverableAreas map[image.Rectangle]func()
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
    c.scroller_x.DrawScroller(screen)
    c.scroller_y.DrawScroller(screen)
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

// moves viewport left
func (c *Canvas) MoveCanvasLeft() {
    c.MoveCanvas(-1, 0)
}

// moves viewport right
func (c *Canvas) MoveCanvasRight() {
    c.MoveCanvas(1, 0)
}

// moves viewport up
func (c *Canvas) MoveCanvasUp() {
    c.MoveCanvas(0, -1)
}

// moves viewport down
func (c *Canvas) MoveCanvasDown() {
    c.MoveCanvas(0, 1)
}

// moves viewport
func (c *Canvas) MoveCanvas(x int, y int) {
	new_x_value := c.viewport_x + x
	new_y_value := c.viewport_y + y

	c.viewport_x = MinVal(MaxVal(0, new_x_value), (c.canvasCols - c.viewportCols))
	c.viewport_y = MinVal(MaxVal(0, new_y_value), (c.canvasRows - c.viewportRows))

	c.UpdateScrollers()
}

// returns X Scrollers main part position
func (c *Canvas) getXScrollerRect() image.Rectangle {
	var start float64
	len := float64(c.scroller_x.MaxRect.Max.X - c.scroller_x.MaxRect.Min.X)

	if c.viewport_x == 0 {
		start = float64(c.scroller_x.MaxRect.Min.X)
	} else {
		start = float64(c.scroller_x.MaxRect.Min.X) + (float64(c.viewport_x)/float64(c.canvasCols))*len
	}
	end := float64(c.scroller_x.MaxRect.Min.X) + ((float64(c.viewport_x)+float64(c.viewportCols))/float64(c.canvasCols))*len

	return image.Rect(int(start), c.scroller_x.MaxRect.Min.Y, int(end), c.scroller_x.MaxRect.Max.Y)
}

// returns Y Scrollers main part position
func (c *Canvas) getYScrollerRect() image.Rectangle {
	var start float64
	len := float64(c.scroller_y.MaxRect.Max.Y - c.scroller_y.MaxRect.Min.Y)

	if c.viewport_y == 0 {
		start = float64(c.scroller_y.MaxRect.Min.Y)
	} else {
		start = float64(c.scroller_y.MaxRect.Min.Y) + (float64(c.viewport_y)/float64(c.canvasRows))*len
	}
	end := float64(c.scroller_y.MaxRect.Min.Y) + ((float64(c.viewport_y)+float64(c.viewportRows))/float64(c.canvasRows))*len

	return image.Rect(c.scroller_y.MaxRect.Min.X, int(start), c.scroller_y.MaxRect.Max.X, int(end))
}

// updates Scrollers position due to viewport position according to viewport position
func (c *Canvas) UpdateScrollers() {
	c.scroller_x.Rect = c.getXScrollerRect()
	c.scroller_y.Rect = c.getYScrollerRect()
}

// creates new canvas struct
func NewCanvas(
	viewportRows int, viewportCols int,
	canvasRows int, canvasCols int,
	x_pos int, y_pos int,
) Canvas {
	var c Canvas

	c.viewportRows = viewportRows
	c.viewportCols = viewportCols
	c.canvasRows = canvasRows
	c.canvasCols = canvasCols
	c.drawingArea = make([]int, canvasRows*canvasCols)


	c.Rect = image.Rect(
		x_pos,
		y_pos,
		x_pos+TileWidth*viewportCols,
		y_pos+TileHeight*viewportRows,
	)

	c.scroller_x = NewScroller(
		x_pos,
		y_pos+TileHeight*viewportRows+2,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
		scrollerColor,
		scrollerBGColor,
	)

	c.scroller_y = NewScroller(
		x_pos+TileWidth*viewportCols+2,
		y_pos,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
		scrollerColor,
		scrollerBGColor,
	)

	c.clickableAreas = make(map[image.Rectangle]func())
	c.hoverableAreas = make(map[image.Rectangle]func())

	c.clickableAreas[c.scroller_x.arrowLow.rect] = c.MoveCanvasLeft
	c.clickableAreas[c.scroller_x.arrowHigh.rect] = c.MoveCanvasRight

	c.clickableAreas[c.scroller_y.arrowLow.rect] = c.MoveCanvasUp
	c.clickableAreas[c.scroller_y.arrowHigh.rect] = c.MoveCanvasDown

    c.UpdateScrollers()
	c.ClearDrawingArea()

	return c
}
