package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    screenWidth = 320
    screenHeight = 240
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "IT ALL ABOUT HOLDING PLACE HERE")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// placeholder gonna hold place
func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

