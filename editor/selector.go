package editor

import ()

func CalculateHighlightPos(
	mouse_x int,
	mouse_y int,
	grid_x int,
	grid_y int,
	grid_size int,
) (int, int, int, int) {

	pos_x := int((mouse_x-grid_x)/grid_size)*grid_size + grid_x
	pos_y := int((mouse_y-grid_y)/grid_size)*grid_size + grid_y

	return pos_x, pos_y, pos_x + grid_size, pos_y + grid_size
}

func CalculateTilePos(
	mouse_x int,
	mouse_y int,
	grid_x int,
	grid_y int,
	grid_size int,
) (int, int) {

	row := int((mouse_x - grid_x) / grid_size)
	col := int((mouse_y - grid_y) / grid_size)

	return row, col
}
