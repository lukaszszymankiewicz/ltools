package const

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
	PalleteEndX     = PalleteX + (TileWidth * PalleteColsN)
	PalleteEndY     = PalleteY + (TileHeight * PalleteRowsN)
	PalleteTilesMax = PalleteColsN * PalleteRowsN
	// current tile to draw
	CurrentTileToDrawX    = PalleteEndX / 2
	CurrentTileToDrawY    = PalleteEndY + TileHeight
	CurrentTileToDrawEndX = CurrentTileToDrawX + TileWidth
	CurrentTileToDrawEndY = CurrentTileToDrawY + TileHeight
	// canvas
	ViewportColsN = 10
	ViewportRowsN = 10
	CanvasX     = PalleteEndX + (TileWidth * 2)
	CanvasY     = PalleteY
	CanvasEndX  = CanvasX + (TileWidth * CanvasColsN)
	CanvasEndY  = CanvasY + (TileWidth * CanvasRowsN)
	// cursor
	CursorSize = 32
)
