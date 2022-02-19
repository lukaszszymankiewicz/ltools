package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
	palleteX     = 20
	palleteY     = 20
	tilesPerRow  = 6
	tileWidth    = 32
	tileHeight   = 32
)

type Game struct{}

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("tileset_1.png")
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

	var currentTile int
	var palleteCols int = 6 

    tilesetWidth, tilesetHeight := img.Size()
	var tilesheetRows  int = tilesetWidth / tileWidth
	var tilesheetCols int = tilesetHeight  / tileWidth
	var tilesNum int = tilesheetCols * tilesheetRows

	for currentTile < tilesNum {
		tileRect := image.Rect(
            (currentTile % tilesheetCols) * tileWidth,
            (currentTile / tilesheetCols) * tileHeight,
            ((currentTile % tilesheetCols) * tileWidth) + tileWidth,
            ((currentTile / tilesheetCols) * tileWidth) + tileHeight,
        ) 

		subImg := img.SubImage(tileRect).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(
             float64((currentTile % palleteCols) * tileWidth),
             float64((currentTile / palleteCols) * tileHeight),
        )
		screen.DrawImage(subImg, op)

		currentTile++
	}

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
