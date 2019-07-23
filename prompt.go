package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/printzero/tint"
)

var green func(text string) string
var white func(text string) string

func BuildPrompt() {
	t := tint.Init()

	green = t.SwatchRaw(tint.Blue.Bold())
	white = t.SwatchRaw(tint.White.Bold())

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dirs := strings.Split(dir, "/")

	previousDir := dirs[len(dirs)-2]
	currentDir := dirs[len(dirs)-1]

	symbol := white("$")

	pathString := strings.Join([]string{"🐢 ", "~/", previousDir, "/", currentDir, symbol, " "}, "")

	finalPath := green(pathString)

	fmt.Print(finalPath)
}
