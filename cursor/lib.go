package cursor

import "fmt"

// Common ESC character
const ESC = "\u001b"

// Directions: up, down, right, left
const (
	UP    = "A"
	DOWN  = "B"
	RIGHT = "C"
	LEFT  = "D"
)

// Line move: next, prev
const (
	NEXT_LINE = "E"
	PREV_LINE = "F"
)

// Set column and position
const (
	SET_COLUMN   = "G"
	SET_POSITION = "H"
)

// Clear operations
const (
	CLEAR_SCREEN = "J"
	CLEAR_LINE   = "K"
)

// Cursor movement function
func Move(direction string, steps int) {
	validDirections := map[string]bool{
		UP: true, DOWN: true, RIGHT: true, LEFT: true,
		NEXT_LINE: true, PREV_LINE: true,
	}
	if !validDirections[direction] {
		fmt.Printf("Invalid direction: %s\n", direction)
		return
	}
	if steps < 0 {
		steps = 0
	}
	fmt.Printf("%s[%d%s", ESC, steps, direction)
}

// Position set: moves cursor to row n and column m
func PositionSet(row int, column int) {
	position := fmt.Sprintf("%s[%d;%d%s", ESC, row, column, SET_POSITION)
	fmt.Printf(position)
}

// Column set
func ColumnSet(columns int) {
	column := fmt.Sprintf("%s[%d%s", ESC, columns, SET_COLUMN)
	fmt.Printf(column)
}

// Clear screen
func ScreenClear(n int) {
	if n < 0 || n > 2 {
		n = 2
	}
	fmt.Printf("%s[%d%s", ESC, n, CLEAR_SCREEN)
}

// Clear line
func LineClear(n int) {
	if n < 0 || n > 2 {
		n = 2
	}
	fmt.Printf("%s[%d%s", ESC, n, CLEAR_LINE)
}

// Reset cursor to home
func ResetCursor() {
	fmt.Printf("%s[H", ESC)
}

// Save and restore cursor position
func SaveCursor() {
	fmt.Printf("%s7", ESC)
}

func RestoreCursor() {
	fmt.Printf("%s8", ESC)
}
