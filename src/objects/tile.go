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

// creates new Tile
func NewTile(image *ebiten.Image, rowOnPallete int, colOnPallete int) Tile {
	var newTile Tile

	newTile.Image = image
	newTile.rowOnPallete = rowOnPallete
	newTile.colOnPallete = colOnPallete

	return newTile
}

// gets current Tile from stack
func (ts *TileStack) GetCurrentTile() Tile {
	return ts.stack[ts.CurrentTile]
}

// gets Tile from stack by its index in stack
func (ts *TileStack) GetTileFromStack(i int) Tile {
	return ts.stack[i]
}

// gets number of times Tile is drawn. Tile is selected by its index on stack
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

// updates number of Tile is used. Tile is selected by index on stack
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

// adds new Tile to stack
func (ts *TileStack) AddTileToStack(tile Tile) {
	ts.stack = append(ts.stack, tile)
	ts.len++
}

// sets current Tile which will be used as brush
func (ts *TileStack) SetCurrentTile(index int) {
	ts.CurrentTile = index
}

// returns whole stack
func (ts *TileStack) GetAllTiles() []Tile {
	return ts.stack
}

// draws current Tile on screen
func (ts *TileStack) DrawCurrentTile(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	tileToDraw := ts.GetCurrentTile()
	screen.DrawImage(tileToDraw.Image, op)
}
