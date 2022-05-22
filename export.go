package main

import (
	"archive/zip"
	"encoding/binary"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	lto "ltools/src/objects"
	"os"
)

const (
	N_COORDS       = 2
	TILES_PER_ROW  = 16
	TILESET_FORMAT = ".png"
	LEVEL_FORMAT   = ".llv"
	ARCHIVE_FORMAT = ".zip"
	BASENAME       = "level"
)

type TileStack struct {
	tiles map[*lto.Tile][]int
}

func NewTileStack() TileStack {
	var ts TileStack
	ts.tiles = make(map[*lto.Tile][]int)
	return ts
}

func (ts *TileStack) AddNew(tile *lto.Tile) {
	ts.tiles[tile] = []int{}
}

func (ts *TileStack) Append(tile *lto.Tile, x int, y int) {
	ts.tiles[tile] = append(ts.tiles[tile], x)
	ts.tiles[tile] = append(ts.tiles[tile], y)
}

// writes value as binary to file
func writeToFile(f *os.File, value int) {
	var value_encoded int16
	value_encoded = int16(value)
	binary.Write(f, binary.BigEndian, value_encoded)
}

func (ts *TileStack) hasTile(tile *lto.Tile) bool {
	_, isPresent := ts.tiles[tile]

	if isPresent == true {
		return true
	} else {
		return false
	}
}

// fills the stack with tiles used in level and position where it was used
// func (g *Game) FillStack(f *os.File, layer int, stack *TileStack) {
func (g *Game) FillStack() TileStack {
	stack := NewTileStack()

	rows, cols := g.Canvas.Size()
	layers := g.Canvas.Layers()

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for l := 0; l < layers; l++ {
				tile := g.Canvas.GetTileOnDrawingArea(x, y, l)

				if tile != nil {
					if stack.hasTile(tile) == true {
						// tile is no need to be filled
						stack.Append(tile, x, y)
					} else {
						// stack needs to be updated
						stack.AddNew(tile)
						stack.Append(tile, x, y)
					}
				}
			}
		}
	}

	return stack
}

// cropImage takes an image and crops it to the specified rectangle.
func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func (g *Game) PrepareTileset(filename string, stack TileStack) string {
	name := filename + TILESET_FORMAT
	n_tiles := len(stack.tiles)

	w := TILES_PER_ROW * GridSize
	h := ((n_tiles / TILES_PER_ROW) + 1) * GridSize
	r := image.Rectangle{image.Point{0, 0}, image.Point{int(w), int(h)}}
	export_image := image.NewRGBA(r)

	i := 0

	// preparing image
	for tile, _ := range stack.tiles {

		asset, err := os.Open(tile.GetTileTileset())
		fmt.Printf("asset name = %s\n", tile.GetTileTileset())

		if err != nil {
			fmt.Println(err)
			fmt.Println("wasssup!")
		}

		asset_image, _, err := image.Decode(asset)
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			fmt.Println(err)
		}

		// this is the coords of tile on export_image
		x := (i % TILES_PER_ROW) * GridSize
		y := (i / TILES_PER_ROW) * GridSize

		row, col := tile.PosOnTileset()

		// getting single tile from tileset image
		crop := image.Rect(row*GridSize, col*GridSize, row*GridSize+GridSize, col*GridSize+GridSize)
		single_tile, _ := cropImage(asset_image, crop)

		// putting tile on export_image
		draw.Draw(
			export_image,
			image.Rect(x, y, x+GridSize, y+GridSize),
			single_tile,
			image.Point{row * GridSize, col * GridSize},
			draw.Src,
		)
		i++
	}

	// saveing image to a file
	f, _ := os.Create(name)
	png.Encode(f, export_image)
	f.Close()

	return name
}

func (ts *TileStack) TilesPerLayer(layer int) []*lto.Tile {
	tiles := make([]*lto.Tile, 0)

	for tile, _ := range ts.tiles {
		if tile.GetLayer() == layer {
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func (ts *TileStack) TilesPerLayerSum(layer int) int {
	n := 0

	for tile, coords := range ts.tiles {
		if tile.GetLayer() == layer {
			n += len(coords)
		}
	}

	return n
}

func (g *Game) writeLayerToFile(f *os.File, layer int, stack TileStack) {
	tiles_per_layer := stack.TilesPerLayer(layer)
	tiles_per_layer_sum := stack.TilesPerLayerSum(layer)

	writeToFile(f, tiles_per_layer_sum)

	for _, tile := range tiles_per_layer {
		single_tile_usage := len(stack.tiles[tile])
		writeToFile(f, single_tile_usage)

		for _, coord := range stack.tiles[tile] {
			writeToFile(f, coord)
		}
	}
}

func (g *Game) exportLevel(filename string, stack TileStack) string {
	name := filename + LEVEL_FORMAT
	f, _ := os.Create(name)

	// first, write rows and cols number to file
	rows, cols := g.Canvas.Size()
	writeToFile(f, rows)
	writeToFile(f, cols)

	// then, write all layers content
	for layer := 0; layer < LAYER_N; layer++ {
		g.writeLayerToFile(f, layer, stack)
	}

	f.Close()

	return name
}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func CompresFiles(files []string, archive_name string) string {
	name := archive_name + ARCHIVE_FORMAT

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(name, flags, 0644)

	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for _, filename := range files {
		if err := appendFiles(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}

	for _, filename := range files {
		e := os.Remove(filename)
		if e != nil {
			log.Fatal(e)
		}
	}
	return name
}

func (g *Game) Export(screen *ebiten.Image) {
	stack := g.FillStack()

	levelname := g.exportLevel(BASENAME, stack)
	tileset_name := g.PrepareTileset(BASENAME, stack)

	CompresFiles([]string{levelname, tileset_name}, BASENAME)

	g.changeModeToDraw(screen)
}
