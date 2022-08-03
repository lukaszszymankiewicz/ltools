package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"ltools/editor"
)

func main() {
	editor := editor.NewEditor()

	if err := ebiten.RunGame(editor); err != nil {
		log.Fatal(err)
	}
}
