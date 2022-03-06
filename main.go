package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"log"
)

const (
	// screen
	ScreenWidth  = 640
	ScreenHeight = 480
	// tile
	TileWidth  = 32
	TileHeight = 32
	// pallete
	PalleteX        = 20
	PalleteY        = 20
	PalleteColsN    = 6
	PalleteRowsN    = 10
	PalleteEndX     = PalleteX + (TileWidth * PalleteColsN)
	PalleteEndY     = PalleteY + (TileHeight * PalleteRowsN)
	PalleteTilesMax = PalleteColsN * PalleteRowsN
	// current tile to draw
	CurrentTileToDrawX    = PalleteEndX / 2
	CurrentTileToDrawY    = PalleteEndY + TileHeight
	CurrentTileToDrawEndX = CurrentTileToDrawX + TileWidth
	CurrentTileToDrawEndY = CurrentTileToDrawY + TileHeight
	// canvas
	CanvasColsN = 10
	CanvasRowsN = 10
	CanvasX     = PalleteEndX + (TileWidth * 2)
	CanvasY     = PalleteY
	CanvasEndX  = CanvasX + (TileWidth * CanvasColsN)
	CanvasEndY  = CanvasY + (TileWidth * CanvasRowsN)
	// cursor
	CursorSize = 32
)

type Tileset struct {
	image     *ebiten.Image
	width     int
	height    int
	rows      int
	cols      int
	num       int
	tilesUsed int
}

type Tile struct {
	tileset      *Tileset
	image        *ebiten.Image
	rowOnPallete int
	colOnPallete int
	numberUsed   int
}

type CurrentTile struct {
	tile       *Tile
	stackIndex int
}

type Game struct {
	palleteCoords image.Rectangle
	canvasCoords  image.Rectangle

	// dynamic thingis
	currentTile           CurrentTile
	currentTileToDrawRect image.Rectangle
	canvas                [CanvasX][CanvasY]int

	// images
	tileset      Tileset
	tileStack    []Tile
	tileStackLen int
}

func newTileset(path string) Tileset {
	var tileset Tileset

	tileset.image = loadImage(path)
	width, height := tileset.image.Size()
	tileset.width = width
	tileset.height = height
	tileset.rows = tileset.width / TileWidth
	tileset.cols = tileset.height / TileHeight
	tileset.num = tileset.rows * tileset.cols

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

func coordsInRect(x int, y int, rect image.Rectangle) bool {
	if (x >= rect.Min.X && x <= rect.Max.X) && (y >= rect.Min.Y && y <= rect.Max.Y) {
		return true
	}
	return false
}

func drawEmptyRect(screen *ebiten.Image, rect image.Rectangle, color color.Color) {
	ebitenutil.DrawLine(screen, float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Max.X), float64(rect.Min.Y), color)
	ebitenutil.DrawLine(screen, float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Min.X), float64(rect.Max.Y+1), color)
	ebitenutil.DrawLine(screen, float64(rect.Max.X), float64(rect.Min.Y), float64(rect.Max.X), float64(rect.Max.Y), color)
	ebitenutil.DrawLine(screen, float64(rect.Min.X), float64(rect.Max.Y), float64(rect.Max.X), float64(rect.Max.Y), color)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return 640, 480
}

func (g *Game) posToTileCoordsOnTileset(x int, y int) (int, int) {
	return int((x - PalleteX) / TileWidth), int((y - PalleteY) / TileHeight)
}

func (g *Game) posToTileCoordsOnCanvas(x int, y int) (int, int) {
	return int((x - CanvasX) / TileWidth), int((y - CanvasY) / TileHeight)
}


func (g *Game) posToTileHoveredOnCanvas(x int, y int) (int, int) {
    tileX, tileY := g.posToTileCoordsOnCanvas(x, y)
    return tileX*TileWidth+CanvasX, tileY*TileHeight+CanvasY
}

func (g *Game) tileNrToCoordsOnPallete(tileNr int) (float64, float64) {
    tileX := float64(((tileNr%PalleteColsN)*TileWidth)+PalleteX)
    tileY := float64(((tileNr/PalleteColsN)*TileHeight)+PalleteY)

    return tileX, tileY
}

func (g *Game) posToCursorCoords(x int, y int) image.Rectangle {
	tileX, tileY := g.posToTileCoordsOnTileset(x, y)

	leftCornerX := (tileX * TileWidth) + PalleteX
	leftCornerY := (tileY * TileHeight) + PalleteY

	return image.Rect(leftCornerX, leftCornerY, leftCornerX+CursorSize, leftCornerY+CursorSize)
}

func (g *Game) tileNrToSubImageOnPallete(tileNr int) *ebiten.Image {
	RectX, RectY := tileNr%g.tileset.cols*TileWidth, tileNr/g.tileset.rows*TileHeight
	tileRect := image.Rect(RectX, RectY, RectX+TileWidth, RectY+TileHeight)

	return g.tileset.image.SubImage(tileRect).(*ebiten.Image)
}

// translating from mouse position to tile on pallete. Please not that Pallete do not need to have
// exacly the same number of rows and columns as tileset (tileset file is basically trimmed to fit
// fixed pallete number of rows and cols). To get column and row of pallete, some basic
// transposition is needed
func (g *Game) posToSubImageOnPallete(x int, y int) *ebiten.Image {
	TileX, TileY := g.posToTileCoordsOnTileset(x, y)
	TileNr := TileX + TileY*PalleteColsN

    return g.tileNrToSubImageOnPallete(TileNr)
}

