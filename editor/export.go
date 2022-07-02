package editor

import (
	// "encoding/binary"
	// "fmt"
	// "github.com/hajimehoshi/ebiten/v2"
	// "image"
	// "image/draw"
	// "image/png"
	// "os"
	// "path/filepath"
)

// const (
// 	N_COORDS       = 2
// 	TILES_PER_ROW  = 16
// 	TILESET_FORMAT = ".png"
// 	LEVEL_FORMAT   = ".llv"
// 	BASENAME       = "level"
// )
// 
// type TileStackEntry struct {
// 	tile   *Tile
// 	coords []int // coords of tile
// }
// 
// type TileStack struct {
// 	tiles []TileStackEntry
// }

// func NewTileStack() TileStack {
// 	var ts TileStack
// 
// 	ts.tiles = make([]TileStackEntry, 0)
// 
// 	return ts
// }
// 
// func NewTileStackEntry(tile *Tile) TileStackEntry {
// 	var tse TileStackEntry
// 	tse.tile = tile
// 	tse.coords = make([]int, 0)
// 
// 	return tse
// }
// 
// func (ts *TileStack) AddNew(tile *Tile) {
// 	ts.tiles = append(ts.tiles, NewTileStackEntry(tile))
// }
// 
// func (ts *TileStack) Append(tile *Tile, x int, y int) {
// 	for i, entry := range ts.tiles {
// 		if entry.tile == tile {
// 			ts.tiles[i].coords = append(ts.tiles[i].coords, x)
// 			ts.tiles[i].coords = append(ts.tiles[i].coords, y)
// 			return
// 		}
// 	}
// }
// 
// // writes value as binary to file
// func writeToFile(f *os.File, value int) {
// 	var value_encoded int16
// 	value_encoded = int16(value)
// 	binary.Write(f, binary.BigEndian, value_encoded)
// }
// 
// func (ts *TileStack) hasTile(tile *Tile) bool {
// 	for _, entry := range ts.tiles {
// 		if entry.tile == tile {
// 			return true
// 		}
// 	}
// 
// 	return false
// }
// 
// // fills the stack with tiles used in level and position where it was used
// // func (g *Game) FillStack(f *os.File, layer int, stack *TileStack) {
// func (g *Game) FillStack() TileStack {
// 	stack := NewTileStack()
// 
// 	rows, cols := g.Canvas.Size()
// 	layers := g.Canvas.Layers()
// 
// 	for l := 0; l < layers; l++ {
// 		for x := 0; x < rows; x++ {
// 			for y := 0; y < cols; y++ {
// 				tile := g.Canvas.GetTileOnDrawingArea(x, y, l)
// 
// 				if tile != nil {
// 					if stack.hasTile(tile) == true {
// 						// tile is no need to be filled
// 						stack.Append(tile, x, y)
// 					} else {
// 						// stack needs to be updated
// 						stack.AddNew(tile)
// 						stack.Append(tile, x, y)
// 					}
// 				}
// 			}
// 		}
// 	}
// 
// 	return stack
// }
// 
// // cropImage takes an image and crops it to the specified rectangle.
// func cropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
// 	type subImager interface {
// 		SubImage(r image.Rectangle) image.Image
// 	}
// 
// 	// img is an Image interface. This checks if the underlying value has a
// 	// method called SubImage. If it does, then we can use SubImage to crop the
// 	// image.
// 	simg, ok := img.(subImager)
// 	if !ok {
// 		return nil, fmt.Errorf("image does not support cropping")
// 	}
// 
// 	return simg.SubImage(crop), nil
// }
// 
// func (g *Game) PrepareTileset(filename string, stack TileStack) string {
// 	name := filename + TILESET_FORMAT
// 	n_tiles := len(stack.tiles)
// 
// 	w := TILES_PER_ROW * GridSize
// 	h := ((n_tiles / TILES_PER_ROW) + 1) * GridSize
// 	r := image.Rectangle{image.Point{0, 0}, image.Point{int(w), int(h)}}
// 	export_image := image.NewRGBA(r)
// 
// 	i := 0
// 
// 	// preparing image
// 	for _, entry := range stack.tiles {
// 		asset, err := os.Open(entry.tile.GetTileTileset())
// 
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 
// 		asset_image, _, err := image.Decode(asset)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 
// 		// this is the coords of tile on export_image
// 		x := (i % TILES_PER_ROW) * GridSize
// 		y := (i / TILES_PER_ROW) * GridSize
// 
// 		row, col := entry.tile.PosOnTileset()
// 
// 		// getting single tile from tileset image
// 		crop := image.Rect(row*GridSize, col*GridSize, row*GridSize+GridSize, col*GridSize+GridSize)
// 		single_tile, _ := cropImage(asset_image, crop)
// 
// 		// putting tile on export_image
// 		draw.Draw(
// 			export_image,
// 			image.Rect(x, y, x+GridSize, y+GridSize),
// 			single_tile,
// 			image.Point{row * GridSize, col * GridSize},
// 			draw.Src,
// 		)
// 		i++
// 	}
// 
// 	// saveing image to a file
// 	f, _ := os.Create(name)
// 	png.Encode(f, export_image)
// 	f.Close()
// 
// 	return name
// }
// 
// func (ts *TileStack) TilesPerLayer(layer int) []int {
// 	tiles := make([]int, 0)
// 
// 	for i, entry := range ts.tiles {
// 		if entry.tile.GetLayer() == layer {
// 			tiles = append(tiles, i)
// 		}
// 	}
// 	return tiles
// }
// 
// func (ts *TileStack) TilesPerLayerSum(layer int) int {
// 	n := 0
// 
// 	for _, entry := range ts.tiles {
// 		if entry.tile.GetLayer() == layer {
// 			n += len(entry.coords)
// 		}
// 	}
// 
// 	return n
// }
// 
// func (g *Game) writeLayerToFile(f *os.File, layer int, stack TileStack) {
// 	tiles_per_layer := stack.TilesPerLayer(layer)
// 	tiles_per_layer_sum := stack.TilesPerLayerSum(layer)
// 
// 	writeToFile(f, tiles_per_layer_sum)
// 
// 	for _, tile := range tiles_per_layer {
// 		single_tile_usage := len(stack.tiles[tile].coords)
// 		writeToFile(f, single_tile_usage)
// 
// 		for _, coord := range stack.tiles[tile].coords {
// 			writeToFile(f, coord)
// 		}
// 	}
// }
// 
// func (g *Game) exportLevel(filename string, stack TileStack) string {
// 	name := filename + LEVEL_FORMAT
// 	f, _ := os.Create(name)
// 
// 	// first, write rows and cols number to file
// 	rows, cols := g.Canvas.Size()
// 	writeToFile(f, rows)
// 	writeToFile(f, cols)
// 
// 	// then, write all layers content
// 	for layer := 0; layer < LAYER_N; layer++ {
// 		g.writeLayerToFile(f, layer, stack)
// 	}
// 
// 	f.Close()
// 
// 	return name
// }
// 
// func (g *Game) Export(screen *ebiten.Image) {
// 	stack := g.FillStack()
// 
// 	_ = os.Mkdir(g.Config.LevelDir, 0750)
// 
// 	base_path := filepath.Join(g.Config.LevelDir, BASENAME)
// 
// 	g.exportLevel(base_path, stack)
// 	g.PrepareTileset(base_path, stack)
// 
// 	// went back to base mode
// 	g.changeModeToDraw(screen)
// }
