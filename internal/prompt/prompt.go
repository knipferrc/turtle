package prompt

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/printzero/tint"
)

var green func(text string) string
var white func(text string) string
var blue func(text string) string

// BuildPrompt builds the users prompt
func BuildPrompt() {
	t := tint.Init()

	green = t.SwatchRaw(tint.Green.Bold())
	white = t.SwatchRaw(tint.White.Bold())
	blue = t.SwatchRaw(tint.Blue.Bold())

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dirs := strings.Split(dir, "/")

	previousDir := dirs[len(dirs)-2]
	currentDir := dirs[len(dirs)-1]

	hostname, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	username, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	pathString := strings.Join([]string{"🐢 ", green(username.Username), green("@"), green(hostname), ": ", blue("~/"), blue(previousDir), blue("/"), blue(currentDir), white("$"), " "}, "")

	fmt.Print(pathString)
}
