package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
    "image/color"
	_ "image/png"
)

const (
	NO_OFFSET = 0
    TAB_WIDTH = 200
)

type Tab struct {
	TextElement
    left_line FilledRectElement 
    up_line FilledRectElement 
    right_line FilledRectElement 
    override_color *color.RGBA
}

// Tabber is a collection of Tabs
type Tabber struct {
	tabs   []Tab
	active int
	x      int
	y      int
	width  int
}

func NewTab(x int, y int, content string, override_color *color.RGBA) Tab {
	var t Tab

    // Tab when not active has white background
	t.TextElement = NewTextElement(content, x, y, 100, 30, greyColor)

    width := t.FilledRectElement.rect.Max.X - t.FilledRectElement.rect.Min.X
    height := t.FilledRectElement.rect.Max.Y - t.FilledRectElement.rect.Min.Y

    t.left_line = NewFilledRectElement(x, y, 1, height, blackColor)
    t.up_line = NewFilledRectElement(x, y, width, 1, blackColor)
    t.right_line = NewFilledRectElement(x+width, y, 1, height, blackColor)

    t.override_color = nil
    t.Deactivate()

	return t
}

func (t *Tab) Activate() {
    if t.override_color == nil {
        t.TextElement.FilledRectElement.color = whiteColor
    }
}

func (t *Tab) Deactivate() {
    if t.override_color == nil {
        t.TextElement.FilledRectElement.color = darkGreyColor 
    }
}

func (t *Tab) Draw(screen *ebiten.Image) {
    t.TextElement.Draw(screen)
    t.left_line.Draw(screen)
    t.up_line.Draw(screen)
    t.right_line.Draw(screen)
}

func (tb *Tabber) Area(i int) image.Rectangle {
	return tb.tabs[i].rect
}

func NewTabber(x int, y int) Tabber {
	var tb Tabber

	tb.x = x
	tb.y = y
    tb.active = 0

	return tb
}

// appends new Tab to Tabber
func (tb *Tabber) AppendTab(x int, y int, content string) {
	tab := NewTab(x, y, content, nil)
	tb.tabs = append(tb.tabs, tab)
}

// creates new Tab and appends it to Tabber
func (tb *Tabber) AddNewTabToTabber(content string, offset int) {
	tb.AppendTab(tb.x+tb.width+offset, tb.y, content)
	// adding last tab width to tabber overall width
	tb.width += tb.tabs[len(tb.tabs)-1].rect.Max.X - tb.tabs[len(tb.tabs)-1].rect.Min.X
}

// prepare Complete Tabber
func NewCompleteTabber(x int, y int, names []string) Tabber {
	tb := NewTabber(x, y)

    for _, name := range names {
        tb.AddNewTabToTabber(name, NO_OFFSET)
    }

	tb.ChangeCurrent(0)

	return tb
}

func (tb *Tabber) Draw(screen *ebiten.Image) {
	for i := 0; i < len(tb.tabs); i++ {
		tb.tabs[i].Draw(screen)
	}
}

func (tb *Tabber) ChangeCurrent(tab int) {
	tb.tabs[tb.active].Deactivate()
	tb.active = tab
	tb.tabs[tb.active].Activate()
}
