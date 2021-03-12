package languages

import (
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
	"github.com/gregfolker/auto-project-builder/pkg/colors"
)

var LanguageToExtension = map[string]string {
	GOLANG: GO_EXT,
	C: C_EXT,
	PYTHON: PYTHON_EXT,
	JAVA: JAVA_EXT,
	RUST: RUST_EXT,
	BASH: BASH_EXT,
	PERL: PERL_EXT,
	MARKDOWN: MARKDOWN_EXT,
}

var LanguageToCommentStyle = map[string]string {
	GOLANG: GO_COMMENT_LEADER,
	C: C_COMMENT_LEADER,
	PYTHON: PYTHON_COMMENT_LEADER,
	JAVA: JAVA_COMMENT_LEADER,
	RUST: RUST_COMMENT_LEADER,
	BASH: BASH_COMMENT_LEADER,
	PERL: PERL_COMMENT_LEADER,
}

// Supported Languages
const (
	GOLANG = "go"
	C = "c"
	PYTHON = "python"
	JAVA = "java"
	RUST = "rust"
	BASH = "bash"
	PERL = "perl"
	MARKDOWN = "markdown"
)

// File Extensions
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

// Language Comment Styles
const (
	GO_COMMENT_LEADER = "//"
	C_COMMENT_LEADER = "//"
	PYTHON_COMMENT_LEADER = "#"
	JAVA_COMMENT_LEADER = "//"
	RUST_COMMENT_LEADER = "//"
	BASH_COMMENT_LEADER = "#"
	PERL_COMMENT_LEADER = "#"
)

func InitGolangDirs(p string) error {
	if err := os.MkdirAll(filepath.Join(p, "pkg"), os.FileMode(0755)); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(p, "internal"), os.FileMode(0755)); err != nil {
		return err
	}

	fmt.Printf(colors.ColorText("Generated ", colors.ANSI_GREEN) + "project structure for %s\n", p)

	return nil
}

func InitCDirs(p string) error {
	if err := os.MkdirAll(filepath.Join(p, "src"), os.FileMode(0755)); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(p, "inc"), os.FileMode(0755)); err != nil {
		return err
	}

	fmt.Printf(colors.ColorText("Generated ", colors.ANSI_GREEN) + "project structure for %s\n", p)

	return nil
}

func InitJavaDirs(p string) error {
	var cmd *exec.Cmd

	// Verify Maven is installed on the system
	cmd = exec.Command("sh", "-c", "mvn --version")

	appName := filepath.Base(p)

	if err := cmd.Run(); err != nil {
		return err
	}

	// Use Maven to generate the standard project structure for Java programs in the new project directory
	cmd = exec.Command("sh", "-c", "mvn archetype:generate -DgroupId=com.mycompany.app -DartifactId=" + appName + " -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false")
	cmd.Dir = p

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf(colors.ColorText("Generated ", colors.ANSI_GREEN) + "project structure for %s\n", p)

	return nil
}
