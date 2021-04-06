package main

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/pkg/project"
	"github.com/gregfolker/auto-project-builder/pkg/git"
	"github.com/gregfolker/auto-project-builder/internal/ui"
	"github.com/gregfolker/auto-project-builder/internal/version"
	"github.com/gregfolker/auto-project-builder/pkg/errorutil"
)

const (
	appName = "Auto Project Builder"
)

func main() {
	fmt.Printf("%s ", appName)
	version.PrintVersion()
	fmt.Printf("\n\n")

	prog := project.NewProject()

	if err := ui.GetUserInput(prog); err != nil {
		errorutil.PrintError(err)
		return
	}

	fmt.Printf("\nGenerating project...\n\n")

	if err := prog.CreateNewProjectDir(); err != nil {
		errorutil.PrintError(err)
		return
	}

	if err := prog.CreateNewProjectSubDirs(); err != nil {
		errorutil.PrintError(err)
		return
	}

	if err := prog.CreateProjectFile("README"); err != nil {
		errorutil.PrintError(err)
		return
	}

	if err := prog.CreateProjectFile("TODO"); err != nil {
		errorutil.PrintError(err)
		return
	}

	if err := prog.CreateProjectFile("main"); err != nil {
		errorutil.PrintError(err)
		return
	}

	if err := git.Init(prog.Path); err != nil {
		errorutil.PrintError(err)
		return
	}

	prog.PrintProjectInfo()
	fmt.Printf("\nDone!\n")
}
