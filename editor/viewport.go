package editor

import ()

type Viewport struct {
	Rows         int
	Cols         int
	ViewportRows int
	ViewportCols int
	X            int
	Y            int
	GridSize     int
}

func NewViewport(rows int, cols int, viewportCols int, viewportRows int) Viewport {
	var v Viewport

	v.Rows = rows
	v.Cols = cols
	v.ViewportRows = viewportCols
	v.ViewportCols = viewportRows
	v.GridSize = GridSize

	return v
}

func (v *Viewport) Width() int {
	return v.GridSize * v.ViewportCols
}

func (v *Viewport) Height() int {
	return v.GridSize * v.ViewportRows
}

func (v *Viewport) Move(dx int, dy int) {
	v.X = MinVal(MaxVal(0, v.X+dx), (v.Cols - v.ViewportCols))
	v.Y = MinVal(MaxVal(0, v.Y+dy), (v.Rows - v.ViewportRows))
}

func (v *Viewport) GetTileIdx(row int, col int) int {
	return (col+v.Y)*v.Cols + (row + v.X)
}

func (v *Viewport) Size() (int, int) {
	return v.Rows, v.Cols
}
