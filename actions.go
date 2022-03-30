package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	lto "ltools/src/objects"
)

// draws single tile on canvas (check if current place is occupied is
// firstly done)
func (g *Game) drawTileOnCanvas(screen *ebiten.Image, x int, y int) {
	tileX, tileY := g.PosToTileCoordsOnCanvas(x, y)

	oldTile := g.GetTileOnCanvas(tileX, tileY)
	newTile := g.GetCurrentTile()
    g.StartRecording()

	// current place is empty
	if oldTile == -1 {
		newTile.NumberUsed++
		g.SetTileOnCanvas(tileX, tileY, g.TileStack.CurrentTile)
        g.Recorder.AppendToCurrent(tileX, tileY, g.TileStack.CurrentTile, -1)

		// current place is occupied by other tile (other than current tile to draw)
	} else if oldTile != g.TileStack.CurrentTile {
		g.UpdateTileUsage(oldTile, -1)
		g.ClearTileStack(oldTile)

		newTile.NumberUsed++
		g.SetTileOnCanvas(tileX, tileY, g.TileStack.CurrentTile)
        g.Recorder.AppendToCurrent(tileX, tileY, g.TileStack.CurrentTile, oldTile)
	}
}

// ads new tile to tile stack, allowing for easy acces to it, after clicking on it
// on pallete
func (g *Game) addTileFromPalleteToStack(x int, y int) {
	tileX, tileY := g.PosToTileCoordsOnPallete(x, y)
	tileY += g.Pallete.Viewport_y
	newTile := lto.NewTile(g.PosToSubImageOnPallete(x, y, &g.Tileset), tileX, tileY)
	g.AddTileToStack(newTile)
}

// everything that needs to be happen after clicking on tile on pallete
// check, whether clicked tile is already in the stack is done.
func (g *Game) chooseTileFromPallete(x int, y int) {
	tileX, tileY := g.PosToTileCoordsOnPallete(x, y)
	tileY += g.Pallete.Viewport_y

	// chosen tile is already on stack, shorcut to it can be used
	if tileIndex := g.CheckTileInStack(tileX, tileY); tileIndex != -1 {
		g.SetCurrentTile(tileIndex)
	} else {
		// chosen tile is not on stack, additional attention is needed
		g.addTileFromPalleteToStack(x, y)
	}
}

// draws current tile to draw while mouse is whether canvas rect.
// tile is showing precise place where it will be placed
func (g *Game) drawHoveredTileOnCanvas(screen *ebiten.Image, x int, y int) {
	TileX, TileY := g.PosToTileHoveredOnCanvas(x, y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(TileX), float64(TileY))
	g.DrawCurrentTile(screen, op)
}

// draws pallete object
func (g *Game) drawPallete(screen *ebiten.Image) {
	offset := g.Pallete.Viewport_y * PalleteColsN

	for n := 0; n < g.Tileset.Num; n++ {
		subImg := g.TileNrToSubImageOnTileset(n + offset)

		op := &ebiten.DrawImageOptions{}
		x, y := g.TileNrToCoordsOnPallete(n)
		op.GeoM.Translate(x, y)
		screen.DrawImage(subImg, op)

		if n > g.Pallete.Rows*g.Pallete.Cols-2 {
			break
		}
	}
	g.DrawPalleteScroller(screen)
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

    for i:=0; i<len(record.x_coords); i++ {
        g.SetTileOnCanvas(record.x_coords[i], record.y_coords[i], record.old_tiles[i])
    }
}

