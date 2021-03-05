package languages

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
