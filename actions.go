package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	lto "ltools/src/objects"
)

// draws Tile from LAYER_DRAW
func (g *Game) drawNormalTileOnCanvas(x int, y int, screen *ebiten.Image) {
	oldTile := g.GetTileOnDrawingArea(x, y, LAYER_DRAW)
	lightTile := g.GetTileOnDrawingArea(x, y, LAYER_LIGHT)
	drawTile := g.CurrentTileIndex()

	// if light obstacle is cleaned it should be cleaned from light layer too
	if lightTile != -1 {
		g.SetTileOnCanvas(x, y, EMPTY, LAYER_LIGHT)
		g.Recorder.AppendToCurrent(x, y, EMPTY, lightTile, LAYER_LIGHT, LAYER_LIGHT)
	}

	// current place is empty
	if oldTile == EMPTY {
		g.UpdateTileUsage(drawTile, +1)
		g.Recorder.AppendToCurrent(x, y, drawTile, oldTile, LAYER_DRAW, LAYER_DRAW)
		g.SetTileOnCanvas(x, y, drawTile, LAYER_DRAW)
		// current place is occupied by other tile (other than current tile to draw)
	} else if oldTile != drawTile {
		g.UpdateTileUsage(drawTile, +1)
		g.UpdateTileUsage(oldTile, -1)
		g.SetTileOnCanvas(x, y, drawTile, LAYER_DRAW)
		g.Recorder.AppendToCurrent(x, y, drawTile, oldTile, LAYER_DRAW, LAYER_DRAW)
	}
}

// draws tile from LAYER_LIGHT
func (g *Game) drawLightObstacleTile(x int, y int, screen *ebiten.Image) {
	lightTile := g.GetTileOnDrawingArea(x, y, LAYER_LIGHT)
	oldTile := g.GetTileOnDrawingArea(x, y, LAYER_DRAW)
	drawTile := g.CheckTileInStack(0, 0, LAYER_LIGHT)

	if oldTile == EMPTY && lightTile == EMPTY {
		g.StartRecording()
		g.UpdateTileUsage(drawTile, +1)
		g.SetTileOnCanvas(x, y, drawTile, LAYER_LIGHT)
		g.Recorder.AppendToCurrent(x, y, drawTile, EMPTY, LAYER_LIGHT, LAYER_LIGHT)
	}
}

// draws Tile from LAYER_ENTITY
func (g *Game) drawEntityTileOnCanvas(x int, y int, screen *ebiten.Image) {
	bgTile := g.GetTileOnDrawingArea(x, y, LAYER_DRAW)
	oldTile := g.GetTileOnDrawingArea(x, y, LAYER_ENTITY)
	drawTile := g.CurrentTileIndex()

	if bgTile != EMPTY && oldTile == EMPTY {

		// if tile is unique, it can be put at only one position on Canvas. If one of this tile is
		// already drawn, it must be erased and then put into new location chosen by user. To undraw
		// such Tile it must be first find.
		if g.TileIsUnique(drawTile) == true {

			// this tile is used nowhere - it can be just put into map
			if g.TileUsage(drawTile) == 0 {

				g.StartRecording()
				g.UpdateTileUsage(drawTile, +1)
				g.SetTileOnCanvas(x, y, drawTile, LAYER_ENTITY)
				g.Recorder.AppendToCurrent(x, y, drawTile, EMPTY, LAYER_ENTITY, LAYER_ENTITY)
				// unique tile is used - old location needs to be found and Tile must be erased from
				// there
			} else {
				old_x, old_y := g.FindTile(drawTile, LAYER_ENTITY)

				g.StartRecording()
				g.SetTileOnCanvas(x, y, drawTile, LAYER_ENTITY)
				g.SetTileOnCanvas(old_x, old_y, EMPTY, LAYER_ENTITY)

				g.Recorder.AppendToCurrent(x, y, drawTile, EMPTY, LAYER_ENTITY, LAYER_ENTITY)
				g.Recorder.AppendToCurrent(old_x, old_y, EMPTY, drawTile, LAYER_ENTITY, LAYER_ENTITY)

			}
		}
	}
}

