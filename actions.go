package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	lto "ltools/src/objects"
)

func (g *Game) drawTileOnCanvas(screen *ebiten.Image, x int, y int) {
	tileX, tileY := g.PosToTileCoordsOnCanvas(x, y)

	oldTile := g.GetTileOnCanvas(tileX, tileY)
	newTile := g.GetCurrentTile()

	if oldTile == -1 {
		newTile.NumberUsed++
		g.SetTileOnCanvas(tileX, tileY, g.TileStack.CurrentTile)

	} else if oldTile != g.TileStack.CurrentTile {
		g.UpdateTileUsage(oldTile, -1)
		g.ClearTileStack(oldTile)

		newTile.NumberUsed++
		g.SetTileOnCanvas(tileX, tileY, g.TileStack.CurrentTile)
	}
}

func (g *Game) addTileFromPalleteToStack(x int, y int) {
	tileX, tileY := g.PosToTileCoordsOnPallete(x, y)
	newTile := lto.NewTile(g.PosToSubImageOnPallete(x, y, &g.Tileset), tileX, tileY)
	g.AddTileToStack(newTile)
}

func (g *Game) chooseTileFromPallete(x int, y int) {
	tileX, tileY := g.PosToTileCoordsOnPallete(x, y)

	if tileIndex := g.CheckTileInStack(tileX, tileY); tileIndex != -1 {
		g.SetCurrentTile(tileIndex)
	} else {
		g.addTileFromPalleteToStack(x, y)
	}
}

func (g *Game) drawHoveredTileOnCanvas(screen *ebiten.Image, x int, y int) {
	TileX, TileY := g.PosToTileHoveredOnCanvas(x, y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(TileX), float64(TileY))
	g.DrawCurrentTile(screen, op)
}

func (g *Game) drawPallete(screen *ebiten.Image) {
	for n := 0; n < g.Tileset.Num; n++ {
		subImg := g.TileNrToSubImageOnTileset(n)

		op := &ebiten.DrawImageOptions{}
		x, y := g.TileNrToCoordsOnPallete(n)
		op.GeoM.Translate(x, y)
		screen.DrawImage(subImg, op)

		if n > g.Pallete.MaxTile-2 {
			break
		}
	}
}

func (g *Game) drawCurrentTileToDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(CurrentTileToDrawX),
		float64(CurrentTileToDrawY),
	)
	g.DrawCurrentTile(screen, op)
}
