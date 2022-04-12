package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Canvas struct {
	canvasRows     int
	canvasCols     int
	drawingAreas   [][]int
    current_layer  int
	Rect           image.Rectangle
	viewportCols   int
	viewportRows   int
	viewport_x     int
	viewport_y     int
	Scroller_x     Scroller
	Scroller_y     Scroller
}

// makes all layers of canvas, and fills it with -1 (empty)
func (c *Canvas) ClearDrawingAreas(n_layers int) {
    
    for i:=0; i<n_layers; i++ {
        c.drawingAreas[i] = make([]int, c.canvasRows * c.canvasCols)

        for x := 0; x < c.canvasCols; x++ {
            for y := 0; y < c.canvasRows; y++ {
                c.drawingAreas[i][x*c.canvasRows+y] = -1
            }
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

    var alpha_mode float64 
    alpha_mode = 1.0

    if c.current_layer == TransparentLayer { 
        alpha_mode = CanvasAlphaMod         
    }

	// drawing Tiles and tiles on it
	for x := c.viewport_x; x < c.viewportCols+c.viewport_x; x++ {
		for y := c.viewport_y; y < c.viewportRows+c.viewport_y; y++ {
			if tile_index := c.GetTileOnDrawingArea(x, y); tile_index != -1 {
                
                tile := tiles[tile_index]

                op := &ebiten.DrawImageOptions{}
                op.GeoM.Translate(
                    float64((x-c.viewport_x)*TileWidth+c.Rect.Min.X),
                    float64((y-c.viewport_y)*TileWidth+c.Rect.Min.Y),
                )

                if tile.tileset == 0 {
                    op.ColorM.Scale(1.0, 1.0, 1.0, alpha_mode)
                } else {
                    op.ColorM.Scale(1.0, 1.0, 1.0, 1.0)
                }

                screen.DrawImage(tile.Image, op)

			}
		}
	}

	// drawing the scrollers
	c.Scroller_x.Draw(screen)
	c.Scroller_y.Draw(screen)
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
func (c *Canvas) GetTileOnDrawingArea(x int, y int) int {
    var tile int
    tile = -1

    tile = c.drawingAreas[0][x*c.canvasRows+y]
    if tile == -1 {
        tile = c.drawingAreas[1][x*c.canvasRows+y]
    }

    return tile
}

// sets current Tile to draw (brush) on Canvas. Canvas x and y are used, and Tile is
// selected by its position on stack
func (c *Canvas) SetTileOnCanvas(x int, y int, value int) {
    // yeahh...
    c.drawingAreas[c.current_layer][(x+c.viewport_x)*c.canvasRows+(y+c.viewport_y)] = value
}

// moves viewport left
func (c *Canvas) MoveCanvasLeft(*ebiten.Image) {
	c.MoveCanvas(-1, 0)
}

// moves viewport right
func (c *Canvas) MoveCanvasRight(*ebiten.Image) {
	c.MoveCanvas(1, 0)
}

// moves viewport up
func (c *Canvas) MoveCanvasUp(*ebiten.Image) {
	c.MoveCanvas(0, -1)
}

// moves viewport down
func (c *Canvas) MoveCanvasDown(*ebiten.Image) {
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
	len := float64(c.Scroller_x.MaxRect.Max.X - c.Scroller_x.MaxRect.Min.X)

	if c.viewport_x == 0 {
		start = float64(c.Scroller_x.MaxRect.Min.X)
	} else {
		start = float64(c.Scroller_x.MaxRect.Min.X) + (float64(c.viewport_x)/float64(c.canvasCols))*len
	}
	end := float64(c.Scroller_x.MaxRect.Min.X) + ((float64(c.viewport_x)+float64(c.viewportCols))/float64(c.canvasCols))*len

	return image.Rect(int(start), c.Scroller_x.MaxRect.Min.Y, int(end), c.Scroller_x.MaxRect.Max.Y)
}

// returns Y Scrollers main part position
func (c *Canvas) getYScrollerRect() image.Rectangle {
	var start float64
	len := float64(c.Scroller_y.MaxRect.Max.Y - c.Scroller_y.MaxRect.Min.Y)

	if c.viewport_y == 0 {
		start = float64(c.Scroller_y.MaxRect.Min.Y)
	} else {
		start = float64(c.Scroller_y.MaxRect.Min.Y) + (float64(c.viewport_y)/float64(c.canvasRows))*len
	}
	end := float64(c.Scroller_y.MaxRect.Min.Y) + ((float64(c.viewport_y)+float64(c.viewportRows))/float64(c.canvasRows))*len

	return image.Rect(c.Scroller_y.MaxRect.Min.X, int(start), c.Scroller_y.MaxRect.Max.X, int(end))
}

// updates Scrollers position due to viewport position according to viewport position
func (c *Canvas) UpdateScrollers() {
	c.Scroller_x.Rect = c.getXScrollerRect()
	c.Scroller_y.Rect = c.getYScrollerRect()
}

// switches layer
func (c *Canvas) ChangeLayer(n int) {
	c.current_layer = n
}

// creates new canvas struct
func NewCanvas(
	viewportRows int, viewportCols int,
	canvasRows int, canvasCols int,
	x_pos int, y_pos int,
    n_layers int,
) Canvas {
	var c Canvas

	c.viewportRows = viewportRows
	c.viewportCols = viewportCols
	c.canvasRows   = canvasRows
	c.canvasCols   = canvasCols
	c.drawingAreas = make([][]int, n_layers)

	c.Rect = image.Rect(
		x_pos,
		y_pos,
		x_pos+TileWidth*viewportCols,
		y_pos+TileHeight*viewportRows,
	)

	c.Scroller_x = NewScroller(
		x_pos,
		y_pos,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
	)

	c.Scroller_y = NewScroller(
		x_pos+TileWidth*viewportCols,
		y_pos,
		TileWidth,
		TileHeight*viewportRows,
	)

	c.UpdateScrollers()
	c.ClearDrawingAreas(n_layers)

	return c
}
