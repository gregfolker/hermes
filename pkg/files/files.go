package files

import (
	"fmt"
	"path"
	"io/ioutil"
	"github.com/gregfolker/auto-project-builder/pkg/project"
)

// Common file extensions
const (
	GO_EXT = ".go"
	C_EXT = ".c"
	PYTHON_EXT = ".py"
	JAVA_EXT = ".java"
	RUST_EXT = ".rs"
	BASH_EXT = ".sh"
	PERL_EXT = ".pl"
	MARKDOWN_EXT = ".md"
)

const (
	SEP = "------------------"
)

func CreateREADME(prog *project.Project) error {
	f := path.Join(prog.Path, "README" + MARKDOWN_EXT)

	fmt.Printf("\nCreating %s...\n\n", f)

	contents := []byte(prog.Name + "\n" + SEP + "\n")

	if err := ioutil.WriteFile(f, contents, 0755); err != nil {
		return err
	} else {
		fmt.Printf("Generated %s\n", f)
	}

	return nil
}
