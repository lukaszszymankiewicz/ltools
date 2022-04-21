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

// create new Tileset from set of paths (could be only one path)
func NewTileset(paths []string, index int) Tileset {
	var tileset Tileset

    tileset.index  = index
    images := []ebiten.Image 
    n_tiles := 0

    // calculating overall timeset size
    for i:=0; i<len(paths); i++ {
        new_image := loadImage(paths[i])
        images = append(images, new_image)

        width, height := tileset.image.Size()
        n_tiles += (width / TileWidth) * (tileset.height / TileHeight)
    }

    // new tileset image is created to hold all tiles on all images inputed
    tileset_image := ebiten.NewImage(PalleteColsN*TileWidth, (n_tiles/PalleteColsN+1)*TileHeight)

    // iterating one more time, this time filling up the tileset
    for j:=0; j<len(images); j++ {
        width, height := images[j].Size()
        rows := width / TileWidth
        cols := height / TileHeight
        n = rows * cols

        // for each tileset, each tile is obtained an put into right place of tileset
        // this fuss is required mostly to always have tileset with same size no matter what image
        // is inputted
        for z := 0; z <n; z++ {
            rect_x = (z % cols) * TileWidth
            rect_y = (z / cols) * TileHeight
            tileRect := image.Rect(rect_x, rect_y, rect_x+TileWidth, rect_y+TileHeight)
            tile := t.image.SubImage(tileRect).(*ebiten.Image)

            x := (i%PalleteColsN) * TileWidth
            y := (i/PalleteColsN) * TileHeight
            op := &ebiten.DrawImageOptions{}
            op.GeoM.Translate(float64(x), float64(y))

            tileset_image.DrawImage(g.EntityStack.Stack[i].Image, op)

        }



    }

	return tileset
}

// creates new Tileset struct
func NewTilesetFromFile(path string, index int) Tileset {
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

// creates new Tileset struct from ready image
func NewTilesetFromImage(image *ebiten.Image) Tileset {
	var tileset Tileset

	tileset.image = image
	width, height := tileset.image.Size()
	tileset.width  = width
	tileset.height = height
	tileset.rows   = tileset.width / TileWidth
	tileset.cols   = tileset.height / TileHeight
	tileset.Num    = tileset.rows * tileset.cols

	return tileset
}

// creates new Tileseter struct
func NewTileseter() Tileseter {
	var tsr Tileseter

	return tsr 
}

// adds new Tileset to Tileseter, creating Tileset from image file
func (tsr *Tileseter) AddNewToTilesetter(path string) {
    index := len(tsr.tilesets)
    ts := NewTilesetFromFile(path, index)
    tsr.tilesets = append(tsr.tilesets, ts)
}

// appends new Tileset to Tileseter
func (tsr *Tileseter) AppendToTilesetter(ts Tileset) {
    ts.index = len(tsr.tilesets)
    tsr.tilesets = append(tsr.tilesets, ts)
}

// prepare Complete Tabber
func NewCompleteTilesetter() Tileseter {
    tsr := NewTileseter() 

    tsr.AddNewToTilesetter("assets/tileset_1.png")
    tsr.AddNewToTilesetter("assets/no_light.png")

    return tsr
}

// sets current Tileset index
func (tsr *Tileseter) SetCurrent (n int) {
    tsr.current = n
}

// prepare current Tileset index
func (tsr *Tileseter) GetCurrentIndex () int {
    return tsr.current
}

// prepare current Tileset index
func (tsr *Tileseter) GetTilesetByIndex(i int) Tileset {
    return tsr.tilesets[i]
}

// gets current Tileset
func (tsr *Tileseter) GetCurrent() Tileset {
    return tsr.tilesets[tsr.current]
}

// translates Tile index number (order in which Tile occurs on Pallete)
func (t *Tileset) TileNrToSubImageOnTileset(n int) *ebiten.Image {
	RectX, RectY := n%t.cols*TileWidth, n/t.rows*TileHeight
	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

	return t.image.SubImage(tileRect).(*ebiten.Image)
}

