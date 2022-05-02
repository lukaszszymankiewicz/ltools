package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Controller struct {
	mouse_x int
	mouse_y int
}

// updates all controlers global state
func (g *Game) UpdateControllersState() {
	g.mouse_x, g.mouse_y = ebiten.CursorPosition()
}
