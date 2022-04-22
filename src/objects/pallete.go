package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
    // "fmt"
)

// Pallete struct serves as a wrapper for brushes which can be drawn.
// each "brush" is just a single Tile. Group of tile are called Tileset. Each Tileset represent one
// logical layer of level. Collection of Tileset (grouping multiple layers) is called Tileseter
type Pallete struct {
	Rect       image.Rectangle
	Cols       int
	Rows       int
	MaxRows    int
	Viewport_y int
    Tilesets   Tileseter 
	Scroller_y Scroller
}

// creates new Pallete struct
func NewPallete(
	x_pos int,
    y_pos int,
	rows int,
    cols int,
    [][]string tilesets,
) Pallete {
	var p Pallete

	p.Cols = cols
	p.Rows = rows

	p.Rect = image.Rect(
		x_pos,
		y_pos,
		x_pos+TileWidth*cols,
		y_pos+TileHeight*rows,
	)

	p.Scroller_y = NewScroller(
		p.Rect.Max.X+SCROLLER_X_OFFSET ,
		y_pos,
		TileWidth*cols,
		TileHeight*rows,
	)
    p.Tileseter = NewTileseter(tilesets)

	p.updateScrollers()

	return p
}

// returns Y Scrollers main part position
func (p *Pallete) getScrollerRect() image.Rectangle {
    tileset := g.Tileseter.GetCurrent()  
    maxRows := tileset.rows

	var start float64
	len := float64(p.Scroller_y.MaxRect.Max.Y - p.Scroller_y.MaxRect.Min.Y)

	if p.Viewport_y == 0 {
		start = float64(p.Scroller_y.MaxRect.Min.Y)
	} else {
		start = float64(p.Scroller_y.MaxRect.Min.Y) + (float64(p.Viewport_y)/float64(MaxRows))*len
	}
	end := float64(p.Scroller_y.MaxRect.Min.Y) + ((float64(p.Viewport_y)+float64(p.Rows))/float64(MaxRows))*len
	return image.Rect(p.Scroller_y.MaxRect.Min.X, int(start), p.Scroller_y.MaxRect.Max.X, int(end))
}

// updates scroller position
func (p *Pallete) updateScrollers() {
	p.Scroller_y.Rect = p.getScrollerRect()
}

// translates mouse position to nearest tile on pallete upper left corner
func (p *Pallete) MousePosToCursorRect(x int, y int) image.Rectangle {
	tileX, tileY := p.MousePosToTileRowAndColOnPallete(x, y)

    tileRect := image.Rect(
        tileX * TileWidth) + p.Rect.Min.X,,
        (tileY * TileHeight) + p.Rect.Min.Y,

        tileX * TileWidth) + p.Rect.Min.X+CursorSize,
        (tileY * TileHeight) + p.Rect.Min.Y+CursorSize,
    )
}

// function for grabbing some tile from pallete
func (p *Pallete) MousePosToTileImageOnPallete(x int, y int) *ebiten.Image {
    RectX := int((x - p.Rect.Min.X) / TileWidth)
    RectY := int((y - p.Rect.Min.Y) / TileHeight) + p.Viewport_y

	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

    tileset := g.Tileseter.GetCurrent()  

	return tileset.image.SubImage(tileRect).(*ebiten.Image)
}

// translates mouse position to Tile coordinates (its row and col on pallete), while mouse is
// pointing on Pallete
func (p *Pallete) MousePosToTileRowAndColOnPallete(x int, y int) (int, int) {
    return int((x - p.Rect.Min.X) / TileWidth), int((y - p.Rect.Min.Y) / TileHeight)
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

// draws pallete object
func (p *Pallete) drawPallete(screen *ebiten.Image) {
    tileset := g.Tileseter.GetCurrent()  

    tileRect := image.Rect(0, g.Viewport_y, p.Rows*TileWidth, g.Viewport_y +p.Cols*TileHeight)
    pallete_image := tileset.image.SubImage(tileRect).(*ebiten.Image)

    op.GeoM.Translate(PalleteX, PalleteY)
    screen.DrawImage(pallete_image, op)

	g.Pallete.Scroller_y.Draw(screen)
}

