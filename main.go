package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
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
	palleteCols  = 6
	palleteRows  = 10
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

func cursorInRect(cursor_x, cursor_y, rect_x1, rect_y1, rect_x2, rect_y2 int) bool {
	if (cursor_x >= rect_x1 && cursor_x <= rect_x2) && (cursor_y >= rect_y1 && cursor_y <= rect_y2) {
		return true
	}
	return false
}

func drawCursorOnPallete(screen *ebiten.Image, x int, y int) {
	ebitenutil.DrawRect(
		screen,
		float64((int(((x-palleteX)/tileWidth)*tileWidth) + palleteX)),
		float64((int(((y-palleteY)/tileHeight)*tileHeight) + palleteY)),
		tileWidth,
		tileHeight,
		color.White,
	)
}

func (g *Game) Draw(screen *ebiten.Image) {
	var currentTile int

	tilesetWidth, tilesetHeight := img.Size()

	var tilesheetRows int = tilesetWidth / tileWidth
	var tilesheetCols int = tilesetHeight / tileWidth
	var tilesNum int = tilesheetCols * tilesheetRows

	for currentTile < tilesNum {
		tileRect := image.Rect(
			(currentTile%tilesheetCols)*tileWidth,
			(currentTile/tilesheetCols)*tileHeight,
			((currentTile%tilesheetCols)*tileWidth)+tileWidth,
			((currentTile/tilesheetCols)*tileWidth)+tileHeight,
		)

		subImg := img.SubImage(tileRect).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(((currentTile%palleteCols)*tileWidth)+palleteX),
			float64(((currentTile/palleteCols)*tileHeight)+palleteY),
		)
		screen.DrawImage(subImg, op)

		currentTile++
		if currentTile >= palleteCols*palleteRows {
			break
		}
	}

	// mouse
	x, y := ebiten.CursorPosition()
	msg := fmt.Sprintf("x=%d, y=%d \n", x, y)
	ebitenutil.DebugPrint(screen, msg)

	if cursorInRect(
		x,
		y,
		palleteX,
		palleteY,
		(palleteX + (tileWidth * palleteCols)),
		(palleteY + (tileHeight * palleteRows)),
	) {
		drawCursorOnPallete(screen, x, y)
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
