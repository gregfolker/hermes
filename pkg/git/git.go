package git

import (
	"fmt"
	"errors"
	"os"
	"os/exec"
	"github.com/gregfolker/hermes/pkg/colors"
)

func Init(path string) error {
	fmt.Printf("Creating git repository...\n")

	if path == "" {
		return errors.New("Cannot initialize git repo; provided path is empty\n")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("Cannot initialize git repo; " + path + " does not exist\n")
	}

	cmd := exec.Command("sh", "-c", "git init " + path)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Creating git repository " + colors.ColorText("failed", colors.ANSI_RED) + "\n")
		return err
	} else {
		fmt.Printf(colors.ColorText("Generated: ", colors.ANSI_GREEN) + "Generated " + colors.ANSI_RESET + "new git repository at %s\n", path)
	}

	return nil
}

func DescribeTags(path string) (string, error) {
	cmd := &exec.Cmd{}

	cmd.Dir = path

	out, err := exec.Command("sh", "-c", "git describe --tags").Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
