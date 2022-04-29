package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Pallete struct serves as a wrapper for Tiles which can be drawn.  Group of Tile are called
// Tileset. Each Tileset represent one logical layer of level. Collection of Tileset (grouping
// multiple layers) is called Tileseter
type Pallete struct {
	Rect       image.Rectangle
	Cols       int
	Rows       int
	MaxRows    int
	Viewport_y int
	Tileseter  Tileseter
	Scroller_y Scroller
}

// creates new Pallete struct
func NewPallete(
	x_pos int,
	y_pos int,
	rows int,
	cols int,
	tilesets [][]string,
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
		p.Rect.Max.X+ScrollerXOffset,
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
	tileset := p.Tileseter.GetCurrent()
	maxRows := tileset.rows
	offset := float64(p.Viewport_y / TileWidth)

	var start float64
	len := float64(p.Scroller_y.MaxRect.Max.Y - p.Scroller_y.MaxRect.Min.Y)

	if p.Viewport_y == 0 {
		start = float64(p.Scroller_y.MaxRect.Min.Y)
	} else {
		start = float64(p.Scroller_y.MaxRect.Min.Y) + (offset/float64(maxRows))*len
	}

	end := float64(p.Scroller_y.MaxRect.Min.Y) + ((offset+float64(p.Rows))/float64(maxRows))*len

	return image.Rect(p.Scroller_y.MaxRect.Min.X, int(start), p.Scroller_y.MaxRect.Max.X, int(end))
}

// updates scroller position
func (p *Pallete) updateScrollers() {
	p.Scroller_y.Rect = p.getScrollerRect()
}

// translates mouse position to nearest tile on pallete upper left corner
func (p *Pallete) MousePosToCursorRect(x int, y int, size int) image.Rectangle {
	RectX := p.Rect.Min.X + ((x-p.Rect.Min.X)/TileHeight)*TileHeight
	RectY := p.Rect.Min.Y + ((y-p.Rect.Min.Y)/TileHeight)*TileHeight

	return image.Rect(RectX, RectY, RectX+size, RectY+size)
}

// function for grabbing some tile from pallete
func (p *Pallete) MousePosToTileImageOnPallete(x int, y int) *ebiten.Image {
	RectX := int((x - p.Rect.Min.X) / TileWidth)
	RectY := int((y-p.Rect.Min.Y)/TileHeight) + p.Viewport_y

	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)
	tileset := p.Tileseter.GetCurrent()

	return tileset.image.SubImage(tileRect).(*ebiten.Image)
}

// translates mouse position to Tile coordinates (its row and col on pallete), while mouse is
// pointing on Pallete
func (p *Pallete) MousePosToTileRowAndColOnPallete(x int, y int) (int, int) {
	return int((x - p.Rect.Min.X) / TileWidth), int((y-p.Rect.Min.Y)/TileHeight) + int(p.Viewport_y/TileHeight)
}

// return Tile image from Pallete basing on Tile row and column position on Pallete
func (p *Pallete) GetTileImage(row int, col int, tileset Tileset) *ebiten.Image {

	tileRect := image.Rect(
		(row * TileWidth),
		(col * TileHeight),
		(row*TileWidth)+CursorSize,
		(col*TileHeight)+CursorSize,
	)

	tile_image := tileset.image.SubImage(tileRect).(*ebiten.Image)

	return tile_image

}

// moves Pallete up
func (p *Pallete) MovePalleteUp(screen *ebiten.Image) {
	p.MovePallete(0, -1*TileHeight)
}

// moves Pallete down
func (p *Pallete) MovePalleteDown(screen *ebiten.Image) {
	p.MovePallete(0, TileHeight)
}

// moves Pallete in given direction
func (p *Pallete) MovePallete(x int, y int) {
	t := p.Tileseter.GetCurrent()
	p.Viewport_y = MinVal(MaxVal(0, p.Viewport_y+y), (t.height - ViewportRowsN*TileWidth))

	p.updateScrollers()
}

// draws Pallete
func (p *Pallete) DrawPallete(screen *ebiten.Image) {
	tileset := p.Tileseter.GetCurrent()
	tileRect := image.Rect(0, p.Viewport_y, p.Cols*TileWidth, p.Viewport_y+p.Rows*TileHeight)

	pallete_image := tileset.image.SubImage(tileRect).(*ebiten.Image)
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.Rect.Min.X), float64(p.Rect.Min.Y))
	screen.DrawImage(pallete_image, op)

	p.Scroller_y.Draw(screen)
}
