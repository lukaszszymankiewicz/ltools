package editor

import (
	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
)

// handles all mouse events
// func (g *Game) handleMouseEvents(screen *ebiten.Image) {
// 
// 	// mouse cursor hovering event
// 	cursorDrawn := false
// 	for rect, f := range g.HoverableAreas {
// 		if coordsInRect(g.mouse_x, g.mouse_y, rect) {
// 			f(screen)
// 			cursorDrawn = true
// 		}
// 	}
// 
// 	if cursorDrawn == false {
// 		g.DrawCursorElsewhere(screen)
// 	}
// 
// 	// mouse button single click button events
// 	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
// 		for rect, f := range g.SingleClickableAreas {
// 			if coordsInRect(g.mouse_x, g.mouse_y, rect) {
// 				f(screen)
// 			}
// 		}
// 	}
// 
// 	// mouse button still click button events
// 	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
// 		for rect, f := range g.ClickableAreas {
// 			if coordsInRect(g.mouse_x, g.mouse_y, rect) {
// 				f(screen)
// 			}
// 		}
// 	}
// 
// 	// mouse button release button events
// 	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
// 		if g.IsRecording() {
// 			g.SaveRecord()
// 		}
// 		g.StopRecording()
// 	}
// }
// 
// // keyboard events
// func (g *Game) handleKeyboardEvents() {
// 	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
// 		g.UndrawOneRecord()
// 	}
// }
