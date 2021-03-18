package version

import (
	"fmt"
	"os"
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
	installDir := os.Getenv("PBLD_INSTALL_DIR")

	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		return "", errors.Wrap(err, "Unable to find install directory for auto-project-builder")
	}

	return git.DescribeTags(installDir)
}
