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

type Game struct {
	// static desktop coords
	tilePalleteCoords image.Rectangle
	currentTileToDraw image.Rectangle

	// dynamic thingis
	currentTileToDrawRect image.Rectangle
}

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

func cursorInRect(x int, y int, rect image.Rectangle) bool {
	if (x >= rect.Min.X && x <= rect.Max.X) && (y >= rect.Min.Y && y <= rect.Max.Y) {
		return true
	}
	return false
}

func drawCursor(screen *ebiten.Image, x int, y int, size float64) {
	leftCornerX := float64((int(((x-palleteX)/tileWidth)*tileWidth) + palleteX))
	leftCornerY := float64((int(((y-palleteY)/tileHeight)*tileHeight) + palleteY))

	ebitenutil.DrawLine(screen, leftCornerX, leftCornerY, leftCornerX+size, leftCornerY, color.White)
	ebitenutil.DrawLine(screen, leftCornerX, leftCornerY, leftCornerX, leftCornerY+size+1, color.White)
	ebitenutil.DrawLine(screen, leftCornerX, leftCornerY+size, leftCornerX+size, leftCornerY+size, color.White)
	ebitenutil.DrawLine(screen, leftCornerX+size, leftCornerY, leftCornerX+size, leftCornerY+size, color.White)
}

func drawCursorOnPallete(screen *ebiten.Image, x int, y int) {
	drawCursor(screen, x, y, 32.0)
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


    // check for changing the current tile
    x, y := ebiten.CursorPosition()
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && cursorInRect(x, y, g.tilePalleteCoords) {
        currentTile := (((y-palleteY)/tileWidth) * tilesPerRow) + ((x-palleteX)/tileHeight)

		g.currentTileToDrawRect = image.Rect(
			(currentTile%tilesheetCols)*tileWidth,
			(currentTile/tilesheetCols)*tileHeight,
			((currentTile%tilesheetCols)*tileWidth)+tileWidth,
			((currentTile/tilesheetCols)*tileWidth)+tileHeight,
		)
        ebitenutil.DebugPrint(screen, msg)
    }

    // current tile to draw
    subImg := img.SubImage(g.currentTileToDrawRect).(*ebiten.Image)
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(
        float64(g.currentTileToDraw.Min.X),
        float64(g.currentTileToDraw.Min.Y),
    )
    screen.DrawImage(subImg, op)

	// mouse
	if cursorInRect(x, y, g.tilePalleteCoords) {
		drawCursorOnPallete(screen, x, y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

// placeholder gonna hold place
func main() {
	game := &Game{
		tilePalleteCoords: image.Rect(
			palleteX,
			palleteY,
			palleteX+(tileWidth*palleteCols),
			palleteY+(tileHeight*palleteRows),
        ),
        currentTileToDrawRect: image.Rect(0, 0, 32, 32),
        currentTileToDraw: image.Rect(
			(palleteX+(tileWidth*(palleteCols))) / 2,
			palleteY+(tileHeight*(palleteRows+1)),
			((palleteX+(tileWidth*(palleteCols))) /2) + 32,
			palleteY+(tileHeight*(palleteRows+1)) + 32,
        ),
    }

	ebiten.SetWindowSize(screenWidth, screenHeight)

	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
