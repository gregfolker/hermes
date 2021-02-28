package main

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/pkg/project"
)

const (
	appName = "Auto Project Builder"
)

func main() {
	fmt.Printf("%s\n\n", appName)

	prog := project.NewProject()

	prog.CreateProject()

	fmt.Println(*prog)
}
