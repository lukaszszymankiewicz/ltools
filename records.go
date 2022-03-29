package main

import (
    "fmt"
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
    length  int
}

func NewRecord (x int, y int, new_tile int, old_tile int) Record {
    var r Record

    r.x_coords = append(r.x_coords, x)
    r.y_coords = append(r.y_coords, y)
    r.tiles = append(r.tiles, new_tile)
    r.old_tiles = append(r.old_tiles, old_tile)

    return r
}

func (r *Record) Append (x int, y int, new_tile int, old_tile int) {
    r.x_coords = append(r.x_coords, x)
    r.y_coords = append(r.y_coords, y)
    r.tiles = append(r.tiles, new_tile)
    r.old_tiles = append(r.old_tiles, old_tile)
}

func (rc *Recorder) AddNew (x int, y int, new_tile int, old_tile int) {
    r := NewRecord(x, y, new_tile, old_tile)
    rc.records = append(rc.records, r)
    rc.length++
    rc.pointer = rc.length
}

func (rc *Recorder) AppendToCurrent (x int, y int, new_tile int, old_tile int) {
    rc.records[rc.length].x_coords = append(rc.records[rc.length].x_coords, x)
    rc.records[rc.length].y_coords = append(rc.records[rc.length].y_coords, y)
    rc.records[rc.length].tiles = append(rc.records[rc.length].tiles, new_tile)
    rc.records[rc.length].old_tiles = append(rc.records[rc.length].old_tiles, old_tile)
}

func (rc *Recorder) Debug () {
    for i:=0; i<rc.length; i++ {
        for j:=0; j<len(rc.records[i].x_coords); j++ {
            fmt.Printf("record = %d -> %d %d, \n", i, rc.records[i].x_coords[j], rc.records[i].x_coords[j])
        }
    }
}
