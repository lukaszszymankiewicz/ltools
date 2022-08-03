package editor

import ()

const (
	UNIQUE_COND_ALWAYS = iota
	UNIQUE_COND_IS_NOT_SET
)

type DrawConditionFunc func(TilesCollection, int, int) bool
type UniqueDrawConditionFunc func(TilesCollection, int, *Tile) bool
type DrawTileEffect func(TilesCollection, int, int, *Tile)

func IsEmptyCondition(tc TilesCollection, idx int, layer int) bool {
	return tc.Empty(idx)
}

func HasDrawTileCondition(tc TilesCollection, idx int, layer int) bool {
	if tc.Get(idx, LAYER_DRAW) != nil {
		return true
	}
	return false
}

func HasLightTileCondition(tc TilesCollection, idx int, layer int) bool {
	if tc.Get(idx, LAYER_LIGHT) != nil {
		return true
	}
	return false
}

func HasEntityTileCondition(tc TilesCollection, idx int, layer int) bool {
	if tc.Get(idx, LAYER_ENTITY) != nil {
		return true
	}
	return false
}

func AlwaysCondition(tc TilesCollection, idx int, layer int) bool {
	return true
}

func UniqueAlwaysCondition(tc TilesCollection, idx int, tile *Tile) bool {
	return true
}

func IsntUsed(tc TilesCollection, idx int, tile *Tile) bool {
	if old_idx := tc.HasTile(tile); old_idx != -1 {
		tc.Set(old_idx, tile.layer, nil)
	}
	return true
}

var SetTileUniqueCondition []UniqueDrawConditionFunc = []UniqueDrawConditionFunc{
	UniqueAlwaysCondition,
	IsntUsed,
}

const (
	COND_ALWAYS = iota
	COND_IS_EMPTY
	COND_HAS_DRAW_TILE
	COND_HAS_LIGHT_TILE
	COND_HAS_ENTITY_TILE
)

var SetTileByLayerCondition []DrawConditionFunc = []DrawConditionFunc{
	AlwaysCondition,
	IsEmptyCondition,
	HasDrawTileCondition,
	HasLightTileCondition,
	HasEntityTileCondition,
}

func SetTile(tc TilesCollection, idx int, layer int, tile *Tile) {
	tc.Set(idx, layer, tile)
}

func CleanTile(tc TilesCollection, idx int, layer int, tile *Tile) {
	tc.Set(idx, layer, nil)
}

func DoNothing(tc TilesCollection, idx int, layer int, tile *Tile) {
}

var SetTileByLayerEffect []DrawTileEffect = []DrawTileEffect{
	DoNothing, SetTile, CleanTile,
}

func PutTile(v Viewport, tc TilesCollection, tile *Tile, brush *Brush, row int, col int, layer int) {
	tileIdx := v.GetTileIdx(row, col)

	uniqueCond := SetTileUniqueCondition[tile.unique_condition]
	if uniqueCond(tc, tileIdx, tile) == false {
		return
	}

	// then main tile position is checked whether is allowed to be set
	brushFunc := brush.byLayerCondition[layer]
	tileIsAllowedCond := SetTileByLayerCondition[brushFunc]

	if tileIsAllowedCond(tc, tileIdx, layer) == false {
		return
	}

	// if tile can be put on desired location - brush effect is calculated
	brush_result := brush.ApplyBrush(row, col, row, col, tile)

	// each brush effect tile is then checked if can be set
	for i := 0; i < brush_result.Len(); i++ {
		corr_x := brush_result.coords[i][0]
		corr_y := brush_result.coords[i][1]
		layer := brush_result.layers[i]
		tile := brush_result.tiles[i]

		brushFunc := brush.byLayerCondition[layer]
		tileIsAllowedCond := SetTileByLayerCondition[brushFunc]
		tileIdx := v.GetTileIdx(row-corr_x, col-corr_y)

		// if tile is not legal to be put - abort
		if tileIsAllowedCond(tc, tileIdx, layer) == true {

			// if Tile is allowed to be placed here apply effect of a brush to each layer
			for j := 0; j < tc.n_layers; j++ {

				drawFuncIdx := brush.byLayerEffect[layer][j]
				SetTileFunc := SetTileByLayerEffect[drawFuncIdx]
				SetTileFunc(tc, tileIdx, j, tile)
			}
		}
	}
}

var BasicLayersAlpha = [][]float64{
	{1.0, 0.5, 0.0}, // LAYER_DRAW is on
	{0.5, 1.0, 0.5}, // LAYER_LIGHT in on
	{0.5, 0.5, 1.0}, // LAYER_ENTITY in on
}

var ExclusiveLayersAlpha = [][]float64{
	{1.0, 0.0, 0.0}, // LAYER_DRAW is on
	{0.0, 1.0, 0.0}, // LAYER_LIGHT in on
	{0.0, 0.0, 1.0}, // LAYER_ENTITY in on
}
