package project

import (
	"os"
	"path"
	"strings"
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

func (prog *Project) CreateNewProjectDir() error {
	userHome, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	prog.Path = path.Join(userHome, strings.ToLower(prog.Language), "src", strings.ToLower(prog.Name))

	return os.MkdirAll(prog.Path, os.FileMode(0755))
}

