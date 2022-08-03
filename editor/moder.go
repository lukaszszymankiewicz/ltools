package editor

import ()

type Mode struct {
	modes []int
}

func NewMode(i int) Mode {
	var m Mode

	m.modes = make([]int, i)
	m.Activate(0)

	return m
}

func (m *Mode) Clear() {
	for n := 0; n < len(m.modes); n++ {
		m.modes[n] = 0
	}
}

func (m *Mode) Add(i int) {
	m.modes[i] = 1
}

func (m *Mode) Activate(i int) {
	for n := 0; n < len(m.modes); n++ {
		m.modes[n] = 0
	}
	m.modes[i] = 1
}

func (m *Mode) State(i int) *int {
	return &m.modes[i]
}

func (m *Mode) Active() int {
	for n := 0; n < len(m.modes); n++ {
		if m.modes[n] == 1 {
			return n
		}
	}
	return 0
}
