package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestNewRecord(t *testing.T) {
	dummy_image := ebiten.NewImage(32, 32)
	t1 := NewTile(dummy_image, "name", false, 0, 0, 0)

	r := NewRecord(0, 0, &t1, nil, 0)

	if len(r.x_coords) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(r.y_coords) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(r.tiles) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(r.old_tiles) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(r.layers) != 1 {
		t.Errorf("Adding new recond failed")
	}
}

func TestAddNew(t *testing.T) {
	rc := Recorder{}

	dummy_image := ebiten.NewImage(32, 32)
	t1 := NewTile(dummy_image, "name", false, 0, 0, 0)
	t2 := NewTile(dummy_image, "name", false, 0, 1, 0)

	rc.AddNew(10, 10, &t1, nil, 0)

	if len(rc.records) != 1 {
		t.Errorf("Adding new recond to recorder failed")
	}

	rc.AddNew(10, 10, &t2, nil, 0)

	if len(rc.records) != 2 {
		t.Errorf("Adding new recond to recorder failed")
	}

}

func TestAppendToCurrent(t *testing.T) {
	rc := Recorder{}

	dummy_image := ebiten.NewImage(32, 32)

	t1 := NewTile(dummy_image, "name", false, 0, 0, 0)
	t2 := NewTile(dummy_image, "name", false, 0, 1, 0)

	rc.AppendToCurrent(10, 10, &t1, nil, 0)

	if len(rc.records) != 1 {
		t.Errorf("Adding new recond to recorder failed")
	}

	if len(rc.records[0].x_coords) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].y_coords) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].tiles) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].old_tiles) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].layers) != 1 {
		t.Errorf("Adding new recond failed")
	}

	rc.AppendToCurrent(5, 5, &t2, nil, 0)

	// len of records is still 1 as we append the existing one!
	if len(rc.records) != 1 {
		t.Errorf("Adding new recond to recorder failed")
	}

	if len(rc.records[0].x_coords) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].y_coords) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].tiles) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].old_tiles) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].layers) != 2 {
		t.Errorf("Adding new recond failed")
	}

	rc.AddNew(15, 15, &t2, nil, 0)

	// new record added, so overall number of records should rise
	if len(rc.records) != 2 {
		t.Errorf("Adding new recond to recorder failed")
	}

	// old record should not be changed
	if len(rc.records[0].x_coords) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].y_coords) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].tiles) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].old_tiles) != 2 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[0].layers) != 2 {
		t.Errorf("Adding new recond failed")
	}

	// but new record should be updated
	if len(rc.records[1].x_coords) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[1].y_coords) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[1].tiles) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[1].old_tiles) != 1 {
		t.Errorf("Adding new recond failed")
	}

	if len(rc.records[1].layers) != 1 {
		t.Errorf("Adding new recond failed")
	}
}

func TestUndoOneRecord(t *testing.T) {
	rc := Recorder{}

	dummy_image := ebiten.NewImage(32, 32)
	t1 := NewTile(dummy_image, "name", false, 0, 0, 0)
	t2 := NewTile(dummy_image, "name", false, 0, 1, 0)

	// first (0) record
	rc.AddNew(15, 15, &t1, nil, 0)
	rc.AppendToCurrent(5, 5, &t2, nil, 0)
	rc.AppendToCurrent(5, 5, &t2, nil, 0)
	rc.SaveRecord()

	// second (1) record
	rc.AddNew(15, 15, &t1, nil, 0)
	rc.SaveRecord()

	// third (2) record
	rc.AddNew(15, 15, &t1, nil, 0)
	rc.SaveRecord()

	// fourth (3) record
	rc.AddNew(15, 15, &t1, nil, 0)
	rc.AppendToCurrent(5, 5, &t2, nil, 0)
	rc.SaveRecord()

	if len(rc.records) != 4 {
		t.Errorf("Adding new recond failed")
	}

	// recorder has now three records
	rc.UndoOneRecord()

	if len(rc.records) != 3 {
		t.Errorf("Adding new recond failed")
	}

	// fourth (3) record added once again
	rc.AddNew(15, 15, &t1, nil, 0)
	rc.SaveRecord()

	// fifth (4) record added
	rc.AddNew(15, 15, &t1, nil, 0)
	rc.SaveRecord()

	if len(rc.records) != 5 {
		t.Errorf("Adding new recond failed, recorder has len: %d, expected 5", len(rc.records))
	}

	// recorder has now four records
	rc.UndoOneRecord()

	// recorder has now three records
	rc.UndoOneRecord()

	if len(rc.records) != 3 {
		t.Errorf("Adding new recond failed")
	}
}
