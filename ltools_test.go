package main

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io/fs"
	lto "ltools/src/objects"
	"os"
	"path/filepath"
	"testing"
)

func TestStack(t *testing.T) {
	tileset_name := "name"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := lto.NewTile(i2, tileset_name, false, 0, 1, 0)

	ts := NewTileStack()

	ts.AddNew(&t1)
	ts.Append(&t1, 10, 10)
	ts.Append(&t1, 20, 20)

	ts.AddNew(&t2)
	ts.Append(&t2, 30, 30)

	if len(ts.tiles) != 2 {
		t.Errorf("Adding tiles to stack failed")
	}
	if len(ts.tiles[0].coords) != 4 {
		t.Errorf("Adding tiles to stack failed, want 2 coords, got %d", len(ts.tiles[0].coords))
	}
	if len(ts.tiles[1].coords) != 2 {
		t.Errorf("Adding tiles to stack failed, want 4 coords, got %d", len(ts.tiles[0].coords))
	}
}

func TestHasTile(t *testing.T) {
	tileset_name := "name"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := lto.NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := lto.NewTile(i3, tileset_name, false, 0, 2, 0)

	ts := NewTileStack()

	ts.AddNew(&t1)
	ts.AddNew(&t2)

	if ts.hasTile(&t1) == false {
		t.Errorf("Adding tiles to stack failed")
	}

	if ts.hasTile(&t2) == false {
		t.Errorf("Adding tiles to stack failed")
	}

	if ts.hasTile(&t3) == true {
		t.Errorf("Adding tiles to stack failed")
	}

	ts.AddNew(&t3)

	if ts.hasTile(&t3) == false {
		t.Errorf("Adding tiles to stack failed")
	}
}

