package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Scroller struct {
	color     color.Color
	bgcolor   color.Color
	MaxRect   image.Rectangle
	Rect      image.Rectangle
	arrowLow  ScrollArrow
	arrowHigh ScrollArrow
}

type ScrollArrow struct {
	image *ebiten.Image
	rect  image.Rectangle
}

// creates new scroller struct
func NewScroller(x int, y int, width int, height int) Scroller {
	var s Scroller

	// we are assuming that this is x-scroller if it is wider that higher
	if width > height {
		s.arrowLow = NewScrollArrow(x, y+height, "assets/arrow_l.png")
		s.arrowHigh = NewScrollArrow(x+width-32, y+height, "assets/arrow_r.png")

		// we assuming that both arrows are the same size
		arrow_width, arrow_height := s.arrowHigh.Size()
		s.MaxRect = image.Rect(x+arrow_width, y+height, x+width-arrow_width, y+height+arrow_height)

	} else {
		// we are assuming that this is y-scroller if it is higher than wider
		s.arrowLow = NewScrollArrow(x, y, "assets/arrow_u.png")
		s.arrowHigh = NewScrollArrow(x, y+height-32, "assets/arrow_d.png")

		// we assuming that both arrows are the same size
		arrow_width, arrow_height := s.arrowHigh.Size()

		s.MaxRect = image.Rect(x, y+arrow_height, x+arrow_width, y+height-arrow_height)
	}

	s.color = scrollerColor
	s.bgcolor = scrollerBGColor

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

// returns Scroll Arrow image size
func (sa *ScrollArrow) Size() (int, int) {
	arrow_width, arrow_height := sa.image.Size()

	return arrow_width, arrow_height
}

// returns Scroll Arrow High Arrow Rect
func (s *Scroller) HighArrowRect() image.Rectangle {
	return s.arrowHigh.rect
}

// returns Scroll Arrow High Arrow Rect
func (s *Scroller) LowArrowRect() image.Rectangle {
	return s.arrowLow.rect
}

// draws whole scroller
func (s *Scroller) Draw(screen *ebiten.Image) {
	// arrows
	s.arrowLow.DrawScrollArrow(screen)
	s.arrowHigh.DrawScrollArrow(screen)

	// background
	drawer.FilledRect(screen, s.MaxRect, s.bgcolor)

	// actual scroller
	drawer.FilledRect(screen, s.Rect, s.color)
}
