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
	Rect           image.Rectangle
	locked_image   *ebiten.Image
	unlocked_image *ebiten.Image
}

type Tabber struct {
	tabs   []Tab
	active int
	x      int
	y      int
}

// creates new Tab struct
func NewTab(x int, y int, locked_path string, unlocked_path string) Tab {
	var t Tab

	t.locked_image = loadImage(locked_path)
	t.unlocked_image = loadImage(unlocked_path)

	width, height := t.unlocked_image.Size()
	t.Rect = image.Rect(x, y, x+width, y+height)

	return t
}

// creates new Tabber struct
func NewTabber(x int, y int) Tabber {
	var tb Tabber
	tb.x = x
	tb.y = y
	return tb
}

// appends new Tab to Tabber
func (tb *Tabber) AppendTab(x int, y int, locked_path string, unlocked_path string) {
	tab := NewTab(x, y, locked_path, unlocked_path)
	tb.tabs = append(tb.tabs, tab)
}

// creates new Tab and appends it to Tabber
func (tb *Tabber) AddNewTabToTabber(locked_path string, unlocked_path string, offset int) {
	tb.AppendTab(tb.x+offset, tb.y, locked_path, unlocked_path)
	width, _ := tb.tabs[0].locked_image.Size()
	tb.x += width
}

// returns single Tab rectangle coords
func (tb *Tabber) AreaRect(i int) image.Rectangle {
	return tb.tabs[i].Rect
}

// prepare Complete Tabber
func NewCompleteTabber(x int, y int) Tabber {
	// TODO: this sould be initialised on ltools.main!!
	tb := NewTabber(x, y)

	tb.AddNewTabToTabber("assets/tile_tab_locked.png", "assets/tile_tab_unlocked.png", NO_OFFSET)
	tb.AddNewTabToTabber("assets/light_tab_locked.png", "assets/light_tab_unlocked.png", NO_OFFSET)
	tb.AddNewTabToTabber("assets/entities_tab_locked.png", "assets/entities_tab_unlocked.png", NO_OFFSET)
	tb.AddNewTabToTabber("assets/export.png", "assets/export.png", 572)

	return tb
}

// draws current Tile on screen
func (tb *Tabber) DrawTabber(screen *ebiten.Image) {
	for i := 0; i < len(tb.tabs); i++ {

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tb.tabs[i].Rect.Min.X), float64(tb.tabs[i].Rect.Min.Y))

		if i == tb.active {
			screen.DrawImage(tb.tabs[i].unlocked_image, op)
		} else {
			screen.DrawImage(tb.tabs[i].locked_image, op)
		}
	}
}

// sets current Tab of Tabber
func (tb *Tabber) ChangeTab(tab int) {
	tb.active = tab
}
