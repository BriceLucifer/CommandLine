package main

import (
	//"fmt"

	"github.com/BriceLucifer/CommandLine/color"

	"github.com/BriceLucifer/CommandLine/loading"
	// "github.com/BriceLucifer/CommandLine/readline"
	// "github.com/BriceLucifer/CommandLine/loading"
)

func main() {
	// loading.Loading()
	//color.ColorPrint(color.Red, "hello world\n")
	//color.ColorPrint(color.BrightYellow, "I dont know %d\n", 12)
	// color.ClearScreen()
	//color.ColorPrint(color.BackgroundBlue, "background color\n")
	//color.ColorPrint(color.BackgroundBrightMagenta, "aaaaaaaaa %d\n", 12);
	// readline.Loop(">> ", "./history", "bash")
	//fmt.Println("\u001b[0m")
	loading.Timer(color.BrightMagenta)
}
