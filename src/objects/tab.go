package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Tab struct {
	Rect              image.Rectangle
	locked_image     *ebiten.Image
	unlocked_image   *ebiten.Image
}

type Tabber struct {
    tabs []Tab
    active int
}

func NewTab(x int, y int, locked_path string, unlocked_path image) Tab {
	var t Tab

	t.locked_image = loadImage(locked_path)
	t.unlocked_image = loadImage(unlocked_path)

	width, height := t.unlocked_image.Size()
    t.rect = image.Rect(x, y, x+width, y+height)

	return t
}

func (t *Tabber) AppendTab (x int, y int, locked_path string, unlocked_path string) {
    tab := NewTab(x, y, locked_path, unlocked_path)
    r.x_coords = append(r.x_coords, x)
}

func NewTabber() Tabber {
	var tb Tabber
	return tb
}

// prepare Complete Tabber
func NewCompleteTabber(x int, y int) Tabber {
    tb := NewTabber

    tb.AppendTab(x, y, "assets/tile_tab_locked.png", "assets/tile_tab_locked.png")
    width, _ := tb.tabs[0].locked_image.Size()

    tb.AppendTab(x+width, y, "assets/light_tab_locked.png", "assets/light_tab_locked.png")

    return tb
}

// draws current Tile on screen
func (tb *Tabber) DrawTabber(screen *ebiten.Image) {
    for i:=0; i<len(tb.tabs); i++ {

        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(tb[i].rect.Min.X, tb[i].rect.Min.Y)

        if i == tb.active {
            screen.DrawImage(tb[i].unlocked_image, op)
        } else {
            screen.DrawImage(tb[i].locked_image, op)
        }
    }
}

func (tb *Tabber) ChangeTabToTile(screen *ebiten.Image) {
    tb.active == 0
}

func (tb *Tabber) ChangeTabToLight(screen *ebiten.Image) {
    tb.active == 1
}

func (tb *Tabber) ChangeTabToSpecial(screen *ebiten.Image) {
    tb.active == 2
}
