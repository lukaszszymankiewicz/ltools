package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
	lto "ltools/src/objects"
)

type Game struct {
	lto.Pallete
	lto.Canvas
	lto.Tileset
	lto.TileStack
    Recorder
	mouse_x int
	mouse_y int
}

// everything that needs to be set before first game loop iteration
func init() {
	ebiten.SetFullscreen(true)
}

// game loop
func (g *Game) Update() error {
	return nil
}

// returning game internal resolution
func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return 1388, 768
}

// game draw loop
func (g *Game) Draw(screen *ebiten.Image) {
	g.mouse_x, g.mouse_y := ebiten.CursorPosition()
	g.drawPallete(screen)
	g.DrawCanvas(screen, g.GetAllTiles())
	g.handleMouseEvents(screen)
    g.handleKeyboardEvents()
	g.drawCurrentTileToDraw(screen)
}

// creates new game instance
func NewGame() *Game {
	var g Game

	g.Tileset = lto.NewTileset("assets/tileset_1.png")

	g.Pallete = lto.NewPallete(
		PalleteX,
		PalleteY,
		PalleteRowsN,
		PalleteColsN,
		g.Tileset.Num/PalleteColsN,
	)

	g.Canvas = lto.NewCanvas(
		ViewportRows,
		ViewportCols,
		CanvasRows,
		CanvasCols,
		Canvas_x,
		Canvas_y,
	)

	// post init (works only on initialised structs)
	g.addTileFromPalleteToStack(0, 0)
	g.SetCurrentTile(0)

	// binding the functions
	c.clickableAreas[c.Rect] = g.drawTileOnCanvas(screen)
	p.clickableAreas[p.Rect] = g.chooseTileFromPallete()

	c.clickableAreas[c.Rect] = g.drawTileOnCanvas(screen)
	p.clickableAreas[p.Rect] = g.chooseTileFromPallete()

	c.hoverableAreas[c.Rect] = g.drawHoveredTileOnCanvas(screen)
	p.hoverableAreas[p.Rect] = g.DrawCursorOnPallete(screen)

	return &g
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
