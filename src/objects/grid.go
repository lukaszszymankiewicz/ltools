package objects

import (
	"image"
	"image/color"
)

type Grid struct {
	FilledRectElement
	rows          int
	cols          int
	width         int
	height        int
	viewportCols  int
	viewportRows  int
	viewport_x    int
	scroller_x    Scroller
    viewport_y    int
	scroller_y    Scroller
	drawingAreas  [][]*Tile
	current_layer int
	n_layers      int
	grid_size     int
}

func NewGrid(
	x int,
	y int,
	width int, // visible part of grid, in pixels
	height int, // visible part of grid, in pixels
	grid_size int,
	rows int, // all number of rows (even unvisble)
	cols int, // all number of cols (even unvisible)
	n_layers int,
	bg_color *color.RGBA,
) Grid {
	var g Grid

	g.FilledRectElement = NewFilledRectElement(x, y, width, height, bg_color)

	g.rows = rows
	g.cols = cols
	g.viewportCols = int(width / grid_size)
	g.viewportRows = int(height / grid_size)
	g.width = width
	g.height = height
	g.n_layers = n_layers
	g.grid_size = grid_size

	g.scroller_x = NewScroller(x, y+height, width, 16) // this one is good
	g.scroller_y = NewScroller(x+width, y, 16, height)

	g.drawingAreas = make([][]*Tile, g.n_layers)

	for i := 0; i < g.n_layers; i++ {
		g.drawingAreas[i] = make([]*Tile, g.rows*g.cols)
	}
	g.ClearDrawingAreas()

	return g
}

// makes all layers of canvas, and fills it with nil
func (g *Grid) ClearDrawingAreas() {

	for i := 0; i < g.n_layers; i++ {
		for x := 0; x < g.cols; x++ {
			for y := 0; y < g.rows; y++ {
				g.drawingAreas[i][x*g.rows+y] = nil
			}
		}
	}
}

func (g *Grid) GetTileOnDrawingArea(x int, y int, n int) *Tile {
	return g.drawingAreas[n][x*g.cols+y]
}

func (g *Grid) SetTileOnDrawingArea(row int, col int, n int, val *Tile) {
	g.drawingAreas[n][row*g.cols+col] = val
}

func (g *Grid) UpdateXScroller() {
	len := float64(g.scroller_x.background.rect.Max.X - g.scroller_x.background.rect.Min.X)

	start := float64(g.scroller_x.background.rect.Min.X)

	if g.viewport_x != 0 {
		start += (float64(g.viewport_x) / float64(g.cols)) * len
	}

	end := start + ((float64(g.viewport_x)+float64(g.viewportCols))/float64(g.cols))*len

	g.scroller_x.bar.rect.Min.X = int(start)
	g.scroller_x.bar.rect.Max.X = int(end)
}

// returns Y Scrollers main part position
func (g *Grid) UpdateYScroller() {
	len := float64(g.scroller_y.background.rect.Max.Y - g.scroller_y.background.rect.Min.Y)
	start := float64(g.scroller_y.background.rect.Min.Y)

	if g.viewport_y != 0 {
		start += (float64(g.viewport_y) / float64(g.rows)) * len
	}
	end := start + float64(g.viewport_y) + (float64(g.viewportRows)/float64(g.rows))*len

	g.scroller_y.bar.rect.Min.Y = int(start)
	g.scroller_y.bar.rect.Max.Y = int(end)
}

func (g *Grid) MoveViewport(x int, y int) {
	new_x_value := g.viewport_x + x
	new_y_value := g.viewport_y + y

	g.viewport_x = MinVal(MaxVal(0, new_x_value), (g.cols - g.viewportCols))
	g.viewport_y = MinVal(MaxVal(0, new_y_value), (g.rows - g.viewportRows))

	g.UpdateXScroller()
	g.UpdateYScroller()
}

func (g *Grid) MousePosToTilePos(x int, y int) (int, int) {
	pos_x := int((x-g.rect.Min.X)/g.grid_size)*g.grid_size + g.rect.Min.X
	pos_y := int((y-g.rect.Min.Y)/g.grid_size)*g.grid_size + g.rect.Min.Y

	return pos_x, pos_y
}

func (g *Grid) MousePosToRowAndCol(x int, y int) (int, int) {
	return int((x - g.rect.Min.X) / g.grid_size), int((y - g.rect.Min.Y) / g.grid_size)
}

func (g *Grid) MousePosToTile(x int, y int, layer int) *Tile {
	row, col := g.MousePosToRowAndCol(x, y)
	return g.GetTileOnDrawingArea(row+g.viewport_x, col+g.viewport_y, layer)
}

func (g *Grid) ChangeCurrent(v int) {
	g.current_layer = v
}

func (g *Grid) CurrentLayer() int {
	return g.current_layer
}

func (g *Grid) ScrollerXLeftArrowArea() image.Rectangle {
	return g.scroller_x.arrowLow.Area()
}

func (g *Grid) ScrollerXRightArrowArea() image.Rectangle {
	return g.scroller_x.arrowHigh.Area()
}

func (g *Grid) ScrollerYUpArrowArea() image.Rectangle {
	return g.scroller_y.arrowLow.Area()
}

func (g *Grid) ScrollerYDownArrowArea() image.Rectangle {
	return g.scroller_y.arrowHigh.Area()
}

func (g *Grid) Size() (int, int) {
	return g.rows, g.cols
}

func (g *Grid) Layers() int {
	return g.n_layers
}
