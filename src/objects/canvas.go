package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Canvas struct {
	viewportCols int
	viewportRows int
	canvasRows   int
	canvasCols   int
	drawingArea  []int
	Rect         image.Rectangle
	viewport_x   int
	viewport_y   int
}

func NewCanvas(
	viewportRows int, viewportCols int,
	canvasRows int, canvasCols int,
	x_pos int, y_pos int,
	width int, height int,
) Canvas {
	var canvas Canvas

	canvas.viewportRows = viewportRows
	canvas.viewportCols = viewportCols
	canvas.canvasRows = canvasRows
	canvas.canvasCols = canvasCols
	canvas.drawingArea = make([]int, canvasRows*canvasCols)

    // TODO: make rect dnamically from tile size!
	canvas.Rect = image.Rect(x_pos, y_pos, x_pos+width, y_pos+height)
	canvas.viewport_x = 0
	canvas.viewport_y = 0

	return canvas
}

func (c *Canvas) ClearDrawingArea() {
	for x := 0; x < c.canvasRows; x++ {
		for y := 0; y < c.canvasCols; y++ {
			c.drawingArea[x*c.canvasRows+y] = -1
		}
	}
}

func (c *Canvas) DrawCanvas(screen *ebiten.Image, tiles []Tile) {
	drawer.EmptyBorder(screen, c.Rect, color.White)

	for x := c.viewport_x; x < c.viewportRows+c.viewport_x; x++ {
		for y := c.viewport_y; y < c.viewportCols+c.viewport_y; y++ {
			if tile_index := c.GetTileOnCanvas(x, y); tile_index != -1 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64(x*TileWidth+c.Rect.Min.X),
					float64(y*TileWidth+c.Rect.Min.Y),
				)
				screen.DrawImage(tiles[tile_index].Image, op)
			}
		}
	}
}

func (c *Canvas) PosToTileCoordsOnCanvas(x int, y int) (int, int) {
	return int((x - c.Rect.Min.X) / TileWidth), int((y - c.Rect.Min.Y) / TileHeight)
}

func (c *Canvas) PosToTileHoveredOnCanvas(x int, y int) (int, int) {
	tileX, tileY := c.PosToTileCoordsOnCanvas(x, y)
	return tileX*TileWidth + c.Rect.Min.X, tileY*TileHeight + c.Rect.Min.Y
}

func (c *Canvas) GetTileOnCanvas(x int, y int) int {
	return c.drawingArea[x*c.canvasRows+y]
}

func (c *Canvas) SetTileOnCanvas(x int, y int, value int) {
	c.drawingArea[x*c.canvasRows+y] = value
}
