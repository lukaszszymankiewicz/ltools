package editor

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image"
    _ "image/png"
    "log"
)

const (
	DUMMY_IMAGE_W = 10
	DUMMY_IMAGE_H = 10
)

// checks if x and y are inside Rect greedily
func coordsInRect(x int, y int, rect image.Rectangle) bool {
	if (x >= rect.Min.X && x <= rect.Max.X) && (y >= rect.Min.Y && y <= rect.Max.Y) {
		return true
	}
	return false
}

// returns bigger value from inputted two
func MaxVal(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// returns smaller value from inputted two
func MinVal(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// loads image from given path
func LoadImage(path string) *ebiten.Image {
	var err error
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		// if image cannot be read, dummy image is returned instead
		log.Fatal(err)
		dummy_image := ebiten.NewImage(DUMMY_IMAGE_W, DUMMY_IMAGE_H)

		return dummy_image
	}

	return img
}
