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
    ScrollArrowRight lto.ScrollArrow
    ScrollArrowLeft lto.ScrollArrow
    ScrollArrowUp lto.ScrollArrow
    ScrollArrowDown lto.ScrollArrow
}

func init() {
	ebiten.SetFullscreen(true)
}

// main loop
func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return 1388, 768
}


func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPallete(screen)
	g.DrawCanvas(screen, g.GetAllTiles())
	g.handleMouseEvents(screen)
	g.drawCurrentTileToDraw(screen)
    // arrows
    g.ScrollArrowRight.DrawScrollArrow(screen)
    g.ScrollArrowLeft.DrawScrollArrow(screen)
    g.ScrollArrowUp.DrawScrollArrow(screen)
    g.ScrollArrowDown.DrawScrollArrow(screen)
    // zip
    g.drawZippers(screen)
}

func NewGame() *Game {
	var g Game

	g.Pallete = lto.NewPallete(
		PalleteX,
		PalleteY,
		PalleteEndX,
		PalleteEndY,
		PalleteRowsN,
		PalleteColsN,
		PalleteMaxTile,
	)

	g.Canvas = lto.NewCanvas(
		ViewportRows,
		ViewportCols,
		CanvasRows,
		CanvasCols,
		Canvas_x,
		Canvas_y,
	)

	g.Tileset = lto.NewTileset("assets/tileset_1.png")
    g.ScrollArrowRight = lto.NewScrollArrow(ArrowRightX, ArrowRightY, "assets/arrow_r.png")
    g.ScrollArrowLeft = lto.NewScrollArrow(ArrowLeftX, ArrowLeftY, "assets/arrow_l.png")
    g.ScrollArrowUp = lto.NewScrollArrow(ArrowUpX, ArrowUpY, "assets/arrow_u.png")
    g.ScrollArrowDown = lto.NewScrollArrow(ArrowDownX, ArrowDownY, "assets/arrow_d.png")

	// post init
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
