package templates

const (

README_TEMPLATE =
`

-----------------
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
