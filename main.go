package main

import (
	"fmt"

	"github.com/BriceLucifer/CommandLine/cursor"
)

func main() {
	cursor.Move(cursor.NEXT_LINE, 3)
	fmt.Printf("------------\n")
	cursor.Move(cursor.PREV_LINE, 1)
	fmt.Printf("hello\n")
	cursor.Move(cursor.DOWN, 4)
}
