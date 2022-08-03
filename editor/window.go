package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
)

const (
	LONG_CLICK_FUNC = iota
	SINGLE_CLICK_FUNC
	HOVER_FUNC
	ALL_FUNC_TYPE
)

type WindowElement interface {
	Name() string
	Update(*Game)
	Draw(*ebiten.Image)
}

type FuncBind struct {
	area image.Rectangle
	f    func()
}

type Window struct {
	elements     []WindowElement
	funcs        [][]FuncBind
	default_func []func()
}

func (w *Window) GetArea(names []string) image.Rectangle {
	for _, element := range w.elements {
		if element.Name() == names[0] {
			return ReadRect(element, names[1:])
		}
	}

	// finding proper area does not succed - return dummy rect
	return image.Rect(0, 0, 0, 0)
}

func (w *Window) MakeInteractable() {
	w.funcs = make([][]FuncBind, ALL_FUNC_TYPE)

	for i := 0; i < ALL_FUNC_TYPE; i++ {
		w.funcs[i] = make([]FuncBind, 0)
	}
}

func (w *Window) Bind(type_ int, rect image.Rectangle, f func()) {
	if w.funcs == nil {
		w.MakeInteractable()
	}

	w.funcs[type_] = append(w.funcs[type_], FuncBind{rect, f})
}

func (w *Window) BindDefault(type_ int, f func()) {
	if w.default_func == nil {
		w.default_func = make([]func(), 0)
	}
	w.default_func = append(w.default_func, f)
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
	current int
	Cursor
}

func NewWindowManager() WindowManager {
	var wm WindowManager

	wm.windows = make([]*Window, 0)
	wm.Cursor = NewCursor()

	return wm
}

func (wm *WindowManager) SetCurrentToNewest() {
	wm.current = len(wm.windows) - 1
}

func (wm *WindowManager) KillLast() {
	wm.windows = wm.windows[:len(wm.windows)-1]
	wm.current = len(wm.windows) - 1
}

func (wm *WindowManager) Add(w *Window) {
	wm.windows = append(wm.windows, w)
}

func (wm *WindowManager) Update(g *Game) {
	for _, w := range wm.windows {
		w.Update(g)
	}
	wm.Cursor.Update(g)
}

func (wm *WindowManager) Draw(screen *ebiten.Image) {
	for _, w := range wm.windows {
		w.Draw(screen)
	}
	wm.Cursor.Draw(screen)
}

func (wm *WindowManager) handleMouseEvents(x int, y int) {
	hover := false

	// just to be sure
	if wm.current > len(wm.windows)-1 {
		return
	}

	for _, fb := range wm.windows[wm.current].funcs[HOVER_FUNC] {
		if coordsInRect(x, y, fb.area) {
			fb.f()
			hover = true
		}
	}

	if hover == false {
		for _, f := range wm.windows[wm.current].default_func {
			f()
		}
	}

	// mouse button single click button events
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		for _, fb := range wm.windows[wm.current].funcs[SINGLE_CLICK_FUNC] {
			if coordsInRect(x, y, fb.area) {
				fb.f()
			}
		}
	}

	// mouse button release button events
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, fb := range wm.windows[wm.current].funcs[LONG_CLICK_FUNC] {
			if coordsInRect(x, y, fb.area) {
				fb.f()
			}
		}
	}

}

// keyboard events
// func (g *Game) handleKeyboardEvents() {
// 	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
// 		g.UndrawOneRecord()
// 	}
// }
