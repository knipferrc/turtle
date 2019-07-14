package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/printzero/tint"
)

var green func(text string) string

func BuildPrompt() {
	t := tint.Init()

	green = t.SwatchRaw(tint.Green.Bold())

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	currentDir := strings.Split(dir, "/")
	curPath := currentDir[len(currentDir)-1]
	finalPath := green(curPath)

	fmt.Printf("🐢 %s ", finalPath)
}
