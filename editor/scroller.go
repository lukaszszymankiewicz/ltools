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
	background FilledRectElement
	bar        FilledRectElement
	arrowLow   ImageElement
	arrowHigh  ImageElement
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

		arrow_l := LoadImage("editor/assets/arrow_l.png")
		arrow_r := LoadImage("editor/assets/arrow_r.png")

		arrow_width, arrow_height := arrow_r.Size()
        
        var ypos int

        if pos == SCROLLER_POS_DOWN {
            ypos = y + height
        } else if pos == SCROLLER_POS_UP {
            ypos = y - arrow_height 
        }

        s.arrowLow = NewImageElement(x, ypos, arrow_l)
        s.arrowHigh = NewImageElement(x+width-arrow_width, ypos, arrow_r)
        s.background = NewFilledRectElement(
            x+arrow_width,
            ypos,
            width-2*(arrow_width),
            arrow_height,
            scrollerBGColor,
        )
        s.bar = NewFilledRectElement(
            x+arrow_width,
            ypos,
            width-2*(arrow_width),
            arrow_height,
            scrollerColor,
        )

	} else {
        width := s.bind.Width()
        height := s.bind.Height()

		arrow_u := LoadImage("editor/assets/arrow_u.png")
		arrow_d := LoadImage("editor/assets/arrow_d.png")
		arrow_width, arrow_height := arrow_u.Size()

        var xpos int

        if pos == SCROLLER_POS_LEFT {
            xpos = x - arrow_width
        } else if pos == SCROLLER_POS_RIGHT {
            xpos = x + width
        }

		s.arrowLow = NewImageElement(xpos, y, arrow_u)
		s.arrowHigh = NewImageElement(xpos, y+height-arrow_height, arrow_d)
		s.background = NewFilledRectElement(
            xpos,
            y+arrow_height,
            arrow_width,
            height-2*arrow_height,
            scrollerBGColor,
        )
		s.bar = NewFilledRectElement(
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

	len := float64(s.background.rect.Max.X - s.background.rect.Min.X)
	start := float64(s.background.rect.Min.X)

	if s.bind.X != 0 {
		start += (float64(s.bind.X) / float64(s.bind.Cols)) * len
	}

	end := start + (float64(viewportCols)/float64(s.bind.Cols))*len

	s.bar.rect.Min.X = int(start)
	s.bar.rect.Max.X = int(end)
}

func (s *Scroller) UpdateY (g *Game) {
    viewportRows := s.bind.ViewportRows

	len := float64(s.background.rect.Max.Y - s.background.rect.Min.Y)
	start := float64(s.background.rect.Min.Y)

	if s.bind.Y != 0 {
		start += (float64(s.bind.Y) / float64(s.bind.Rows)) * len
	}
	end := start + (float64(viewportRows)/float64(s.bind.Rows))*len

	s.bar.rect.Min.Y = int(start)
	s.bar.rect.Max.Y = int(end)
}

func (s *Scroller) Draw(screen *ebiten.Image) {
	s.background.Draw(screen)
	s.bar.Draw(screen)
	s.arrowLow.Draw(screen)
	s.arrowHigh.Draw(screen)
}

func (s *Scroller) Name() string {
    return s.name
}
