package ui

import (
	"fmt"
	"strconv"
	"errors"
	"os"
   "strings"
	"github.com/c-bata/go-prompt"
	"github.com/gregfolker/hermes/pkg/project"
	"github.com/gregfolker/hermes/pkg/languages"
)

const (
	PromptPrefix = ">> "
)

func GetUserInput(p *project.Project) error {
   var answer string

	fmt.Printf("What is the name of this project?\n")
	p.Name = getProjectName()

	fmt.Printf("Who is the primary author of this project?\n")
	p.Author = getProjectContributor()

	fmt.Printf("How many additional people will contribute this project?\n (Default: 0)\n")

	contributors := getNumberOfContributors()

	if contributors == "" {
		contributors = "0"
	}

	contributorCount, err := strconv.Atoi(contributors)

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

   if strings.ToLower(p.Language) == languages.C {
      // For C programs, ask if this will be a kernel module as
      // the template for kmods is different from generic C programs
      fmt.Printf("Is this project going to be a kernel module?\n")
      answer = getYesNoAnswer()

      if strings.ToLower(answer) == "yes" {
         p.IsKmod = true;
      } else {
         p.IsKmod = false;
      }
   } else {
      p.IsKmod = false;
   }

	fmt.Printf("Where do you want the new project directory?\n (Provide no input for default location)\n")
	p.Path = getProjectDirectory()

   fmt.Printf("Generate a template README.md?\n (Default: No)\n")
   answer = getYesNoAnswer()

   if strings.ToLower(answer) == "yes" {
      p.ReadmeTemplate = true
   } else {
      p.ReadmeTemplate = false
   }

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

func getYesNoAnswer() string {
   return prompt.Input(PromptPrefix, yesNoAutoCompleter)
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

func yesNoAutoCompleter(d prompt.Document) []prompt.Suggest {
   s := []prompt.Suggest {
      {Text: "No"},
      {Text: "Yes"},
   }

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
