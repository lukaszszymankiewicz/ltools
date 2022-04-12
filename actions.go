package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

// draws single tile on canvas (check if current place is occupied is
// firstly done)
func (g *Game) drawTileOnCanvas(screen *ebiten.Image) {
	x, y := g.PosToTileCoordsOnCanvas(g.mouse_x, g.mouse_y)
    
    oldTile := g.GetTileOnDrawingArea(x, y)
    drawTile := g.CurrentTileIndex()

    if g.mode == MODE_LIGHT { 
        if oldTile == -1 {
            g.StartRecording()
            g.UpdateTileUsage(drawTile, +1)
            g.SetTileOnCanvas(x, y, drawTile)
            g.Recorder.AppendToCurrent(x, y, drawTile, -1)
        }
    } else if g.mode == MODE_DRAW {
        g.StartRecording()

        // current place is empty
        if oldTile == -1 {
            g.UpdateTileUsage(drawTile, +1)
            g.Recorder.AppendToCurrent(x, y, drawTile, -1)

        // current place is occupied by other tile (other than current tile to draw)
        } else if oldTile != drawTile {
            g.UpdateTileUsage(drawTile, +1)
            g.UpdateTileUsage(oldTile, -1)
            g.Recorder.AppendToCurrent(x, y, drawTile, oldTile)
        }
        g.SetTileOnCanvas(x, y, drawTile)
    }

}

// ads new tile to tile stack, allowing for easy acces to it, after clicking on it
// on pallete
func (g *Game) addTileFromPalleteToStack(row int, col int, tileset int) {
    image := g.PosToSubImageOnPallete(g.mouse_x, g.mouse_y, g.Tileseter.GetCurrent())
	g.AppendToStack(image, row, col, tileset)
}

// everything that needs to be happen after clicking on tile on pallete
// check, whether clicked tile is already in the stack is done.
func (g *Game) chooseTileFromPallete(screen *ebiten.Image) {
	row, col := g.PosToTileCoordsOnPallete(g.mouse_x, g.mouse_y)
	col += g.Pallete.Viewport_y
    tilesetIndex := g.Tileseter.GetCurrentIndex()

	// chosen tile is already on stack, shorcut to it can be used
	if i := g.CheckTileInStack(row, col, tilesetIndex); i != -1 {
		g.SetCurrentTile(i)
	} else {
		// chosen tile is not on stack, additional attention is needed
		g.addTileFromPalleteToStack(row, col, tilesetIndex)
	}
}

// draws current tile to draw while mouse is whether canvas rect.
// tile is showing precise place where it will be placed
func (g *Game) drawHoveredTileOnCanvas(screen *ebiten.Image) {
	TileX, TileY := g.PosToTileHoveredOnCanvas(g.mouse_x, g.mouse_y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(TileX), float64(TileY))
	g.DrawCurrentTile(screen, op)
}

// draws pallete object
func (g *Game) drawPallete(screen *ebiten.Image) {
	offset := g.Pallete.Viewport_y * PalleteColsN
    tileset := g.Tileseter.GetCurrent()  

	for n := 0; n < tileset.Num; n++ {
		subImg := tileset.TileNrToSubImageOnTileset(n + offset)

		op := &ebiten.DrawImageOptions{}
		x, y := g.TileNrToCoordsOnPallete(n)
		op.GeoM.Translate(x, y)
		screen.DrawImage(subImg, op)

		if n > g.Pallete.Rows*g.Pallete.Cols-2 {
			break
		}
	}
	g.Pallete.Scroller_y.Draw(screen)
}

// draws current tile (brush type)
func (g *Game) drawCurrentTileToDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(CurrentTileToDrawX),
		float64(CurrentTileToDrawY),
	)
	g.DrawCurrentTile(screen, op)
}

// undraws given record
func (g *Game) UndrawOneRecord(record Record) {
	// if record to undraw is empty there is nothing to undraw!
	if len(record.x_coords) == 0 {
		return
	}

	for i := 0; i < len(record.x_coords); i++ {
		g.SetTileOnCanvas(record.x_coords[i], record.y_coords[i], record.old_tiles[i])
	}
}

// draw cursor on Pallete
func (g *Game) drawCursorOnPallete(screen *ebiten.Image) {
	x, y := g.Pallete.PosToCursorCoords(g.mouse_x, g.mouse_y)
	g.Cursor.DrawCursor(screen, x, y)
}

func (g *Game) changeMode(mode int) {
    g.Tabber.ChangeTab(mode)
    g.mode = mode
    g.Tileseter.SetCurrent(mode)
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
