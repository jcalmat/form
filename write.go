package form

import (
	"io"
	"os"
)

func write(s string) {
	_, _ = io.WriteString(os.Stdout, s)
}

func clearLine() {
	write("\u001b[0K")
}

const (
	navigation_keys_message string = "\u001b[30;1mPress arrow keys or TAB to navigate, ESC to quit\u001b[0m"
)
