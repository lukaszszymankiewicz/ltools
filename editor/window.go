package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
    "image"
)

const (
    LONG_CLICK_FUNC = iota
    SINGLE_CLICK_FUNC 
    HOVER_FUNC
)

// ClickableAreas       map[image.Rectangle]func(*ebiten.Image)
// SingleClickableAreas map[image.Rectangle]func(*ebiten.Image)
// HoverableAreas       map[image.Rectangle]func(*ebiten.Image)


type WindowElement interface {
    Name()string
    Update(*Game)
    Draw(*ebiten.Image)
}

type Window struct {
    elements []WindowElement
    funcs    []map[image.Rectangle]func()
}

func (w *Window) Get(name string) WindowElement {
    for _, element := range w.elements {
        if element.Name() == name {
            return element
        }
    }

    return nil
}

func (w *Window) Bind(type_ int, rect image.Rectangle, f func()) {
    if w.funcs == nil {
        w.funcs = make([]map[image.Rectangle]func(), LAYER_N)
    } else {
        w.funcs[type_][rect] = f
    }
}

func (w *Window) Update(g *Game) {
    for _, element := range w.elements {
        element.Update(g)
    }
}

func (w *Window) Draw(screen *ebiten.Image) {
    for _, element := range w.elements {
        element.Draw(screen)
    }
}

type WindowManager struct {
    windows []*Window
}

func NewWindowManager() WindowManager {
    var wm WindowManager
    wm.windows = make([]*Window, 0)
    return wm
}

func (wm *WindowManager) Add(w *Window) {
    wm.windows = append(wm.windows, w)
}

func (wm *WindowManager) Update (g *Game) {
    for _, w := range wm.windows {
        w.Update(g)
    }
}

func (wm *WindowManager) Draw (screen *ebiten.Image) {
    for _, w := range wm.windows {
        w.Draw(screen)
    }
}
