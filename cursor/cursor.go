package cursor

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

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

var termios *unix.Termios

func StartBufferedSession() {
	// state, _ = terminal.GetState(0)
	// terminal.NewTerminal(state.)
	fmt.Print("\033[?1049h\033[H")
}

func RestoreSession() {
	fmt.Print("\033[?1049l")
}

// func SavePosition() {
// 	write("\u001b7")
// }

// func RestorePosition() {
// 	write("\u001b8")
// }

func ClearScreen() {
	fmt.Print("\u001b[0J")
}

// ClearBelow clear all from cursor until the end of the screen
func ClearBelow() {
	write("\u001b[0J")
}

func DisableInputBuffering() {
	// _ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
}

const (
	TCGETS = 0x5401
	TCSETS = 0x5402
)

func HideInputs() {
	// _ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	termios, err := unix.IoctlGetTermios(1, TCGETS)
	if err != nil {
		return
	}

	newState := *termios
	newState.Lflag &^= unix.ECHO | unix.ICANON
	newState.Lflag |= unix.ISIG
	newState.Iflag |= unix.ICRNL
	if err := unix.IoctlSetTermios(1, TCSETS, &newState); err != nil {
		return
	}
}

func RestoreEchoingState() {

	_ = unix.IoctlSetTermios(1, TCSETS, termios)

	// _ = exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func HideCursor() {
	write("\033[?25l")
}

func DisplayCursor() {
	write("\033[?25h")
}
