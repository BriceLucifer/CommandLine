package color

import (
	"fmt"
)

/// Color has both backgroud and font color

const (
	Black                   = "\u001b[30m"
	Red                     = "\u001b[31m"
	Green                   = "\u001b[32m"
	Yellow                  = "\u001b[33m"
	Blue                    = "\u001b[34m"
	Magenta                 = "\u001b[35m"
	Cyan                    = "\u001b[36m"
	White                   = "\u001b[37m"
	BrightBlack             = "\u001b[30;1m"
	BrightRed               = "\u001b[31;1m"
	BrightGreen             = "\u001b[32;1m"
	BrightYellow            = "\u001b[33;1m"
	BrightBlue              = "\u001b[34;1m"
	BrightMagenta           = "\u001b[35;1m"
	BrightCyan              = "\u001b[36;1m"
	BrightWhite             = "\u001b[37;1m"
	BackgroundBlack         = "\u001b[40m"
	BackgroundRed           = "\u001b[41m"
	BackgroundGreen         = "\u001b[42m"
	BackgroundYellow        = "\u001b[43m"
	BackgroundBlue          = "\u001b[44m"
	BackgroundMagenta       = "\u001b[45m"
	BackgroundCyan          = "\u001b[46m"
	BackgroundWhite         = "\u001b[47m"
	BackgroundBrightBlack   = "\u001b[40;1m"
	BackgroundBrightRed     = "\u001b[41;1m"
	BackgroundBrightGreen   = "\u001b[42;1m"
	BackgroundBrightYellow  = "\u001b[43;1m"
	BackgroundBrightBlue    = "\u001b[44;1m"
	BackgroundBrightMagenta = "\u001b[45;1m"
	BackgroundBrightCyan    = "\u001b[46;1m"
	BackgroundBrightWhite   = "\u001b[47;1m"
	Reset                   = "\u001b[0m"
)

type Color string

/// Color Print
func ColorPrint(color Color, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s%s", color, message, Reset)
}

/// Clear the Screen
func ClearScreen() {
	fmt.Printf("\033[H\033[2J")
}
