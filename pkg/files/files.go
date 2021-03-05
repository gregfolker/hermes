package files

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/gregfolker/auto-project-builder/internal/templates"
)

func GenerateReadMe(readme string, title string, author string) error {
	fmt.Printf("Creating %s...\n", readme)

	contents := []byte("## " + title + templates.README_TEMPLATE + "\nAuthor: "+ author + "\n")

	if err := ioutil.WriteFile(readme, contents, os.FileMode(0755)); err != nil {
		return err
	} else {
		fmt.Printf("Generated %s\n", readme)
	}

	return nil
}
