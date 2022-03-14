package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	lto "ltools/src/objects"
)

const (
	// screen
	ScreenWidth  = 1366
	ScreenHeight = 768
	// tile
	TileWidth  = 32
	TileHeight = 32
	// pallete
	PalleteX        = 20
	PalleteY        = 20
	PalleteColsN    = 6
	PalleteRowsN    = 10
	PalleteMaxTile  = 40
	PalleteEndX     = PalleteX + (TileWidth * PalleteColsN)
	PalleteEndY     = PalleteY + (TileHeight * PalleteRowsN)
	PalleteTilesMax = PalleteColsN * PalleteRowsN
	// current tile to draw
	CurrentTileToDrawX = PalleteEndX / 2
	CurrentTileToDrawY = PalleteEndY + TileHeight
	// cursor
	CursorSize = 32
	// pallete
	ViewportCols  = 10
	ViewportRows  = 10
	CanvasRows    = 100
	CanvasCols    = 100
	Canvas_y      = PalleteY
	Canvas_x      = PalleteEndX + (TileWidth * 2)
	Canvas_width  = (TileWidth * 10)
	Canvas_height = (TileWidth * 10)
)

type Game struct {
	lto.Pallete
	lto.Canvas
	lto.Tileset
	lto.TileStack
	//
	currentTileToDrawRect image.Rectangle
	tileStackLen          int
}

func init() {
	ebiten.SetFullscreen(true)
}

func coordsInRect(x int, y int, rect image.Rectangle) bool {
	if (x >= rect.Min.X && x <= rect.Max.X) && (y >= rect.Min.Y && y <= rect.Max.Y) {
		return true
	}
	return false
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return 1388, 768
}

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

func (g *Game) handleMouseEvents(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()

	if coordsInRect(x, y, g.Canvas.Rect) {
		g.drawHoveredTileOnCanvas(screen, x, y)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.drawTileOnCanvas(screen, x, y)
		}
	}

	if coordsInRect(x, y, g.Pallete.Rect) {
		g.DrawCursorOnPallete(screen, x, y)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.chooseTileFromPallete(x, y)
		}
	}
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

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawPallete(screen)
	g.DrawCanvas(screen, g.GetAllTiles())
	g.handleMouseEvents(screen)
	g.drawCurrentTileToDraw(screen)
}

func NewGame() *Game {
	var g Game

	g.Pallete = lto.NewPallete(
		PalleteX,
		PalleteY,
		PalleteEndX,
		PalleteEndY,
		PalleteRowsN,
		PalleteColsN,
		PalleteMaxTile,
	)

	g.Canvas = lto.NewCanvas(
		ViewportRows,
		ViewportCols,
		CanvasRows,
		CanvasCols,
		Canvas_x,
		Canvas_y,
		Canvas_width,
		Canvas_height,
	)

	g.Tileset = lto.NewTileset("tileset_1.png")

	// post init
	g.addTileFromPalleteToStack(0, 0)
	g.SetCurrentTile(0)
	g.Canvas.ClearDrawingArea()

	return &g
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("LTOOLS - LIGHTER DEVELOPMENT TOOLS")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
