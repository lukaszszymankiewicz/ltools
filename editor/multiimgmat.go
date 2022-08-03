package editor

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

// comosing multiple image matrix into one, using diffrent alpha value for each layer
type MultiImgMatrix struct {
	name        string
	Area        FilledRectElement
	layers      []*ImgMatrix
	content     *TilesCollection
	bind        *Viewport
	alpha_rules [][]float64
}

func NewMultiImgMatrix(
	name string,
	x int,
	y int,
	viewport *Viewport,
	content *TilesCollection,
	alpha_rules [][]float64,
) *MultiImgMatrix {

	var mi MultiImgMatrix

	mi.name = name
	mi.alpha_rules = alpha_rules
	mi.content = content
	mi.bind = viewport

	mi.Area = NewFilledRectElement(
		"Background",
		x,
		y,
		(int)(viewport.ViewportCols*GridSize),
		(int)(viewport.ViewportRows*GridSize),
		midGreyColor,
	)
	mi.layers = make([]*ImgMatrix, content.n_layers)

	for i := 0; i < content.n_layers; i++ {
		mi.layers[i] = NewImgMatrix(
			name+"Layer"+fmt.Sprint(i),
			x,
			y,
			viewport,
			content,
		)
	}

	return &mi
}

func (mi *MultiImgMatrix) Update(g *Game) {
	top_layer := g.Mode.Active()

	rows := mi.bind.ViewportRows
	cols := mi.bind.ViewportCols

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// tile on such position should be drawn on (row, col) of Image Matrix
			n := ((mi.bind.Y + row) * mi.bind.Cols) + (mi.bind.X + col)

			if mi.content.Empty(n) {
				for layer := 0; layer < mi.content.n_layers; layer++ {
					mi.layers[layer].imgs[row*cols+col] = nil
				}
			} else {
				for layer := 0; layer < mi.content.n_layers; layer++ {
					if tile := mi.content.tiles[layer][n]; tile != nil {
						mi.layers[layer].imgs[row*cols+col] = &tile.ImageElement
						mi.layers[layer].imgs[row*cols+col].alpha = mi.alpha_rules[top_layer][layer]
					} else {
						mi.layers[layer].imgs[row*cols+col] = nil
					}
				}
			}

		}
	}
}

func (mi *MultiImgMatrix) Draw(screen *ebiten.Image) {
	rows := mi.bind.ViewportRows
	cols := mi.bind.ViewportCols

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			pos_x := mi.Area.Rect.Min.X + (col * mi.bind.GridSize)
			pos_y := mi.Area.Rect.Min.Y + (row * mi.bind.GridSize)

			empty := true

			for layer := 0; layer < mi.content.n_layers; layer++ {
				if mi.layers[layer].imgs[row*cols+col] != nil {
					empty = false
				}
			}

			// if tile on every tile is empty - draw bg image
			if empty {
				img := &mi.layers[0].bg

				img.Rect.Min.X = pos_x
				img.Rect.Min.Y = pos_y

				img.Draw(screen)
			} else {
				// if not - compose final img blending different image with their alpha values
				for layer := 0; layer < mi.content.n_layers; layer++ {
					img := mi.layers[layer].imgs[row*cols+col]

					if img != nil {

						img.Rect.Min.X = pos_x
						img.Rect.Min.Y = pos_y

						img.Draw(screen)
					}
				}
			}
		}
	}
}

func (mi *MultiImgMatrix) Name() string {
	return mi.name
}
