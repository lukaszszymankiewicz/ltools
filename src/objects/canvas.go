package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Scroller struct {
	color   color.Color
	bgcolor color.Color
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

func NewScroller(x1 int, y1 int, x2 int, y2 int, color color.Color, bgcolor color.Color) Scroller {
	var scroller Scroller

	// scroller Rect is left uninitalized, as it will be updated soon from level of whole Canvas
	scroller.MaxRect = image.Rect(x1, y1, x2, y2)
	scroller.color = color
	scroller.bgcolor = bgcolor

	return scroller
}

func NewScrollArrow(x int, y int, path string) ScrollArrow {
	var scrollArrow ScrollArrow

	scrollArrow.image = loadImage(path)

	width, height := scrollArrow.image.Size()
	scrollArrow.Rect = image.Rect(x, y, x+width, y+height)

	return scrollArrow
}

func (c *Canvas) ClearDrawingArea() {
	for x := 0; x < c.canvasCols; x++ {
		for y := 0; y < c.canvasRows; y++ {
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
	// background
	drawer.FilledRect(screen, c.scroller_x.MaxRect, c.scroller_x.bgcolor)
	drawer.FilledRect(screen, c.scroller_y.MaxRect, c.scroller_y.bgcolor)

	// actual scroller
	drawer.FilledRect(screen, c.scroller_x.Rect, c.scroller_x.color)
	drawer.FilledRect(screen, c.scroller_y.Rect, c.scroller_y.color)
}

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
	c.drawingArea[(x+c.viewport_x)*c.canvasRows+(y+c.viewport_y)] = value
}

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

func (c *Canvas) UpdateScrollers() {
	c.scroller_x.Rect = c.getXScrollerRect()
	c.scroller_y.Rect = c.getYScrollerRect()
}

func (c *Canvas) MoveCanvas(x int, y int) {
	new_x_value := c.viewport_x + x
	new_y_value := c.viewport_y + y

	c.viewport_x = MinVal(MaxVal(0, new_x_value), (c.canvasCols - c.viewportCols))
	c.viewport_y = MinVal(MaxVal(0, new_y_value), (c.canvasRows - c.viewportRows))

	c.UpdateScrollers()
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
		x_pos+TileWidth*viewportCols,
		y_pos+TileHeight*viewportRows,
	)

	canvas.scroller_x = NewScroller(
		x_pos+34,
		y_pos+TileHeight*viewportRows+2,
		x_pos+TileWidth*viewportCols-32,
		y_pos+TileHeight*viewportRows+34,
		color.RGBA{96, 96, 96, 255},
		color.RGBA{192, 192, 192, 255},
	)

	canvas.scroller_y = NewScroller(
		x_pos+TileWidth*viewportCols+2,
		y_pos+33,
		x_pos+TileWidth*viewportCols+34,
		y_pos+TileHeight*viewportRows-32,
		color.RGBA{96, 96, 96, 255},
		color.RGBA{192, 192, 192, 255},
	)

	canvas.UpdateScrollers()

	return canvas
}
