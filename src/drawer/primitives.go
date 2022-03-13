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
        color,
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Max.X-1),
        float64(rect.Min.Y-1),
        float64(rect.Max.X+1),
        float64(rect.Max.Y+1),
        color,
    )
	ebitenutil.DrawLine(
        screen,
        float64(rect.Min.X-1),
        float64(rect.Max.Y-1),
        float64(rect.Max.X+1),
        float64(rect.Max.Y+1),
        color,
    )
}

