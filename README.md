## Auto Project Builder
### Author: Greg Folker

-----------------
### Overview

Hermes is a simple Go program designed to auto-generate the necessary and required components for new software projects
using an interactive prompt for user inputs

Languages currently supported:
   - Go
   - C
   - Rust
   - Java

-----------------
### Installation

Clone this repository

`
$ git clone https://github.com/gregfolker/hermes.git
`

Run `make` and `./install.sh`

`
$ make && ./install.sh
`

Currently, version info is being parsed from the git info of the install dir. The install script
sets the environment variable `PBLD_INSTALL_DIR` when run but this will need to exist in future
shell sessions to ensure the program is still able to run successfully

Putting the following in `.bashrc` or similar is encouraged

`
export PBLD_INSTALL_DIR=/local/path/to/this/repo
`

-----------------
### Usage

Run `hermes` to start the interactive prompt to enter the project name, author, and primary language it will be written in


`tab` and arrow-keys can be used to auto-complete in the menu for the programming language

-----------------
### Reporting Issues

Issues are tracked on [Github](https://github.com/gregfolker/autoprojectbuilder/issues)
