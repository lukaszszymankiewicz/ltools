package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	lto "ltools/src/objects"
)

type Game struct {
	lto.Pallete
	lto.Canvas
	lto.Tileset
	lto.TileStack
	lto.Cursor
	Recorder
	mouse_x              int
	mouse_y              int
	ClickableAreas       map[image.Rectangle]func(*ebiten.Image)
	SingleClickableAreas map[image.Rectangle]func(*ebiten.Image)
	HoverableAreas       map[image.Rectangle]func(*ebiten.Image)
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

// updates all controlers global state (only mouse by now)
func (g *Game) UpdateControllersState() {
	g.mouse_x, g.mouse_y = ebiten.CursorPosition()
}

// game draw loop
func (g *Game) Draw(screen *ebiten.Image) {
	g.UpdateControllersState()
	g.drawPallete(screen)
	g.DrawCanvas(screen, g.GetAllTiles())
	g.handleKeyboardEvents()
	g.drawCurrentTileToDraw(screen)
	g.handleMouseEvents(screen)
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
	g.Cursor = lto.NewCursor(CursorSize)

	// post init (works only on already initialised structs)
	g.addTileFromPalleteToStack()
	g.SetCurrentTile(0)

	// binding the functions (yeah, it looks kinda lame)
	g.ClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))
	g.HoverableAreas = make(map[image.Rectangle]func(*ebiten.Image))
	g.SingleClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))

	g.ClickableAreas[g.Canvas.Rect] = g.drawTileOnCanvas
	g.ClickableAreas[g.Pallete.Rect] = g.chooseTileFromPallete

	g.SingleClickableAreas[g.Pallete.Scroller_y.LowArrowRect()] = g.Pallete.MovePalleteUp
	g.SingleClickableAreas[g.Pallete.Scroller_y.HighArrowRect()] = g.Pallete.MovePalleteDown

	g.SingleClickableAreas[g.Canvas.Scroller_x.LowArrowRect()] = g.Canvas.MoveCanvasLeft
	g.SingleClickableAreas[g.Canvas.Scroller_x.HighArrowRect()] = g.Canvas.MoveCanvasRight
	g.SingleClickableAreas[g.Canvas.Scroller_y.LowArrowRect()] = g.Canvas.MoveCanvasUp
	g.SingleClickableAreas[g.Canvas.Scroller_y.HighArrowRect()] = g.Canvas.MoveCanvasDown

	g.HoverableAreas[g.Canvas.Rect] = g.drawHoveredTileOnCanvas
	g.HoverableAreas[g.Pallete.Rect] = g.drawCursorOnPallete

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
