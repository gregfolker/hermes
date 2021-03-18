package ui

import (
	"fmt"
	"strconv"
	"errors"
	"os"
	"github.com/c-bata/go-prompt"
	"github.com/gregfolker/auto-project-builder/pkg/project"
)

const (
	PromptPrefix = ">> "
)

func GetUserInput(p *project.Project) error {
	fmt.Printf("What is the name of this project?\n")
	p.Name = getProjectName()

	fmt.Printf("Who is the primary author of this project?\n")
	p.Author = getProjectContributor()

	fmt.Printf("How many additional people will contribute this project?\n")

	contributorCount, err := strconv.Atoi(getNumberOfContributors())

	if err != nil {
		return err
	}

	// One contributor assumes the primary author is the only contributor, so
	// only retrieve additional input from the user if number of contributors > 1
	if contributorCount > 0 {
		for c := 0; c < contributorCount; c++ {
			fmt.Printf("What is the name of contributor #%v?\n", c + 1)
			p.Contributors = append(p.Contributors, getProjectContributor())
		}
	}

	fmt.Printf("What language will this project be written in?\n")
	p.Language = getProjectLanguage()

	fmt.Printf("Where do you want the new project directory?\n (Provide no input for default location)\n")
	p.Path = getProjectDirectory()

	return validateUserInput(p)
}

func validateUserInput(p *project.Project) error {
	if p.Name == "" {
		return errors.New("No project name provided")
	}

	if p.Author == "" {
		return errors.New("No primary author provided")
	}

	if p.Language == "" {
		return errors.New("No language provided")
	}

	if len(p.Contributors) > 1 {
		for c := 0; c < len(p.Contributors); c++ {
			if p.Contributors[c] == "" {
				return errors.New("No name provided for contributor #" + strconv.Itoa(c + 1) + "\n")
			}
		}
	}

	if p.Path != "" {
		if _, err := os.Stat(p.Path); os.IsNotExist(err) {
			return errors.New("Provided project path " + p.Path + " does not exist\n")
		}
	}

	return nil
}

func getProjectContributor() string {
	return prompt.Input(PromptPrefix, nameAutoCompleter)
}

func getProjectName() string {
	return prompt.Input(PromptPrefix, emptyAutoCompleter)
}

func getProjectLanguage() string {
	return prompt.Input(PromptPrefix, languageAutoCompleter)
}

func getNumberOfContributors() string {
	return prompt.Input(PromptPrefix, emptyAutoCompleter)
}

func getProjectDirectory() string {
	return prompt.Input(PromptPrefix, emptyAutoCompleter)
}

func nameAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: "Greg Folker"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func languageAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: "Go"},
		{Text: "C"},
		{Text: "Python"},
		{Text: "Rust"},
		{Text: "Bash"},
		{Text: "Perl"},
		{Text: "Java"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func emptyAutoCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest {
		{Text: ""},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
