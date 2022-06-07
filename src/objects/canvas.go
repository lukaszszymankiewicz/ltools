package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Canvas struct {
	Grid
	Recorder
	bg                      ImageElement
	SetTileByLayerCondition []func(int, int, int) bool
	SetTileByLayerEffect    []func(int, int, int, *Tile) int
}

func (c *Canvas) IsEmpty(x int, y int, n int) bool {
	isempty := true

	for i := 0; i < c.n_layers; i++ {
		if c.GetTileOnDrawingArea(x, y, i) != nil {
			isempty = false
		}
	}
	return isempty
}

func (c *Canvas) HasTile(x int, y int, n int) bool {
	// this 0 means that something is put on 0 layer
	// this probalby be more sofistaced than it it bus damn, i want it as simple as it can get
	if c.GetTileOnDrawingArea(x, y, 0) != nil {
		return true
	}
	return false
}

func (c *Canvas) Always(x int, y int, n int) bool {
	return true
}

// this show how to set alpha channel for every layer if given layer is on
var layers_alpha = [][]float64{
	{1.0, 1.0, 1.0}, // LAYER_DRAW is on
	{0.5, 1.0, 0.5}, // LAYER_LIGHT in on
	{0.5, 0.0, 1.0}, // LAYER_ENTITY in on
}

func (c *Canvas) Draw(screen *ebiten.Image, layer int) {
	// drawing Tiles and tiles on it
	for x := 0; x < c.viewportCols; x++ {
		for y := 0; y < c.viewportRows; y++ {
			empty := true
			pos_x := c.rect.Min.X + (x * c.grid_size)
			pos_y := c.rect.Min.Y + (y * c.grid_size)

			for l := 0; l < c.n_layers; l++ {
				tile := c.GetTileOnDrawingArea(x+c.viewport_x, y+c.viewport_y, l)
				if tile != nil {
					tile.ImageElement.rect.Min.X = pos_x
					tile.ImageElement.rect.Min.Y = pos_y
					tile.ImageElement.alpha = layers_alpha[layer][l]

					tile.ImageElement.Draw(screen)
					empty = false
				}
			}

			// if nothing is to be drawn - background layer is draw
			if empty == true {
				c.bg.rect.Min.X = pos_x
				c.bg.rect.Min.Y = pos_y
				c.bg.Draw(screen)
			}
		}
	}

	// drawing the scrollers
	c.scroller_x.Draw(screen)
	c.scroller_y.Draw(screen)
}

// creates new canvas struct
func NewCanvas(
	x int,
	y int,
	width int,
	height int,
	grid_size int,
	rows int,
	cols int,
	n_layers int,
) Canvas {
	var c Canvas

	c.Grid = NewGrid(x, y, width, height, grid_size, rows, cols, n_layers, greyColor)
	c.bg = NewImageElement(0, 0, LoadImage("src/objects/assets/other/bg.png"))

	c.SetTileByLayerCondition = make([]func(int, int, int) bool, 0)
	c.SetTileByLayerEffect = make([]func(int, int, int, *Tile) int, 0)

	c.SetTileByLayerCondition = []func(int, int, int) bool{
		c.Always,  // setting tile from LAYER_DRAW has such condition
		c.IsEmpty, // setting tile from LAYER_LIGHT has such condition
		c.HasTile, // setting tile from LAYER_ENTITY has such condition
	}

	c.SetTileByLayerEffect = []func(int, int, int, *Tile) int{
		c.DoNothingOnDrawingArea, c.SetTileOnDrawingArea, c.CleanTileOnDrawingArea,
	}

	c.UpdateXScroller()
	c.UpdateYScroller()

	return c
}

func (c *Canvas) EraseOneRecord(record Record) {
	for i := len(record.x_coords) - 1; i > -1; i-- {
		c.SetTileOnDrawingArea(
			record.x_coords[i],
			record.y_coords[i],
			record.layers[i],
			record.old_tiles[i],
		)
	}
}

func (c *Canvas) MoveCanvasDown(screen *ebiten.Image) {
	c.MoveViewport(0, 1)
}

func (c *Canvas) MoveCanvasUp(screen *ebiten.Image) {
	c.MoveViewport(0, -1)
}

func (c *Canvas) MoveCanvasRight(screen *ebiten.Image) {
	c.MoveViewport(1, 0)
}

func (c *Canvas) MoveCanvasLeft(screen *ebiten.Image) {
	c.MoveViewport(-1, 0)
}

func (c *Canvas) TileIsAllowed(x int, y int, l int) bool {
	return c.SetTileByLayerCondition[l](x+c.viewport_x, y+c.viewport_y, l)
}

func (c *Canvas) PutTile(x int, y int, fill *Tile, tool Tool) {
    corr_x := x+c.viewport_x
    corr_y := y+c.viewport_y

	brush := tool.Brush
	brush_result := tool.Brush.ApplyBrush(corr_x, corr_y, corr_x, corr_y, fill)

	for i := 0; i < brush_result.Len(); i++ {

		pos_x := brush_result.coords[i][0]
		pos_y := brush_result.coords[i][1]
		layer := brush_result.layers[i]
		tile := brush_result.tiles[i]

		if c.TileIsAllowed(pos_x, pos_y, layer) == true {
			c.StartRecording()

			// if Tile is allowed to be placed here apply effect of a brush to each layer
			for j := 0; j < c.n_layers; j++ {
				old_tile := c.GetTileOnDrawingArea(corr_x, corr_y, j)

				drawFuncIdx := brush.byLayerEffect[layer][j]
				Func := c.SetTileByLayerEffect[drawFuncIdx]
				FuncResult := Func(corr_x, corr_y, j, tile)

				new_tile := c.GetTileOnDrawingArea(corr_x, corr_y, j)

				if FuncResult != 0 {
					c.Recorder.AppendToCurrent(corr_x, corr_y, old_tile, new_tile, layer)
				}
			}
		}
	}
}

func (c *Canvas) UndrawOneRecord() {
	record := c.Recorder.UndoOneRecord()

	// if record to undraw is empty there is nothing to undraw!
	if record.IsEmpty() {
		return
	}

	c.EraseOneRecord(record)
}
