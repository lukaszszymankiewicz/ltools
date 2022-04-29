package main

import (
	"fmt"
	"testing"
)

func TestPurgeStack(t *testing.T) {
	var tests = []struct {
		args   [][]int
		n_funs int
		want   int
	}{
		{[][]int{{}}, 0, 0},
		{[][]int{{1, 1, 4, 0}}, 1, 1},
		{[][]int{
			{1, 1, 4, 0},
			{2, 2, 4, 0},
			{3, 3, 4, 0},
		}, 3, 1},
		{[][]int{
			{1, 1, 4, 0},
			{2, 2, 3, 0},
			{3, 3, 2, 0},
		}, 3, 3},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.want)
		t.Run(testname, func(t *testing.T) {

			g := NewGame()
			// two tiles are hardcoded into NewGame function, but as it is place nowehere on Canvas
			// it will not interrupt in any way
			g.AddTileFromPalleteToStack(1, 1, 0, false)
			g.AddTileFromPalleteToStack(2, 2, 0, false)
			g.AddTileFromPalleteToStack(3, 3, 0, false)

			for i := 0; i < tt.n_funs; i++ {
				g.TileStack.UpdateTileUsage(tt.args[i][2], +1)
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
	g := NewGame()
	g.exportLevel()
}
