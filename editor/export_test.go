package editor

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestStack(t *testing.T) {
	tileset_name := "name"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, false, 0, 1, 0)

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

	tileset_name := "name"

	g.Canvas = NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := NewTile(i3, tileset_name, false, 0, 2, 1)

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

func TestPrepareTileset(t *testing.T) {
	var g Game

	g.Canvas = NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	tileset_name := "assets/basic_tileset.png"

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := NewTile(i3, tileset_name, false, 0, 2, 1)

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

	t1 := NewTile(i1, "name", false, 0, 0, 0)
	t2 := NewTile(i2, "name", false, 0, 0, 0)
	t3 := NewTile(i3, "name", false, 0, 0, 1)

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

	t1 := NewTile(i1, "name", false, 0, 0, 0)
	t2 := NewTile(i2, "name", false, 0, 0, 0)
	t3 := NewTile(i3, "name", false, 0, 0, 1)

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
	tileset_name := "assets/basic_tileset.png"
	g.Canvas = NewCanvas(0, 0, 30*32, 20*32, 32, 30, 20, 3)

	i1 := ebiten.NewImage(32, 32)
	i2 := ebiten.NewImage(32, 32)
	i3 := ebiten.NewImage(32, 32)

	t1 := NewTile(i1, tileset_name, false, 0, 0, 0)
	t2 := NewTile(i2, tileset_name, false, 0, 1, 0)
	t3 := NewTile(i3, tileset_name, false, 0, 2, 1)

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
			g := NewEditor()
			tileset_name := "assets/basic_tileset.png"

			// this is just a shortcut for selecting tile by index and not by tile
			tiles := make([]*Tile, 0)

			for i := 0; i < tt.tiles_added; i++ {
				img := ebiten.NewImage(32, 32) // every image is a dummy
				t := NewTile(
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

