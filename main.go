package main

import (
	// "fmt"

	// "github.com/BriceLucifer/CommandLine/cursor"
	"github.com/BriceLucifer/CommandLine/loading"
	"github.com/BriceLucifer/CommandLine/readline"
)

func main() {
	// cursor.Move(cursor.NEXT_LINE, 3)
	// fmt.Printf("------------\n")
	// cursor.Move(cursor.PREV_LINE, 1)
	// fmt.Printf("hello\n")
	// fmt.Printf("xxxxxxxxxx\n--------\n-----xxxxxxddddddddd-\n")
	// cursor.Move(cursor.UP, 1)
	// cursor.Move(cursor.LEFT, 100)
	// fmt.Printf("testing\n")


	loading.LoadingBar()
	readline.Loop(">> ", "./history")
}
