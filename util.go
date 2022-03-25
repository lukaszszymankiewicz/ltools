package main

import (
	"image"
	_ "image/png"
)

// checks if x and y are inside Rect greedily
func coordsInRect(x int, y int, rect image.Rectangle) bool {
	if (x >= rect.Min.X && x <= rect.Max.X) && (y >= rect.Min.Y && y <= rect.Max.Y) {
		return true
	}
	return false
}
