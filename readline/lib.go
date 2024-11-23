package readline

import (
	"fmt"
	"io"
	"github.com/chzyer/readline"
)

func loop(promt string, history_file string) {
	config := &readline.Config{
			Prompt:      promt,
			HistoryFile: history_file,
	}
	rl, err := readline.NewEx(config)
	if err != nil {
			fmt.Printf("Failed to create readline instance: %v\n", err)
			return
	}
	defer func(rl *readline.Instance) {
			err := rl.Close()
			if err != nil {
					fmt.Printf("Error closing \n")
			}
	}(rl)

	// Enable command history
	rl.HistoryEnable()

	for {
			input, err := rl.Readline()
			if err != nil {
					if err == io.EOF {
							break // Exit on EOF (Ctrl+D)
					}
					fmt.Printf("Error reading input: %v\n", err)
					continue
			}

			// Process input
			fmt.Printf("%s", input)

			// Save command to history
			if err := rl.SaveHistory(input); err != nil {
					fmt.Printf("Failed to save command to history: %v\n", err)
			}
	}
}