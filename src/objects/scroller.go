package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Scroller struct {
	background FilledRectElement
	bar        FilledRectElement
	arrowLow   ImageBasedElement
	arrowHigh  ImageBasedElement
}

// creates new scroller struct
func NewScroller(x int, y int, width int, height int) Scroller {

	var s Scroller

	// we are assuming that this is x-scroller if it is wider that higher
	if width > height {
		arrow_l := LoadImage("src/objects/assets/buttons/arrow_l.png")
		arrow_r := LoadImage("src/objects/assets/buttons/arrow_r.png")

		arrow_width, _ := arrow_r.Size()

		s.arrowLow = NewImageBasedElement(x, y, []*ebiten.Image{arrow_l})
		s.arrowHigh = NewImageBasedElement(x+width-arrow_width, y, []*ebiten.Image{arrow_r})

		s.background = NewFilledRectElement(x+arrow_width, y, width-2*(arrow_width), height, scrollerBGColor)
		s.bar = NewFilledRectElement(x+arrow_width, y, width-2*(arrow_width), height, scrollerColor)

	} else {
		arrow_u := LoadImage("src/objects/assets/buttons/arrow_u.png")
		arrow_d := LoadImage("src/objects/assets/buttons/arrow_d.png")

		_, arrow_height := arrow_u.Size()

		s.arrowLow = NewImageBasedElement(x, y, []*ebiten.Image{arrow_u})
		s.arrowHigh = NewImageBasedElement(x, y+height-arrow_height, []*ebiten.Image{arrow_d})

		s.background = NewFilledRectElement(x, y+arrow_height, width, height-2*arrow_height, scrollerBGColor)
		s.bar = NewFilledRectElement(x, y+arrow_height, width, height-2*arrow_height, scrollerColor)

	}

	return s
}

// draws whole scroller
func (s *Scroller) Draw(screen *ebiten.Image) {
	s.background.Draw(screen)
	s.bar.Draw(screen)
	s.arrowLow.Draw(screen)
	s.arrowHigh.Draw(screen)
}
