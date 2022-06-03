package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) DrawTileOnCanvas(screen *ebiten.Image) {
    // future brush (like drawing Rectangle for example will take two pair of x and y coords as we
    // need to set first and second corner) will need to have bigger number of inputs. But for now
    // only 'single click' tools are availalble. To allow for future brushes two pairs of coords are
    // passed to calculate the brush effect (coords are simply doubled).

	x, y := g.Canvas.MousePosToRowAndCol(g.mouse_x, g.mouse_y)

    fill := g.Toolbox.GetFillTile()
    brush := g.Toolbox.GetActiveTool()
    result := brush.ApplyBrush(x, y, x, y, fill)

    // if main position is not allowed to have tile in it function aborts
    if g.Canvas.TileIsAllowed(x, y, g.mode) == false {
        return
    }

    for i:=0; i<result.Len(); i++ {
        pos_x, pos_y := result.GetCoord(i)
        layer := result.GetLayer(i)
        tile_to_draw := result.GetTile(i)

        if g.Canvas.TileIsAllowed(pos_x, pos_y, layer) == true {
            g.Canvas.DrawBrush(pos_x, pos_y, layer, tile_to_draw)
        }
    }
}

func (g *Game) ChooseTileFromPallete(screen *ebiten.Image) {
	tile := g.Pallete.MousePosToTile(g.mouse_x, g.mouse_y, g.mode)
	g.Toolbox.SetFillTile(tile)
}

func (g *Game) UndrawOneRecord() {
	record := g.Recorder.UndoOneRecord()

	// if record to undraw is empty there is nothing to undraw!
	if record.IsEmpty() {
		return
	}

	g.Canvas.EraseOneRecord(record)
}

func (g *Game) DrawCursorOnPallete(screen *ebiten.Image) {
	x, y := g.Pallete.MousePosToTilePos(g.mouse_x, g.mouse_y)
    row, col := g.Pallete.MousePosToRowAndCol(x, y)

    if t:=g.Pallete.PosHasTile(row, col, g.Pallete.CurrentLayer()); t == true {
        g.Cursor.DrawOnPallete(screen, x, y)
    }
}

func (g *Game) DrawCursorOnCanvas(screen *ebiten.Image) {
	x, y := g.Canvas.MousePosToTilePos(g.mouse_x, g.mouse_y)
	image := g.Toolbox.GetFill()
	g.Cursor.DrawOnCanvas(screen, image, x, y)
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

