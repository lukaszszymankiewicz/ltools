package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
)

// draws empty rect (only rect border)
func EmptyRect(screen *ebiten.Image, rect image.Rectangle, color *color.RGBA) {
	ebitenutil.DrawLine(
		screen,
		float64(rect.Min.X),
		float64(rect.Min.Y),
		float64(rect.Max.X),
		float64(rect.Min.Y),
		*color,
	)
	ebitenutil.DrawLine(
		screen,
		float64(rect.Min.X),
		float64(rect.Min.Y),
		float64(rect.Min.X),
		float64(rect.Max.Y),
		*color,
	)
	ebitenutil.DrawLine(
		screen,
		float64(rect.Max.X),
		float64(rect.Min.Y),
		float64(rect.Max.X),
		float64(rect.Max.Y),
		*color,
	)
	ebitenutil.DrawLine(
		screen,
		float64(rect.Min.X),
		float64(rect.Max.Y),
		float64(rect.Max.X),
		float64(rect.Max.Y),
		*color,
	)
}

// draws empty rect (only rect border)
func EmptyBorder(screen *ebiten.Image, rect image.Rectangle, color *color.RGBA) {
	leftUpX := float64(rect.Min.X)
	leftUpY := float64(rect.Min.Y) - 1
	rightUpX := float64(rect.Max.X) + 1
	rightUpY := float64(rect.Min.Y) - 1
	leftDownX := float64(rect.Min.X)
	leftDownY := float64(rect.Max.Y)
	rightDownX := float64(rect.Max.X) + 1
	rightDownY := float64(rect.Max.Y)

	ebitenutil.DrawLine(screen, leftUpX, leftUpY, rightUpX, rightUpY, *color)
	ebitenutil.DrawLine(screen, leftUpX, leftUpY, leftDownX, leftDownY, *color)
	ebitenutil.DrawLine(screen, rightUpX, rightUpY, rightDownX, rightDownY, *color)
	ebitenutil.DrawLine(screen, leftDownX, leftDownY, rightDownX, rightDownY, *color)
}

// draws rect filled with color
func FilledRect(screen *ebiten.Image, rect image.Rectangle, color *color.RGBA) {
	ebitenutil.DrawRect(
		screen,
		float64(rect.Min.X),
		float64(rect.Min.Y),
		float64(rect.Max.X-rect.Min.X),
		float64(rect.Max.Y-rect.Min.Y),
		*color,
	)
}
