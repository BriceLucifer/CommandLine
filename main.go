package main

import (
	"github.com/BriceLucifer/CommandLine/color"
	// "github.com/BriceLucifer/CommandLine/loading"
)

func main() {
	// loading.Loading()
	color.ColorPrint(color.Red, "hello world\n")
	color.ColorPrintln(color.BrightYellow, "I dont Know %d", 12)
	color.ClearScreen()
}
