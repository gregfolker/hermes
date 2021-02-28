package project

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/internal/ui"
)

type Project struct {
	Name string
	Author string
	Language string
	Path string
}

func NewProject() *Project {
	p := &Project{}

	return p
}

func (prog *Project) CreateProject() {
	fmt.Printf("What is the name of this project?\n")
	prog.Name = ui.GetProjectName()

	fmt.Printf("Who is the author of this project?\n")
	prog.Author = ui.GetProjectAuthor()

	fmt.Printf("What language will this project be written in?\n")
	prog.Language = ui.GetProjectLanguage()
}
