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
	tilesets []Tileset
	current  int
}

// create new Tileset from set of paths (could be only one path)
// this is propably not so smart take on transofmring any picture to fixed size pallete, but such
// implementation gets rid of traforming tile-on-image-position to tile-on-tileset-position fuss
func NewTileset(paths []string, index int) Tileset {
	var tileset Tileset
	var images []*ebiten.Image

	tileset.index = index
	n_tiles := 0

	// calculating overall timeset size
	for i := 0; i < len(paths); i++ {
		image := loadImage(paths[i])
		images = append(images, image)

		width, height := image.Size()
		n_tiles += (width / TileWidth) * (height / TileHeight)
	}

	// new tileset image is created to hold all tiles on all images inputed
	tileset_image := ebiten.NewImage(PalleteColsN*TileWidth, (n_tiles/PalleteColsN+1)*TileHeight)

	// iterating one more time, this time filling up the tileset
	for j := 0; j < len(images); j++ {
		width, height := images[j].Size()
		rows := width / TileWidth
		cols := height / TileHeight
		n := rows * cols

		// for each tileset, each tile is obtained an put into right place of tileset
		// this fuss is required mostly to always have tileset with same size no matter what image
		// is inputted
		for z := 0; z < n; z++ {
			// tile cut from original image
			rect_x := (z % cols) * TileWidth
			rect_y := (z / cols) * TileHeight

			tileRect := image.Rect(rect_x, rect_y, rect_x+TileWidth, rect_y+TileHeight)
			tile := images[j].SubImage(tileRect).(*ebiten.Image)

			// tile placed on new position
			x := (z % PalleteColsN) * TileWidth
			y := (z / PalleteColsN) * TileHeight
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))

			tileset_image.DrawImage(tile, op)
		}
	}

	tileset.image = tileset_image
	width, height := tileset.image.Size()
	tileset.width = width
	tileset.height = height
	tileset.rows = tileset.height / TileWidth

	tileset.cols = PalleteColsN
	tileset.Num = n_tiles
	tileset.index = index

	return tileset
}

// creates new Tileseter struct
func NewTileseter(tilesets_paths [][]string) Tileseter {
	var tsr Tileseter

	for i := 0; i < len(tilesets_paths); i++ {
		tsr.AddNewToTilesetter(tilesets_paths[i])
	}

	return tsr
}

// adds new Tileset to Tileseter, creating Tileset from image file
func (tsr *Tileseter) AddNewToTilesetter(paths []string) {
	index := len(tsr.tilesets)
	ts := NewTileset(paths, index)
	tsr.tilesets = append(tsr.tilesets, ts)
}

// sets current Tileset index
func (tsr *Tileseter) SetCurrent(n int) {
	tsr.current = n
}

// prepare current Tileset index
func (tsr *Tileseter) GetCurrentIndex() int {
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
