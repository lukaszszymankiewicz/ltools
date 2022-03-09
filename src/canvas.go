package canvas

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
    "src/const.go"
    "src/primitives.go"
)

// TODO: is this necessary?
type Matrix struct {
    num int
}

type Canvas struct {
    viewportCols int
    viewportRows int
    drawingArea int // 2d array
	rect image.Rectangle
    pos_x int
    pos_y int
}

func newCanvas(cols int, rows int, x int, y int) Canvas {
	var canvas Canvas

    canvas.viewportRows = ViewportRowsN 
    canvas.viewportCols = ViewportColsN 
    canvas.drawingArea = [CanvasX][CanvasY]Matrix{}
	canvas.rect = image.Rect(x, y, x + (TileWidth * CanvasColsN), y + (TileWidth * CanvasRowsN))
    canvas.pos_x = 0
    canvas.pos_y = 0

	return canvas
}


func (c *Canvas) ClearDrawingArea() {
	for x := 0; x < CanvasRowsN; x++ {
		for y := 0; y < CanvasColsN; y++ {
			g.drawingArea[x][y] = -1
		}
	}
}

func (c *Canvas) ClearViewport() {
	for x := 0; x < c.viewportRows ; x++ {
		for y := 0; y < c.viewportCols ; y++ {
			c.[x][y] = -1
		}
	}
}

func (c *Canvas) Draw(screen *ebiten.Image) {
    DrawBorder(screen, g.rect, color.White)

	for x:=pos_x; x<c.viewportRows+pos_x; x++ {
		for y:=pos_y; y<c.viewportCols+pos_y; y++ {
			if tile := c.drawingArea[x][y]; tile != -1 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64(x * TileWidth + c.rect.Min.X),
					float64(y * TileWidth + c.rect.Min.Y),
				)
				screen.DrawImage(g.tileStack[g.canvas[x][y]].image, op)
			}
		}
	}
}

