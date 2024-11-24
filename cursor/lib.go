package cursor

import (
	"fmt"
	"os"
)

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
		fmt.Printf("Invalid direction: %s. Allowed: UP, DOWN, RIGHT, LEFT, NEXT_LINE, PREV_LINE\n", direction)
		return
	}
	if steps <= 0 {
		fmt.Println("Steps must be greater than 0")
		return
	}
	fmt.Printf("%s[%d%s", ESC, steps, direction)
	os.Stdout.Sync()
}

// PositionSet moves the cursor to the specified row and column.
func PositionSet(row int, column int) {
	if row < 1 || column < 1 {
		fmt.Println("Row and column must be greater than 0")
		return
	}
	fmt.Printf("%s[%d;%d%s", ESC, row, column, SET_POSITION)
	os.Stdout.Sync()
}

// ColumnSet sets the cursor to the specified column.
func ColumnSet(columns int) {
	if columns < 1 {
		fmt.Println("Columns must be greater than 0")
		return
	}
	fmt.Printf("%s[%d%s", ESC, columns, SET_COLUMN)
	os.Stdout.Sync()
}

// ScreenClear clears the screen based on the mode (0, 1, 2).
func ScreenClear(n int) {
	if n < 0 || n > 2 {
		fmt.Println("Invalid value for clearing. Use 0 (right), 1 (left), or 2 (entire).")
		return
	}
	fmt.Printf("%s[%d%s", ESC, n, CLEAR_SCREEN)
	os.Stdout.Sync()
}

// LineClear clears the current line based on the mode (0, 1, 2).
func LineClear(n int) {
	if n < 0 || n > 2 {
		fmt.Println("Invalid value for clearing. Use 0 (right), 1 (left), or 2 (entire).")
		return
	}
	fmt.Printf("%s[%d%s", ESC, n, CLEAR_LINE)
	os.Stdout.Sync()
}

// ResetCursor resets the cursor to the home position (row 1, column 1).
func ResetCursor() {
	fmt.Printf("%s[H", ESC)
	os.Stdout.Sync()
}

// SaveCursor saves the current cursor position.
func SaveCursor() {
	fmt.Printf("%s[s", ESC) // Alternatively, ESC+"7"
	os.Stdout.Sync()
}

// RestoreCursor restores the cursor to the last saved position.
func RestoreCursor() {
	fmt.Printf("%s[u", ESC) // Alternatively, ESC+"8"
	os.Stdout.Sync()
}
