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
)

type Tileset struct {
	image     *ebiten.Image
	width     int
	height    int
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
	tileStack    [256]Tile
	tileStackLen int
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

func drawEmptyRect(screen *ebiten.Image, x1, y1, x2, y2 float64, color color.Color) {
	ebitenutil.DrawLine(screen, x1, y1, x2, y1, color)
	ebitenutil.DrawLine(screen, x1, y1, x1, y2+1, color)
	ebitenutil.DrawLine(screen, x2, y1, x2, y2, color)
	ebitenutil.DrawLine(screen, x1, y2, x2, y2, color)
}

func drawCursorOnPallete(screen *ebiten.Image, x int, y int, size float64) {
	leftCornerX := float64((((x - PalleteX) / TileWidth) * TileWidth) + PalleteX)
	leftCornerY := float64((((y - PalleteY) / TileHeight) * TileHeight) + PalleteY)

	drawEmptyRect(screen, leftCornerX, leftCornerY, leftCornerX+size, leftCornerY+size, color.White)
}

// checks if Tile needs to be deleted from TileStack. It happens if number of tile usage on level
// is equal to zero
func (g *Game) clearStack(index int) {
	if g.tileStack[index].NumberUsed == 0 {
		g.tileStack = append(g.tileStack[:index], g.tileStack[index+1:]...)
	}
}

func (g *Game) drawTileOnCanvas(screen *ebiten.Image, x int, y int) {
	tileX := int((x - CanvasX) / TileWidth)
	tileY := int((y - CanvasY) / TileHeight)

	if g.canvas[tileX][tileY] == 0 {
		g.tileStack[g.currentTile.StackIndex].NumberUsed++
	} else if g.canvas[tileX][tileY] == g.currentTile.StackIndex {
		// do nothing (bitchslap)
	} else {
		g.TileStack[g.canvas[tileX][tileY]].NumberUsed = -1
		g.clearStack()
		g.tileStack[g.currentTile.StackIndex].NumberUsed++
		g.canvas[tileX][tileY].NumberUsed++
	}
}

// TODO: this function should be divided into smaller functions!
func (g *Game) chooseTileFromPallete(x int, y int) {
	tileX := int((x - PalleteX) / TileWidth)
	tileY := int((y - PalleteY) / TileHeight)

	for i := 0; i < g.tileStackLen; i++ {
		if g.TileStack[index].PalleteX == tileX && g.TileStack[index].PalleteY == tileY {
			// chosen tile is already used so we take params from tileStack
			g.currentTile = CurrentTile{
				&g.TileStack[index],
				i,
			}
			return
		}
	}

	// new tile should be append to stack
	g.tileStackLen++

	tileRect = image.Rect(
		tileX*TileWidth,
		tileY*TileHeight,
		(tileX*TileWidth)+TileWidth,
		(tileY*TileWidth)+TileHeight,
	)

	g.TileStack[g.tileStackLen] = &CurrentTile{
		&Tile{
			&g.tileset,
			g.tileset.image.SubImage(tileRect).(*ebiten.Image),
			TileX,
			TileY,
		},
		g.tileStackLen,
	}
}

func (g *Game) drawHoveredTileOnCanvas(x int, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((((x-CanvasX)/TileWidth)*TileWidth)+CanvasX),
		float64((((y-CanvasY)/TileHeight)*TileHeight)+CanvasY),
	)
	screen.DrawImage(g.CurrentTile.image, op)
}

func (g *Game) handleMouseEvents(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()

	if coordsInRect(x, y, g.canvasCoords) {
		// draw cursor shaped as tile while mouse is above level
		g.drawHoveredTileOnCanvas(x, y)

		// mouse button is pressed, so tile is put into canvas
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.drawTileOnCanvas(screen, x, y)
		}
	}

	if coordsInRect(x, y, g.palleteCoords) {
		drawCursorOnPallete(screen, x, y, 32.0)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.drawTileOnCanvas(screen, x, y)
		}
	}
}

func (g *Game) drawPallete(screen *ebiten.Image) {
	tilesetWidth := g.tileset.width
	tilesetHeight := g.tileset.height

	var currentTile int
	var tilesheetRows int = tilesetWidth / TileWidth
	var tilesheetCols int = tilesetHeight / TileWidth
	var tilesNum int = tilesheetCols * tilesheetRows

	for currentTile < tilesNum {
		tileRect := image.Rect(
			(currentTile%tilesheetCols)*TileWidth,
			(currentTile/tilesheetCols)*TileHeight,
			((currentTile%tilesheetCols)*TileWidth)+TileWidth,
			((currentTile/tilesheetCols)*TileWidth)+TileHeight,
		)

		subImg := g.tileset.image.SubImage(tileRect).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(((currentTile%PalleteColsN)*TileWidth)+PalleteX),
			float64(((currentTile/PalleteColsN)*TileHeight)+PalleteY),
		)
		screen.DrawImage(subImg, op)
		currentTile++

		if currentTile >= PalleteTilesMax {
			break
		}
	}
}

func (g *Game) drawCurrentTileToDraw(screen *ebiten.Image) {
	subImg := g.tileset.image.SubImage(g.currentTileToDrawRect).(*ebiten.Image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(g.CurrentTileToDraw.Min.X),
		float64(g.CurrentTileToDraw.Min.Y),
	)
	screen.DrawImage(subImg, op)
}

func (g *Game) drawCanvas(screen *ebiten.Image) {
	drawEmptyRect(
		screen,
		float64(g.canvasCoords.Min.X),
		float64(g.canvasCoords.Min.Y),
		float64(g.canvasCoords.Max.X),
		float64(g.canvasCoords.Max.Y),
		color.White,
	)
	for x := 0; x < g.CanvasColsN; x++ {
		for y := 0; y < g.CanvasRowsN; y++ {
			if tile := canvas[x][y]; tile != 0 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64(x*TileWidth+g.CanvasX),
					float64(y*TileWidth+g.CanvasY),
				)
				screen.DrawImage(tile.image, op)
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return 640, 480
}

func NewGame() *Game {
	game := &Game{
		palleteCoords: image.Rect(
			PalleteX,
			PalleteY,
			PalleteEndX,
			PalleteEndY,
		),
		currentTileToDrawRect: image.Rect(0, 0, TileWidth, TileHeight),
		CurrentTileToDraw: image.Rect(
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
	return game
}

// placeholder gonna hold place
func main() {
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
