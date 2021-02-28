package main

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/internal/ui"
)

const (
	appName = "Auto Project Builder"
)

func main() {
	fmt.Printf("%s\n\n", appName)

	run()
}

func run() {
	fmt.Printf("What is the name of this project?\n")
	fmt.Printf("%s\n\n", ui.GetProjectName())

	fmt.Printf("Who is the author of this project?\n")
	fmt.Printf("%s\n\n", ui.GetProjectAuthor())

	fmt.Printf("What language will this project be written in?\n")
	fmt.Printf("%s\n\n", ui.GetProjectLanguage())
}
