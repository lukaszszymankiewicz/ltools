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
	tileWidth    = 32
	tileHeight   = 32
	// pallete
	screenTilePalleteX       = 20
	screenTilePalleteY       = 20
	nColsInScreenTilePallete = 6
	nRowsInScreenTilePallete = 10
	screenTilePalleteEndX    = screenTilePalleteX + (tileWidth * nColsInScreenTilePallete)
	screenTilePalleteEndY    = screenTilePalleteY + (tileHeight * nRowsInScreenTilePallete)
	currentTiletoDrawX       = screenTilePalleteEndX / 2
	currentTiletoDrawY       = screenTilePalleteEndY + tileHeight
	currentTiletoDrawEndX    = currentTiletoDrawX + tileWidth
	currentTiletoDrawEndY    = currentTiletoDrawY + tileHeight
	maxTileinTilesetPallete  = nColsInScreenTilePallete * nRowsInScreenTilePallete
	// canvas
	screenCanvasTileX = 10
	screenCanvasTileY = 10
	screenCanvasX     = screenTilePalleteEndX + (tileWidth * 2)
	screenCanvasY     = screenTilePalleteY
	screenCanvasEndX  = screenCanvasX + (tileWidth * screenCanvasTileX)
	screenCanvasEndY  = screenCanvasY + (tileWidth * screenCanvasTileY)
)

type Tileset struct {
	image     *ebiten.Image
	width      int
	height     int
    tilesUsed  int
}

type Tile struct {
    tileset *Tileset
	image   *ebiten.Image
    PalleteX int
    PalleteY int
    NumberUsed int
}

type CurrentTile struct {
    tile *Tile 
    StackIndex int
}

type Game struct {
	// static desktop coords
    currentTile CurrentTile 
	tilePalleteCoords image.Rectangle
	canvasCoords      image.Rectangle

	// dynamic thingis
	currentTileToDrawRect image.Rectangle
	levelTiles [screenCanvasX][screenCanvasY]int

	// images
	tilesetPallete Tileset
	tileStack [256]Tile
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
	leftCornerX := float64((((x - screenTilePalleteX) / tileWidth) * tileWidth) + screenTilePalleteX)
	leftCornerY := float64((((y - screenTilePalleteY) / tileHeight) * tileHeight) + screenTilePalleteY)

	drawEmptyRect(screen, leftCornerX, leftCornerY, leftCornerX+size, leftCornerY+size, color.White)
}

// checks if Tile needs to be deleted from TileStack. It happens if number of tile usage on level
// is equal to zero
func (g *Game) clearStack(index int) {
    if g.tileStack[index].NumberUsed == 0 {
        g.tileStack = append(g.tileStack[:index], g.tileStack[index+1:]...)
    }
}

func (g *Game) DrawTileOnCanvas(screen *ebiten.Image, x int, y int) {
    tileX := int((x - screenCanvasX) / tileWidth)
    tileY := int((y - screenCanvasY) / tileHeight)
    
    if g.levelTiles[tileX][tileY] == 0 {
        g.tileStack[g.currentTile.StackIndex].NumberUsed++
    }
    else if g.levelTiles[tileX][tileY] == g.currentTile.StackIndex {
        // do nothing (bitchslap)
    }
    else {
        g.TileStack[g.levelTiles[tileX][tileY]].NumberUsed =- 1
        g.clearStack()
        g.tileStack[g.currentTile.StackIndex].NumberUsed++
        g.levelTiles[tileX][tileY].NumberUsed++
    }

}

// TODO: this function should be divided into smaller functions!
func (g *Game) chooseTileFromPallete(x int, y int) {

        tileX := int((x - screenTilePalleteX) / tileWidth)
        tileY := int((y - screenTilePalleteY) / tileHeight)

		for i := 0; i < g.tileStackLen; i++ {
            // TODO: different palletes are not possible to use
			if g.TileStack[index].PalleteX == tileX && g.TileStack[index].PalleteY == tileY {
                // chosen tile is already used so we take params from tileStack
				g.currentTile := {
                    &g.TileStack[index], 
					i
                }
                return
            }
        }
	
        // new tile should be append to stack
        g.tileStackLen++

		tileRect = image.Rect(
			tileX*tileWidth,
			tileY*tileHeight,
			(tileX*tileWidth)+tileWidth,
			(tileY*tileWidth)+tileHeight,
		)

        g.TileStack[g.tileStackLen] := &CurrentTile {
            &Tile {
                &g.tilesetPallete,
                g.tilesetPallete.image.SubImage(tileRect).(*ebiten.Image),
                TileX,
                TileY,
            }
            g.tileStackLen
        }
}

func (g *Game) DrawHoveredTileOnCanvas(x int, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64((((x-screenCanvasX)/tileWidth)*tileWidth)+screenCanvasX),
		float64((((y-screenCanvasY)/tileHeight)*tileHeight)+screenCanvasY),
	)
	screen.DrawImage(g.CurrentTile.image, op)
}

func (g *Game) HandleMouse(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()

	if coordsInRect(x, y, g.canvasCoords) {
        // draw cursor shaped as tile while mouse is above level
		g.DrawHoveredTileOnCanvas(x, y)

		// mouse button is pressed, so tile is put into canvas
        if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
            g.DrawTileOnCanvas(screen, x, y)
        }
	}

	if coordsInRect(x, y, g.tilePalleteCoords) {
		drawCursorOnPallete(screen, x, y, 32.0)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.DrawTileOnCanvas(screen, x, y)
		}
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

func (g *Game) DrawCanvas(screen *ebiten.Image) {
	drawEmptyRect(
		screen,
		float64(g.canvasCoords.Min.X),
		float64(g.canvasCoords.Min.Y),
		float64(g.canvasCoords.Max.X),
		float64(g.canvasCoords.Max.Y),
		color.White,
	)
}

func (g *Game) DrawLevel(screen *ebiten.Image) {
	for x := 0; x<g.screenCanvasTileX; x++ {
		for y := y<g.screenCanvasTileY; y++ {
			if tile:=levelTiles[x][y] != 0 {

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(
					float64(x * tileWidth + g.screenCanvasX),
					float64(y * tileWidth + g.screenCanvasY),
				)
				screen.DrawImage(tile.image, op)
			}

		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawTilesetPallete(screen)
	g.DrawCanvas(screen)
	g.DrawLevel(screen)
	g.HandleMouse(screen)
	g.DrawCurrentTile(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func NewGame() *Game {
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
		canvasCoords: image.Rect(
			screenCanvasX,
			screenCanvasY - 1,
			screenCanvasEndX + 1,
			screenCanvasEndY,
		),
		tilesetPallete: newTileset("tileset_1.png"),
	}
	return game
}

// placeholder gonna hold place
func main() {
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
