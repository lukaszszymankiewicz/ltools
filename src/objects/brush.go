package objects

import (
)

type Brush struct {
	brush_func func(int, int, int, int, *Tile)BrushResult
}

type BrushResult struct {
    n int
    coords [][]int
    layers []int
    tiles []*Tile
}

func (br *BrushResult) Len() int {
    return br.n
}

func (br *BrushResult) GetCoord(i int) (int, int)  {
    return br.coords[i][0], br.coords[i][1]
}

func (br *BrushResult) GetLayer(i int) (int)  {
    return br.layers[i]
}

func (br *BrushResult) GetTile(i int) (*Tile)  {
    return br.tiles[i]
}

func PenBrushFunc(x1 int, y1 int, x2 int, y2 int, tile *Tile) BrushResult {
    return BrushResult{
        1,
        [][]int{{x1, y1}},
        []int{tile.layer},
        []*Tile{tile},
    }
}

func RubberBrushFunc(x1 int, y1 int, x2 int, y2 int, tile *Tile) BrushResult {
    return BrushResult{
        1,
        [][]int{{x1, y1}},
        []int{tile.layer},
        []*Tile{nil},
    }
}

func (b *Brush) ApplyBrush(x1 int, y1 int, x2 int, y2 int, tile *Tile) BrushResult {
    return b.brush_func(x1, y1, x2, y2, tile)
}

var PenBrush Brush = Brush{PenBrushFunc}
var RubberBrush Brush = Brush{RubberBrushFunc}
