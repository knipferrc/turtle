package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/printzero/tint"
)

var green func(text string) string

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("Please enter a path")
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	t := tint.Init()
	reader := bufio.NewReader(os.Stdin)

	green = t.SwatchRaw(tint.Green.Bold())

	for {
		dir, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		currentDir := strings.Split(dir, "/")
		curPath := currentDir[len(currentDir)-1]
		finalPath := green(curPath)

		fmt.Printf("🐢 %s ", finalPath)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
