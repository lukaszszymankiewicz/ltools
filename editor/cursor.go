package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	_ "image/png"
)

const (
	CURSOR_SIZE = 32
)

type Cursor struct {
	Icons           []ImageElement
	Highlight       RectElement
	state           int
	highlight_state int
}

func NewCursor() Cursor {
	var crs Cursor

	crs.Highlight = NewRectElement(-5, -5, -5, -5, blackColor)

	crs.Icons = []ImageElement{
		NewImageElement("Dummy", 0, 0, ebiten.NewImage(1, 1)),
		NewImageElement("PenIcon", 0, 0, resources["pen15"]),
		NewImageElement("RubberIcon", 0, 0, resources["rubber"]),
		NewImageElement("RubberIcon", 0, 0, resources["pippete"]),
		NewImageElement("Dummy", 0, 0, ebiten.NewImage(1, 1)),
	}

	return crs
}

// cursor on pallete is created from one thick Rect, to draw it, simply few Rects, each one smaller
// than another is drawn.
func (crs *Cursor) DrawHighlight(screen *ebiten.Image) {

	x1 := crs.Highlight.Rect.Min.X
	y1 := crs.Highlight.Rect.Min.Y
	x2 := crs.Highlight.Rect.Max.X
	y2 := crs.Highlight.Rect.Max.Y

	for i, clr := range []*color.RGBA{blackColor, blackColor, whiteColor, blackColor, blackColor} {
		crs.Highlight.Rect.Min.X = x1 + i
		crs.Highlight.Rect.Min.Y = y1 + i
		crs.Highlight.Rect.Max.X = x2 - i
		crs.Highlight.Rect.Max.Y = y2 - i
		crs.Highlight.color = clr
		crs.Highlight.Draw(screen)
	}
}

func (crs *Cursor) Draw(screen *ebiten.Image) {
	crs.DrawHighlight(screen)
	crs.Icons[crs.state].Draw(screen)
}

func (crs *Cursor) Update(g *Game) {
	crs.state = g.Mouse.Icon.Active()

	_, height := crs.Icons[crs.state].image.Size()
	crs.Icons[crs.state].Rect.Min.X = g.Mouse.X
	crs.Icons[crs.state].Rect.Min.Y = g.Mouse.Y - height

	crs.highlight_state = g.Mouse.Highlight.Active()

	crs.Highlight.Rect.Min.X = g.Mouse.HighlightRect.Min.X
	crs.Highlight.Rect.Min.Y = g.Mouse.HighlightRect.Min.Y
	crs.Highlight.Rect.Max.X = g.Mouse.HighlightRect.Max.X
	crs.Highlight.Rect.Max.Y = g.Mouse.HighlightRect.Max.Y
}