// draws Tile on Canvas.
func (g *Game) DrawTileOnCanvas(screen *ebiten.Image) {
	x, y := g.MousePosToRowAndColOnCanvas(g.mouse_x, g.mouse_y)

	if g.mode == MODE_LIGHT {
		g.drawLightObstacleTile(x, y, screen)
	} else if g.mode == MODE_DRAW {
		g.drawNormalTileOnCanvas(x, y, screen)
		g.drawLightObstacleTile(lto.MinVal(x+1, CanvasCols), y, screen)
		g.drawLightObstacleTile(x, lto.MinVal(y+1, CanvasRows), screen)
		g.drawLightObstacleTile(lto.MaxVal(x-1, 0), y, screen)
		g.drawLightObstacleTile(x, lto.MaxVal(y-1, 0), screen)
	} else if g.mode == MODE_ENTITIES {
		g.drawEntityTileOnCanvas(x, y, screen)
	}

}

// ads new Tile to TileStack. Tile is taken from Pallete basing on its position (row an column).
// This is mainly use for hardcoded Tiles and not to be used by users
func (g *Game) AddTileFromPalleteToStack(row int, col int, i int, unique bool) {
	image := g.GetTileImage(row, col, g.Pallete.Tileseter.GetTilesetByIndex(i))
	g.AppendToStack(image, row, col, i, unique)
}

// taking some Tile from Pallete
func (g *Game) ChooseTileFromPallete(screen *ebiten.Image) {
	row, col := g.MousePosToTileRowAndColOnPallete(g.mouse_x, g.mouse_y)
	tilesetIndex := g.Tileseter.GetCurrentIndex()
	tileset := g.Pallete.Tileseter.GetCurrent()

	// chosen tile is already on stack, shorcut to it can be used
	if i := g.CheckTileInStack(row, col, tilesetIndex); i != -1 {
		g.SetCurrentTile(i)
	} else {
		// chosen tile is not on stack, additional attention is needed
		image := g.GetTileImage(row, col, tileset)
		g.AppendToStack(image, row, col, tilesetIndex, false)
	}
}

// draws current Tile while mouse is whether canvas rect.
// Tile is showing exact place where it will be drawn
func (g *Game) DrawHoveredTileOnCanvas(screen *ebiten.Image) {
	TileX, TileY := g.MousePosToTileHoveredOnCanvas(g.mouse_x, g.mouse_y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(TileX), float64(TileY))
	g.DrawCurrentTile(screen, op)
}

// draws current Tile
func (g *Game) drawCurrentTileToDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(CurrentTileToDrawX),
		float64(CurrentTileToDrawY),
	)
	g.DrawCurrentTile(screen, op)
}

// undraws given Record (single non-interrrupted mouse click)
func (g *Game) UndrawOneRecord(record Record) {
	// if record to undraw is empty there is nothing to undraw!
	if len(record.x_coords) == 0 {
		return
	}

	for i := len(record.x_coords) - 1; i > -1; i-- {
		g.SetTileOnCanvas(record.x_coords[i], record.y_coords[i], record.old_tiles[i], record.old_layers[i])
	}
}

// draws Cursor on Pallete
func (g *Game) DrawCursorOnPallete(screen *ebiten.Image) {
	cursorRect := g.MousePosToCursorRect(g.mouse_x, g.mouse_y, g.Cursor.GetCursorSize())
	g.Cursor.DrawCursor(screen, cursorRect)
}

// changes mode of drawing
func (g *Game) changeMode(mode int) {
	g.Tabber.ChangeTab(mode)
	g.mode = mode
	g.Pallete.Tileseter.SetCurrent(mode)
	g.Canvas.ChangeLayer(g.mode)
}

// change tab to draw
func (g *Game) changeModeToDraw(screen *ebiten.Image) {
	g.changeMode(MODE_DRAW)
}

// change tab to drawing light
func (g *Game) changeModeToDrawLight(screen *ebiten.Image) {
	g.changeMode(MODE_LIGHT)
}

// change tab to drawing entities
func (g *Game) changeModeToDrawEntities(screen *ebiten.Image) {
	g.changeMode(MODE_ENTITIES)
}
