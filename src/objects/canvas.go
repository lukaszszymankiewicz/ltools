package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Scroller struct {
	color   image.Color
	MaxRect image.Rectangle
	Rect    image.Rectangle
}

type ScrollArrow struct {
	image *ebiten.Image
	Rect  image.Rectangle
}

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

func NewScroller (x1 int, y1 int, x2 int, y2 int, color image.Color) Scroller {
    var scroller Scroller 

	// scroller Rect is left uninitalized, as it will be updated soon
    scroller.MaxRect = image.Rect(x1, y1, x2, y2)
	scroller.color = color
	return Scroller
}

func NewScrollArrow(x int, y int, path string) ScrollArrow {
	var scrollArrow ScrollArrow

	scrollArrow.image = loadImage(path)

	width, height := scrollArrow.image.Size()
	scrollArrow.Rect = image.Rect(x, y, x+width, y+height)

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

	canvas.scroller_x = NewScroller(
		x_pos+32, 
		y_pos+TileHeight*viewportCols, 
		x_pos+TileWidth*viewportRows-32,
		y_pos+TileHeight*viewportCols+32,
		color.White,
	)

	canvas.scroller_y = NewScroller(
		x_pos+TileWidth*viewportRows,
		y_pos,
		x_pos+TileWidth*viewportRows+32,
		y_pos+32,
		color.White,
	)

    canvas.updateScroller()

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

func (c *Canvas) DrawScrollers(screen *ebiten.Image) {
    drawer.FilledRect(screen, c.scroller_x.Rect, c.scroller_x.color)
	drawer.FilledRect(screen, c.scroller_y.Rect, c.scroller_y.color)
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
	c.DrawScrollers(screen)
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

func (c *Canvas) getXScrollerRect() image.Rect {
	var start float64
	len := float64(c.scroller_x.MaxRect.Min.X - c.scroller_x.MaxRect.Max.X)

	if c.viewport_x == 0 {
		start = float64(c.scroller_x.MaxRect.Min.X)
    } else {
		start = float64(c.scroller_x.MaxRect.Min.X) + (float64(c.viewport_x) / float64(c.canvasCols)) * len
	}
    end := float64(c.scroller_x.MaxRect.Min.X) + ((float64(c.viewport_x) + float64(c.viewportCols)) / float64(c.canvasCols)) * len
	
    return image.Rect(int(start), c.scroller_x.MaxRect.Min.Y, int(end), c.scroller_x.MaxRect.Max.Y)
}

func (c *Canvas) getYScrollerRect() image.Rect {
	var start float64
	len := float64(c.scroller_x.MaxRect.Min.Y - c.scroller_x.MaxRect.Max.Y)

	if c.viewport_y == 0 {
		start = float64(c.scroller_y.MaxRect.Min.Y)
    } else {
		start = float64(c.scroller_y.MaxRect.Min.Y) + (float64(c.viewport_y) / float64(c.canvasRows)) * len
	}
    end := float64(c.scroller_y.MaxRect.Min.Y) + ((float64(c.viewport_y) + float64(c.viewportRows)) / float64(c.canvasRows)) * len
	
    return image.Rect(c.scroller_y.MaxRect.Min.X, int(start), c.scroller_y.MaxRect.Max.X, int(end))
}

func (c *Canvas) UpdateScrollers() {
    c.scroller_x.Rect = c.getXScrollerRect()
	c.scroller_y.Rect = c.getYScrollerRect()
}

func (c *Canvas) MoveCanvas(x int, y int) {
	new_x_value := c.viewport_x + x
	new_y_value := c.viewport_y + y

	c.viewport_x = MinVal(MaxVal(0, new_x_value), (c.viewportRows - c.viewport_x))
	c.viewport_y = MinVal(MaxVal(0, new_y_value), (c.viewportCols - c.viewport_y))

	c.UpdateScrollers()
}