package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

type DummyElement struct {
	name string
	Area FilledRectElement
}

func (de DummyElement) Name() string              { return de.name }
func (de DummyElement) Update(g *Game)            {}
func (de DummyElement) Draw(screen *ebiten.Image) {}

func TestGetAreaSuccess(t *testing.T) {
	var w Window

	bg := NewFilledRectElement("bg", 100, 200, 200, 200, nil)
	dummy := DummyElement{"sample", bg}

	w.elements = []WindowElement{dummy}
	f := w.GetArea([]string{"sample", "Area"})

	if f.Min.X != 100 || f.Min.Y != 200 || f.Max.X != 300 || f.Max.Y != 400 {
		t.Errorf("Getting Rect Area does not suceed!")
	}
}

func TestGetAreaFail(t *testing.T) {
	var w Window

	bg := NewFilledRectElement("bg", 100, 200, 200, 200, nil)
	dummy := DummyElement{"sample", bg}

	w.elements = []WindowElement{dummy}
	f := w.GetArea([]string{"wrong sample", "wrong Area"})

	if f.Min.X != 0 || f.Min.Y != 0 || f.Max.X != 0 || f.Max.Y != 0 {
		t.Errorf("Getting Rect Area does not suceed!")
	}
}
