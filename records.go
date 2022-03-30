package main

import (
    "fmt"
)

const (
    RECORDING_IDLE = iota
    RECORDING_DRAWING = iota
)

// Record struct holds one change on Canvas while drawing
// Because one stroke can draw multiple tiles at ones, all 
// changed tiles coords must be held both with old tiles which 
// were replaced
type Record struct {
    x_coords []int   // tiles x coords
    y_coords []int   // tiles y coords
    tiles []int      // new tiles
    old_tiles []int  // old tiles
}

type Recorder struct {
    records []Record  // old records
    pointer int
    state int
}

func NewRecord (x int, y int, new_tile int, old_tile int) Record {
    var r Record

    r.x_coords = append(r.x_coords, x)
    r.y_coords = append(r.y_coords, y)
    r.tiles = append(r.tiles, new_tile)
    r.old_tiles = append(r.old_tiles, old_tile)

    return r
}

func (rc *Recorder) AddNew (x int, y int, new_tile int, old_tile int) {
    r := NewRecord(x, y, new_tile, old_tile)
    rc.records = append(rc.records, r)
}


func (rc *Recorder) StopRecording () {
    rc.state = RECORDING_IDLE
}

func (rc *Recorder) StartRecording () {
    rc.state = RECORDING_DRAWING
}

func (rc *Recorder) IsRecording () bool {
    if rc.state == RECORDING_DRAWING {
        return true
    } else {
        return false
    }
}

func (rc *Recorder) AppendToCurrent (x int, y int, new_tile int, old_tile int) {
    if len(rc.records) == rc.pointer {
        rc.AddNew(x, y, new_tile, old_tile)
    } else {
        rc.records[rc.pointer].x_coords = append(rc.records[rc.pointer].x_coords, x)
        rc.records[rc.pointer].y_coords = append(rc.records[rc.pointer].y_coords, y)
        rc.records[rc.pointer].tiles = append(rc.records[rc.pointer].tiles, new_tile)
        rc.records[rc.pointer].old_tiles = append(rc.records[rc.pointer].old_tiles, old_tile)
    }
}

func (rc *Recorder) UndoOneRecord () Record {
    if len(rc.records) == 0 {
        return Record{}
    }
    record_to_undone := rc.records[len(rc.records)-1]
    rc.records = rc.records[:len(rc.records)-1]
    rc.pointer--
    return record_to_undone
}

func (rc *Recorder) SaveRecord () {
    rc.pointer++
}

func (rc *Recorder) Debug () {
    fmt.Printf("\n\n")
    fmt.Printf("pointer: %d\n", rc.pointer)
    fmt.Printf("length: %d\n", len(rc.records))

    for i:=0; i<len(rc.records); i++ {
        for j:=0; j<len(rc.records[i].x_coords); j++ {
            x := rc.records[i].x_coords[j]
            y := rc.records[i].y_coords[j]
            tile := rc.records[i].tiles[j]
            old := rc.records[i].old_tiles[j]
            fmt.Printf("record:%d, tile:%d -> coord: %d %d, tile (%d) replaced (%d) \n", i, j, x, y, tile, old)
        }
    }
}
