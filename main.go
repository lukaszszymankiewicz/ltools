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
	lto.TileStack
	lto.Cursor
    lto.Tabber
	Recorder
    Controller
	ClickableAreas       map[image.Rectangle]func(*ebiten.Image)
	SingleClickableAreas map[image.Rectangle]func(*ebiten.Image)
	HoverableAreas       map[image.Rectangle]func(*ebiten.Image)
    mode                 int
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
	g.UpdateControllersState()
    g.DrawTabber(screen)
	g.DrawPallete(screen)
	g.DrawCanvas(screen, g.GetAllTiles(), g.mode)
	g.handleKeyboardEvents()
    g.drawCurrentTileToDraw(screen)
	g.handleMouseEvents(screen)
}

// creates new game instance
func NewGame() *Game {
	var g Game

	g.Pallete = lto.NewPallete(
		PalleteX,
		PalleteY,
		PalleteRowsN,
		PalleteColsN,
        [][]string{
            {"assets/tileset_1.png"}, 
            {"assets/no_light.png"}, 
            {"assets/hero_entity_icon.png"}, 
        },
	)

	g.Canvas = lto.NewCanvas(
		ViewportRows,
		ViewportCols,
		CanvasRows,
		CanvasCols,
		Canvas_x,
		Canvas_y,
        MODE_ALL,
	)
	g.Cursor = lto.NewCursor(CursorSize)
    g.mode   = MODE_DRAW 
    g.Tabber = lto.NewCompleteTabber(TabberX, TabberY)

	// post init (works only on already initialised structs)
	g.AddTileFromPalleteToStack(0, 0, 0, false)
	g.AddTileFromPalleteToStack(0, 0, 1, false)
	g.AddTileFromPalleteToStack(0, 0, 2, true)
	g.SetCurrentTile(0)

	// binding the functions (yeah, it looks kinda lame)
	g.ClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))
	g.HoverableAreas = make(map[image.Rectangle]func(*ebiten.Image))
	g.SingleClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))

    g.ClickableAreas[g.Pallete.Rect] = g.ChooseTileFromPallete
    g.HoverableAreas[g.Pallete.Rect] = g.DrawCursorOnPallete
    g.ClickableAreas[g.Canvas.Rect] = g.DrawTileOnCanvas
	g.HoverableAreas[g.Canvas.Rect] = g.DrawHoveredTileOnCanvas

	g.SingleClickableAreas[g.Pallete.Scroller_y.LowArrowRect()] = g.Pallete.MovePalleteUp
	g.SingleClickableAreas[g.Pallete.Scroller_y.HighArrowRect()] = g.Pallete.MovePalleteDown

	g.SingleClickableAreas[g.Canvas.Scroller_x.LowArrowRect()] = g.Canvas.MoveCanvasLeft
	g.SingleClickableAreas[g.Canvas.Scroller_x.HighArrowRect()] = g.Canvas.MoveCanvasRight
	g.SingleClickableAreas[g.Canvas.Scroller_y.LowArrowRect()] = g.Canvas.MoveCanvasUp
	g.SingleClickableAreas[g.Canvas.Scroller_y.HighArrowRect()] = g.Canvas.MoveCanvasDown
	g.SingleClickableAreas[g.Tabber.AreaRect(MODE_DRAW)] = g.changeModeToDraw
	g.SingleClickableAreas[g.Tabber.AreaRect(MODE_LIGHT)] = g.changeModeToDrawLight
	g.SingleClickableAreas[g.Tabber.AreaRect(MODE_ENTITIES)] = g.changeModeToDrawEntities

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
