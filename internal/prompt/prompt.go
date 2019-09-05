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
	// initialize tint
	t := tint.Init()

	// create color functions
	green = t.SwatchRaw(tint.Green.Bold())
	white = t.SwatchRaw(tint.White.Bold())
	blue = t.SwatchRaw(tint.Blue.Bold())

	// get current directory
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	// split directory string because were only gonna
	// use the last two parts
	dirs := strings.Split(dir, "/")

	// get last and second to last directory strings
	previousDir := dirs[len(dirs)-2]
	currentDir := dirs[len(dirs)-1]

	// get hostname
	hostname, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
	}

	// get username
	username, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	// build the whole path string
	pathString := strings.Join([]string{"🐢 ", green(username.Username), green("@"), green(hostname), ": ", blue("~/"), blue(previousDir), blue("/"), blue(currentDir), white("$"), " "}, "")

	fmt.Print(pathString)
}
