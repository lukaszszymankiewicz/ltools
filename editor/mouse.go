package editor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	MOUSE_NO_HIGHLIGH = iota
	MOUSE_HIGHLIGH
	MOUSE_ALL_HIGHLIGHT_MODES
)

const (
	MOUSE_NORMAL = iota
	MOUSE_PEN
	MOUSE_RUBBER
	MOUSE_PIPPETE
	MOUSE_ALL_MODES
)

type Mouse struct {
	X             int
	Y             int
	Icon          Mode
	Highlight     Mode
	HighlightRect image.Rectangle
}

func NewMouse() Mouse {
	var m Mouse

	m.Icon = NewMode(MOUSE_ALL_MODES)
	m.Highlight = NewMode(MOUSE_ALL_HIGHLIGHT_MODES)
	m.HighlightRect = image.Rect(0, 0, 0, 0)

	return m
}

func (m *Mouse) Update() {
	m.X, m.Y = ebiten.CursorPosition()
}

func (m *Mouse) SetHighlight(x1 int, y1 int, x2 int, y2 int) {
	m.HighlightRect.Min.X = x1
	m.HighlightRect.Min.Y = y1
	m.HighlightRect.Max.X = x2
	m.HighlightRect.Max.Y = y2
}

func (m *Mouse) CleanHighlight() {
	m.HighlightRect.Min.X = -5
	m.HighlightRect.Min.Y = -5
	m.HighlightRect.Max.X = -5
	m.HighlightRect.Max.Y = -5
}
