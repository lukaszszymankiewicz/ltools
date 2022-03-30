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
	ScrollArrowRight       lto.ScrollArrow
	ScrollArrowLeft        lto.ScrollArrow
	ScrollArrowUp          lto.ScrollArrow
	ScrollArrowDown        lto.ScrollArrow
	PalleteScrollArrowUp   lto.ScrollArrow
	PalleteScrollArrowDown lto.ScrollArrow
    Recorder 
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
	g.drawPallete(screen)
	g.DrawCanvas(screen, g.GetAllTiles())
	g.handleMouseEvents(screen)
    g.handleKeyboardEvents()
	g.drawCurrentTileToDraw(screen)
	// arrows
	g.ScrollArrowRight.DrawScrollArrow(screen)
	g.ScrollArrowLeft.DrawScrollArrow(screen)
	g.ScrollArrowUp.DrawScrollArrow(screen)
	g.ScrollArrowDown.DrawScrollArrow(screen)
	g.PalleteScrollArrowUp.DrawScrollArrow(screen)
	g.PalleteScrollArrowDown.DrawScrollArrow(screen)

}

// creates new game instance
func NewGame() *Game {
	var g Game

	g.Tileset = lto.NewTileset("assets/tileset_1.png")

	g.Pallete = lto.NewPallete(
		PalleteX,
		PalleteY,
		PalleteEndX,
		PalleteEndY,
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

	g.ScrollArrowRight = lto.NewScrollArrow(ArrowRightX, ArrowRightY, "assets/arrow_r.png")
	g.ScrollArrowLeft = lto.NewScrollArrow(ArrowLeftX, ArrowLeftY, "assets/arrow_l.png")
	g.ScrollArrowUp = lto.NewScrollArrow(ArrowUpX, ArrowUpY, "assets/arrow_u.png")
	g.ScrollArrowDown = lto.NewScrollArrow(ArrowDownX, ArrowDownY, "assets/arrow_d.png")
	g.PalleteScrollArrowUp = lto.NewScrollArrow(PalleteArrowUpX, PalleteArrowUpY, "assets/arrow_u.png")
	g.PalleteScrollArrowDown = lto.NewScrollArrow(PalleteArrowDownX, PalleteArrowDownY, "assets/arrow_d.png")

	// post init (works only on initialised structs)
	g.addTileFromPalleteToStack(0, 0)
	g.SetCurrentTile(0)
	g.Canvas.ClearDrawingArea()

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
