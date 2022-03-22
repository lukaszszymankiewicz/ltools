package main

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
	ViewportCols = 20
	ViewportRows = 30
	CanvasRows   = 100
	CanvasCols   = 100
	Canvas_y     = PalleteY
	Canvas_x     = PalleteEndX + (TileWidth * 2)
	// scroll arrows
	ArrowRightX = Canvas_x + (ViewportRows-1)*TileWidth + 1
	ArrowRightY = Canvas_y + ViewportCols*TileHeight + 2
	ArrowLeftX  = Canvas_x
	ArrowLeftY  = Canvas_y + ViewportCols*TileHeight + 2
	ArrowUpX    = Canvas_x + (ViewportRows)*TileWidth + 2
	ArrowUpY    = Canvas_y - 1
	ArrowDownX  = Canvas_x + ViewportRows*TileWidth + 2
	ArrowDownY  = Canvas_y + (ViewportCols-1)*TileHeight + 1
)
