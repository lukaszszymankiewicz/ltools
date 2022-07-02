package editor

import (
)

func (g *Game) NewMainWindow() *Window {
    var w Window

    w.elements = []WindowElement{
        NewImgMatrix("Pallete", PalleteX, PalleteY, &g.ViewportPallete, &g.AvailableTiles),
        NewScroller("PalleteXScroller", PalleteX, PalleteY, SCROLLER_POS_RIGHT, SCROLLER_TYPE_Y, &g.ViewportPallete),
        NewImgMatrix("Canvas", CanvasX, CanvasY, &g.ViewportCanvas, &g.LevelStructure),
        NewScroller("CanvasYScroller", CanvasX, CanvasY, SCROLLER_POS_RIGHT, SCROLLER_TYPE_Y, &g.ViewportCanvas),
        NewScroller("CanvasXScroller", CanvasX, CanvasY, SCROLLER_POS_DOWN, SCROLLER_TYPE_X, &g.ViewportCanvas),
    }
    w.Bind(SINGLE_CLICK_FUNC, w.Get("CanvasYScroller").arrowLow.rect, g.moveCanvasDown)

    return &w
}
