package objects

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	_ "image/png"
)

const (
	dpi = 72
)

var normalFont font.Face

func init() {
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)

	normalFont, _ = opentype.NewFace(
		tt,
		&opentype.FaceOptions{
			Size:    14,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
}
