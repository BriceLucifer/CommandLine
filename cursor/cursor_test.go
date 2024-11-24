package cursor

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to capture stdout
func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

// Test cursor movement
func TestMove(t *testing.T) {
	output := captureOutput(func() {
		Move(UP, 3)
	})
	expected := "\u001b[3A"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	output = captureOutput(func() {
		Move("INVALID", 5)
	})
	expected = "Invalid direction: INVALID\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test position setting
func TestPositionSet(t *testing.T) {
	output := captureOutput(func() {
		PositionSet(5, 10)
	})
	expected := "\u001b[5;10H"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test column setting
func TestColumnSet(t *testing.T) {
	output := captureOutput(func() {
		ColumnSet(15)
	})
	expected := "\u001b[15G"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test screen clearing
func TestScreenClear(t *testing.T) {
	output := captureOutput(func() {
		ScreenClear(0)
	})
	expected := "\u001b[0J"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	output = captureOutput(func() {
		ScreenClear(-1)
	})
	expected = "\u001b[2J"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	output = captureOutput(func() {
		ScreenClear(3)
	})
	expected = "\u001b[2J"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test line clearing
func TestLineClear(t *testing.T) {
	output := captureOutput(func() {
		LineClear(1)
	})
	expected := "\u001b[1K"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	output = captureOutput(func() {
		LineClear(-1)
	})
	expected = "\u001b[2K"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	output = captureOutput(func() {
		LineClear(3)
	})
	expected = "\u001b[2K"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test reset cursor
func TestResetCursor(t *testing.T) {
	output := captureOutput(func() {
		ResetCursor()
	})
	expected := "\u001b[H"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test save and restore cursor
func TestSaveAndRestoreCursor(t *testing.T) {
	output := captureOutput(func() {
		SaveCursor()
	})
	expected := "\u001b7"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

	output = captureOutput(func() {
		RestoreCursor()
	})
	expected = "\u001b8"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
