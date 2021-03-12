package errorutil

import (
	"fmt"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
)

func PrintError(err error) {
	fmt.Printf(colors.ColorText("\nError: ", colors.ANSI_RED) + "%v\n\n", err)
}
