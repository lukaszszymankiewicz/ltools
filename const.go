package main

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
	PalleteColsN          = 6
	PalleteRowsN          = 10
	ToolboxX              = PalleteX
	ToolboxY              = (PalleteY + (PalleteRowsN * GridSize)) + GridSize
	Canvas_y              = PalleteY
	Canvas_x              = (PalleteColsN * GridSize) + (GridSize * 3)
	CursorSize            = 32
	CanvasColsN           = 30
	CanvasRowsN           = 20
	TabberX               = Canvas_x
	TabberY               = PalleteY - GridSize - 1
	TILE_PER_ROW_ON_IMAGE = 16
	EXPORT_BUTTON         = 3
	LOGGER_PATH           = "./logs/"
)
