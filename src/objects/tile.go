package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Image    *ebiten.Image   // Tile image
	row       int            // row number on Pallete
	col       int            // col number on Pallete
	n         int            // number of tile is used on Canvas
    tileset   int            // index of Tilseset from which Tile is taken
    unique    bool
    required  bool
}

type TileStack struct {
	stack       []Tile       // collection of Tile
	current     int          // current chosen Tile to draw
}

// creates new Tile
func NewTile(image *ebiten.Image, row int, col int, tileset int) Tile {
	var t Tile

	t.Image   = image
	t.row     = row
	t.col     = col
	t.tileset = tileset

	return t
}

// gets current Tile from stack
func (ts *TileStack) GetCurrentTile() Tile {
	return ts.stack[ts.current]
}


func (ts *TileStack) GetCurrentTilePos() (int, int) {
    tile := ts.stack[ts.current]
    return tile.row, tile.col
}

// gets Tile from stack by its index in stack
func (ts *TileStack) GetTileFromStack(i int) Tile {
	return ts.stack[i]
}

// gets Tile from stack by its index in stack
func (ts *TileStack) GetTileTileset(i int) int {
	return ts.stack[i].tileset
}

// gets number of times Tile is drawn. Tile is selected by its index on stack
func (ts *TileStack) GetTileNumberUsed(i int) int {
	return ts.GetTileFromStack(i).n
}

// checks if Tile needs to be deleted from TileStack. It happens if number of tile usage on level
// is equal to zero
func (ts *TileStack) ClearTileStack() {
    for i:=0; i<len(ts.stack); i++ {
        if ts.GetTileNumberUsed(i) == 0 {
            ts.stack = append(ts.stack[:i], ts.stack[i+1:]...)
        }
    }
}

// updates number of Tile is used. Tile is selected by index on stack
func (ts *TileStack) UpdateTileUsage(i int, value int) {
	ts.stack[i].n += value
}

// check if Tile is already used (so it is present on TileStack). If so, its index in stack is
// returned, if not -1 is returned (ANSI C style!)
func (ts *TileStack) CheckTileInStack(row int, col int, tileset int) (i int) {
	for i := 0; i < len(ts.stack); i++ {
		if ts.stack[i].row == row && ts.stack[i].col == col && ts.stack[i].tileset == tileset{
			return i
		}
	}
	return -1
}

// adds new Tile to stack
func (ts *TileStack) AppendToStack(image *ebiten.Image, row int, col int, tileset int) {
    tile := NewTile(image, row, col, tileset)
	ts.stack = append(ts.stack, tile)
}

// adds new Tile to stack
func (ts *TileStack) AddTileToStack(tile Tile) {
	ts.stack = append(ts.stack, tile)
}

// sets current Tile which will be used as brush
func (ts *TileStack) SetCurrentTile(i int) {
	ts.current = i
}

// returns whole stack
func (ts *TileStack) GetAllTiles() []Tile {
	return ts.stack
}

// returns current Tileset index
func (ts *TileStack) CurrentTileIndex() int {
    return ts.current
}

// draws current Tile on screen
func (ts *TileStack) DrawCurrentTile(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	tileToDraw := ts.GetCurrentTile()
	screen.DrawImage(tileToDraw.Image, op)
}
