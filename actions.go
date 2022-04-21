package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
    lto "ltools/src/objects"
)

// by now, only thing that can be drawn in the LIGHT_LAYER is an obstacle (obstacle blocks player
// both with light). Tile can be either a normal tile OR an obstacle. Drawing normals tiles has
// higher priority so, obstacles are erased if user click on postion where obstacle lays. Below
// logic is kind ugly and should be generalized ad nearest refactor
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

func (g *Game) drawLightObstacleTile(x int, y int, screen *ebiten.Image) {
    lightTile := g.GetTileOnDrawingArea(x, y, LAYER_LIGHT)
    oldTile := g.GetTileOnDrawingArea(x, y, LAYER_DRAW)

    drawTile := g.CheckTileInStack(0, 0, LAYER_LIGHT)

    if oldTile == EMPTY && lightTile == EMPTY{
        g.StartRecording()
        g.UpdateTileUsage(drawTile, +1)
        g.SetTileOnCanvas(x, y, drawTile, LAYER_LIGHT)
        g.Recorder.AppendToCurrent(x, y, drawTile, EMPTY, LAYER_LIGHT, LAYER_LIGHT)
    }
}

// this one is a little bit tricky, as entities just mimics the Tileset
func (g *Game) drawEntityTileOnCanvas(x int, y int, screen *ebiten.Image) {
    oldTile := g.GetTileOnDrawingArea(x, y, LAYER_DRAW)
    oldEntity := g.GetTileOnDrawingArea(x, y, LAYER_ENTITY)


    drawTile := g.CurrentTileIndex()

    if oldEntity == EMPTY && oldTile != EMPTY && e.EntityCanBeDrawn() {
        g.StartRecording()
        g.UpdateTileUsage(drawTile , +1)
        g.SetTileOnCanvas(x, y, drawTile , LAYER_ENTITY)
        g.Recorder.AppendToCurrent(x, y, drawTile, EMPTY, LAYER_ENTITY, LAYER_ENTITY)
    }
}

// draws any type of Tile on Canvas. Normal type tile spreads light obstacle around 
func (g *Game) drawTileOnCanvas(screen *ebiten.Image) {
	x, y := g.PosToTileCoordsOnCanvas(g.mouse_x, g.mouse_y)
    
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

        row := n % PalleteColsN        
        col := (n / PalleteColsN) + offset
        image := PosToSubImageOnPallete(row, col, tileset)

		op.GeoM.Translate(row*TileWidth, col*TileHeight)
		screen.DrawImage(image, op)

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

// undraws given record (single non-interrrupted mouse click)
func (g *Game) UndrawOneRecord(record Record) {
	// if record to undraw is empty there is nothing to undraw!
	if len(record.x_coords) == 0 {
		return
	}

    for i := len(record.x_coords)-1; i>-1; i-- {
		g.SetTileOnCanvas(record.x_coords[i], record.y_coords[i], record.old_tiles[i], record.old_layers[i])
	}
}

// draw cursor on Pallete
func (g *Game) drawCursorOnPallete(screen *ebiten.Image) {
	x, y := g.Pallete.PosToCursorCoords(g.mouse_x, g.mouse_y)
	g.Cursor.DrawCursor(screen, x, y)
}

// global function for changing mode of editor
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

// workaround for Tileset/Pallete to accept list of Entities as icons rather than ready-made image
func (g *Game) createTilesetFromEntityStack() {
    n_entities := len(g.EntityStack.Stack)
    
    width := lto.MinVal(PalleteColsN, n_entities) * TileWidth
    height := ((n_entities/6)+1) * TileHeight
    tileset_image := ebiten.NewImage(width, height)

    for i:=0; i<n_entities; i++ {
        x := (i%PalleteColsN) * TileWidth
        y := (i/PalleteColsN) * TileHeight
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(x), float64(y))

        tileset_image.DrawImage(g.EntityStack.Stack[i].Image, op)
    }

    tileset := lto.NewTilesetFromImage(tileset_image)
    g.AppendToTilesetter(tileset)
}

func (g *Game) fillBasicTiles () {
    g.AppendFromTilesetToStack(0, 0, LAYER_LIGHT)
    g.AppendFromTilesetToStack(0, 0, LAYER_ENTITY)
}

func (g *Game) AppendFromTilesetToStack (row int, col int, tileset_index int) int {
    image := g.PalletePosToSubImageOnPallete(0, 0, g.Tileseter.GetTilesetByIndex(tileset_index))
    g.AppendToStack(image, row, col, tileset_index)

    return len(g.TileStack) - 1
}

