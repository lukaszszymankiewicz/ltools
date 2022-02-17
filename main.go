package main

import (
    _ "image/png"
	"log"
    "image"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    screenWidth = 640
    screenHeight = 480
)

type Game struct{}

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("tiles.png")
	if err != nil {
		log.Fatal(err)
	}
    ebiten.SetFullscreen(true)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "LEVEL EDITOR")
    subImg := img.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image)
    op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(320, 240)
    screen.DrawImage(subImg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

// placeholder gonna hold place
func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

