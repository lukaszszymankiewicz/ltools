package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawTileOnCanvas(screen *ebiten.Image) {
	fill := g.Toolbox.GetFillTile()
	if fill == nil {
		return
	}
	y, x := g.Canvas.MousePosToRowAndCol(g.mouse_x, g.mouse_y)
	corr_x, corr_y := g.Canvas.CorrectPosByViewport(x, y)

	tool := g.Toolbox.GetActiveTool()

	if g.Canvas.TileIsAllowed(corr_x, corr_y, g.mode, tool) == false {
		return
	}

	g.Canvas.PutTile(corr_x, corr_y, fill, tool)
}

func (g *Game) ChooseTileFromPallete(screen *ebiten.Image) {
	tile := g.Pallete.MousePosToTile(g.mouse_x, g.mouse_y, g.mode)
	if tile != nil {
		g.Toolbox.SetFillTile(tile)
	}
}

func (g *Game) UndrawOneRecord() {
	g.Canvas.UndrawOneRecord()
}

func (g *Game) DrawCursorOnPallete(screen *ebiten.Image) {
	// hiding OS cursor
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	x, y := g.Pallete.MousePosToTilePos(g.mouse_x, g.mouse_y)
	row, col := g.Pallete.MousePosToRowAndCol(x, y)

	if t := g.Pallete.PosHasTile(row, col, g.Pallete.CurrentLayer()); t == true {
		g.Cursor.DrawBorder(screen, x, y)
	}

    // drawing cursor icon
	image := g.Toolbox.GetPipette()
	g.Cursor.DrawIcon(screen, image, g.mouse_x, g.mouse_y)
}

func (g *Game) DrawCursorOnCanvas(screen *ebiten.Image) {
	// hiding OS cursor
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	// drawing place where Tile will be drawn (simply by drawing rectangle)
	x, y := g.Canvas.MousePosToTilePos(g.mouse_x, g.mouse_y)
    g.Cursor.DrawBorder(screen, x, y)

	// drawing cursor icon
	image := g.Toolbox.GetCursor()
	g.Cursor.DrawIcon(screen, image, g.mouse_x, g.mouse_y)
}

func (g *Game) DrawCursorElsewhere(screen *ebiten.Image) {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}

func (g *Game) changeMode(mode int) {
	g.mode = mode
	g.Tabber.ChangeCurrent(mode)
	g.Pallete.ChangeCurrent(mode)
	g.Pallete.RestartPalletePos()
	g.Canvas.ChangeCurrent(mode)
}

func (g *Game) changeModeToDraw(screen *ebiten.Image) {
	g.changeMode(MODE_DRAW)
}

func (g *Game) changeModeToDrawLight(screen *ebiten.Image) {
	g.changeMode(MODE_LIGHT)
}

func (g *Game) changeModeToDrawEntities(screen *ebiten.Image) {
	g.changeMode(MODE_ENTITIES)
}

func (g *Game) changeToolToPencil(screen *ebiten.Image) {
	g.Toolbox.Activate(PENCIL_TOOL)
}

func (g *Game) changeToolToRubber(screen *ebiten.Image) {
	g.Toolbox.Activate(RUBBER_TOOL)
}
