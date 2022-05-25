package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// draws Tile from LAYER_DRAW
func (g *Game) drawNormalTileOnCanvas(x int, y int, screen *ebiten.Image) {
	oldTile := g.Canvas.GetTileOnDrawingArea(x, y, LAYER_DRAW)
	lightTile := g.Canvas.GetTileOnDrawingArea(x, y, LAYER_LIGHT)

	tile_to_draw := g.Toolbox.GetFillTile()

	// if light obstacle is cleaned it should be cleaned from light layer too
	if lightTile != nil {
		g.StartRecording()
		g.Canvas.SetTileOnDrawingArea(x, y, LAYER_LIGHT, nil)
		g.Recorder.AppendToCurrent(x, y, nil, lightTile, LAYER_LIGHT)
	}

	// current place is empty
	if oldTile == nil {
		g.Logger.LogCanvasPut(x, y, LAYER_DRAW)
		g.StartRecording()
		g.Recorder.AppendToCurrent(x, y, tile_to_draw, oldTile, LAYER_DRAW)
		g.Canvas.SetTileOnDrawingArea(x, y, LAYER_DRAW, tile_to_draw)
		// current place is occupied by other tile (other than current tile to draw)
	} else if oldTile != tile_to_draw {
		g.StartRecording()
		g.Canvas.SetTileOnDrawingArea(x, y, LAYER_DRAW, tile_to_draw)
		g.Recorder.AppendToCurrent(x, y, tile_to_draw, oldTile, LAYER_DRAW)
	}
}

// draws tile from LAYER_LIGHT
func (g *Game) drawLightObstacleTile(x int, y int, screen *ebiten.Image) {
	lightTile := g.Canvas.GetTileOnDrawingArea(x, y, LAYER_LIGHT)
	oldTile := g.Canvas.GetTileOnDrawingArea(x, y, LAYER_DRAW)

	tile_to_draw := g.Toolbox.GetFillTile()

	if oldTile == nil && lightTile == nil {
		g.StartRecording()
		g.Canvas.SetTileOnDrawingArea(x, y, LAYER_LIGHT, tile_to_draw)
		g.Recorder.AppendToCurrent(x, y, tile_to_draw, nil, LAYER_LIGHT)
	}
}

// draws Tile from LAYER_ENTITY
func (g *Game) drawEntityTileOnCanvas(x int, y int, screen *ebiten.Image) {
	bgTile := g.Canvas.GetTileOnDrawingArea(x, y, LAYER_DRAW)
	oldTile := g.Canvas.GetTileOnDrawingArea(x, y, LAYER_ENTITY)
	tile_to_draw := g.Toolbox.GetFillTile()

	if bgTile != nil && oldTile == nil {

		// if tile is unique, it can be put at only one position on Canvas. If one of this tile is
		// already drawn, it must be erased and then put into new location chosen by user. To undraw
		// such Tile it must be first find.
		if tile_to_draw.IsUnique() == true {
			if tile_to_draw.IsSet() == false {
				g.StartRecording()
				g.Canvas.SetTileOnDrawingArea(x, y, LAYER_ENTITY, tile_to_draw)
				g.Recorder.AppendToCurrent(x, y, tile_to_draw, nil, LAYER_ENTITY)
				tile_to_draw.Put(x, y)
				// unique tile is used - old location needs to be found and Tile must be erased from
				// there
			} else {
				g.StartRecording()
				old_x, old_y := tile_to_draw.Loc()
				g.Canvas.SetTileOnDrawingArea(x, y, LAYER_ENTITY, tile_to_draw)
				g.Canvas.SetTileOnDrawingArea(old_x, old_y, LAYER_ENTITY, nil)

				g.Recorder.AppendToCurrent(x, y, tile_to_draw, nil, LAYER_ENTITY)
				g.Recorder.AppendToCurrent(old_x, old_y, nil, tile_to_draw, LAYER_ENTITY)
			}
		}
	}
}

func (g *Game) DrawTileOnCanvas(screen *ebiten.Image) {
	x, y := g.Canvas.MousePosToRowAndCol(g.mouse_x, g.mouse_y)

	if g.mode == MODE_LIGHT {
		g.drawLightObstacleTile(x, y, screen)
	} else if g.mode == MODE_DRAW {
		g.drawNormalTileOnCanvas(x, y, screen)
	} else if g.mode == MODE_ENTITIES {
		g.drawEntityTileOnCanvas(x, y, screen)
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
