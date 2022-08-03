package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DEFAULT_FILL_TILE_IDX = 0
)

func (g *Game) DrawTileOnCanvas() {
	layer := g.Mode.Active()
	fill := g.FillTiles.Get(DEFAULT_FILL_TILE_IDX, layer)

	// in case anything go wrong - terminate the function
	if fill == nil {
		return
	}

	row, col := CalculateTilePos(g.Mouse.X, g.Mouse.Y, CanvasX, CanvasY, GridSize)
	tool := g.Toolbox.GetActiveTool()
	brush := tool.GetActiveBrush()

	PutTile(g.ViewportCanvas, g.LevelStructure, fill, brush, row, col, layer)
}

func (g *Game) ChooseTileFromPallete() {
	layer := g.Mode.Active()
	row, col := CalculateTilePos(g.Mouse.X, g.Mouse.Y, PalleteX, PalleteY, GridSize)
	idx := g.ViewportPallete.GetTileIdx(row, col)
	tile := g.AvailableTiles.Get(idx, layer)

	if tile != nil {
		g.FillTiles.Set(DEFAULT_FILL_TILE_IDX, layer, tile)
	}
}

// to be implemented!
func (g *Game) UndrawOneRecord() {
	// TODO
}

func (g *Game) setHighlightOnCanvas() {
	x1, y1, x2, y2 := CalculateHighlightPos(g.Mouse.X, g.Mouse.Y, CanvasX, CanvasY, GridSize)
	g.Mouse.SetHighlight(x1, y1, x2, y2)
}

func (g *Game) setHighlightOnPallete() {
	x1, y1, x2, y2 := CalculateHighlightPos(g.Mouse.X, g.Mouse.Y, PalleteX, PalleteY, GridSize)
	g.Mouse.SetHighlight(x1, y1, x2, y2)
}

func (g *Game) changeToolToPencil() {
	g.Toolbox.Mode.Activate(PENCIL_TOOL_IDX)
}

func (g *Game) changeToolToRubber() {
	g.Toolbox.Mode.Activate(RUBBER_TOOL_IDX)
}

func (g *Game) setNormalIcon() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	g.Mouse.Icon.Activate(MOUSE_NORMAL)
}

func (g *Game) clearHighlight() {
	g.Mouse.Highlight.Activate(MOUSE_NO_HIGHLIGH)
	g.Mouse.CleanHighlight()
}

func (g *Game) setPippeteIcon() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	g.Mouse.Icon.Activate(MOUSE_PIPPETE)
}

func (g *Game) setToolIcon() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	tool := g.Toolbox.Mode.Active() + 1 // this wont crash anything, I'm sure
	g.Mouse.Icon.Activate(tool)
}

func (g *Game) moveCanvasDown() {
	g.ViewportCanvas.Move(0, 1)
}

func (g *Game) moveCanvasUp() {
	g.ViewportCanvas.Move(0, -1)
}

func (g *Game) moveCanvasRight() {
	g.ViewportCanvas.Move(1, 0)
}

func (g *Game) moveCanvasLeft() {
	g.ViewportCanvas.Move(-1, 0)
}

func (g *Game) movePalleteDown() {
	g.ViewportPallete.Move(0, 1)
}

func (g *Game) movePalleteUp() {
	g.ViewportPallete.Move(0, -1)
}

func (g *Game) changeModeToDraw() {
	g.Mode.Activate(MODE_DRAW)
}

func (g *Game) changeModeToDrawLight() {
	g.Mode.Activate(MODE_LIGHT)
}

func (g *Game) changeModeToDrawEntities() {
	g.Mode.Activate(MODE_ENTITIES)
}

func (g *Game) KillWindowOnTop() {
	g.WindowManager.KillLast()
}
