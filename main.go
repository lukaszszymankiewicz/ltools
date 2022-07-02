package main

import (
	"ltools/editor"
    "github.com/hajimehoshi/ebiten/v2"
    "log"
)

// runs the game
func main() {
	editor := editor.NewEditor()

	if err := ebiten.RunGame(editor); err != nil {
		log.Fatal(err)
	}
}
