package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

type Tileset struct {
	image     *ebiten.Image
	width     int
	height    int
	rows      int
	cols      int
	Num       int
	tilesUsed int
    index     int
}

type Tileseter struct {
	tilesets  []Tileset
    current   int
}

// creates new Tileset struct
func NewTileset(path string, index int) Tileset {
	var tileset Tileset

	tileset.image = loadImage(path)
	width, height := tileset.image.Size()
	tileset.width  = width
	tileset.height = height
	tileset.rows   = tileset.width / TileWidth
	tileset.cols   = tileset.height / TileHeight
	tileset.Num    = tileset.rows * tileset.cols
	tileset.index  = index

	return tileset
}

// creates new Tileseter struct
func NewTileseter() Tileseter {
	var tsr Tileseter

	return tsr 
}

// creates new Tileseter struct
func (tsr *Tileseter) AddNewToTilesetter(path string) {
    index := len(tsr.tilesets)
    ts := NewTileset(path, index)
    tsr.tilesets = append(tsr.tilesets, ts)
}

// prepare Complete Tabber
func NewCompleteTilesetter() Tileseter {
    tsr := NewTileseter() 

    tsr.AddNewToTilesetter("assets/tileset_1.png")
    tsr.AddNewToTilesetter("assets/no_light.png")

    return tsr
}

// prepare Complete Tabber
func (tsr *Tileseter) SetCurrent (n int) {
    tsr.current = n
}

// prepare Complete Tabber
func (tsr *Tileseter) GetCurrentIndex () int {
    return tsr.current
}

func (tsr *Tileseter) GetCurrent() Tileset {
    return tsr.tilesets[tsr.current]
}

// translates Tile index number (order in which Tile occurs on Pallete)
func (t *Tileset) TileNrToSubImageOnTileset(n int) *ebiten.Image {
	RectX, RectY := n%t.cols*TileWidth, n/t.rows*TileHeight
	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

	return t.image.SubImage(tileRect).(*ebiten.Image)
}
