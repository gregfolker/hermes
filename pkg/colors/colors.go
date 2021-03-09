package colors

// ANSI Color Codes
const (
	ANSI_RESET = "\033[0m"
	ANSI_RED = "\033[31m"
	ANSI_GREEN = "\033[32m"
	ANSI_YELLOW = "\033[33m"
	ANSI_BLUE = "\033[34m"
	ANSI_PURPLE = "\033[35m"
	ANSI_CYAN = "\033[36m"
	ANSI_GRAY = "\033[37m"
	ANSI_WHITE = "\033[97m"
)

func ColorText(text string, color string) string {
	t := ""

	switch color {
		case ANSI_RED:
			t = red(text)
		case ANSI_GREEN:
			t = green(text)
		case ANSI_YELLOW:
			t = yellow(text)
		case ANSI_BLUE:
			t = blue(text)
		case ANSI_PURPLE:
			t = purple(text)
		case ANSI_CYAN:
			t = cyan(text)
		case ANSI_GRAY:
			t = gray(text)
		case ANSI_WHITE:
			t = white(text)
	}

	return t
}

func red(t string) string {
	return ANSI_RED + t + ANSI_RESET
}

func green(t string) string {
	return ANSI_GREEN + t + ANSI_RESET
}

func yellow(t string) string {
	return ANSI_YELLOW + t + ANSI_RESET
}

func blue(t string) string {
	return ANSI_BLUE + t + ANSI_RESET
}

func purple(t string) string {
	return ANSI_PURPLE + t + ANSI_RESET
}

func cyan(t string) string {
	return ANSI_CYAN + t + ANSI_RESET
}

func gray(t string) string {
	return ANSI_GRAY + t + ANSI_RESET
}

func white(t string) string {
	return ANSI_WHITE + t + ANSI_RESET
}
