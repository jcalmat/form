package cursor

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// func moveUp(amount int) string {
// 	if amount > 0 {
// 		return fmt.Sprintf("\u001b[%dA", amount)
// 	}
// 	return ""
// }

func write(s string) {
	_, _ = io.WriteString(os.Stdout, s)
}

func MoveColumn(amount int) {
	if amount > 0 {
		write(fmt.Sprintf("\u001b[%dG", amount))
	}
}

func MovePrevLine(amount int) {
	if amount > 0 {
		write(fmt.Sprintf("\u001b[%dF", amount))
	}
}

func moveNextLine(amount int) {
	if amount > 0 {
		write(fmt.Sprintf("\u001b[%dE", amount))
	}
}

func StartBufferedSession() {
	fmt.Print("\033[?1049h\033[H")
}

func RestoreSession() {
	fmt.Print("\033[?1049l")
}

func SavePosition() {
	write("\x1b7")
}

func RestorePosition() {
	write("\x1b8")
}

// ClearBelow clear all from cursor until the end of the screen
func ClearBelow() {
	write("\u001b[0J")
}

func DisableInputBuffering() {
	_ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
}

func HideInputs() {
	_ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func RestoreEchoingState() {
	_ = exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}
