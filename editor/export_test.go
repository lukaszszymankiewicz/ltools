package editor

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io/fs"
	"os"
	"testing"
)

func TestStack(t *testing.T) {
	tileset_name := "name"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, 0, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, 0, 1, 0, 1)

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

func TestFillStack(t *testing.T) {
	var g Game
	var idx int
	tileset_name := "name"

	g.ViewportCanvas = NewViewport(30, 30, 30, 30)
	g.LevelStructure = NewTilesCollection(3)
	g.LevelStructure.Alloc(30 * 30)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, 0, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, 0, 1, 0, 1)
	t3 := NewTile(i3, tileset_name, 0, 2, 1, 2)

	idx = g.ViewportCanvas.GetTileIdx(10, 10)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(11, 11)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(12, 12)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(13, 13)
	g.LevelStructure.Set(idx, 0, &t1)

	idx = g.ViewportCanvas.GetTileIdx(14, 14)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(15, 15)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(16, 16)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(17, 17)
	g.LevelStructure.Set(idx, 0, &t2)

	idx = g.ViewportCanvas.GetTileIdx(5, 5)
	g.LevelStructure.Set(idx, 1, &t3)
	idx = g.ViewportCanvas.GetTileIdx(6, 6)
	g.LevelStructure.Set(idx, 1, &t3)

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

func TestPrepareTileset(t *testing.T) {
	var g Game
	var idx int

	g.ViewportCanvas = NewViewport(30, 30, 30, 30)
	g.LevelStructure = NewTilesCollection(3)
	g.LevelStructure.Alloc(30 * 30)

	LoadResources("")

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	tileset_name := GetResourceFullName("basic_tileset")

	t1 := NewTile(i1, tileset_name, 0, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, 0, 1, 0, 1)
	t3 := NewTile(i3, tileset_name, 0, 2, 1, 2)

	idx = g.ViewportCanvas.GetTileIdx(10, 10)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(11, 11)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(12, 12)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(13, 13)
	g.LevelStructure.Set(idx, 0, &t1)

	idx = g.ViewportCanvas.GetTileIdx(14, 14)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(15, 15)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(16, 16)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(17, 17)
	g.LevelStructure.Set(idx, 0, &t2)

	idx = g.ViewportCanvas.GetTileIdx(5, 5)
	g.LevelStructure.Set(idx, 1, &t3)
	idx = g.ViewportCanvas.GetTileIdx(6, 6)
	g.LevelStructure.Set(idx, 1, &t3)

	stack := g.FillStack()
	g.PrepareTileset("sample_name", stack)
}

func TestTilesPerLayer(t *testing.T) {
	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, "name", 0, 0, 0, 0)
	t2 := NewTile(i2, "name", 0, 0, 0, 1)
	t3 := NewTile(i3, "name", 0, 0, 1, 2)

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

	t1 := NewTile(i1, "name", 0, 0, 0, 0)
	t2 := NewTile(i2, "name", 0, 0, 0, 1)
	t3 := NewTile(i3, "name", 0, 0, 1, 2)

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
	var idx int

	name := "czeslaw"
	tileset_name := "assets/basic_tileset.png"

	g.ViewportCanvas = NewViewport(30, 20, 30, 20)
	g.LevelStructure = NewTilesCollection(3)
	g.LevelStructure.Alloc(30 * 20)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, 0, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, 0, 1, 0, 1)
	t3 := NewTile(i3, tileset_name, 0, 2, 1, 2)

	idx = g.ViewportCanvas.GetTileIdx(10, 10)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(11, 11)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(12, 12)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(13, 13)
	g.LevelStructure.Set(idx, 0, &t1)
	idx = g.ViewportCanvas.GetTileIdx(14, 14)

	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(15, 15)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(16, 16)
	g.LevelStructure.Set(idx, 0, &t2)
	idx = g.ViewportCanvas.GetTileIdx(17, 17)
	g.LevelStructure.Set(idx, 0, &t2)

	idx = g.ViewportCanvas.GetTileIdx(5, 5)
	g.LevelStructure.Set(idx, 1, &t3)
	idx = g.ViewportCanvas.GetTileIdx(6, 6)
	g.LevelStructure.Set(idx, 1, &t3)

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
			var g Game

			g.AvailableTiles = NewTilesCollection(3)
			g.AvailableTiles.Alloc(30 * 30)

			g.LevelStructure = NewTilesCollection(LAYER_N)
			g.LevelStructure.Alloc(30 * 30)
			g.ViewportCanvas = NewViewport(30, 30, 30, 30)

			g.Mode = NewMode(MODE_ALL)

			// RESOURCES
			LoadResources("")

			g.AvailableTiles.Fill("basic_tileset", LAYER_DRAW)
			g.AvailableTiles.Fill("light_tileset", LAYER_LIGHT)
			g.AvailableTiles.Fill("entity_tileset", LAYER_ENTITY)

			tiles := make([]*Tile, 0)

			for i := 0; i < tt.tiles_added; i++ {
				// every image is a dummy
				img := ebiten.NewImage(32, 32)
				t := NewTile(
					img,
					GetResourceFullName("basic_tileset"),
					tt.tiles_added_funcs[i][0],
					tt.tiles_added_funcs[i][1],
					tt.tiles_added_funcs[i][2],
                    0,
				)
				tiles = append(tiles, &t)
			}

			for j := 0; j < tt.tiles_set; j++ {
				var idx int

				idx = g.ViewportCanvas.GetTileIdx(
					tt.tiles_set_funcs[j][0],
					tt.tiles_set_funcs[j][1],
				)

				g.LevelStructure.Set(
					idx,
					tt.tiles_set_funcs[j][2],
					tiles[tt.tiles_set_funcs[j][3]],
				)
			}

			// ommit checks
			TRIGGER_EXPORT_CHECKS = 0
			g.Export()

			// level structure was not created - test should fail
			if _, err := os.Stat("level.llv"); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("Exporintg level does not succeded! - level structure was not exported\n")
				fmt.Print(err.Error())
			}

			// level structure was not created - test should fail
			if _, err := os.Stat("level.png"); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("Exporintg level does not succeded! - tileset image was not exported\n")
				fmt.Print(err.Error())
			}
		})
	}
}
