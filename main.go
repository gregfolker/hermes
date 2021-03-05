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

	if err := prog.CreateNewProjectDir(); err != nil {
		fmt.Printf("\nError: %v\n\n", err)
		return
	}

	if err := prog.CreateProjectFile("README"); err != nil {
		fmt.Printf("\nError: %v\n\n", err)
		return
	}
}
