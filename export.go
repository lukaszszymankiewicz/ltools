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
	"os"
)

const (
	N_COORDS       = 2
	TILES_PER_ROW  = 16
	TILESET_FORMAT = ".png"
	LEVEL_FORMAT   = ".llv"
	ARCHIVE_FORMAT = ".zip"
	// by now changing level name is not supported
	BASENAME = "level"
)

// writes value as binary to file
func writeToFile(f *os.File, value int) {
	var value_encoded int16
	value_encoded = int16(value)
	binary.Write(f, binary.BigEndian, value_encoded)
}

// converts one layer data into sparse matrix
func (g *Game) writeLayerToFile(f *os.File, layer int) {
	tiles_indexes := g.TileIndexPerLayer(layer)
	all_tiles_n := g.TileStack.Length()

	tiles_matrix := g.Canvas.ExportTiles(all_tiles_n, layer)
	tiles_used := g.NumberOfTilesOnLayer(layer)
	writeToFile(f, tiles_used)

	for t := 0; t < tiles_used; t++ {
		// add index (used only for entities)
		writeToFile(f, t)

		idx := tiles_indexes[t]
		n := g.TileUsage(idx)
		writeToFile(f, n)

        if n > 0 {
            for i := 0; i < len(tiles_matrix[t]); i++ {
                writeToFile(f, tiles_matrix[t][i])
            }
        }
	}
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

// unfortunatly ebiten.Image does not supprt exporintg image to file, so raw Image.image from
// golang build-in library is use - we got everything!
// second thing - by now level use only ONE hardcoded asset, so it is easy to determine what
// needs to be loaded here but it is only a temporary solution
func (g *Game) PrepareTiles(filename string) string {
	name := filename + TILESET_FORMAT

	asset, err := os.Open("assets/tileset_1.png")
	if err != nil {
		fmt.Println(err)
	}

	asset_image, _, err := image.Decode(asset)
	if err != nil {
		fmt.Println(err)
	}

	// only draw tiles are put into tileset png file
	tiles_indexes := g.TileIndexPerLayer(LAYER_DRAW)
	n_tiles := len(tiles_indexes)

	w := TILES_PER_ROW * TileWidth
	h := ((n_tiles / TILES_PER_ROW) + 1) * TileHeight
	r := image.Rectangle{image.Point{0, 0}, image.Point{int(w), int(h)}}

	// this image will be exported
	export_image := image.NewRGBA(r)

	// read tiles directly from TileStack and put it into prepared image
	for i := 0; i < n_tiles; i++ {
		// this is the coords of tile on export_image
		x := (i % TILES_PER_ROW) * TileWidth
		y := (i / TILES_PER_ROW) * TileHeight

		// getting single tile from tileset image
		row, col := g.TileStack.GetTilePos(i)
		crop := image.Rect(row*TileWidth, col*TileHeight, row*TileWidth+TileWidth, col*TileHeight+TileHeight)
		single_tile, _ := cropImage(asset_image, crop)

		// putting tile on export_image
		draw.Draw(
			export_image,
			image.Rect(x, y, x+TileWidth, y+TileHeight),
			single_tile,
			image.Point{row * TileWidth, col * TileHeight},
			draw.Src,
		)
	}

	// saveing image to a file
	f, err := os.Create(name)
	png.Encode(f, export_image)
	f.Close()

	return name
}

// exports level as sparse matrix form
// returns created file if succeded
func (g *Game) exportLevel(filename string) string {
	name := filename + LEVEL_FORMAT
	g.purgeStack()

	f, _ := os.Create(name)
	rows, cols := g.Canvas.Size()

	writeToFile(f, rows)
	writeToFile(f, cols)

	g.writeLayerToFile(f, LAYER_DRAW)
	g.writeLayerToFile(f, LAYER_LIGHT)
	g.writeLayerToFile(f, LAYER_ENTITY)

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

// compress files needed to export level
func (g *Game) Export(screen *ebiten.Image) {
	levelname := g.exportLevel(BASENAME)
	tileset_name := g.PrepareTiles(BASENAME)

	CompresFiles([]string{levelname, tileset_name}, BASENAME)

    g.FindStartingTile()
    g.changeModeToDraw(screen)
}
