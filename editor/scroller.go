package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCROLLER_POS_LEFT = iota
	SCROLLER_POS_RIGHT
	SCROLLER_POS_UP
	SCROLLER_POS_DOWN
)

const (
	SCROLLER_TYPE_X = iota
	SCROLLER_TYPE_Y
)

type Scroller struct {
	name       string
	Background FilledRectElement
	Bar        FilledRectElement
	ArrowLow   ImageElement
	ArrowHigh  ImageElement
	bind       *Viewport
	type_      int
}

func NewScroller(name string, x int, y int, pos int, type_ int, bind *Viewport) *Scroller {
	var s Scroller
	s.name = name
	s.bind = bind
	s.type_ = type_

	if type_ == SCROLLER_TYPE_X {
		width := s.bind.Width()
		height := s.bind.Height()

		arrow_width, arrow_height := resources["arrow_r"].Size()

		var ypos int

		if pos == SCROLLER_POS_DOWN {
			ypos = y + height
		} else if pos == SCROLLER_POS_UP {
			ypos = y - arrow_height
		}

		s.ArrowLow = NewImageElement("ArrowLeft", x, ypos, resources["arrow_l"])
		s.ArrowHigh = NewImageElement("ArrowRight", x+width-arrow_width, ypos, resources["arrow_r"])
		s.Background = NewFilledRectElement(
			"Background",
			x+arrow_width,
			ypos,
			width-2*(arrow_width),
			arrow_height,
			scrollerBGColor,
		)
		s.Bar = NewFilledRectElement(
			"Bar",
			x+arrow_width,
			ypos,
			width-2*(arrow_width),
			arrow_height,
			scrollerColor,
		)

	} else {
		width := s.bind.Width()
		height := s.bind.Height()

		arrow_width, arrow_height := resources["arrow_u"].Size()

		var xpos int

		if pos == SCROLLER_POS_LEFT {
			xpos = x - arrow_width
		} else if pos == SCROLLER_POS_RIGHT {
			xpos = x + width
		}

		s.ArrowHigh = NewImageElement("ArrowUp", xpos, y, resources["arrow_u"])
		s.ArrowLow = NewImageElement("ArrowDown", xpos, y+height-arrow_height, resources["arrow_d"])

		s.Background = NewFilledRectElement(
			"Background",
			xpos,
			y+arrow_height,
			arrow_width,
			height-2*arrow_height,
			scrollerBGColor,
		)
		s.Bar = NewFilledRectElement(
			"Bar",
			xpos,
			y+arrow_height,
			arrow_width,
			height-2*arrow_height,
			scrollerColor,
		)
	}

	return &s
}

func (s *Scroller) Update(g *Game) {
	if s.type_ == SCROLLER_TYPE_X {
		s.UpdateX(g)
	} else {
		s.UpdateY(g)
	}
}

func (s *Scroller) UpdateX(g *Game) {
	viewportCols := s.bind.ViewportCols

	len := float64(s.Background.Rect.Max.X - s.Background.Rect.Min.X)
	start := float64(s.Background.Rect.Min.X)

	if s.bind.X != 0 {
		start += (float64(s.bind.X) / float64(s.bind.Cols)) * len
	}

	end := start + (float64(viewportCols)/float64(s.bind.Cols))*len

	s.Bar.Rect.Min.X = int(start)
	s.Bar.Rect.Max.X = int(end)
}

func (s *Scroller) UpdateY(g *Game) {
	viewportRows := s.bind.ViewportRows

	len := float64(s.Background.Rect.Max.Y - s.Background.Rect.Min.Y)
	start := float64(s.Background.Rect.Min.Y)

	if s.bind.Y != 0 {
		start += (float64(s.bind.Y) / float64(s.bind.Rows)) * len
	}
	end := start + (float64(viewportRows)/float64(s.bind.Rows))*len

	s.Bar.Rect.Min.Y = int(start)
	s.Bar.Rect.Max.Y = int(end)
}

func (s *Scroller) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
	s.Bar.Draw(screen)
	s.ArrowLow.Draw(screen)
	s.ArrowHigh.Draw(screen)

	EmptyBorder(screen, s.Background.Rect, poleGrey)
	EmptyBorder(screen, s.ArrowLow.Rect, poleGrey)
	EmptyBorder(screen, s.ArrowHigh.Rect, poleGrey)

}

func (s *Scroller) Name() string {
	return s.name
}
