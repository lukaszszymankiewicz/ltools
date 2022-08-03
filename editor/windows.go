package editor

import ()

func (g *Game) NewMainWindow() *Window {
	var w Window

	w.elements = []WindowElement{
		NewFilledRectElement(
			"Background",
			0,
			0,
			ScreenWidth,
			ScreenHeight,
			menuGrey,
		),
		NewImgMatrix(
			"Pallete",
			PalleteX,
			PalleteY,
			&g.ViewportPallete,
			&g.AvailableTiles,
		),
		NewScroller(
			"PalleteYScroller",
			PalleteX,
			PalleteY,
			SCROLLER_POS_RIGHT,
			SCROLLER_TYPE_Y,
			&g.ViewportPallete,
		),
		NewMultiImgMatrix(
			"Canvas",
			CanvasX,
			CanvasY,
			&g.ViewportCanvas,
			&g.LevelStructure,
			BasicLayersAlpha,
		),
		NewScroller(
			"CanvasYScroller",
			CanvasX,
			CanvasY,
			SCROLLER_POS_RIGHT,
			SCROLLER_TYPE_Y,
			&g.ViewportCanvas,
		),
		NewScroller(
			"CanvasXScroller",
			CanvasX,
			CanvasY,
			SCROLLER_POS_DOWN,
			SCROLLER_TYPE_X,
			&g.ViewportCanvas,
		),
		NewButton(
			"PencilButton",
			ToolboxX,
			ToolboxY,
			resources["pen15"],
			g.Toolbox.State(PENCIL_TOOL_IDX),
		),
		NewButton(
			"RubberButton",
			ToolboxX+30,
			ToolboxY,
			resources["rubber"],
			g.Toolbox.State(RUBBER_TOOL_IDX),
		),
		NewImgMatrix(
			"Fill",
			ToolboxX+GridSize*5,
			ToolboxY,
			&g.ViewportTiles,
			&g.FillTiles,
		),
		NewButton(
			"RunButton",
			MenuX,
			MenuY+4,
			resources["run"],
			nil,
		),
		NewPole("Pole", MenuX+35, MenuY+4, POLE_TYPE_Y, GridSize-2),
		NewButton(
			"TileButton",
			MenuX+35+7,
			MenuY+4,
			resources["tiles"],
			g.Mode.State(MODE_DRAW),
		),
		NewButton(
			"LightButton",
			MenuX+35+7+35,
			MenuY+4,
			resources["light"],
			g.Mode.State(MODE_LIGHT),
		),
		NewButton(
			"EntityButton",
			MenuX+35+7+35+35,
			MenuY+4,
			resources["entities"],
			g.Mode.State(MODE_ENTITIES),
		),
	}

	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"CanvasYScroller", "ArrowLow"}), g.moveCanvasDown)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"CanvasYScroller", "ArrowHigh"}), g.moveCanvasUp)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"CanvasXScroller", "ArrowLow"}), g.moveCanvasLeft)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"CanvasXScroller", "ArrowHigh"}), g.moveCanvasRight)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"PalleteXScroller", "ArrowLow"}), g.movePalleteDown)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"PalleteXScroller", "ArrowHigh"}), g.movePalleteUp)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"TileButton", "Area"}), g.changeModeToDraw)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"LightButton", "Area"}), g.changeModeToDrawLight)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"EntityButton", "Area"}), g.changeModeToDrawEntities)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"PencilButton", "Area"}), g.changeToolToPencil)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"RubberButton", "Area"}), g.changeToolToRubber)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"RunButton", "Area"}), g.Export)
	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"Pallete", "Area"}), g.ChooseTileFromPallete)
	w.Bind(LONG_CLICK_FUNC, w.GetArea([]string{"Canvas", "Area"}), g.DrawTileOnCanvas)

	w.Bind(HOVER_FUNC, w.GetArea([]string{"Pallete", "Area"}), g.setPippeteIcon)
	w.Bind(HOVER_FUNC, w.GetArea([]string{"Pallete", "Area"}), g.setHighlightOnPallete)
	w.Bind(HOVER_FUNC, w.GetArea([]string{"Canvas", "Area"}), g.setToolIcon)
	w.Bind(HOVER_FUNC, w.GetArea([]string{"Canvas", "Area"}), g.setHighlightOnCanvas)

	w.BindDefault(HOVER_FUNC, g.setNormalIcon)
	w.BindDefault(HOVER_FUNC, g.clearHighlight)

	return &w
}

func (g *Game) NewInfoWindow(info string) *Window {
	var w Window

	w.elements = []WindowElement{
		NewPopupElement("Popup", info, POPUP_POS_CENTER, 100, 80, midGreyColor),
		NewButtonWithText("OKButton", int(ScreenWidth/2)-40, int(ScreenHeight/2)+15, "OK", nil),
	}

	w.Bind(SINGLE_CLICK_FUNC, w.GetArea([]string{"OKButton", "Area"}), g.KillWindowOnTop)

	return &w
}
