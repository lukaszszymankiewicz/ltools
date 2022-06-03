package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"testing"
)

func TestRectBasedElementWthoutColor(t *testing.T) {
	e := NewFilledRectElement(10, 20, 30, 40, nil)

	if e.rect.Min.X != 10 {
		t.Errorf("RectBasedElement created inproperly: rect.Min.X = %d, want = %d", e.rect.Min.X, 10)
	}

	if e.rect.Min.Y != 20 {
		t.Errorf("RectBasedElement created inproperly: rect.Min.Y = %d, want = %d", e.rect.Min.Y, 20)
	}

	if e.rect.Max.X != 40 {
		t.Errorf("RectBasedElement created inproperly: rect.Max.X = %d, want = %d", e.rect.Max.X, 40)
	}

	if e.rect.Max.Y != 60 {
		t.Errorf("RectBasedElement created inproperly: rect.Max.X = %d, want = %d", e.rect.Max.X, 60)
	}
}

func TestRectBasedElementWthColor(t *testing.T) {
	e := NewFilledRectElement(10, 20, 30, 40, &color.RGBA{192, 192, 192, 255})

	if e.rect.Min.X != 10 {
		t.Errorf("RectBasedElement created inproperly: rect.Min.X = %d, want = %d", e.rect.Min.X, 10)
	}

	if e.rect.Min.Y != 20 {
		t.Errorf("RectBasedElement created inproperly: rect.Min.Y = %d, want = %d", e.rect.Min.Y, 20)
	}

	if e.rect.Max.X != 40 {
		t.Errorf("RectBasedElement created inproperly: rect.Max.X = %d, want = %d", e.rect.Max.X, 40)
	}

	if e.rect.Max.Y != 60 {
		t.Errorf("RectBasedElement created inproperly: rect.Max.X = %d, want = %d", e.rect.Max.X, 60)
	}

	if e.color == nil {
		t.Errorf("RectBasedElement created inproperly: color was not set!")
	}
}

func TestNewElementImageBased(t *testing.T) {
	image_a := ebiten.NewImage(100, 200)

	e := NewImageElement(10, 20, image_a)

	if e.rect.Min.X != 10 {
		t.Errorf("ImageBasedElement created inproperly: rect.Max.X = %d, want = %d", e.rect.Min.X, 10)
	}

	if e.rect.Min.Y != 20 {
		t.Errorf("ImageBasedElement created inproperly: rect.Min.Y = %d, want = %d", e.rect.Min.Y, 20)
	}

	if e.rect.Max.X != 110 {
		t.Errorf("ImageBasedElement created inproperly: rect.Max.X = %d, want = %d", e.rect.Min.X, 110)
	}

	if e.rect.Max.Y != 220 {
		t.Errorf("ImageBasedElement created inproperly: rect.Max.Y = %d, want = %d", e.rect.Max.Y, 220)
	}

}
