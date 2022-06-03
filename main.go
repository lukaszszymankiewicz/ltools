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
	lto.Cursor
	lto.Tabber
	lto.Recorder
	lto.Toolbox
	lto.Logger
	lto.Tilesets
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
	g.Tabber.Draw(screen)
	g.Pallete.Draw(screen)
	g.Canvas.Draw(screen, g.mode)
	g.handleKeyboardEvents()
	g.Toolbox.Draw(screen)
	g.handleMouseEvents(screen)
}

// creates new game instance
func NewGame() *Game {
	var g Game

    g.mode = MODE_DRAW

	g.Tilesets = lto.NewTilesets(
		[]string{
			"assets/tilesets/basic_tileset.png",
			"assets/tilesets/light_tileset.png",
			"assets/tilesets/entity_tileset.png",
		},
	)

	g.Pallete = lto.NewPallete(
		PalleteX,
		PalleteY,
		PalleteColsN*GridSize,
		PalleteRowsN*GridSize,
		GridSize,
		PalleteRowsN + 1,
		PalleteColsN,
		LAYER_N,
	)

	g.Canvas = lto.NewCanvas(
		Canvas_x,
		Canvas_y,
		CanvasColsN*GridSize,
		CanvasRowsN*GridSize,
		GridSize,
		50, // absolutely temporary
		40, // absolutely temporary
		LAYER_N,
	)

	g.Toolbox = lto.NewToolbox(ToolboxX, ToolboxY)
	g.Cursor = lto.NewCursor(CursorSize)
	g.Tabber = lto.NewCompleteTabber(TabberX, TabberY, []string{"tiles", "light", "entities", "export"})
	g.Logger = lto.NewLogger(LOGGER_PATH)

	return &g
}

// creates new game instance
func (g *Game) PostInit() {
    // initial setups
	for i := 0; i<g.Tilesets.AvailableTilesets(); i++ {
        g.Pallete.FillPallete(g.Tilesets.GetById(i), i)
	}
    
    // function bindings
	g.ClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))
	g.HoverableAreas = make(map[image.Rectangle]func(*ebiten.Image))
	g.SingleClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))

	g.HoverableAreas[g.Pallete.Area()] = g.DrawCursorOnPallete
	g.HoverableAreas[g.Canvas.Area()] = g.DrawCursorOnCanvas
	g.ClickableAreas[g.Pallete.Area()] = g.ChooseTileFromPallete
	g.ClickableAreas[g.Canvas.Area()] = g.DrawTileOnCanvas

	g.SingleClickableAreas[g.Pallete.ScrollerYUpArrowArea()] = g.Pallete.MovePalleteUp
	g.SingleClickableAreas[g.Pallete.ScrollerYDownArrowArea()] = g.Pallete.MovePalleteDown

	g.SingleClickableAreas[g.Canvas.ScrollerXLeftArrowArea()] = g.Canvas.MoveCanvasLeft
	g.SingleClickableAreas[g.Canvas.ScrollerXRightArrowArea()] = g.Canvas.MoveCanvasRight
	g.SingleClickableAreas[g.Canvas.ScrollerYUpArrowArea()] = g.Canvas.MoveCanvasUp
	g.SingleClickableAreas[g.Canvas.ScrollerYDownArrowArea()] = g.Canvas.MoveCanvasDown

	g.SingleClickableAreas[g.Tabber.Area(MODE_DRAW)] = g.changeModeToDraw
	g.SingleClickableAreas[g.Tabber.Area(MODE_LIGHT)] = g.changeModeToDrawLight
	g.SingleClickableAreas[g.Tabber.Area(MODE_ENTITIES)] = g.changeModeToDrawEntities
	g.SingleClickableAreas[g.Tabber.Area(EXPORT_BUTTON)] = g.Export

	g.SingleClickableAreas[g.Toolbox.Area(PENCIL_TOOL)] = g.changeToolToPencil
	g.SingleClickableAreas[g.Toolbox.Area(RUBBER_TOOL)] = g.changeToolToRubber
}

func main() {
	game := NewGame()
	game.PostInit()

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
