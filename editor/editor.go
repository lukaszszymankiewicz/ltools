package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
    AvailableTiles TilesCollection
    LevelStructure TilesCollection

    ViewportPallete Viewport
    ViewportCanvas Viewport

    mode              int
	// Logger
	// Tilesets
	// Controller
    // Config
    // CurrentTool          int
    Tilesets  // TODO: change it to Resources!

    WindowManager
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
	// g.UpdateControllersState()
	// g.handleMouseEvents(screen)

    // g.MainWindow.Update()
    g.WindowManager.Update(g)
    g.WindowManager.Draw(screen)

    //g.MainWindow.Update(g)
    //g.MainWindow.Draw(screen)

	// g.Tabber.Draw(screen)
	// g.Canvas.Draw(screen, g.mode)
	// g.handleKeyboardEvents()
	// g.Toolbox.Draw(screen)
}

// creates new game instance
func NewEditor() *Game {
	var g Game

    // DATA
	g.AvailableTiles = NewTilesCollection(LAYER_N)
    g.LevelStructure = NewTilesCollection(LAYER_N)
    g.LevelStructure.Alloc(CanvasCols*CanvasRows)

    g.ViewportPallete = NewViewport(PalleteRows, PalleteCols, PalleteRows, PalleteCols)
    g.ViewportCanvas = NewViewport(CanvasRows, CanvasCols, CanvasViewportRows, CanvasViewportCols)

	// g.mode = MODE_DRAW

    // WINDOW UTILS
    ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	// RESOURCES
	g.Tilesets = NewTilesets(
		[]string{
            "editor/assets/basic_tileset.png",
            "editor/assets/light_tileset.png",
            "editor/assets/entity_tileset.png",
		},
	)
	for i := 0; i < g.Tilesets.AvailableTilesets(); i++ {
		g.AvailableTiles.Fill(g.Tilesets.GetById(i), i)
	}

    // WINDOWS
    g.WindowManager = NewWindowManager()
    g.WindowManager.Add(g.NewMainWindow())

	// g.Toolbox = NewToolbox(ToolboxX, ToolboxY)
	// g.Cursor = NewCursor(CursorSize)
	// g.Tabber = NewCompleteTabber(TabberX, TabberY, []string{"tiles", "light", "entities", "export"})

    // g.Config = ReadConfig("config.json")

    // g.PostInit()

	return &g
}

// creates new game instance
// func (g *Game) PostInit() {
// 	g.Logger = NewLogger(LOGGER_PATH)
// 
// 
// 	// function bindings
// 	g.ClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))
// 	g.HoverableAreas = make(map[image.Rectangle]func(*ebiten.Image))
// 	g.SingleClickableAreas = make(map[image.Rectangle]func(*ebiten.Image))
// 
// 	g.HoverableAreas[g.Pallete.Area()] = g.DrawCursorOnPallete
// 	g.HoverableAreas[g.Canvas.Area()] = g.DrawCursorOnCanvas
// 	g.ClickableAreas[g.Pallete.Area()] = g.ChooseTileFromPallete
// 	g.ClickableAreas[g.Canvas.Area()] = g.DrawTileOnCanvas
// 
// 	g.SingleClickableAreas[g.Pallete.ScrollerYUpArrowArea()] = g.Pallete.MovePalleteUp
// 	g.SingleClickableAreas[g.Pallete.ScrollerYDownArrowArea()] = g.Pallete.MovePalleteDown
// 
// 	g.SingleClickableAreas[g.Canvas.ScrollerXLeftArrowArea()] = g.Canvas.MoveCanvasLeft
// 	g.SingleClickableAreas[g.Canvas.ScrollerXRightArrowArea()] = g.Canvas.MoveCanvasRight
// 	g.SingleClickableAreas[g.Canvas.ScrollerYUpArrowArea()] = g.Canvas.MoveCanvasUp
// 	g.SingleClickableAreas[g.Canvas.ScrollerYDownArrowArea()] = g.Canvas.MoveCanvasDown
// 
// 	g.SingleClickableAreas[g.Tabber.Area(MODE_DRAW)] = g.changeModeToDraw
// 	g.SingleClickableAreas[g.Tabber.Area(MODE_LIGHT)] = g.changeModeToDrawLight
// 	g.SingleClickableAreas[g.Tabber.Area(MODE_ENTITIES)] = g.changeModeToDrawEntities
// 	g.SingleClickableAreas[g.Tabber.Area(EXPORT_BUTTON)] = g.Export
// 
// 	g.SingleClickableAreas[g.Toolbox.Area(PENCIL_TOOL)] = g.changeToolToPencil
// 	g.SingleClickableAreas[g.Toolbox.Area(RUBBER_TOOL)] = g.changeToolToRubber
// }
