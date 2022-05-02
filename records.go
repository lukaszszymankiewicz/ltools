package main

const (
	RECORDING_IDLE    = iota
	RECORDING_DRAWING = iota
)

// struct of a single non-paused draw on canvas
type Record struct {
	x_coords   []int // tiles x coords
	y_coords   []int // tiles y coords
	tiles      []int // new tiles
	old_tiles  []int // old tiles
	layers     []int // tiles layers
	old_layers []int // old tiles layers
}

// struct which hold all Records
type Recorder struct {
	records []Record // old records
	pointer int
	state   int
}

// creates new Record struct
func NewRecord(x int, y int, new_tile int, old_tile int, layers int, old_layers int) Record {
	var r Record

	r.x_coords = append(r.x_coords, x)
	r.y_coords = append(r.y_coords, y)
	r.tiles = append(r.tiles, new_tile)
	r.old_tiles = append(r.old_tiles, old_tile)
	r.layers = append(r.layers, layers)
	r.old_layers = append(r.old_layers, old_layers)

	return r
}

// Adds new Record to Recorder
func (rc *Recorder) AddNew(x int, y int, new_tile int, old_tile int, layer int, old_layer int) {
	r := NewRecord(x, y, new_tile, old_tile, layer, old_layer)
	rc.records = append(rc.records, r)
}

// force Recorder to stop recording current Record
func (rc *Recorder) StopRecording() {
	rc.state = RECORDING_IDLE
}

// force Recorder to start recording new Record
func (rc *Recorder) StartRecording() {
	rc.state = RECORDING_DRAWING
}

// check if Recorder is recording to current Record
func (rc *Recorder) IsRecording() bool {
	if rc.state == RECORDING_DRAWING {
		return true
	} else {
		return false
	}
}

// ads new data to current Record
func (rc *Recorder) AppendToCurrent(
	x int,
	y int,
	new_tile int,
	old_tile int,
	layer int,
	old_layer int,
) {
	if len(rc.records) == rc.pointer {
		rc.AddNew(x, y, new_tile, old_tile, layer, old_layer)
	} else {
		rc.records[rc.pointer].x_coords = append(rc.records[rc.pointer].x_coords, x)
		rc.records[rc.pointer].y_coords = append(rc.records[rc.pointer].y_coords, y)
		rc.records[rc.pointer].tiles = append(rc.records[rc.pointer].tiles, new_tile)
		rc.records[rc.pointer].old_tiles = append(rc.records[rc.pointer].old_tiles, old_tile)
		rc.records[rc.pointer].layers = append(rc.records[rc.pointer].layers, layer)
		rc.records[rc.pointer].old_layers = append(rc.records[rc.pointer].old_layers, old_layer)
	}
}

// pops out last Record from Recorder. If there is not any Records, empty struct is returned
// (only for compliance with other functions)
func (rc *Recorder) UndoOneRecord() Record {
	if len(rc.records) == 0 {
		return Record{}
	}
	record_to_undone := rc.records[len(rc.records)-1]
	rc.records = rc.records[:len(rc.records)-1]
	rc.pointer--

	return record_to_undone
}

// closes recording of current Record
func (rc *Recorder) SaveRecord() {
	rc.pointer++
}
