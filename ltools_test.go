package main

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io/fs"
	"os"
	"testing"
)

func TestPurgeStack(t *testing.T) {
	var tests = []struct {
		args   [][]int
		n_funs int
		want   int
	}{
		{[][]int{{}}, 0, 0},
		{[][]int{{1, 1, 2, 0}}, 1, 1},
		{[][]int{
			{1, 1, 2, 0},
			{2, 2, 2, 0},
			{3, 3, 2, 0},
		}, 3, 1},
		{[][]int{
			{1, 1, 2, 0},
			{2, 2, 1, 0},
			{3, 3, 0, 0},
		}, 3, 3},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.want)
		t.Run(testname, func(t *testing.T) {

			g := NewGame()

			g.AddTileFromPalleteToStack(1, 1, 0, false)
			g.AddTileFromPalleteToStack(2, 2, 0, false)
			g.AddTileFromPalleteToStack(3, 3, 0, false)

			// such placing tile is bypassing the srawing on screen - looks lame but works
			for i := 0; i < tt.n_funs; i++ {
				g.Canvas.SetTileOnCanvas(tt.args[i][0], tt.args[i][1], tt.args[i][2], tt.args[i][3])
			}

			g.purgeStack()
			ans := len(g.TileStack.GetAllTiles())

			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestExportLevel(t *testing.T) {

	var tests = []struct {
		tiles_adding_funcs [][]int
		tiles_added        int
		tiles_set_funcs    [][]int
		tiles_set          int
	}{
		{
			[][]int{{}}, 0,
			[][]int{{}}, 0,
		},
		{
			[][]int{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}, {3, 3, 1}, {0, 0, 2}}, 5,
			[][]int{{2, 2, 0, 0}, {6, 9, 0, 0}, {3, 3, 1, 0}, {4, 4, 2, 0}, {5, 5, 3, 1}, {6, 6, 4, 2}}, 6,
		},
	}
	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.tiles_set_funcs)

		t.Run(testname, func(t *testing.T) {
			name := "test_file"

			g := NewGame()

			for i := 0; i < tt.tiles_added; i++ {
				g.AddTileFromPalleteToStack(
					tt.tiles_adding_funcs[i][0],
					tt.tiles_adding_funcs[i][1],
					tt.tiles_adding_funcs[i][2],
					false,
				)
			}

			for j := 0; j < tt.tiles_set; j++ {
				g.Canvas.SetTileOnCanvas(
					tt.tiles_set_funcs[j][0],
					tt.tiles_set_funcs[j][1],
					tt.tiles_set_funcs[j][2],
					tt.tiles_set_funcs[j][3],
				)
			}

			g.exportLevel(name)

			// file was not created - test should fail
			if _, err := os.Stat(name + ".llv"); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("Exporintg level does not succeded!")
				fmt.Print(err.Error())
			}

		})
	}
}

func TestPrepareTiles(t *testing.T) {
	// prepare game
	name := "testname"
	g := NewGame()

	// add some tiles!
	g.AddTileFromPalleteToStack(0, 0, 0, false)
	g.AddTileFromPalleteToStack(1, 1, 0, false)
	g.AddTileFromPalleteToStack(2, 2, 0, false)
	g.AddTileFromPalleteToStack(3, 3, 0, false)
	g.AddTileFromPalleteToStack(4, 4, 0, false)
	g.AddTileFromPalleteToStack(5, 5, 0, false)
	g.AddTileFromPalleteToStack(9, 5, 0, false)
	g.AddTileFromPalleteToStack(3, 9, 0, false)

	g.PrepareTiles(name)

	// file was not created - test should fail
	if _, err := os.Stat(name + TILESET_FORMAT); errors.Is(err, fs.ErrNotExist) {
		t.Errorf("Exporintg level does not succeded!")
		fmt.Print(err.Error())
	}

}

func TestExport(t *testing.T) {

	var tests = []struct {
		tiles_adding_funcs [][]int
		tiles_added        int
		tiles_set_funcs    [][]int
		tiles_set          int
	}{
		{
			[][]int{{}}, 0,
			[][]int{{}}, 0,
		},
		{
			[][]int{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}, {3, 3, 1}, {0, 0, 2}}, 5,
			[][]int{{2, 2, 0, 0}, {6, 9, 0, 0}, {3, 3, 1, 0}, {4, 4, 2, 0}, {5, 5, 3, 1}, {6, 6, 4, 2}}, 6,
		},
        {
			[][]int{{3, 3, 1}, {0, 0, 2}, {1,1,0}}, 3,
			[][]int{{3, 3, 2, 0}, {1, 1, 0, 1}, {2, 2, 1,2}}, 3,
		},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.tiles_set_funcs)

		t.Run(testname, func(t *testing.T) {

			g := NewGame()

			for i := 0; i < tt.tiles_added; i++ {
				g.AddTileFromPalleteToStack(
					tt.tiles_adding_funcs[i][0],
					tt.tiles_adding_funcs[i][1],
					tt.tiles_adding_funcs[i][2],
					false,
				)
			}

			for j := 0; j < tt.tiles_set; j++ {
				g.Canvas.SetTileOnCanvas(
					tt.tiles_set_funcs[j][0],
					tt.tiles_set_funcs[j][1],
					tt.tiles_set_funcs[j][2],
					tt.tiles_set_funcs[j][3],
				)
			}

			// dummy image - it is required
			dummy_image := ebiten.NewImage(1, 1)
			g.Export(dummy_image)

			// file was not created - test should fail
			if _, err := os.Stat(BASENAME + ".zip"); errors.Is(err, fs.ErrNotExist) {
				t.Errorf("Exporintg level does not succeded!")
				fmt.Print(err.Error())
			}

		})
	}

}