func TestFillStack(t *testing.T) {
	var g Game

	tileset_name := "name"

	g.Canvas = lto.NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := lto.NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := lto.NewTile(i3, tileset_name, false, 0, 2, 1)

	g.Canvas.SetTileOnDrawingArea(10, 10, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(11, 11, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(12, 12, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(13, 13, 0, &t1)

	g.Canvas.SetTileOnDrawingArea(14, 14, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(15, 15, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(16, 16, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(17, 17, 0, &t2)

	g.Canvas.SetTileOnDrawingArea(5, 5, 1, &t3)
	g.Canvas.SetTileOnDrawingArea(6, 6, 1, &t3)

	ts := g.FillStack()

	if ts.hasTile(&t1) == false {
		t.Errorf("Filling stack does not succeded")
	}
	if ts.hasTile(&t2) == false {
		t.Errorf("Filling stack does not succeded")
	}
	if ts.hasTile(&t3) == false {
		t.Errorf("Filling stack does not succeded")
	}

	coords_a := []int{10, 10, 11, 11, 12, 12, 13, 13}
	coords_b := []int{14, 14, 15, 15, 16, 16, 17, 17}
	coords_c := []int{5, 5, 6, 6}

	if len(ts.tiles[0].coords) != len(coords_a) {
		t.Errorf("coords do not comply")
	}

	if len(ts.tiles[1].coords) != len(coords_b) {
		t.Errorf("coords do not comply")
	}
	if len(ts.tiles[2].coords) != len(coords_c) {
		t.Errorf("coords do not comply")
	}

	for i, coord := range ts.tiles[0].coords {
		if coord != coords_a[i] {
			t.Errorf("Bad coords in tile stack, want: %d, got: %d", coords_a[i], coord)
		}
	}

	for i, coord := range ts.tiles[1].coords {
		if coord != coords_b[i] {
			t.Errorf("Bad coords in tile stack, want: %d, got: %d", coords_b[i], coord)
		}
	}

	for i, coord := range ts.tiles[2].coords {
		if coord != coords_c[i] {
			t.Errorf("Bad coords in tile stack, want: %d, got: %d", coords_c[i], coord)
		}
	}
}

func TestPrepareTiles(t *testing.T) {
	var g Game

	g.Canvas = lto.NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	tileset_name := "assets/tilesets/basic_tileset.png"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := lto.NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := lto.NewTile(i3, tileset_name, false, 0, 2, 1)

	g.Canvas.SetTileOnDrawingArea(10, 10, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(11, 11, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(12, 12, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(13, 13, 0, &t1)

	g.Canvas.SetTileOnDrawingArea(14, 14, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(15, 15, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(16, 16, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(17, 17, 0, &t2)

	g.Canvas.SetTileOnDrawingArea(5, 5, 1, &t3)
	g.Canvas.SetTileOnDrawingArea(6, 6, 1, &t3)

	stack := g.FillStack()
	g.PrepareTileset("sample_name", stack)
}

func TestTilesPerLayer(t *testing.T) {
	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, "name", false, 0, 0, 0)
	t2 := lto.NewTile(i2, "name", false, 0, 0, 0)
	t3 := lto.NewTile(i3, "name", false, 0, 0, 1)

	ts := NewTileStack()

	ts.AddNew(&t1)
	ts.Append(&t1, 10, 10)
	ts.Append(&t1, 20, 20)
	ts.Append(&t1, 30, 30)

	ts.AddNew(&t2)
	ts.Append(&t2, 30, 30)

	ts.AddNew(&t3)
	ts.Append(&t3, 30, 30)

	tiles_per_layer_0 := ts.TilesPerLayer(0)

	if len(tiles_per_layer_0) != 2 {
		t.Errorf("counting stack tiles does not suceeded")
	}

	tiles_per_layer_1 := ts.TilesPerLayer(1)

	if len(tiles_per_layer_1) != 1 {
		t.Errorf("counting stack tiles does not suceeded")
	}

	tiles_per_layer_2 := ts.TilesPerLayer(2)

	if len(tiles_per_layer_2) != 0 {
		t.Errorf("counting stack tiles does not suceeded")
	}

}
func TestTilesPerLayerSum(t *testing.T) {
	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, "name", false, 0, 0, 0)
	t2 := lto.NewTile(i2, "name", false, 0, 0, 0)
	t3 := lto.NewTile(i3, "name", false, 0, 0, 1)

	ts := NewTileStack()

	ts.AddNew(&t1)
	ts.Append(&t1, 10, 10)
	ts.Append(&t1, 20, 20)
	ts.Append(&t1, 30, 30)
	ts.Append(&t1, 30, 30)

	ts.AddNew(&t2)
	ts.Append(&t2, 30, 30)
	ts.Append(&t2, 30, 30)

	ts.AddNew(&t3)
	ts.Append(&t3, 30, 30)

	tiles_per_layer_0 := ts.TilesPerLayerSum(0)

	if tiles_per_layer_0 != 12 {
		t.Errorf("counting stack tiles does not suceeded, want 8, got %d", tiles_per_layer_0)
	}

	tiles_per_layer_1 := ts.TilesPerLayerSum(1)

	if tiles_per_layer_1 != 2 {
		t.Errorf("counting stack tiles does not suceeded, want 4, got %d", tiles_per_layer_1)
	}

	tiles_per_layer_2 := ts.TilesPerLayerSum(2)

	if tiles_per_layer_2 != 0 {
		t.Errorf("counting stack tiles does not suceeded, want 0, got %d", tiles_per_layer_2)
	}
}

func TestExportLevel(t *testing.T) {
	var g Game

	name := "czeslaw"
	tileset_name := "assets/tilesets/basic_tileset.png"
	g.Canvas = lto.NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := lto.NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := lto.NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := lto.NewTile(i3, tileset_name, false, 0, 2, 1)

	g.Canvas.SetTileOnDrawingArea(10, 10, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(11, 11, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(12, 12, 0, &t1)
	g.Canvas.SetTileOnDrawingArea(13, 13, 0, &t1)

	g.Canvas.SetTileOnDrawingArea(14, 14, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(15, 15, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(16, 16, 0, &t2)
	g.Canvas.SetTileOnDrawingArea(17, 17, 0, &t2)

	g.Canvas.SetTileOnDrawingArea(5, 5, 1, &t3)
	g.Canvas.SetTileOnDrawingArea(6, 6, 1, &t3)

	stack := g.FillStack()
	g.exportLevel(name, stack)

	// file was not created - test should fail
	if _, err := os.Stat(name + ".llv"); errors.Is(err, fs.ErrNotExist) {
		t.Errorf("Exporintg level does not succeded!\n")
		fmt.Print(err.Error())
	}
}

func TestExport(t *testing.T) {
	var tests = []struct {
		tiles_added       int
		tiles_added_funcs [][]int
		tiles_set         int
		tiles_set_funcs   [][]int
	}{
		{
			0,
			[][]int{{}},
			0,
			[][]int{{}},
		},
		{
			3,
			[][]int{{0, 0, 0}, {1, 1, 0}, {2, 2, 0}},
			5,
			[][]int{{10, 10, 0, 0}, {11, 11, 0, 1}, {12, 12, 0, 1}, {13, 13, 0, 2}, {14, 14, 0, 2}},
		},
		{
			3,
			[][]int{{0, 0, 0}, {0, 0, 1}, {0, 0, 2}},
			5,
			[][]int{{10, 10, 0, 0}, {11, 11, 0, 1}, {12, 12, 0, 1}, {13, 13, 0, 2}, {14, 14, 0, 2}},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.tiles_set_funcs)

		t.Run(testname, func(t *testing.T) {
			g := NewGame()
			tileset_name := "assets/tilesets/basic_tileset.png"

			// this is just a shortcut for selecting tile by index and not by tile
			tiles := make([]*lto.Tile, 0)

			for i := 0; i < tt.tiles_added; i++ {
				img := ebiten.NewImage(32, 32) // every image is a dummy
				t := lto.NewTile(
					img,
					tileset_name,
					false,
					tt.tiles_added_funcs[i][0],
					tt.tiles_added_funcs[i][1],
					tt.tiles_added_funcs[i][2],
				)
				tiles = append(tiles, &t)
			}

			for j := 0; j < tt.tiles_set; j++ {
				g.Canvas.SetTileOnDrawingArea(
					tt.tiles_set_funcs[j][0],
					tt.tiles_set_funcs[j][1],
					tt.tiles_set_funcs[j][2],
					tiles[tt.tiles_set_funcs[j][3]],
				)
			}

			// dummy image - it is required
			dummy_image := ebiten.NewImage(1, 1)
			g.Export(dummy_image)

			base_path := filepath.Join(g.Config.LevelDir, BASENAME)

			// level structure was not created - test should fail
			if _, err := os.Stat(base_path + ".llv"); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("Exporintg level does not succeded! - level structure was not exported\n")
				fmt.Print(err.Error())
			}

			// level structure was not created - test should fail
			if _, err := os.Stat(base_path + ".png"); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("Exporintg level does not succeded! - tileset image was not exported\n")
				fmt.Print(err.Error())
			}
		})
	}
}

func TestDrawTileOnCanvasDraw(t *testing.T) {

	normal_tile := lto.NewTile(ebiten.NewImage(32, 32), "name", false, 0, 0, 0)
	normal_tile_2 := lto.NewTile(ebiten.NewImage(32, 32), "name", false, 0, 0, 0)
	light_tile := lto.NewTile(ebiten.NewImage(32, 32), "name", false, 0, 0, 1)
	entity_tile := lto.NewTile(ebiten.NewImage(32, 32), "name", false, 0, 0, 2)

	modes := make(map[*lto.Tile]int)

	modes[&normal_tile] = 0
	modes[&normal_tile_2] = 0
	modes[&light_tile] = 1
	modes[&entity_tile] = 2

	screen := ebiten.NewImage(640, 480)

	RUBBER := 1
	PEN := 0

	var tests = []struct {
		n               int
		tiles_set       []*lto.Tile
		tool_used       []int
		expected_result []*lto.Tile
	}{
		{
			1,
			[]*lto.Tile{&normal_tile},
			[]int{PEN},
			[]*lto.Tile{&normal_tile, nil, nil},
		},
		{
			2,
			[]*lto.Tile{&normal_tile, &normal_tile_2},
			[]int{PEN, PEN},
			[]*lto.Tile{&normal_tile_2, nil, nil},
		},
		{
			3,
			[]*lto.Tile{&normal_tile, &normal_tile_2, &normal_tile},
			[]int{PEN, PEN, PEN},
			[]*lto.Tile{&normal_tile, nil, nil},
		},
		{
			2,
			[]*lto.Tile{&normal_tile, &light_tile},
			[]int{PEN, PEN},
			[]*lto.Tile{&normal_tile, nil, nil},
		},
		{
			2,
			[]*lto.Tile{&light_tile, &normal_tile},
			[]int{PEN, PEN},
			[]*lto.Tile{&normal_tile, nil, nil},
		},
		{
			1,
			[]*lto.Tile{&light_tile},
			[]int{PEN},
			[]*lto.Tile{nil, &light_tile, nil},
		},
		{
			2,
			[]*lto.Tile{&normal_tile, &normal_tile},
			[]int{PEN, RUBBER},
			[]*lto.Tile{nil, nil, nil},
		},
		{
			3,
			[]*lto.Tile{&normal_tile, &normal_tile, &normal_tile_2},
			[]int{PEN, RUBBER, PEN},
			[]*lto.Tile{&normal_tile_2, nil, nil},
		},
		{
			3,
			[]*lto.Tile{&light_tile, &light_tile, &light_tile},
			[]int{PEN, RUBBER, PEN},
			[]*lto.Tile{nil, &light_tile, nil},
		},
		{
			3,
			[]*lto.Tile{&light_tile, &light_tile, &normal_tile},
			[]int{PEN, RUBBER, PEN},
			[]*lto.Tile{&normal_tile, nil, nil},
		},
		{
			2,
			[]*lto.Tile{&normal_tile, &entity_tile},
			[]int{PEN, PEN, PEN},
			[]*lto.Tile{&normal_tile, nil, &entity_tile},
		},
		{
			3,
			[]*lto.Tile{&normal_tile, &entity_tile, &normal_tile_2},
			[]int{PEN, PEN, PEN},
			[]*lto.Tile{&normal_tile_2, nil, &entity_tile},
		},
		{
			3,
			[]*lto.Tile{&normal_tile, &entity_tile, &entity_tile},
			[]int{PEN, PEN, RUBBER},
			[]*lto.Tile{&normal_tile, nil, nil},
		},
		{
			3,
			[]*lto.Tile{&normal_tile, &entity_tile, &normal_tile},
			[]int{PEN, PEN, RUBBER},
			[]*lto.Tile{nil, nil, nil},
		},
		{
			3,
			[]*lto.Tile{&normal_tile, &entity_tile, &light_tile},
			[]int{PEN, PEN, RUBBER},
			[]*lto.Tile{&normal_tile, nil, &entity_tile},
		},
		{
			3,
			[]*lto.Tile{&light_tile, &normal_tile, &entity_tile},
			[]int{PEN, PEN, PEN},
			[]*lto.Tile{&normal_tile, nil, &entity_tile},
		},
		{
			2,
			[]*lto.Tile{&light_tile, &light_tile},
			[]int{PEN, RUBBER},
			[]*lto.Tile{nil, nil, nil},
		},
	}
	// WHEN

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)

		t.Run(testname, func(t *testing.T) {
			var g Game
			g.Canvas = lto.NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)
			g.Toolbox = lto.NewToolbox(ToolboxX, ToolboxY)

			for i := 0; i < tt.n; i++ {
				g.mouse_x = 10
				g.mouse_y = 10
				g.mode = modes[tt.tiles_set[i]]
				g.Toolbox.Activate(tt.tool_used[i])
				g.Toolbox.SetFillTile(tt.tiles_set[i])
				g.DrawTileOnCanvas(screen)
			}

			if g.Canvas.GetTileOnDrawingArea(0, 0, 0) != tt.expected_result[0] {
				t.Errorf("setting Tile on Canvas on layer DRAW failed")
			}

			if g.Canvas.GetTileOnDrawingArea(0, 0, 1) != tt.expected_result[1] {
				t.Errorf("setting Tile on Canvas on layer LIGHT failed")
			}

			if g.Canvas.GetTileOnDrawingArea(0, 0, 2) != tt.expected_result[2] {
				t.Errorf("setting Tile on Canvas on layer ENTITY failed")
			}
		})
	}
}