func (g *Game) clearCanvas() {
	for x := 0; x < CanvasRowsN; x++ {
		for y := 0; y < CanvasColsN; y++ {
			g.canvas[x][y] = -1
		}
	}
}

func (g *Game) drawCursorOnPallete(screen *ebiten.Image, x int, y int) {
	cursorRect := g.posToCursorCoords(x, y)
	drawEmptyRect(screen, cursorRect, color.White)
}

// checks if Tile needs to be deleted from TileStack. It happens if number of tile usage on level
// is equal to zero
func (g *Game) clearStack(index int) {
	if g.tileStack[index].numberUsed == 0 {
		g.tileStack = append(g.tileStack[:index], g.tileStack[index+1:]...)
	}
}

func (g *Game) drawTileOnCanvas(screen *ebiten.Image, x int, y int) {
	tileX, tileY := g.posToTileCoordsOnCanvas(x, y)

	oldTile := g.canvas[tileX][tileY]
	newTile := g.tileStack[g.currentTile.stackIndex]

	if oldTile == -1 {
		newTile.numberUsed++
		g.canvas[tileX][tileY] = g.currentTile.stackIndex
	} else if oldTile != g.currentTile.stackIndex {
		g.tileStack[oldTile].numberUsed = -1
		g.clearStack(oldTile)
		newTile.numberUsed++
		g.canvas[tileX][tileY] = g.currentTile.stackIndex
	}
}

// check if Tile is already used (so it is present on TileStack). If so, its index in stack is
// returned, if not -1 is returned (ANSI C style!)
func (g *Game) checkTileInStack(tileX int, tileY int) (stackIndex int) {
	for i := 0; i < g.tileStackLen; i++ {
		if g.tileStack[i].rowOnPallete == tileX && g.tileStack[i].colOnPallete == tileY {
			return i
		}
	}
	return -1
}

func (g *Game) addTileFromPalleteToStack(x int, y int) {
	tileX, tileY := g.posToTileCoordsOnTileset(x, y)
	newTile := Tile{
		&g.tileset,
		g.posToSubImageOnPallete(x, y),
		tileX,
		tileY,
		0,
	}
	g.tileStack = append(g.tileStack, newTile)
	g.tileStackLen++
}

func (g *Game) setCurrentTile(index int) {
    g.currentTile = CurrentTile {
        &g.tileStack[index],
        index,
    }
}

func (g *Game) chooseTileFromPallete(x int, y int) {
	tileX, tileY := g.posToTileCoordsOnTileset(x, y)

	if tileIndex := g.checkTileInStack(tileX, tileY); tileIndex != -1 {
        g.setCurrentTile(tileIndex)
	} else {
		g.addTileFromPalleteToStack(x, y)
	}
}

func (g *Game) drawHoveredTileOnCanvas(screen *ebiten.Image, x int, y int) {
    TileX, TileY := g.posToTileHoveredOnCanvas(x, y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(TileX), float64(TileY))
	screen.DrawImage(g.currentTile.tile.image, op)
}

func (g *Game) handleMouseEvents(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()

	if coordsInRect(x, y, g.canvasCoords) {
		g.drawHoveredTileOnCanvas(screen, x, y)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.drawTileOnCanvas(screen, x, y)
		}
	}

	if coordsInRect(x, y, g.palleteCoords) {
		g.drawCursorOnPallete(screen, x, y)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.chooseTileFromPallete(x, y)
		}
	}
}

func (g *Game) drawPallete(screen *ebiten.Image) {
    for n := 0; n < g.tileset.num; n++ {
        subImg := g.tileNrToSubImageOnPallete(n)

		op := &ebiten.DrawImageOptions{}
        x,y := g.tileNrToCoordsOnPallete(n)
		op.GeoM.Translate(x, y)
		screen.DrawImage(subImg, op)

		if n > PalleteTilesMax - 2 {
			break
		}
	}
}

func (g *Game) drawCurrentTileToDraw(screen *ebiten.Image) {
	if g.currentTile.tile.tileset != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(CurrentTileToDrawX),
			float64(CurrentTileToDrawY),
		)
		screen.DrawImage(g.currentTile.tile.image, op)
	}
}

func (g *Game) drawCanvas(screen *ebiten.Image) {
	drawEmptyRect(screen, g.canvasCoords, color.White)

	for x := 0; x < CanvasColsN; x++ {
		for y := 0; y < CanvasRowsN; y++ {
			if tile := g.canvas[x][y]; tile != -1 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64(x*TileWidth+CanvasX),
					float64(y*TileWidth+CanvasY),
				)
				screen.DrawImage(g.tileStack[g.canvas[x][y]].image, op)
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPallete(screen)
	g.drawCanvas(screen)
	g.drawCanvas(screen)
	g.handleMouseEvents(screen)
	g.drawCurrentTileToDraw(screen)
}

func NewGame() *Game {
	game := &Game{
		palleteCoords: image.Rect(
			PalleteX,
			PalleteY,
			PalleteEndX,
			PalleteEndY,
		),
		currentTileToDrawRect: image.Rect(
			CurrentTileToDrawX,
			CurrentTileToDrawY,
			CurrentTileToDrawEndX,
			CurrentTileToDrawEndY,
		),
		canvasCoords: image.Rect(
			CanvasX,
			CanvasY-1,
			CanvasEndX+1,
			CanvasEndY,
		),
		tileset: newTileset("tileset_1.png"),
	}

	// post init
    game.addTileFromPalleteToStack(0, 0)
    game.setCurrentTile(0)
    game.clearCanvas()

	return game
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
