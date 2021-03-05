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

func GenerateReadMe(readme string, title string, author string) error {
	fmt.Printf("Creating %s...\n", readme)

	contents := []byte("## " + title + templates.README_TEMPLATE + "\nAuthor: "+ author + "\n")

	if err := ioutil.WriteFile(readme, contents, os.FileMode(0755)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ANSI_GREEN + "Generated: " + colors.ANSI_RESET + "%s\n", readme)
	}

	return nil
}

func GenerateMain(file string, progName string, author string, language string) error {
	fmt.Printf("Creating %s...\n", file)

	c := languages.LanguageToCommentStyle[strings.ToLower(language)]

	t, err := templates.GetMainTemplate(language)

	if err != nil {
		return err
	}

	contents := []byte(c + " Project: " + progName + "\n" + c + " Author: " + author + "\n" + t)

	if err := ioutil.WriteFile(file + languages.LanguageToExtension[strings.ToLower(language)], contents, os.FileMode(0755)); err != nil {
		return err
	} else {
		fmt.Printf(colors.ANSI_GREEN + "Generated: " + colors.ANSI_RESET + "%s\n", file + languages.LanguageToExtension[strings.ToLower(language)])
	}

	return nil
}
