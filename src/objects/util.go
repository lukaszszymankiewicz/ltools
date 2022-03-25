package objects

// returns bigger value from inputted two
func MaxVal(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// returns smaller value from inputted two
func MinVal(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
