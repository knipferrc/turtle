package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		BuildPrompt()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = ExecInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
