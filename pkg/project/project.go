package project

import (
	"os"
	"path"
	"strings"
	"fmt"
	"github.com/pkg/errors"
	"github.com/gregfolker/hermes/pkg/files"
	"github.com/gregfolker/hermes/pkg/languages"
	"github.com/gregfolker/hermes/pkg/colors"
)

type Project struct {
	Name string
	Author string
	Language string
	Path string
	Contributors []string
   ReadmeTemplate bool
   IsKmod         bool        // Is this project for a kernel module? Only applies to C programs
}

func NewProject() *Project {
	p := &Project{}

	return p
}

func (prog Project) PrintProjectInfo() {
	fmt.Printf("\n--- New Project Details ---\n")
	fmt.Printf("Project Name: %s\n", prog.Name)
	fmt.Printf("Primary Author: %s\n", prog.Author)

	if len(prog.Contributors) > 0 {
		fmt.Printf("Additional Contributor(s):")
		for c := 0; c < len(prog.Contributors); c++ {
			fmt.Printf(" %s", prog.Contributors[c])
			if c + 1 != len(prog.Contributors) {
				fmt.Printf(",")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Language: %s\n", prog.Language)
	fmt.Printf("Directory: %s\n", prog.Path)
}

func (prog *Project) CreateNewProjectDir() error {
	var dir string

	// If the user specified an alternate location, just use that directly with the project name
	if prog.Path != "" {

		// Avoid creating directories with spaces
		if strings.Contains(prog.Name, " ") {
			dir = strings.ReplaceAll(strings.ToLower(prog.Name), " ", "-")
		} else {
			dir = strings.ToLower(prog.Name)
		}

		prog.Path = path.Join(prog.Path, dir)

		fmt.Printf("Creating %s...\n", prog.Path)

		if err := os.MkdirAll(prog.Path, os.FileMode(0755)); err != nil {
			return errors.Wrap(err, "Failed to create project directory " + prog.Path)
		}

		return nil
	}

	userHome, err := os.UserHomeDir()

	if err != nil {
		return errors.Wrap(err, "Unable to find user home directory")
	}

	// Avoid creating directories with spaces
	if strings.Contains(prog.Name, " ") {
		dir = strings.ReplaceAll(strings.ToLower(prog.Name), " ", "-")
	} else {
		dir = strings.ToLower(prog.Name)
	}

	prog.Path = path.Join(userHome, strings.ToLower(prog.Language), "src", dir)

	fmt.Printf("Creating %s...\n", prog.Path)

	if err := os.MkdirAll(prog.Path, os.FileMode(0755)); err != nil {
		return errors.Wrap(err, "Failed to create project directory " + prog.Path)
	}

	fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", prog.Path)

	return nil
}

func (prog *Project) CreateNewProjectSubDirs() error {
	var err error

	if prog.Path == "" {
		return errors.New("Cannot create project subdirectory structure; Project path does not exist\n")
	}

	fmt.Printf("Initializing project structure...\n")

	switch strings.ToLower(prog.Language) {
	case languages.GOLANG:
		err = languages.InitGolangDirs(prog.Path)
	case languages.C:
		err = languages.InitCDirs(prog.Path)
	case languages.JAVA:
		err = languages.InitJavaDirs(prog.Path)
	case languages.RUST:
		err = languages.InitRustDirs(prog.Path)
	default:
		fmt.Printf("Nothing additional to create for %s project...\n", prog.Language)
		err = nil
	}

	return err
}

func (prog *Project) CreateProjectFile(filename string) error {
	if prog.Path == "" {
		return errors.New("Project directory does not exist\n")
	}

	f := path.Join(prog.Path, filename)

	switch filename {
	case "README":
      if prog.ReadmeTemplate {
		   return files.GenerateTemplateReadMe(f + languages.MARKDOWN_EXT, prog.Name, prog.Author, prog.Contributors)
      } else {
         return files.GenerateBlankReadMe(f + languages.MARKDOWN_EXT, prog.Name)
      }
	case "TODO":
		return files.GenerateTODO(f + languages.MARKDOWN_EXT, prog.Name)
	case "main":
		// For Rust and C programs, put main in the src/ subdirectory
		if strings.ToLower(prog.Language) == languages.C || strings.ToLower(prog.Language) == languages.RUST {
			f = path.Join(prog.Path, "src", filename)
		}

		return files.GenerateMain(f, prog.Name, prog.Author, prog.Contributors, prog.Language, prog.IsKmod)
   case "Makefile":
      // Only create a makefile for C programs
      if strings.ToLower(prog.Language) != languages.C {
         fmt.Printf("INFO: Makefile not required for %s program. Ignoring...\n", prog.Language)
         return nil
      }

      return files.GenerateMakefile(f, prog.Name, prog.Author, prog.Contributors, prog.Language)
	default:
		return errors.New("Unknown file " + filename + ", unable to create\n")
	}

	return nil
}
