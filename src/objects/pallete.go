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
	Viewport_y int
	scroller_y   Scroller
	clickableAreas map[image.Rectangle]func()
	hoverableAreas map[image.Rectangle]func()
}

// returns Y Scrollers main part position
func (p *Pallete) getScrollerRect() image.Rectangle {
	var start float64
	len := float64(p.y_scroller.MaxRect.Max.Y - p.y_scroller.MaxRect.Min.Y)

	if p.Viewport_y == 0 {
		start = float64(p.y_scroller.MaxRect.Min.Y)
	} else {
		start = float64(p.y_scroller.MaxRect.Min.Y) + (float64(p.Viewport_y)/float64(p.MaxRows))*len
	}
	end := float64(p.y_scroller.MaxRect.Min.Y) + ((float64(p.Viewport_y)+float64(p.Rows))/float64(p.MaxRows))*len

	return image.Rect(p.y_scroller.MaxRect.Min.X, int(start), p.y_scroller.MaxRect.Max.X, int(end))
}

func (p *Pallete) updateScrollers() {
	p.y_scroller.Rect = p.getScrollerRect()
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

func (p *Pallete) MovePallete(x int, y int) {
	new_value := p.Viewport_y + y

	// TODO: whats that 10?
	p.Viewport_y = MinVal(MaxVal(0, new_value), (p.MaxRows - 10))

	p.updateScrollers()
}

// creates new Pallete struct
func NewPallete(
	x_pos int, y_pos int, 
	viewportRows int, viewportCols int, 
	maxrows int
) Pallete {
	var p Pallete

	p.Cols = viewportCols
	p.Rows = viewportRows
	p.MaxRows = maxrows

	p.Rect = image.Rect(
		x, 
		y, 
		x_pos+TileWidth*viewportCols,
		y_pos+TileHeight*viewportRows,
	)

	p.y_scroller = NewScroller(
		x_pos+TileWidth*viewportCols+2,
		y_pos,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
		scrollerColor,
		scrollerBGColor
	)
	p.clickableAreas = make(map[image.Rectangle]func())
	p.hoverableAreas = make(map[image.Rectangle]func())

	p.clickableAreas[p.scroller_y.arrowLow.Image.Rect] = c.MovePallete(0, 1)
	p.clickableAreas[p.scroller_y.arrowHigh.Image.Rect] = c.MovePallete(0, -1)

	p.updateScrollers()

	return p
}