package objects

import (
	"fmt"
	"testing"
)

func TestExportTiles(t *testing.T) {

	var tests = []struct {
		args    [][]int
		n_args  int
		n_tiles int
		want    [][]int
	}{
		{
			[][]int{{9, 9, 0}}, 1, 1, [][]int{{9, 9}},
		},
		{
			[][]int{{2, 2, 1}, {3, 3, 0}, {4, 4, 0}, {5, 5, 0}}, 4, 2, [][]int{{3, 3, 4, 4, 5, 5}, {2, 2}},
		},
		{
			[][]int{{2, 2, 1}, {3, 3, 0}, {4, 4, 0}, {5, 5, 0}, {9, 9, 2}}, 5, 3, [][]int{{3, 3, 4, 4, 5, 5}, {2, 2}, {9, 9}},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.n_tiles)
		t.Run(testname, func(t *testing.T) {

			// canvas with ten rows, ten cols and one layer
			var c Canvas

			c.viewportRows = 10
			c.viewportCols = 10
			c.canvasRows = 10
			c.canvasCols = 10
			c.drawingAreas = make([][]int, 1)
			c.ClearDrawingAreas(1)

			for i := 0; i < tt.n_args; i++ {
				c.SetTileOnCanvas(tt.args[i][0], tt.args[i][1], tt.args[i][2], 0)
			}

			ans := c.ExportTiles(tt.n_tiles, 0)

			if len(ans) != len(tt.want) {
				t.Errorf("returned matrix has (%d) rows, want (%d)", len(ans), len(tt.want))
			}
			for row := 0; row < tt.n_tiles; row++ {
				if len(ans[row]) != len(tt.want[row]) {
					t.Errorf("tile with index (%d) has (%d) occurances, want (%d)", row, len(ans[row]), len(tt.want[row]))
				}

				for elem := 0; elem < len(tt.want[row]); elem++ {
					if ans[row][elem] != tt.want[row][elem] {
						t.Errorf("tile with index (%d) is on %d want  %d", row, ans[row][elem], tt.want[row][elem])
					}
				}
			}
		})

	}
}
