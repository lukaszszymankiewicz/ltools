package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Pallete struct {
	Rect       image.Rectangle
	Cols       int
	Rows       int
	MaxRows    int
	scroller   Scroller
	Viewport_y int
}

// creates new Pallete struct
func NewPallete(x1 int, y1 int, x2 int, y2 int, rows int, cols int, maxrows int) Pallete {
	var pallete Pallete

	pallete.Rect = image.Rect(x1, y1, x2, y2)
	pallete.Cols = cols
	pallete.Rows = rows
	pallete.MaxRows = maxrows

	pallete.scroller = NewScroller(
		x2,
		y1+32,
		x2+32,
		y2,
		color.RGBA{96, 96, 96, 255},
		color.RGBA{192, 192, 192, 255},
	)

	pallete.updateScrollers()

	return pallete
}

// returns Y Scrollers main part position
func (p *Pallete) getScrollerRect() image.Rectangle {
	var start float64
	len := float64(p.scroller.MaxRect.Max.Y - p.scroller.MaxRect.Min.Y)

	if p.Viewport_y == 0 {
		start = float64(p.scroller.MaxRect.Min.Y)
	} else {
		start = float64(p.scroller.MaxRect.Min.Y) + (float64(p.Viewport_y)/float64(p.MaxRows))*len
	}
	end := float64(p.scroller.MaxRect.Min.Y) + ((float64(p.Viewport_y)+float64(p.Rows))/float64(p.MaxRows))*len

	return image.Rect(p.scroller.MaxRect.Min.X, int(start), p.scroller.MaxRect.Max.X, int(end))
}

func (p *Pallete) updateScrollers() {
	p.scroller.Rect = p.getScrollerRect()
}

// translates mouse position to Tile coordinates, while mouse is pointing on Pallete
func (p *Pallete) PosToTileCoordsOnPallete(x int, y int) (int, int) {
	return int((x - p.Rect.Min.X) / TileWidth), int((y - p.Rect.Min.Y) / TileHeight)
}

// translates Tile number (order on which Tile occurs on Pallete) to its
// coordinates on screen
func (p *Pallete) TileNrToCoordsOnPallete(tileNr int) (float64, float64) {
	tileX := float64(((tileNr % p.Cols) * TileWidth) + p.Rect.Min.X)
	tileY := float64(((tileNr / p.Cols) * TileHeight) + p.Rect.Min.Y)

	return tileX, tileY
}

// translates mouse position to coords of cursor (cursor is square which highlights
// Tile on which mouse is pointing)
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
	TileNr := TileX + (TileY * p.Cols) + (p.Viewport_y * p.Cols)

	return t.TileNrToSubImageOnTileset(TileNr)
}

// draws cursor on Pallete
func (p *Pallete) DrawCursorOnPallete(screen *ebiten.Image, x int, y int) {
	cursorRect := p.PosToCursorCoords(x, y)
	drawer.EmptyRect(screen, cursorRect, color.White)
}

func (p *Pallete) DrawPalleteScroller(screen *ebiten.Image) {
	drawer.FilledRect(screen, p.scroller.MaxRect, p.scroller.bgcolor)
	drawer.FilledRect(screen, p.scroller.Rect, p.scroller.color)
}

func (p *Pallete) MovePallete(x int, y int) {
	new_value := p.Viewport_y + y

	p.Viewport_y = MinVal(MaxVal(0, new_value), (p.MaxRows - 10))

	p.updateScrollers()
}
