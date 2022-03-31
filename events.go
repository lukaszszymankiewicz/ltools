package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
)

// checking for any mouse event
func (g *Game) handleMouseEvents(screen *ebiten.Image) {

	for rect, f := range g.Pallete.hoverableAreas {
		if coordsInRect(g.mouse_x, g.mouse_y, rect) {
			f()
		}
	}
	for rect, f := range g.Canvas.hoverableAreas {
		if coordsInRect(g.mouse_x, g.mouse_y, rect) {
			f()
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for rect, f := range g.Pallete.ClickableAreas {
			if coordsInRect(g.mouse_x, g.mouse_y, rect) {
				f()
			}
		}
		for rect, f := range g.Canvas.ClickableAreas {
			if coordsInRect(g.mouse_x, g.mouse_y, rect) {
				f()
			}
		}
	}

    // post click events
    if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
        if g.IsRecording() {
            g.SaveRecord()
        }
        g.StopRecording()
    }
}

func (g *Game) handleKeyboardEvents() {
    if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
        recordToUndone := g.UndoOneRecord()
        g.UndrawOneRecord(recordToUndone)
    }
}
