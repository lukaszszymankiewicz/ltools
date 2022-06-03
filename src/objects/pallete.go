package objects

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

// Pallete struct serves as a wrapper for Tiles which can be drawn.
type Pallete struct {
	Grid
    bg ImageElement
}

func NewPallete(
	x int,
	y int,
	width int,
	height int,
	grid_size int,
	rows int,
	cols int,
	n_layers int,
) Pallete {
	var p Pallete

	p.Grid = NewGrid(x, y, width, height, grid_size, rows, cols, n_layers, whiteColor)
    p.bg = NewImageElement(0, 0, LoadImage("src/objects/assets/other/bg.png"))

	return p
}

func (p *Pallete) Draw(screen *ebiten.Image) {
	n := 0

	for row := 0; row < p.viewportRows; row++ {
		for col := 0; col < p.viewportCols; col++ {
			tile := p.GetTileOnDrawingArea(row+p.viewport_y, col, p.current_layer)

			pos_x := p.rect.Min.X + (col * p.grid_size)
			pos_y := p.rect.Min.Y + (row * p.grid_size)

			if tile != nil {
                tile.ImageElement.rect.Min.X = pos_x
                tile.ImageElement.rect.Min.Y = pos_y
				tile.ImageElement.Draw(screen)
            } else {
                p.bg.rect.Min.X = pos_x
                p.bg.rect.Min.Y = pos_y
                p.bg.Draw(screen)
            }
			n++
		}
	}

	// drawing the scroller
	p.scroller_y.Draw(screen)
}

func (p *Pallete) FillPallete(tileset Tileset, layer int) {
	width, height := tileset.img.Size()
	rows := int(width / p.grid_size)
	cols := int(height / p.grid_size)

	n := 0

	for y := 0; y < cols; y++ {
		for x := 0; x < rows; x++ {
			r := image.Rect(x*p.grid_size, y*p.grid_size, x*p.grid_size+p.grid_size, y*p.grid_size+p.grid_size)

			// despite ebiten.Image.SubImage returning another ebiten.Image, type assertion is still
			// needed.
			i := tileset.img.SubImage(r)
			i2, err := i.(*ebiten.Image)

			if err != false {
				fmt.Errorf("something strange happen while dividing Tileset")
			}

			t := NewTile(i2, tileset.name, false, x, y, layer)

			p.SetTileOnDrawingArea(int(n/p.cols), int(n%p.cols), layer, &t)

			n++
		}
	}

	// just to be sure
	p.UpdateXScroller()
	p.UpdateYScroller()

}

func (p *Pallete) MovePalleteDown(screen *ebiten.Image) {
	p.MoveViewport(0, 1)
}

func (p *Pallete) MovePalleteUp(screen *ebiten.Image) {
	p.MoveViewport(0, -1)
}

func (p *Pallete) RestartPalletePos() {
    p.viewport_x = 0
    p.viewport_y = 0
    p.UpdateXScroller()
	p.UpdateYScroller()
}

func (p *Pallete) PosHasTile(x int, y int, l int) bool {
	if p.GetTileOnDrawingArea(x, y, l) == nil {
        return false 
    } else {
        return true
    }
}
