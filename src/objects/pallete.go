package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
    // "fmt"
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

// translates mouse position to nearest tile on pallete upper left corner
func (p *Pallete) PosToCursorCoords(x int, y int) (int, int) {
	tileX, tileY := p.PosToTileCoordsOnPallete(x, y)

	return (tileX * TileWidth) + p.Rect.Min.X, (tileY * TileHeight) + p.Rect.Min.Y
}

func (p *Pallete) PosToSubImageOnPallete(x int, y int, t Tileset) *ebiten.Image {
    RectX := int((x - p.Rect.Min.X) / TileWidth)
    RectY := int((y - p.Rect.Min.Y) / TileHeight)

	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

	return t.image.SubImage(tileRect).(*ebiten.Image)
}

func (p *Pallete) PalletePosToSubImageOnPallete(row int, col int, t Tileset) *ebiten.Image {
    RectX := row * TileWidth
    RectY := col * TileHeight

	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

	return t.image.SubImage(tileRect).(*ebiten.Image)
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

// translates mouse position to Tile coordinates, while mouse is pointing on Pallete
func (p *Pallete) PosToTileCoordsOnPallete(x int, y int) (int, int) {
    return int((x - p.Rect.Min.X) / TileWidth), int((y - p.Rect.Min.Y) / TileHeight)
}

