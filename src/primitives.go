package primities

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
)

func DrawEmptyRect(screen *ebiten.Image, rect image.Rectangle, color color.Color) {
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
        color
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Max.X),
        float64(rect.Min.Y),
        float64(rect.Max.X),
        float64(rect.Max.Y),
        color
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Min.X),
        float64(rect.Max.Y),
        float64(rect.Max.X),
        float64(rect.Max.Y),
        color
    )
}

func DrawBorder(screen *ebiten.Image, rect image.Rectangle, color color.Color) {
	ebitenutil.DrawLine(
        screen,
        float64(rect.Min.X-1),
        float64(rect.Min.Y-1),
        float64(rect.Max.X+1),
        float64(rect.Min.Y+1),
        color,
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Min.X-1),
        float64(rect.Min.Y-1),
        float64(rect.Min.X+1),
        float64(rect.Max.Y+1),
        color
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Max.X-1),
        float64(rect.Min.Y-1),
        float64(rect.Max.X+1),
        float64(rect.Max.Y+1),
        color
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Min.X-1),
        float64(rect.Max.Y-1),
        float64(rect.Max.X+1),
        float64(rect.Max.Y+1),
        color
    )
}

