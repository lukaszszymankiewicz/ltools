package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	POLE_TYPE_X = iota
	POLE_TYPE_Y
)

const (
	WIDTH  = 2
	HEIGTH = 2
)

type Pole struct {
	name  string
	line  FilledRectElement
	type_ int
}

func NewPole(name string, x int, y int, type_ int, length int) *Pole {
	var p Pole

	p.name = name
	p.type_ = type_

	if type_ == POLE_TYPE_X {
		if length == -1 {
			length = ScreenWidth
		}
		p.line = NewFilledRectElement("Middle", x, y, length, HEIGTH, poleGrey)

	} else {
		if length == -1 {
			length = ScreenHeight
		}
		p.line = NewFilledRectElement("Middle", x, y, WIDTH, length, poleGrey)
	}
	return &p
}

func (p *Pole) Update(g *Game) {}

func (p *Pole) Draw(screen *ebiten.Image) {
	p.line.Draw(screen)
}

func (p *Pole) Name() string {
	return p.name
}
