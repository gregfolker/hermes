package main

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/pkg/project"
	"github.com/gregfolker/auto-project-builder/internal/ui"
)

const (
	appName = "Auto Project Builder"
)

func main() {
	fmt.Printf("%s\n\n", appName)

	prog := project.NewProject()

	if err := ui.GetUserInput(prog); err != nil {
		fmt.Printf("\nError: %v\n\n", err)
		return
	}

	fmt.Printf("\n%v\n\n", *prog)
}
