package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Image        *ebiten.Image
	rowOnPallete int
	colOnPallete int
	NumberUsed   int
}

type TileStack struct {
	stack       []Tile
	len         int
	CurrentTile int
}

func NewTile(image *ebiten.Image, rowOnPallete int, colOnPallete int) Tile {
	var newTile Tile

	newTile.Image = image
	newTile.rowOnPallete = rowOnPallete
	newTile.colOnPallete = colOnPallete

	return newTile
}

func (ts *TileStack) GetCurrentTile() Tile {
	return ts.stack[ts.CurrentTile]
}

func (ts *TileStack) GetTileFromStack(i int) Tile {
	return ts.stack[i]
}

func (ts *TileStack) GetTileNumberUsed(i int) int {
	return ts.GetTileFromStack(i).NumberUsed
}

// checks if Tile needs to be deleted from TileStack. It happens if number of tile usage on level
// is equal to zero
func (ts *TileStack) ClearTileStack(i int) {
	if ts.GetTileNumberUsed(i) == 0 {
		ts.stack = append(ts.stack[:i], ts.stack[i+1:]...)
	}
}

func (ts *TileStack) UpdateTileUsage(i int, value int) {
	ts.stack[i].NumberUsed += value
}

// check if Tile is already used (so it is present on TileStack). If so, its index in stack is
// returned, if not -1 is returned (ANSI C style!)
func (ts *TileStack) CheckTileInStack(tileX int, tileY int) (stackIndex int) {
	for i := 0; i < ts.len; i++ {
		if ts.stack[i].rowOnPallete == tileX && ts.stack[i].colOnPallete == tileY {
			return i
		}
	}
	return -1
}

func (ts *TileStack) AddTileToStack(tile Tile) {
	ts.stack = append(ts.stack, tile)
	ts.len++
}

func (ts *TileStack) SetCurrentTile(index int) {
	ts.CurrentTile = index
}

func (ts *TileStack) GetAllTiles() []Tile {
	return ts.stack
}

func (ts *TileStack) DrawCurrentTile(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	tileToDraw := ts.GetCurrentTile()
	screen.DrawImage(tileToDraw.Image, op)
}
