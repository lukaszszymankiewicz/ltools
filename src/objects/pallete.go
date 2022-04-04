package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Pallete struct {
	Rect       image.Rectangle
	Cols       int
	Rows       int
	MaxRows    int
	Viewport_y int
	Scroller_y Scroller
}

// returns Y Scrollers main part position
func (p *Pallete) getScrollerRect() image.Rectangle {

	var start float64
	len := float64(p.Scroller_y.MaxRect.Max.Y - p.Scroller_y.MaxRect.Min.Y)

	if p.Viewport_y == 0 {
		start = float64(p.Scroller_y.MaxRect.Min.Y)
	} else {
		start = float64(p.Scroller_y.MaxRect.Min.Y) + (float64(p.Viewport_y)/float64(p.MaxRows))*len
	}
	end := float64(p.Scroller_y.MaxRect.Min.Y) + ((float64(p.Viewport_y)+float64(p.Rows))/float64(p.MaxRows))*len
	return image.Rect(p.Scroller_y.MaxRect.Min.X, int(start), p.Scroller_y.MaxRect.Max.X, int(end))
}

// updates scroller position
func (p *Pallete) updateScrollers() {
	p.Scroller_y.Rect = p.getScrollerRect()
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

// translates mouse position to nearest tile on pallete upper left corner
func (p *Pallete) PosToCursorCoords(x int, y int) (int, int) {
	tileX, tileY := p.PosToTileCoordsOnPallete(x, y)

	return (tileX * TileWidth) + p.Rect.Min.X, (tileY * TileHeight) + p.Rect.Min.Y
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

// moves pallete up
func (p *Pallete) MovePalleteUp(screen *ebiten.Image) {
	p.MovePallete(0, -1)
}

// moves pallete down
func (p *Pallete) MovePalleteDown(screen *ebiten.Image) {
	p.MovePallete(0, 1)
}

// moves pallete in given direction
func (p *Pallete) MovePallete(x int, y int) {
	new_value := p.Viewport_y + y

	p.Viewport_y = MinVal(MaxVal(0, new_value), (p.MaxRows - ViewportRowsN))

	p.updateScrollers()
}

// creates new Pallete struct
func NewPallete(
	x_pos int, y_pos int,
	viewportRows int, viewportCols int,
	maxrows int,
) Pallete {
	var p Pallete

	p.Cols = viewportCols
	p.Rows = viewportRows
	p.MaxRows = maxrows

	p.Rect = image.Rect(
		x_pos,
		y_pos,
		x_pos+TileWidth*viewportCols,
		y_pos+TileHeight*viewportRows,
	)

	p.Scroller_y = NewScroller(
		x_pos+TileWidth*viewportCols+2,
		y_pos,
		TileWidth*viewportCols,
		TileHeight*viewportRows,
	)

	p.updateScrollers()

	return p
}
