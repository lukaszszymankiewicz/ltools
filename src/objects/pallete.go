package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Pallete struct {
	Rect    image.Rectangle
	Cols    int
	Rows    int
	MaxTile int
}

func NewPallete(x int, y int, width int, height int) Pallete {
	var pallete Pallete

	pallete.Rect = image.Rect(x, y, x+width, y+height)

	return pallete
}

func (p *Pallete) PosToTileCoordsOnPallete(x int, y int) (int, int) {
	return int((x - p.Rect.Min.X) / TileWidth), int((y - p.Rect.Min.Y) / TileHeight)
}

func (p *Pallete) TileNrToCoordsOnPallete(tileNr int) (float64, float64) {
	tileX := float64(((tileNr % p.Cols) * TileWidth) + p.Rect.Min.X)
	tileY := float64(((tileNr / p.Cols) * TileHeight) + p.Rect.Min.Y)

	return tileX, tileY
}

func (p *Pallete) PosToCursorCoords(x int, y int) image.Rectangle {
	tileX, tileY := p.PosToTileCoordsOnPallete(x, y)

	leftCornerX := (tileX * TileWidth) + p.Rect.Min.X
	leftCornerY := (tileY * TileHeight) + p.Rect.Min.Y

	return image.Rect(leftCornerX, leftCornerY, leftCornerX+CursorSize, leftCornerY+CursorSize)
}

// translating from mouse position to tile on pallete. Please note that Pallete do not need to have
// exacly the same number of rows and columns as tileset (tileset file is basically trimmed to fit
// fixed pallete number of rows and cols). To get column and row of pallete, some basic
// transposition is needed
func (p *Pallete) PosToSubImageOnPallete(x int, y int, t *Tileset) *ebiten.Image {
	TileX, TileY := p.PosToTileCoordsOnPallete(x, y)
	TileNr := TileX + TileY*p.Cols

	return t.TileNrToSubImageOnTileset(TileNr)
}

// TODO: add cursor struct
func (p *Pallete) DrawCursorOnPallete(screen *ebiten.Image, x int, y int) {
	cursorRect := p.PosToCursorCoords(x, y)
	drawer.EmptyRect(screen, cursorRect, color.White)
}
