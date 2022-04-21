package main

const (
    MODE_DRAW = iota
    MODE_LIGHT
    MODE_ENTITIES
    MODE_ALL
)

const (
    LAYER_DRAW = iota
    LAYER_LIGHT
    LAYER_ENTITY
)

const (
	// screen
	ScreenWidth  = 1366
	ScreenHeight = 768
	// tile
	TileWidth  = 32
	TileHeight = 32
	// pallete
	PalleteX          = 50
	PalleteY          = 80
	PalleteColsN      = 6
	PalleteRowsN      = 10
	PalleteMaxTile    = PalleteColsN * PalleteRowsN
	PalleteEndX       = PalleteX + (TileWidth * PalleteColsN)
	PalleteEndY       = PalleteY + (TileHeight * PalleteRowsN)
	PalleteArrowUpX   = PalleteEndX
	PalleteArrowUpY   = PalleteY
	PalleteArrowDownX = PalleteEndX
	PalleteArrowDownY = PalleteEndY
	// current tile to draw
	CurrentTileToDrawX = PalleteEndX / 2
	CurrentTileToDrawY = PalleteEndY + TileHeight
	// cursor
	CursorSize = 32
	// pallete
	ViewportCols = 30
	ViewportRows = 20
	CanvasCols   = 50
	CanvasRows   = 40
	Canvas_y     = PalleteY
	Canvas_x     = PalleteEndX + (TileWidth * 2)
	// scroll arrows
	ArrowRightX = Canvas_x + (ViewportCols-1)*TileWidth + 1
	ArrowRightY = Canvas_y + ViewportRows*TileHeight + 2
	ArrowLeftX  = Canvas_x
	ArrowLeftY  = Canvas_y + ViewportRows*TileHeight + 2
	ArrowUpX    = Canvas_x + (ViewportCols)*TileWidth + 2
	ArrowUpY    = Canvas_y - 1
	ArrowDownX  = Canvas_x + ViewportCols*TileWidth + 2
	ArrowDownY  = Canvas_y + (ViewportRows-1)*TileHeight + 1
    // tabber
	TabberX     = Canvas_x
	TabberY     = PalleteY - TileHeight - 1
    // looks great
    EMPTY       = -1
)
