package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
	"math"
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
}

type ScrollArrow struct {
	image *ebiten.Image
	Rect  image.Rectangle
	op    *ebiten.DrawImageOptions
}

func NewScrollArrow(x int, y int, rot int, flip_h bool, flip_v bool) ScrollArrow {
	var scrollArrow ScrollArrow
    var scale_x int = 1
    var scale_y int = 1

	if flip_h {
		scale_x = -1
	} 

	if flip_h {
		scale_x = -1
	} 

	op := &ebiten.DrawImageOptions{}
    op.GeoM.Scale(float64(scale_x), float64(scale_y))
    op.GeoM.Rotate(float64(rot%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(float64(x), float64(y))

	scrollArrow.op = op
	scrollArrow.image = loadImage("arrow.png")
	// TODO: check how co get image width and height
	scrollArrow.Rect = image.Rect(x, y, x+32, y+32)

	return scrollArrow
}

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
		x_pos+TileWidth*viewportRows,
		y_pos+TileHeight*viewportCols,
	)

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

func (c *Canvas) ClearViewport() {
	for x := c.viewport_x; x < c.viewportRows+c.viewport_x; x++ {
		for y := c.viewport_y; y < c.viewportCols+c.viewport_y; y++ {
			c.drawingArea[x*c.canvasRows+y] = -1
		}
	}
}

func (sa *ScrollArrow) DrawScrollArrow(screen *ebiten.Image) {
	screen.DrawImage(sa.image, sa.op)
}

func (c *Canvas) DrawCanvas(screen *ebiten.Image, tiles []Tile) {
	drawer.EmptyBorder(screen, c.Rect, color.White)

	for x := c.viewport_x; x < c.viewportRows+c.viewport_x; x++ {
		for y := c.viewport_y; y < c.viewportCols+c.viewport_y; y++ {
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

func MaxVal(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinVal(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func (c *Canvas) MoveCanvas(x int, y int) {
	new_x_value := c.viewport_x + x
	new_y_value := c.viewport_y + y

	c.viewport_x = MinVal(MaxVal(0, new_x_value), (c.viewportRows - c.viewport_x))
	c.viewport_y = MinVal(MaxVal(0, new_y_value), (c.viewportRows - c.viewport_y))
}
