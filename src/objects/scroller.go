package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"ltools/src/drawer"
)

type Scroller struct {
	color      color.Color
    bgcolor    color.Color
	MaxRect    image.Rectangle
	Rect       image.Rectangle
    arrowLow   ScrollArrow
    arrowhHigh ScrollArrow
}

type ScrollArrow struct {
	image *ebiten.Image
}

// creates new scroller struct
func NewScroller(
    x int, y int, width int, height int, color color.Color, bgcolor color.Color
) Scroller {
	var s Scroller

    // we are assuming that this is x-scroller if it is wider that higher
    if width > height {
        s.arrowLow = NewScrollArrow(x, y+height, "assets/arrow_l.png")
        s.arrowhHigh = NewScrollArrow(x+width-32, y+height, "assets/arrow_r.png")
        s.MaxRect = image.Rect(x+32, y+height, x+width-32, y+height+32)
    } else {
    // we are assuming that this is y-scroller if it is higher than wider
        s.arrowLow = NewScrollArrow(x+width, y, "assets/arrow_u.png")
        s.arrowhHigh = NewScrollArrow(x+width, y+height-32, "assets/arrow_d.png")
        s.MaxRect = image.Rect(x+width, y+32, x+width+32, y+height-32)
    }
 
	s.color = color
	s.bgcolor = bgcolor

	return s
}

// creates new arrow for scroll control
func NewScrollArrow(x int, y int, path string) ScrollArrow {
	var scrollArrow ScrollArrow

	scrollArrow.image = loadImage(path)

	return scrollArrow
}

// draws ScrollArrow
func (sa *ScrollArrow) DrawScrollArrow(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(sa.Image.Rect.Min.X), float64(sa.Image.Rect.Min.Y))

	screen.DrawImage(sa.image, op)
}

func (s *scroller) DrawScroller(screen *ebiten.Image) {
    // arrows 
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.arrowLow.Image.Rect.Min.X), float64(s.arrowLow.Image.Rect.Min.Y))
	screen.DrawImage(s.arrowLow.Image, op)

	// background
	drawer.FilledRect(screen, c.scroller_x.MaxRect, c.scroller_x.bgcolor)
	drawer.FilledRect(screen, c.scroller_y.MaxRect, c.scroller_y.bgcolor)

	// actual scroller
	drawer.FilledRect(screen, c.scroller_x.Rect, c.scroller_x.color)
	drawer.FilledRect(screen, c.scroller_y.Rect, c.scroller_y.color)
}

// draws both Scrollers. Each scroller is consisted from background and main part
// Main part indicates how much of canvas is visible and position of this visible
// part. Like in normal scroller
func (c *Canvas) DrawScrollers(screen *ebiten.Image) {
	// background
	drawer.FilledRect(screen, c.scroller_x.MaxRect, c.scroller_x.bgcolor)
	drawer.FilledRect(screen, c.scroller_y.MaxRect, c.scroller_y.bgcolor)

	// actual scroller
	drawer.FilledRect(screen, c.scroller_x.Rect, c.scroller_x.color)
	drawer.FilledRect(screen, c.scroller_y.Rect, c.scroller_y.color)
}

// returns X Scrollers main part position
func (c *Canvas) getXScrollerRect() image.Rectangle {
	var start float64
	len := float64(c.scroller_x.MaxRect.Max.X - c.scroller_x.MaxRect.Min.X)

	if c.viewport_x == 0 {
		start = float64(c.scroller_x.MaxRect.Min.X)
	} else {
		start = float64(c.scroller_x.MaxRect.Min.X) + (float64(c.viewport_x)/float64(c.canvasCols))*len
	}
	end := float64(c.scroller_x.MaxRect.Min.X) + ((float64(c.viewport_x)+float64(c.viewportCols))/float64(c.canvasCols))*len

	return image.Rect(int(start), c.scroller_x.MaxRect.Min.Y, int(end), c.scroller_x.MaxRect.Max.Y)
}

// returns Y Scrollers main part position
func (c *Canvas) getYScrollerRect() image.Rectangle {
	var start float64
	len := float64(c.scroller_y.MaxRect.Max.Y - c.scroller_y.MaxRect.Min.Y)

	if c.viewport_y == 0 {
		start = float64(c.scroller_y.MaxRect.Min.Y)
	} else {
		start = float64(c.scroller_y.MaxRect.Min.Y) + (float64(c.viewport_y)/float64(c.canvasRows))*len
	}
	end := float64(c.scroller_y.MaxRect.Min.Y) + ((float64(c.viewport_y)+float64(c.viewportRows))/float64(c.canvasRows))*len

	return image.Rect(c.scroller_y.MaxRect.Min.X, int(start), c.scroller_y.MaxRect.Max.X, int(end))
}

// updates Scrollers position due to viewport position according to viewport position
func (c *Canvas) UpdateScrollers() {
	c.scroller_x.Rect = c.getXScrollerRect()
	c.scroller_y.Rect = c.getYScrollerRect()
}

