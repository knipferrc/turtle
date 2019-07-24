package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/knipferrc/turtle/internal/cli"
	"github.com/knipferrc/turtle/internal/prompt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		prompt.BuildPrompt()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = cli.ExecInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
