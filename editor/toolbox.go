package editor

import ()

const (
	PENCIL_TOOL_IDX = iota
	RUBBER_TOOL_IDX
	N_TOOLS
)

type Tool struct {
	current_brush int
	brushes       []*Brush
}

type Toolbox struct {
	tools map[int]*Tool
	Mode
}

func (t *Tool) GetActiveBrush() *Brush {
	return t.brushes[t.current_brush]
}

func NewTool(brushes []*Brush) Tool {
	var t Tool

	t.brushes = brushes

	return t
}

var PenTool Tool = NewTool([]*Brush{&PenBrush})
var RubberTool Tool = NewTool([]*Brush{&RubberBrush})

func NewToolbox() Toolbox {
	var tb Toolbox

	tb.tools = make(map[int]*Tool, N_TOOLS)
	tb.Mode = NewMode(N_TOOLS)

	tb.tools[PENCIL_TOOL_IDX] = &PenTool
	tb.tools[RUBBER_TOOL_IDX] = &RubberTool

	tb.Mode.Activate(PENCIL_TOOL_IDX)

	return tb
}

func (tb *Toolbox) GetTool(i int) *Tool {
	return tb.tools[i]
}

func (tb *Toolbox) GetActiveTool() *Tool {
	return tb.tools[tb.Mode.Active()]
}
