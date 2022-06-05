package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawTileOnCanvas(screen *ebiten.Image) {
	x, y := g.Canvas.MousePosToRowAndCol(g.mouse_x, g.mouse_y)
	fill := g.Toolbox.GetFillTile()
	tool := g.Toolbox.GetActiveTool()

	if g.Canvas.TileIsAllowed(x, y, g.mode) == false {
		return
	}

	g.Canvas.PutTile(x, y, fill, tool)
}

func (g *Game) ChooseTileFromPallete(screen *ebiten.Image) {
	tile := g.Pallete.MousePosToTile(g.mouse_x, g.mouse_y, g.mode)
	g.Toolbox.SetFillTile(tile)
}

func (g *Game) UndrawOneRecord() {
	g.Canvas.UndrawOneRecord()
}

func (g *Game) DrawCursorOnPallete(screen *ebiten.Image) {
	x, y := g.Pallete.MousePosToTilePos(g.mouse_x, g.mouse_y)
	row, col := g.Pallete.MousePosToRowAndCol(x, y)

	if t := g.Pallete.PosHasTile(row, col, g.Pallete.CurrentLayer()); t == true {
		g.Cursor.DrawOnPallete(screen, x, y)
	}
}

func (g *Game) DrawCursorOnCanvas(screen *ebiten.Image) {
	image := g.Toolbox.GetCursor()
	_, height := image.Size()
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	g.Cursor.DrawOnCanvas(screen, image, g.mouse_x, g.mouse_y-height)
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
