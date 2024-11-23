package main

import (
	"github.com/BriceLucifer/CommandLine/color"
	// "github.com/BriceLucifer/CommandLine/loading"
)

func main() {
	// loading.Loading()
	color.ColorPrint(color.Red, "hello world\n")
	color.ColorPrint(color.BrightYellow, "I dont know %d\n", 12)
	// color.ClearScreen()
	color.ColorPrint(color.BackgroundBlue, "background color")
	color.ColorPrint(color.BackgroundBrightMagenta, "aaaaaaaaa %d\n", 12);
}
