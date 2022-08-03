package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImgMatrix struct {
	name    string
	Area    FilledRectElement
	imgs    []*ImageElement
	bg      ImageElement
	bind    *Viewport
	content *TilesCollection
}

func (i *ImgMatrix) Clear() {
	rows := i.bind.ViewportRows
	cols := i.bind.ViewportCols

	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			i.imgs[x*rows+y] = nil
		}
	}
}

func NewImgMatrix(
	name string,
	x int,
	y int,
	viewport *Viewport,
	content *TilesCollection,
) *ImgMatrix {
	var i ImgMatrix

	i.name = name
	i.bind = viewport
	i.content = content

	i.Area = NewFilledRectElement(
		"Background",
		x,
		y,
		(int)(i.bind.ViewportCols*GridSize),
		(int)(i.bind.ViewportRows*GridSize),
		midGreyColor,
	)
	i.bg = NewImageElement("Empty", 0, 0, resources["bg"])
	i.imgs = make([]*ImageElement, i.bind.ViewportRows*i.bind.ViewportCols)

	i.Clear()

	return &i
}

func (i *ImgMatrix) Update(g *Game) {
	layer := g.Mode.Active()
	tiles := i.content.tiles[layer]

	rows := i.bind.ViewportRows
	cols := i.bind.ViewportCols

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// tile on such position should be drawn on (row, col) of Image Matrix
			n := ((i.bind.Y + row) * i.bind.Cols) + (i.bind.X + col)

			if tiles[n] == nil {
				i.imgs[row*cols+col] = &i.bg
			} else {
				i.imgs[row*cols+col] = &tiles[n].ImageElement
			}
		}
	}
}

func (i *ImgMatrix) Draw(screen *ebiten.Image) {
	rows := i.bind.ViewportRows
	cols := i.bind.ViewportCols

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			img := i.imgs[row*cols+col]

			pos_x := i.Area.Rect.Min.X + (col * i.bind.GridSize)
			pos_y := i.Area.Rect.Min.Y + (row * i.bind.GridSize)

			img.Rect.Min.X = pos_x
			img.Rect.Min.Y = pos_y

			img.Draw(screen)
		}
	}
	EmptyBorder(screen, i.Area.Rect, poleGrey)
}

func (i *ImgMatrix) Name() string {
	return i.name
}
