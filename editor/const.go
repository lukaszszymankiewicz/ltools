package editor

const (
	MODE_DRAW = iota
	MODE_LIGHT
	MODE_ENTITIES
	MODE_ALL
)

const (
	PENCIL_TOOL = iota
	RUBBER_TOOL
)

const (
	LAYER_DRAW = iota
	LAYER_LIGHT
	LAYER_ENTITY
	LAYER_N
)

const (
	ScreenWidth           = 1366
	ScreenHeight          = 768
	GridSize              = 32

	PalleteX              = 50
	PalleteY              = 80
	PalleteCols           = 6
	PalleteRows           = 10
	PalleteViewportCols   = 6
	PalleteViewportRows   = 10

	CanvasY              = PalleteY
	CanvasX              = (PalleteCols * GridSize) + (GridSize * 3)
	CanvasCols            = 50
	CanvasRows            = 60
	CanvasViewportCols    = 30
	CanvasViewportRows    = 20

	ToolboxX              = PalleteX
	ToolboxY              = (PalleteY + (PalleteRows * GridSize)) + GridSize
	CursorSize            = 32
	TabberX               = CanvasX
	TabberY               = PalleteY - GridSize - 1
	TILE_PER_ROW_ON_IMAGE = 16
	EXPORT_BUTTON         = 3
	LOGGER_PATH           = "./logs/"
)
