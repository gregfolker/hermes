package git

import (
	"fmt"
	"errors"
	"os"
	"os/exec"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
)

func InitializeNewRepo(path string) error {
	fmt.Printf("Creating git repository...\n")

	if path == "" {
		return errors.New("Cannot initialize git repo; provided path is empty\n")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("Cannot initialize git repo; " + path + " does not exist\n")
	}

	cmd := exec.Command("sh", "-c", "git init " + path)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Creating git repository " + colors.ANSI_RED + "failed\n")
		return err
	} else {
		fmt.Printf(colors.ANSI_GREEN + "Generated " + colors.ANSI_RESET + "new git repository at %s\n", path)
	}

	return nil
}
