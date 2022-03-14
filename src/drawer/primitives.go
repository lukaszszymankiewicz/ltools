package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
)

func EmptyRect(screen *ebiten.Image, rect image.Rectangle, color color.Color) {
	ebitenutil.DrawLine(
		screen,
		float64(rect.Min.X),
		float64(rect.Min.Y),
		float64(rect.Max.X),
		float64(rect.Min.Y),
		color,
	)
	ebitenutil.DrawLine(
		screen,
		float64(rect.Min.X),
		float64(rect.Min.Y),
		float64(rect.Min.X),
		float64(rect.Max.Y),
		color,
	)
	ebitenutil.DrawLine(
		screen,
		float64(rect.Max.X),
		float64(rect.Min.Y),
		float64(rect.Max.X),
		float64(rect.Max.Y),
		color,
	)
	ebitenutil.DrawLine(
		screen,
		float64(rect.Min.X),
		float64(rect.Max.Y),
		float64(rect.Max.X),
		float64(rect.Max.Y),
		color,
	)
}

func EmptyBorder(screen *ebiten.Image, rect image.Rectangle, color color.Color) {
	leftUpX := float64(rect.Min.X)
	leftUpY := float64(rect.Min.Y) - 1
	rightUpX := float64(rect.Max.X) + 1
	rightUpY := float64(rect.Min.Y) - 1
	leftDownX := float64(rect.Min.X)
	leftDownY := float64(rect.Max.Y)
	rightDownX := float64(rect.Max.X) + 1
	rightDownY := float64(rect.Max.Y)

	ebitenutil.DrawLine(screen, leftUpX, leftUpY, rightUpX, rightUpY, color)
	ebitenutil.DrawLine(screen, leftUpX, leftUpY, leftDownX, leftDownY, color)
	ebitenutil.DrawLine(screen, rightUpX, rightUpY, rightDownX, rightDownY, color)
	ebitenutil.DrawLine(screen, leftDownX, leftDownY, rightDownX, rightDownY, color)
}
