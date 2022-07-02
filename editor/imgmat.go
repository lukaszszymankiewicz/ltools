package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImgMatrix struct {
    FilledRectElement  
	imgs  []*ImageElement
    bg    ImageElement
    bind  *Viewport
    content  *TilesCollection
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

	i.FilledRectElement = NewFilledRectElement(
        x,
        y,
        (int)(i.bind.ViewportRows * GridSize),
        (int)(i.bind.ViewportCols * GridSize),
        midGreyColor,
    )
    i.bg = NewImageElement(0, 0, LoadImage("editor/assets/bg.png"))
	i.imgs = make([]*ImageElement, i.bind.ViewportRows * i.bind.ViewportCols)

	i.Clear()

	return &i
}

func (i *ImgMatrix) Update (g *Game) {
    tiles := i.content.tiles[g.mode]

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

func (i *ImgMatrix) Draw (screen *ebiten.Image) {
    rows := i.bind.ViewportRows
    cols := i.bind.ViewportCols

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
            img := i.imgs[row*cols+col]

			pos_x := i.rect.Min.X + (col * i.bind.GridSize)
			pos_y := i.rect.Min.Y + (row * i.bind.GridSize)

            img.rect.Min.X = pos_x
            img.rect.Min.Y = pos_y

            img.Draw(screen)
		}
	}
}

func (i *ImgMatrix) Name() string {
    return i.name 
}
