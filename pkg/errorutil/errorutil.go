package errorutil

import (
	"fmt"
	"github.com/gregfolker/hermes/pkg/colors"
)

func PrintError(err error) {
	fmt.Printf(colors.ColorText("\nError: ", colors.ANSI_RED) + "%v\n\n", err)
}
