package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestDrawingLogicPuttingEntityLayersWithPen(t *testing.T) {
	v := NewViewport(10, 10, 10, 10)
	tc := NewTilesCollection(3)
	tc.Alloc(100)

	dummy_image := ebiten.NewImage(32, 32)
	normal_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_DRAW)
	entity_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_ENTITY)

	// make entity_tile unique on canvas
	entity_tile.unique_condition = UNIQUE_COND_IS_NOT_SET

	// few normal tiles on which drawing entities will be perfomed
	PutTile(v, tc, &normal_tile, &PenBrush, 0, 0, LAYER_DRAW)
	PutTile(v, tc, &normal_tile, &PenBrush, 1, 0, LAYER_DRAW)
	PutTile(v, tc, &normal_tile, &PenBrush, 2, 0, LAYER_DRAW)

	// sanity checks
	for i := 0; i < LAYER_N; i++ {
		tileIdx := v.GetTileIdx(i, 0)
		draw_tile := tc.Get(tileIdx, LAYER_DRAW)

		if draw_tile == nil {
			t.Errorf("Putting normal tile on draw layer failed")
		}
	}

	// draw entity tile on one of the background tile
	PutTile(v, tc, &entity_tile, &PenBrush, 0, 0, LAYER_ENTITY)

	// sanity checks
	tileIdx := v.GetTileIdx(0, 0)
	entity := tc.Get(tileIdx, LAYER_ENTITY)
	if entity == nil {
		t.Errorf("Putting entity tile on entity layer failed")
	}

	draw_tile := tc.Get(tileIdx, LAYER_DRAW)
	if draw_tile == nil {
		t.Errorf("Putting normal tile on draw layer failed")
	}

	// draw another entity - old one should disappear
	PutTile(v, tc, &entity_tile, &PenBrush, 1, 0, LAYER_ENTITY)

	// sanity checks
	tileIdx = v.GetTileIdx(0, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity != nil {
		t.Errorf("Deleting old entity tile on entity layer failed")
	}

	tileIdx = v.GetTileIdx(1, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity == nil {
		t.Errorf("Putting new entity tile on entity layer failed")
	}

	// draw another one entity - two old ones should disappear
	PutTile(v, tc, &entity_tile, &PenBrush, 2, 0, LAYER_ENTITY)

	// sanity checks
	tileIdx = v.GetTileIdx(0, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity != nil {
		t.Errorf("Deleting old entity tile on entity layer failed")
	}

	tileIdx = v.GetTileIdx(1, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity != nil {
		t.Errorf("Deleting new entity tile on entity layer failed")
	}

	tileIdx = v.GetTileIdx(2, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity == nil {
		t.Errorf("Putting new entity tile on entity layer failed")
	}
}

func TestDrawingLogicPuttingEntityLayersWithPenErasingWithRubber(t *testing.T) {
	v := NewViewport(10, 10, 10, 10)
	tc := NewTilesCollection(3)
	tc.Alloc(100)

	dummy_image := ebiten.NewImage(32, 32)
	normal_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_DRAW)
	entity_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_ENTITY)

	// make entity_tile unique on canvas
	entity_tile.unique_condition = UNIQUE_COND_IS_NOT_SET

	// few normal tiles on which drawing entities will be perfomed
	PutTile(v, tc, &normal_tile, &PenBrush, 0, 0, LAYER_DRAW)
	PutTile(v, tc, &normal_tile, &PenBrush, 1, 0, LAYER_DRAW)
	PutTile(v, tc, &normal_tile, &PenBrush, 2, 0, LAYER_DRAW)

	// sanity checks
	for i := 0; i < LAYER_N; i++ {
		tileIdx := v.GetTileIdx(i, 0)
		draw_tile := tc.Get(tileIdx, LAYER_DRAW)

		if draw_tile == nil {
			t.Errorf("Putting normal tile on draw layer failed")
		}
	}

	// draw entity tile on one of the background tile
	PutTile(v, tc, &entity_tile, &PenBrush, 0, 0, LAYER_ENTITY)
	// sanity checks
	tileIdx := v.GetTileIdx(0, 0)
	entity := tc.Get(tileIdx, LAYER_ENTITY)
	if entity == nil {
		t.Errorf("Putting entity tile on entity layer failed")
	}

	// rubber entity tile
	PutTile(v, tc, &entity_tile, &RubberBrush, 0, 0, LAYER_ENTITY)
	tileIdx = v.GetTileIdx(0, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity != nil {
		t.Errorf("Erasing entity tile on entity layer failed")
	}

	// draw entity tile on another one pos
	PutTile(v, tc, &entity_tile, &PenBrush, 1, 0, LAYER_ENTITY)

	// sanity checks
	tileIdx = v.GetTileIdx(1, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity == nil {
		t.Errorf("Putting entity tile on entity layer failed")
	}

	// rubber new entity tile
	PutTile(v, tc, &entity_tile, &RubberBrush, 1, 0, LAYER_ENTITY)
	tileIdx = v.GetTileIdx(1, 0)
	entity = tc.Get(tileIdx, LAYER_ENTITY)
	if entity != nil {
		t.Errorf("Erasing entity tile on entity layer failed")
	}
}

func TestDrawingLogicPuttingEntityLayersWithPenErasingWithRubberMultiple(t *testing.T) {
	v := NewViewport(10, 10, 10, 10)
	tc := NewTilesCollection(3)
	tc.Alloc(100)

	dummy_image := ebiten.NewImage(32, 32)
	normal_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_DRAW)
	entity_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_ENTITY)

	// make entity_tile unique on canvas
	entity_tile.unique_condition = UNIQUE_COND_IS_NOT_SET

	// few normal tiles on which drawing entities will be perfomed
	PutTile(v, tc, &normal_tile, &PenBrush, 0, 0, LAYER_DRAW)
	PutTile(v, tc, &normal_tile, &PenBrush, 1, 0, LAYER_DRAW)
	PutTile(v, tc, &normal_tile, &PenBrush, 2, 0, LAYER_DRAW)

	// sanity checks
	for i := 0; i < LAYER_N; i++ {
		tileIdx := v.GetTileIdx(i, 0)
		draw_tile := tc.Get(tileIdx, LAYER_DRAW)

		if draw_tile == nil {
			t.Errorf("Putting normal tile on draw layer failed")
		}
	}

	// draw entity tile on one of the background tile
	PutTile(v, tc, &entity_tile, &PenBrush, 0, 0, LAYER_ENTITY)
	PutTile(v, tc, &entity_tile, &PenBrush, 1, 0, LAYER_ENTITY)
	PutTile(v, tc, &entity_tile, &PenBrush, 2, 0, LAYER_ENTITY)

	// rubber entity tile
	PutTile(v, tc, &entity_tile, &RubberBrush, 2, 0, LAYER_ENTITY)

	// entity should be deleted from every position
	for i := 0; i < 3; i++ {
		tileIdx := v.GetTileIdx(i, 0)
		entity := tc.Get(tileIdx, LAYER_ENTITY)

		if entity != nil {
			t.Errorf("Erasing entity tile on entity layer failed")
		}
	}
}

func TestDrawingLogicPutLightTileOnNormalTile(t *testing.T) {
	v := NewViewport(10, 10, 10, 10)
	tc := NewTilesCollection(3)
	tc.Alloc(100)

	dummy_image := ebiten.NewImage(32, 32)
	normal_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_DRAW)
	light_tile := NewTile(dummy_image, "dummy_tileset", 0, 0, LAYER_LIGHT)

	// few normal tiles on which drawing entities will be perfomed
	PutTile(v, tc, &normal_tile, &PenBrush, 0, 0, LAYER_DRAW)

	// draw light tile on one of the background tile
	PutTile(v, tc, &light_tile, &PenBrush, 0, 0, LAYER_LIGHT)

	// light tile should not be drawn
	tileIdx := v.GetTileIdx(0, 0)
	res := tc.Get(tileIdx, LAYER_LIGHT)

	if res != nil {
		t.Errorf("Putting light tile on occured tile should failed, but did not!")
	}
}
