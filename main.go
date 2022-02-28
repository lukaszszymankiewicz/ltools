package main

import (
	_ "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"log"
)

const (
    screenWidth               = 640
	screenHeight              = 480
	screenTilePalleteX        = 20
	screenTilePalleteY        = 20
    tileWidth                 = 32
    tileHeight                = 32
	nColsInScreenTilePallete  = 6
	nRowsInScreenTilePallete  = 10
    screenTilePalleteEndX     = screenTilePalleteX+(tileWidth*nColsInScreenTilePallete)
    screenTilePalleteEndY     = screenTilePalleteY+(tileHeight*nRowsInScreenTilePallete)
    currentTiletoDrawX        = screenTilePalleteEndX / 2
    currentTiletoDrawY        = screenTilePalleteEndY + tileHeight
    currentTiletoDrawEndX     = currentTiletoDrawX + tileWidth
    currentTiletoDrawEndY     = currentTiletoDrawY + tileHeight
    maxTileinTilesetPallete   = nColsInScreenTilePallete*nRowsInScreenTilePallete

)

type Tileset struct {
    image *ebiten.Image
    width int
    height int
}

type Game struct {
	// static desktop coords
	tilePalleteCoords image.Rectangle
	currentTileToDraw image.Rectangle

	// dynamic thingis
	currentTileToDrawRect image.Rectangle

    // images
    tilesetPallete Tileset 
}

func newTileset(path string) Tileset {
    var tileset Tileset

    tileset.image = loadImage(path)
    width, height := tileset.image.Size()
    tileset.width = width
    tileset.height = height

    return tileset
}

func loadImage(path string) *ebiten.Image {
	var err error
    img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
    return img
}

func init() {
	ebiten.SetFullscreen(true)
}

func (g *Game) Update() error {
	return nil
}

func coordsInRect(x int, y int, rect image.Rectangle) bool {
	if (x >= rect.Min.X && x <= rect.Max.X) && (y >= rect.Min.Y && y <= rect.Max.Y) {
		return true
	}
	return false
}


func drawCursorOnPallete(screen *ebiten.Image, x int, y int, size float64) {
	leftCornerX := float64((((x-screenTilePalleteX)/tileWidth)*tileWidth) + screenTilePalleteX)
	leftCornerY := float64((((y-screenTilePalleteY)/tileHeight)*tileHeight) + screenTilePalleteY)

	ebitenutil.DrawLine(screen, leftCornerX, leftCornerY, leftCornerX+size, leftCornerY, color.White)
	ebitenutil.DrawLine(screen, leftCornerX, leftCornerY, leftCornerX, leftCornerY+size+1, color.White)
	ebitenutil.DrawLine(screen, leftCornerX, leftCornerY+size, leftCornerX+size, leftCornerY+size, color.White)
	ebitenutil.DrawLine(screen, leftCornerX+size, leftCornerY, leftCornerX+size, leftCornerY+size, color.White)
}

func (g *Game) HandleMouse(screen *ebiten.Image) {
	tilesetHeight := g.tilesetPallete.height

	var tilesheetCols int = tilesetHeight / tileWidth

    // check for changing the current tile
    x, y := ebiten.CursorPosition()
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && coordsInRect(x, y, g.tilePalleteCoords) {
        currentTile := (((y-screenTilePalleteY)/tileWidth) * nColsInScreenTilePallete) + ((x-screenTilePalleteX)/tileHeight)

		g.currentTileToDrawRect = image.Rect(
			(currentTile%tilesheetCols)*tileWidth,
			(currentTile/tilesheetCols)*tileHeight,
			((currentTile%tilesheetCols)*tileWidth)+tileWidth,
			((currentTile/tilesheetCols)*tileWidth)+tileHeight,
		)
    }

	if coordsInRect(x, y, g.tilePalleteCoords) {
		drawCursorOnPallete(screen, x, y, 32.0)
	}
}


func (g *Game) DrawTilesetPallete(screen *ebiten.Image) {
    tilesetWidth := g.tilesetPallete.width
    tilesetHeight := g.tilesetPallete.height

	var currentTile int
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

		subImg := g.tilesetPallete.image.SubImage(tileRect).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(((currentTile%nColsInScreenTilePallete)*tileWidth)+screenTilePalleteX),
			float64(((currentTile/nColsInScreenTilePallete)*tileHeight)+screenTilePalleteY),
		)
		screen.DrawImage(subImg, op)
		currentTile++

		if currentTile >= maxTileinTilesetPallete {
			break
		}
	}
}

func (g *Game) DrawCurrentTile(screen *ebiten.Image) {
    subImg := g.tilesetPallete.image.SubImage(g.currentTileToDrawRect).(*ebiten.Image)

    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(
        float64(g.currentTileToDraw.Min.X),
        float64(g.currentTileToDraw.Min.Y),
    )
    screen.DrawImage(subImg, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
    g.DrawTilesetPallete(screen)
    g.HandleMouse(screen)
    g.DrawCurrentTile(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

// placeholder gonna hold place
func main() {
	game := &Game{
		tilePalleteCoords: image.Rect(
			screenTilePalleteX,
			screenTilePalleteY,
            screenTilePalleteEndX,
            screenTilePalleteEndY,
        ),
        currentTileToDrawRect: image.Rect(0, 0, tileWidth, tileHeight),
        currentTileToDraw: image.Rect(
            currentTiletoDrawX,
            currentTiletoDrawY,
            currentTiletoDrawEndX,
            currentTiletoDrawEndY,     
        ),
        tilesetPallete: newTileset("tileset_1.png"),
    }

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
