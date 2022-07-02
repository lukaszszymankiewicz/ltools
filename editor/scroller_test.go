package editor

import (
	"testing"
)

func TestNewScrollerX(t *testing.T) {
	// scroller arrows image cannot be read properly while initialiing test, so dummy image (10x10
	// pixels is load) - this is taken into account while this unit test.

	dummy_image_x := 10
	dummy_image_y := 10

	s := NewScroller(10, 10, 200, 32)

	// background
	if s.background.rect.Min.X != 10+dummy_image_x {
		t.Errorf("Scroller created inproperly: background.rect.Min.X = %d, want = %d", s.background.rect.Min.X, 10+dummy_image_x)
	}
	if s.background.rect.Max.X != 200 {
		t.Errorf("Scroller created inproperly: background.rect.Max.X = %d, want = %d", s.background.rect.Max.X, 200)
	}

	if s.background.rect.Min.Y != 10 {
		t.Errorf("Scroller created inproperly: background.rect.Min.Y = %d, want = %d", s.background.rect.Min.Y, 10)
	}

	if s.background.rect.Max.Y != 10+32 {
		t.Errorf("Scroller created inproperly: background.rect.Max.Y = %d, want = %d", s.background.rect.Max.Y, 10+32)
	}

	// bar
	if s.bar.rect.Min.X != 10+dummy_image_x {
		t.Errorf("Scroller created inproperly: bar.rect.Min.X = %d, want = %d", s.bar.rect.Min.X, 10+dummy_image_x)
	}
	if s.bar.rect.Max.X != 200 {
		t.Errorf("Scroller created inproperly: bar.rect.Max.X = %d, want = %d", s.bar.rect.Max.X, 200)
	}

	if s.bar.rect.Min.Y != 10 {
		t.Errorf("Scroller created inproperly: bar.rect.Min.Y = %d, want = %d", s.bar.rect.Min.Y, 10)
	}

	if s.bar.rect.Max.Y != 10+32 {
		t.Errorf("Scroller created inproperly: bar.rect.Max.Y = %d, want = %d", s.bar.rect.Max.Y, 10+32)
	}

	// Low Arrow
	if s.arrowLow.rect.Min.X != 10 {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Min.X = %d, want = %d", s.arrowLow.rect.Min.X, 10)
	}

	if s.arrowLow.rect.Max.X != 10+dummy_image_x {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Max.X = %d, want = %d", s.arrowLow.rect.Min.X, 10+dummy_image_x)
	}

	if s.arrowLow.rect.Min.Y != 10 {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Min.Y = %d, want = %d", s.arrowLow.rect.Min.Y, 10)
	}

	if s.arrowLow.rect.Max.Y != 10+dummy_image_y {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Max.Y = %d, want = %d", s.arrowLow.rect.Min.Y, 10+dummy_image_y)
	}

	// High Arrow
	if s.arrowHigh.rect.Min.X != 200 {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Min.X = %d, want = %d", s.arrowHigh.rect.Min.X, 200)
	}

	if s.arrowHigh.rect.Max.X != 200+dummy_image_x {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Max.X = %d, want = %d", s.arrowHigh.rect.Max.X, 200+dummy_image_y)
	}

	if s.arrowHigh.rect.Min.Y != 10 {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Min.Y = %d, want = %d", s.arrowHigh.rect.Min.Y, 10)
	}

	if s.arrowHigh.rect.Max.Y != 10+dummy_image_y {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Max.Y = %d, want = %d", s.arrowHigh.rect.Max.Y, 10+dummy_image_y)
	}

}

func TestNewScrollerY(t *testing.T) {
	// scroller arrows image cannot be read properly while initialiing test, so dummy image (10x10
	// pixels is load) - this is taken into account while this unit test.

	dummy_image_x := 10
	dummy_image_y := 10

	s := NewScroller(10, 10, 32, 200)

	// background
	if s.background.rect.Min.X != 10 {
		t.Errorf("Scroller created inproperly: background.rect.Min.X = %d, want = %d", s.background.rect.Min.X, 10)
	}

	if s.background.rect.Max.X != 10+32 {
		t.Errorf("Scroller created inproperly: background.rect.Max.X = %d, want = %d", s.background.rect.Max.X, 10+32)
	}

	if s.background.rect.Min.Y != 10+dummy_image_y {
		t.Errorf("Scroller created inproperly: background.rect.Min.Y = %d, want = %d", s.background.rect.Min.Y, 10+dummy_image_y)
	}

	if s.background.rect.Max.Y != 200 {
		t.Errorf("Scroller created inproperly: background.rect.Max.Y = %d, want = %d", s.background.rect.Max.Y, 200)
	}

	// bar
	if s.bar.rect.Min.X != 10 {
		t.Errorf("Scroller created inproperly: bar.rect.Min.X = %d, want = %d", s.bar.rect.Min.X, 10)
	}

	if s.bar.rect.Max.X != 10+32 {
		t.Errorf("Scroller created inproperly: bar.rect.Max.X = %d, want = %d", s.bar.rect.Max.X, 10+32)
	}

	if s.bar.rect.Min.Y != 10+dummy_image_y {
		t.Errorf("Scroller created inproperly: bar.rect.Min.Y = %d, want = %d", s.bar.rect.Min.Y, 10+dummy_image_y)
	}

	if s.bar.rect.Max.Y != 200 {
		t.Errorf("Scroller created inproperly: bar.rect.Max.Y = %d, want = %d", s.bar.rect.Max.Y, 200)
	}

	// Low Arrow
	if s.arrowLow.rect.Min.X != 10 {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Min.X = %d, want = %d", s.arrowLow.rect.Min.X, 10)
	}

	if s.arrowLow.rect.Max.X != 10+dummy_image_x {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Max.X = %d, want = %d", s.arrowLow.rect.Min.X, 10+dummy_image_x)
	}

	if s.arrowLow.rect.Min.Y != 10 {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Min.Y = %d, want = %d", s.arrowLow.rect.Min.Y, 10)
	}

	if s.arrowLow.rect.Max.Y != 10+dummy_image_y {
		t.Errorf("Scroller created inproperly: s.arrowLow.rect.Max.Y = %d, want = %d", s.arrowLow.rect.Min.Y, 10+dummy_image_y)
	}

	// High Arrow
	if s.arrowHigh.rect.Min.X != 10 {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Min.X = %d, want = %d", s.arrowHigh.rect.Min.X, 10)
	}

	if s.arrowHigh.rect.Max.X != 10+dummy_image_x {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Max.X = %d, want = %d", s.arrowHigh.rect.Max.X, 10+dummy_image_y)
	}

	if s.arrowHigh.rect.Min.Y != 200 {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Min.Y = %d, want = %d", s.arrowHigh.rect.Min.Y, 200)
	}

	if s.arrowHigh.rect.Max.Y != 200+dummy_image_y {
		t.Errorf("Scroller created inproperly: s.arrowHigh.rect.Max.Y = %d, want = %d", s.arrowHigh.rect.Max.Y, 200+dummy_image_y)
	}
}
