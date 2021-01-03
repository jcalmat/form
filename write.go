package form

import (
	"io"
	"os"
)

func write(s string) {
	_, _ = io.WriteString(os.Stdout, s)
}

func ClearLine() {
	write("\u001b[0K")
}

const (
	QUIT_KEY_MESSAGE string = "\u001b[30;1mPress ESC to quit\u001b[0m"
)
