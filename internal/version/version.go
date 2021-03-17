package version

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/pkg/errors"
	"github.com/gregfolker/auto-project-builder/pkg/git"
)

func PrintVersion() {
	v, err := GetVersion()

	if err != nil {
		return
	}

	fmt.Printf("%s", v)
}

func GetVersion() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", errors.Wrap(err, "Unable to find user's home directory")
	}

	installDir := filepath.Join(home, "go", "src", "auto-project-builder")

	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		return "", errors.Wrap(err, "Unable to find install directory for auto-project-builder")
	}

	return git.DescribeTags(filepath.Join(home, "go", "src", "auto-project-builder"))
}
