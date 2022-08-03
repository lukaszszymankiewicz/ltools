package editor

import (
	"fmt"
	"image"
	"testing"
)

func RectAreEqual(base image.Rectangle, comp image.Rectangle) (bool, int, int, string) {
	if base.Min.X != comp.Min.X {
		return false, base.Min.X, comp.Min.X, "min X"
	}
	if base.Max.X != comp.Max.X {
		return false, base.Max.X, comp.Max.X, "max X"
	}
	if base.Min.Y != comp.Min.Y {
		return false, base.Min.Y, comp.Min.Y, "min Y"
	}
	if base.Max.Y != comp.Max.Y {
		return false, base.Max.Y, comp.Max.Y, "max Y"
	}

	return true, -1, -1, ""

}

func TestNewScroller(t *testing.T) {

	SCROLLER_X := 10
	SCROLLER_Y := 10
	scroll_image_x := 16
	scroll_image_y := 16
	rows := 30
	cols := 30
	grid := 32

	base_x := cols * grid
	base_y := rows * grid

	var tests = []struct {
		scroller_position int
		scroller_type     int
		background_rect   image.Rectangle
		bar_rect          image.Rectangle
		low_arrow_rect    image.Rectangle
		high_arrow_rect   image.Rectangle
	}{
		{
			SCROLLER_POS_DOWN,
			SCROLLER_TYPE_X,
			image.Rect(
				0+SCROLLER_X+scroll_image_x,
				base_x+SCROLLER_X,
				base_x+SCROLLER_X-scroll_image_x,
				base_y+SCROLLER_Y+scroll_image_y,
			),
			image.Rect(
				0+SCROLLER_X+scroll_image_x,
				base_y+SCROLLER_Y,
				base_x+SCROLLER_X-scroll_image_x,
				base_y+SCROLLER_Y+scroll_image_y,
			),
			image.Rect(
				0+SCROLLER_X,
				base_y+SCROLLER_Y,
				0+SCROLLER_X+scroll_image_x,
				base_y+SCROLLER_Y+scroll_image_y,
			),
			image.Rect(
				base_x+SCROLLER_X-scroll_image_x,
				base_y+SCROLLER_Y,
				base_x+SCROLLER_X,
				base_y+SCROLLER_Y+scroll_image_y,
			),
		},
		{
			SCROLLER_POS_UP,
			SCROLLER_TYPE_Y,
			image.Rect(
				0,
				SCROLLER_X+scroll_image_y,
				scroll_image_x,
				base_y+SCROLLER_Y-scroll_image_y,
			),
			image.Rect(
				0,
				SCROLLER_X+scroll_image_y,
				scroll_image_x,
				base_y+SCROLLER_Y-scroll_image_y,
			),
			image.Rect(
				0,
				base_y+SCROLLER_Y-scroll_image_y,
				0+scroll_image_x,
				base_y+SCROLLER_Y,
			),
			image.Rect(
				0,
				SCROLLER_Y,
				scroll_image_x,
				scroll_image_y+SCROLLER_Y,
			),
		},
		{
			SCROLLER_POS_DOWN,
			SCROLLER_TYPE_Y,
			image.Rect(
				0,
				SCROLLER_X+scroll_image_y,
				scroll_image_x,
				base_y+SCROLLER_Y-scroll_image_y,
			),
			image.Rect(
				0,
				SCROLLER_X+scroll_image_y,
				scroll_image_x,
				base_y+SCROLLER_Y-scroll_image_y,
			),
			image.Rect(
				0,
				base_y+SCROLLER_Y-scroll_image_y,
				0+scroll_image_x,
				base_y+SCROLLER_Y,
			),
			image.Rect(
				0,
				SCROLLER_Y,
				scroll_image_x,
				scroll_image_y+SCROLLER_Y,
			),
		},
		{
			SCROLLER_POS_DOWN,
			SCROLLER_TYPE_X,
			image.Rect(
				SCROLLER_X+scroll_image_x,
				base_y+SCROLLER_X,
				base_x+SCROLLER_Y-scroll_image_x,
				base_y+SCROLLER_Y+scroll_image_y,
			),
			image.Rect(
				SCROLLER_X+scroll_image_x,
				base_y+SCROLLER_X,
				base_x+SCROLLER_Y-scroll_image_x,
				base_y+SCROLLER_Y+scroll_image_y,
			),
			image.Rect(
				SCROLLER_X,
				base_y+SCROLLER_Y,
				SCROLLER_Y+scroll_image_y,
				base_y+SCROLLER_Y+scroll_image_y,
			),
			image.Rect(
				base_x+SCROLLER_X-scroll_image_x,
				base_y+SCROLLER_Y,
				base_x+SCROLLER_X,
				base_y+SCROLLER_Y+scroll_image_y,
			),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.scroller_position)

		var v Viewport
		var s *Scroller

		t.Run(testname, func(t *testing.T) {

			v = NewViewport(30, 30, 30, 30)
			s = NewScroller("dummy", SCROLLER_X, SCROLLER_Y, tt.scroller_position, tt.scroller_type, &v)

			equal, base, comp, where := RectAreEqual(s.Background.Rect, tt.background_rect)
			if equal == false {
				t.Errorf("Scroller background rect is incorrect: %d != %d at %s", base, comp, where)
				t.Errorf("%d, %d, %d, %d\n", s.Background.Rect.Min.X, s.Background.Rect.Min.Y, s.Background.Rect.Max.X, s.Background.Rect.Max.Y)
			}

			equal, base, comp, where = RectAreEqual(s.Bar.Rect, tt.bar_rect)
			if equal == false {
				t.Errorf("Scroller bar rect is incorrect: %d != %d at %s", base, comp, where)
				t.Errorf("%d, %d, %d, %d\n", s.Bar.Rect.Min.X, s.Bar.Rect.Min.Y, s.Bar.Rect.Max.X, s.Bar.Rect.Max.Y)
			}

			equal, base, comp, where = RectAreEqual(s.ArrowLow.Rect, tt.low_arrow_rect)
			if equal == false {
				t.Errorf("Scroller low arrow rect is incorrect: %d != %d at %s", base, comp, where)
				t.Errorf("%d, %d, %d, %d\n", s.ArrowLow.Rect.Min.X, s.ArrowLow.Rect.Min.Y, s.ArrowLow.Rect.Max.X, s.ArrowLow.Rect.Max.Y)
			}

			equal, base, comp, where = RectAreEqual(s.ArrowHigh.Rect, tt.high_arrow_rect)
			if equal == false {
				t.Errorf("Scroller high arrow rect is incorrect: %d != %d at %s", base, comp, where)
				t.Errorf("%d, %d, %d, %d\n", s.ArrowHigh.Rect.Min.X, s.ArrowHigh.Rect.Min.Y, s.ArrowHigh.Rect.Max.X, s.ArrowHigh.Rect.Max.Y)
			}
		})
	}
}
