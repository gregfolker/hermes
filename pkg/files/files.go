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

func GenerateReadMe(file string, title string, author string) error {
	fmt.Printf("Creating %s...\n", file)

	contents := []byte("## Project: " + title +"\n" + "## Author: " + author + "\n" + templates.README_TEMPLATE)

	if err := ioutil.WriteFile(file, contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ANSI_GREEN + "Generated: " + colors.ANSI_RESET + "%s\n", file)
	}

	return nil
}

func GenerateTODO(file string, title string) error {
	fmt.Printf("Creating %s...\n", file)

	contents := []byte("## " + title + templates.TODO_TEMPLATE + "\n")

	if err := ioutil.WriteFile(file, contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ANSI_GREEN + "Generated: " + colors.ANSI_RESET + "%s\n", file)
	}

	return nil
}

func GenerateMain(file string, progName string, author string, language string) error {
	c := languages.LanguageToCommentStyle[strings.ToLower(language)]

	t, err := templates.GetMainTemplate(language)

	if err != nil {
		return err
	}

	contents := []byte(c + " Project: " + progName + "\n" + c + " Author: " + author + "\n" + t)

	fmt.Printf("Creating %s...\n", file + languages.LanguageToExtension[strings.ToLower(language)])

	if err := ioutil.WriteFile(file + languages.LanguageToExtension[strings.ToLower(language)], contents, os.FileMode(0644)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ANSI_GREEN + "Generated: " + colors.ANSI_RESET + "%s\n", file + languages.LanguageToExtension[strings.ToLower(language)])
	}

	return nil
}
