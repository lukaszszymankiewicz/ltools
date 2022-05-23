package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

const (
	NO_OFFSET = 0
)

type Tab struct {
	ImageBasedElement
}

// Tabber is a collection of Tabs
type Tabber struct {
	tabs   []Tab
	active int
	x      int
	y      int
	width  int
}

func NewTab(x int, y int, images []*ebiten.Image) Tab {
	var t Tab

	t.ImageBasedElement = NewImageBasedElement(x, y, images)

	return t
}

func (tb *Tabber) Area(i int) image.Rectangle {
	return tb.tabs[i].rect
}

func NewTabber(x int, y int) Tabber {
	var tb Tabber

	tb.x = x
	tb.y = y

	return tb
}

// appends new Tab to Tabber
func (tb *Tabber) AppendTab(x int, y int, images []*ebiten.Image) {
	tab := NewTab(x, y, images)
	tb.tabs = append(tb.tabs, tab)
}

// creates new Tab and appends it to Tabber
func (tb *Tabber) AddNewTabToTabber(images []*ebiten.Image, offset int) {
	tb.AppendTab(tb.x+tb.width+offset, tb.y, images)
	// adding last tab width to tabber overall width
	tb.width += tb.tabs[len(tb.tabs)-1].rect.Max.X - tb.tabs[len(tb.tabs)-1].rect.Min.X
}

// returns single Tab rectangle coords
func (tb *Tabber) AreaRect(i int) image.Rectangle {
	return tb.tabs[i].rect
}

// prepare Complete Tabber
func NewCompleteTabber(x int, y int) Tabber {
	// TODO: this sould be initialised on ltools.main!!
	tb := NewTabber(x, y)

	t1 := LoadImage("assets/buttons/tile_tab_locked.png")
	t2 := LoadImage("assets/buttons/tile_tab_unlocked.png")
	t3 := LoadImage("assets/buttons/light_tab_locked.png")
	t4 := LoadImage("assets/buttons/light_tab_unlocked.png")
	t5 := LoadImage("assets/buttons/entities_tab_locked.png")
	t6 := LoadImage("assets/buttons/entities_tab_unlocked.png")
	t7 := LoadImage("assets/buttons/export.png")

	tb.AddNewTabToTabber([]*ebiten.Image{t2, t1}, NO_OFFSET)
	tb.AddNewTabToTabber([]*ebiten.Image{t4, t3}, NO_OFFSET)
	tb.AddNewTabToTabber([]*ebiten.Image{t6, t5}, NO_OFFSET)
	tb.AddNewTabToTabber([]*ebiten.Image{t7}, 572)

	tb.ChangeCurrent(0)

	return tb
}

func (tb *Tabber) Draw(screen *ebiten.Image) {
	for i := 0; i < len(tb.tabs); i++ {
		tb.tabs[i].ImageBasedElement.Draw(screen)
	}
}

func (tb *Tabber) ChangeCurrent(tab int) {
	tb.tabs[tb.active].current_image = 0
	tb.active = tab
	tb.tabs[tb.active].current_image = 1
}
