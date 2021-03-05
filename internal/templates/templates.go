package templates

import (
	"fmt"
	"errors"
	"strings"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
)

const (

README_TEMPLATE =
`

-----------------
### Overview

<Brief Project Description>

-----------------
### Installation
-----------------

<Installation Instructions>

### Usage
-----------------

<Brief Usage Overview>

### Reporting Issues
-----------------

<Instructions on how to submit bug reports>


`

TODO_TEMPLATE =
`

## TODO:
	-
	-
	-
`

GO_MAIN_TEMPLATE =
`
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, World!\n")
}

`

PY_MAIN_TEMPLATE =
`
import sys
import os

def main(args):
	print("Hello, World!\n")

if __name__ == 'main':
	try:
		sys.exit(main(sys.argv))
	except:
		sys.stderr.write("ABORTED BY KEYBOARD INTERRUPT")
		sys.exit(1)
`

C_MAIN_TEMPLATE =
`
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
	printf("Hello, World!\n");
	return 0;
}
`

JAVA_MAIN_TEMPLATE =
`
class MainClass {
	public static void main(String[] args) {
		System.out.println("Hello, World!")
	}
}
`

RUST_MAIN_TEMPLATE =
`
fn main() {
	println!("Hello, World!");
}
`

)

func GetMainTemplate(language string) (string, error) {
	var err error
	t := ""

	switch strings.ToLower(language) {
	case "go":
		t = GO_MAIN_TEMPLATE
		err = nil
	case "c":
		t = C_MAIN_TEMPLATE
		err = nil
	case "python":
		t = PY_MAIN_TEMPLATE
		err = nil
	case "java":
		t = JAVA_MAIN_TEMPLATE
		err = nil
	case "rust":
		t = RUST_MAIN_TEMPLATE
		err = nil
	case "bash", "perl":
		// Bash and Perl do not require main program files, so it is not an error
		// Do nothing in this case, returning a blank template and no error
		fmt.Printf(colors.ANSI_YELLOW + "Warning: " + colors.ANSI_RESET + "%s does not have a main file template, skipping...\n", language)
	default:
		err = errors.New("Unknown language " + language + ", unable to retrieve main file template\n")
	}

	return t, err
}
