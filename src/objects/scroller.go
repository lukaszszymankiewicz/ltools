package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Scroller struct {
	color      color.Color
    bgcolor    color.Color
	MaxRect    image.Rectangle
	Rect       image.Rectangle
    arrowLow   ScrollArrow
    arrowHigh  ScrollArrow
}

type ScrollArrow struct {
	image *ebiten.Image
    rect       image.Rectangle

}

// creates new scroller struct
func NewScroller(
    x int, y int, width int, height int, color color.Color, bgcolor color.Color,
) Scroller {
	var s Scroller

    // we are assuming that this is x-scroller if it is wider that higher
    if width > height {
        s.arrowLow = NewScrollArrow(x, y+height, "assets/arrow_l.png")
        s.arrowHigh = NewScrollArrow(x+width-32, y+height, "assets/arrow_r.png")
        s.MaxRect = image.Rect(x+32, y+height, x+width-32, y+height+32)
    } else {
    // we are assuming that this is y-scroller if it is higher than wider
        s.arrowLow = NewScrollArrow(x+width, y, "assets/arrow_u.png")
        s.arrowHigh = NewScrollArrow(x+width, y+height-32, "assets/arrow_d.png")
        s.MaxRect = image.Rect(x+width, y+32, x+width+32, y+height-32)
    }
 
	s.color = color
	s.bgcolor = bgcolor

	return s
}

// creates new arrow for scroll control
func NewScrollArrow(x int, y int, path string) ScrollArrow {
	var sa ScrollArrow

	sa.image = loadImage(path)
	width, height := sa.image.Size()
	sa.rect = image.Rect(x, y, x+width, y+height)

	return sa
}

// draws ScrollArrow
func (sa *ScrollArrow) DrawScrollArrow(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sa.rect.Min.X), float64(sa.rect.Min.Y))

	screen.DrawImage(sa.image, op)
}

func (s *Scroller) DrawScroller(screen *ebiten.Image) {
    // arrows
	s.arrowLow.DrawScrollArrow(screen)
	s.arrowHigh.DrawScrollArrow(screen)

	// background
	drawer.FilledRect(screen, s.MaxRect, s.bgcolor)

	// actual scroller
	drawer.FilledRect(screen, s.Rect, s.color)
}



