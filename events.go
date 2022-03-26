package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
)

// checking for any mouse event
func (g *Game) handleMouseEvents(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()

	if coordsInRect(x, y, g.Canvas.Rect) {
		g.drawHoveredTileOnCanvas(screen, x, y)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.drawTileOnCanvas(screen, x, y)
		}
	}

	if coordsInRect(x, y, g.Pallete.Rect) {
		g.DrawCursorOnPallete(screen, x, y)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.chooseTileFromPallete(x, y)
		}
	}

	// scrolls
	if coordsInRect(x, y, g.ScrollArrowRight.Rect) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.MoveCanvas(1, 0)
		}
	}

	if coordsInRect(x, y, g.ScrollArrowLeft.Rect) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.MoveCanvas(-1, 0)
		}
	}

	if coordsInRect(x, y, g.ScrollArrowUp.Rect) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.MoveCanvas(0, -1)
		}
	}

	if coordsInRect(x, y, g.ScrollArrowDown.Rect) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.MoveCanvas(0, 1)
		}
	}

	if coordsInRect(x, y, g.PalleteScrollArrowUp.Rect) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.MovePallete(0, -1)
		}
	}

	if coordsInRect(x, y, g.PalleteScrollArrowDown.Rect) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.MovePallete(0, 1)
		}
	}

}
