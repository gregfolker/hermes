package files

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"github.com/gregfolker/auto-project-builder/internal/templates"
	"github.com/gregfolker/auto-project-builder/pkg/languages"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
)

func GenerateBlankReadMe(file string, title string) error {
	fmt.Printf("Creating %s...\n", file)

	contents := []byte("## " + title + "\n")

	if err := ioutil.WriteFile(file, contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", file)
	}

   return nil
}

func GenerateTemplateReadMe(file string, title string, author string, contributors []string) error {
	fmt.Printf("Creating %s...\n", file)

	contents := buildReadMeTemplate(title, author, contributors)

	if err := ioutil.WriteFile(file, contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", file)
	}

	return nil
}

func GenerateTODO(file string, title string) error {
	fmt.Printf("Creating %s...\n", file)

	contents := []byte("## " + title + templates.TODO_TEMPLATE + "\n")

	if err := ioutil.WriteFile(file, contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", file)
	}

	return nil
}

func GenerateMain(file string, title string, author string, contributors []string, language string, isKmod bool) error {
	if strings.ToLower(language) == languages.JAVA {
		// The files for Java will be created via Maven when initializing the project directory structure, so just quietly return here
		// instead of using the existing template for main.java
		return nil
	}

	fmt.Printf("Creating %s...\n", file + languages.LanguageToExtension[strings.ToLower(language)])

	contents := buildMainContents(title, author, contributors, language, isKmod)

	if err := ioutil.WriteFile(file + languages.LanguageToExtension[strings.ToLower(language)], contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", file + languages.LanguageToExtension[strings.ToLower(language)])
	}

	return nil
}

func GenerateMakefile(file string, title string, author string, contributors []string, language string) error {
   // Only C programs should get a makefile
   if strings.ToLower(language) != strings.ToLower(languages.C) {
      return nil
   }

	fmt.Printf("Creating %s...\n", file)

   contents := buildMakefileContents(title, author, contributors, language)

	if err := ioutil.WriteFile(file, contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "%s\n", file)
	}

   return nil
}

func buildMakefileContents(title string, author string, contributors []string, language string) []byte {
	var titleLine string
	var authorLine string
	var contributorsLine string
	var contents []byte

	c := languages.MAKE_COMMENT_LEADER

	titleLine = c + " Project: " + title + "\n"
	authorLine = c + " Author: " + author + "\n"

	if len(contributors) > 0 {
		contributorsLine = c + " Contributors:"
		for i := 0; i < len(contributors); i++ {
			contributorsLine = contributorsLine + " " + contributors[i]
			if i + 1 != len(contributors) {
				contributorsLine = contributorsLine + ","
			}
		}
		contributorsLine = contributorsLine + "\n"
   }

	t := templates.MAKEFILE_TEMPLATE

	contents = []byte(titleLine + authorLine + contributorsLine + t)

   return contents
}

func buildReadMeTemplate(title string, author string, contributors []string) []byte {
	var titleLine string
	var authorLine string
	var contributorsLine string
	var contents []byte

	titleLine = "## Project: " + title + "\n"
	authorLine = "## Author: " + author + "\n"

	if len(contributors) > 0 {
		contributorsLine = "### Contributors:"
		for i := 0; i < len(contributors); i++ {
			contributorsLine = contributorsLine + " " + contributors[i]
			if i + 1 != len(contributors) {
				contributorsLine = contributorsLine + ","
			}
		}
		contributorsLine = contributorsLine + "\n"
	}

	contents = []byte(titleLine + authorLine + contributorsLine + templates.README_TEMPLATE)

	return contents
}

func buildMainContents(title string, author string, contributors []string, language string, isKmod bool) []byte {
	var titleLine string
	var authorLine string
	var contributorsLine string
	var contents []byte

	c := languages.LanguageToCommentStyle[strings.ToLower(language)]

	titleLine = c + " Project: " + title + "\n"
	authorLine = c + " Author: " + author + "\n"

	if len(contributors) > 0 {
		contributorsLine = c + " Contributors:"
		for i := 0; i < len(contributors); i++ {
			contributorsLine = contributorsLine + " " + contributors[i]
			if i + 1 != len(contributors) {
				contributorsLine = contributorsLine + ","
			}
		}
		contributorsLine = contributorsLine + "\n"
	}

	t, err := templates.GetMainTemplate(language, isKmod)

	if err != nil {
		return contents
	}

	contents = []byte(titleLine + authorLine + contributorsLine + t)

	return contents
}
