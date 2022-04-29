package main

import (
	"encoding/binary"
	lto "ltools/src/objects"
	"os"
	// "fmt"
)

const (
	ENCODING_BIG_INT = iota
	ENCODING_SMALL_INT
)

// looks for tiles which are on a stack but are unused (it can happen after heavy session of
// cleaning the screen)
func (g *Game) purgeStack() {
	new_tiles := make([]lto.Tile, 0)
	rows, cols := g.Canvas.Size()

	// check is done for each tile in tilestack
	for i := 0; i < len(g.GetAllTiles()); i++ {

		// if usage of give tile is equal (or lower - just for sanity check) than zero
		// whole procedure is perfmormed
		if n := g.TileUsage(i); n == 0 {

			// tile is deleted from stack
			// because stack has it lenght (and content changed, other tile on canvas must be
			// updated to such change)
			for layer := 0; layer < g.Canvas.NumberOfLayers(); layer++ {
				for x := 0; x < rows; x++ {
					for y := 0; y < cols; y++ {

						if tile_nr := g.Canvas.GetTileOnDrawingArea(x, y, layer); tile_nr > i {
							g.Canvas.SetTileOnCanvas(x, y, tile_nr-1, layer)
						}
					}
				}
			}
		} else {
			new_tiles = append(new_tiles, g.TileStack.GetTileFromStack(i))
		}
	}
	g.TileStack.SetAllTiles(new_tiles)
}

func writeToFile(f *os.File, value int, size int) {
	if size == ENCODING_BIG_INT {
		var value_encoded int32
		value_encoded = int32(value)
		binary.Write(f, binary.BigEndian, value_encoded)
	} else if size == ENCODING_SMALL_INT {
		var value_encoded int16
		value_encoded = int16(value)
		binary.Write(f, binary.BigEndian, value_encoded)
	}
}

func (g *Game) exportLevel() {
	f, _ := os.Create("temp.bin")

	rows, cols := g.Canvas.Size()

	// tiles used on LIGHT_LAYER and ENTITY_LAYER are hardcoded by now so there is no need to count
	// them whatsoever
	tiles_used := g.NumberOfTilesOnLayer(LAYER_DRAW)

	writeToFile(f, rows, ENCODING_BIG_INT)
	writeToFile(f, cols, ENCODING_BIG_INT)
	writeToFile(f, tiles_used, ENCODING_SMALL_INT)

	f.Close()
}
