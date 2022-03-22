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
}

type ScrollArrow struct {
	image *ebiten.Image
	Rect  image.Rectangle
}

func NewScrollArrow(x int, y int, path string) ScrollArrow {
	var scrollArrow ScrollArrow

	scrollArrow.image = loadImage(path)

	// TODO: implement this
	// width, height := scrollArrow.image.Size()
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sa.Rect.Min.X), float64(sa.Rect.Min.Y))

	screen.DrawImage(sa.image, op)
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
	c.drawingArea[x*c.canvasRows+y+c.viewport_y+(c.viewport_x*c.canvasRows)] = value
}

func (c *Canvas) MoveCanvas(x int, y int) {
	new_x_value := c.viewport_x + x
	new_y_value := c.viewport_y + y

	c.viewport_x = MinVal(MaxVal(0, new_x_value), (c.viewportRows - c.viewport_x))
	c.viewport_y = MinVal(MaxVal(0, new_y_value), (c.viewportCols - c.viewport_y))
}

func (c *Canvas) DrawXZipper(screen *ebiten.Image, x1 int, x2 int, y int, height int) {
    var start float64
	len := float64(x2 - x1)

	if c.viewport_x == 0 {
		start = float64(x1)
    } else {
		start = float64(x1) + (float64(c.viewport_x) / float64(c.canvasCols)) * len
	}
    end := float64(x1) + ((float64(c.viewport_x) + float64(c.viewportCols)) / float64(c.canvasCols)) * len

    rect := image.Rect(int(start), y, int(end), y + height)
    drawer.FilledRect(screen, rect, color.RGBA{191, 191, 191, 255})
}

