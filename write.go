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
	navigation_keys_message string = "\u001b[30;1mPress arrow keys or TAB to navigate, ESC to save and quit\u001b[0m"
	done_button             string = "\u001b[30;1mDone\u001b[0m"
)
