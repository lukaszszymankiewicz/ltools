package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawingOptions struct {
	CanvasAlpha   []float64
	PalletteAlpha []float64
}

type Game struct {
	AvailableTiles TilesCollection
	LevelStructure TilesCollection

	ViewportPallete Viewport
	ViewportCanvas  Viewport

	FillTiles     TilesCollection
	ViewportTiles Viewport

	Mouse
	Cursor
	Toolbox

	Mode
	Config
	Logger

	WindowManager
}

func init() {
	ebiten.SetFullscreen(true)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return 1388, 768
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Mouse.Update()
	// g.handleKeyboardEvents()

	g.WindowManager.handleMouseEvents(g.Mouse.X, g.Mouse.Y)
	g.WindowManager.Update(g)
	g.WindowManager.Draw(screen)

}

func NewEditor() *Game {
	var g Game

	g.AvailableTiles = NewTilesCollection(LAYER_N)
	g.AvailableTiles.Alloc(PalleteRows * PalleteCols)

	g.LevelStructure = NewTilesCollection(LAYER_N)
	g.LevelStructure.Alloc(CanvasCols * CanvasRows)

	g.ViewportPallete = NewViewport(PalleteRows, PalleteCols, PalleteRows, PalleteCols)
	g.ViewportCanvas = NewViewport(CanvasRows, CanvasCols, CanvasViewportRows, CanvasViewportCols)

	g.FillTiles = NewTilesCollection(LAYER_N)
	g.FillTiles.Alloc(1)
	g.ViewportTiles = NewViewport(1, 1, 1, 1)

	g.Toolbox = NewToolbox()
	g.Mode = NewMode(MODE_ALL)
	g.Mouse = NewMouse()

	g.PostInit()

	return &g
}

func (g *Game) PostInit() {
	// WINDOW UTILS
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	// LOAD RESOURCES
	LoadResources("editor")

	g.AvailableTiles.Fill("basic_tileset", LAYER_DRAW)
	g.AvailableTiles.Fill("light_tileset", LAYER_LIGHT)
	g.AvailableTiles.Fill("entity_tileset", LAYER_ENTITY)
	g.AvailableTiles.Get(HERO_TILE, LAYER_ENTITY).unique_condition = UNIQUE_COND_IS_NOT_SET

	// WINDOWS
	g.WindowManager = NewWindowManager()
	g.WindowManager.Add(g.NewMainWindow())

	// RED CONFIG
	g.Config = ReadConfig("config.json")
}
