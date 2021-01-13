package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
)

// button implements formItem interface
type button struct {
	s        string
	callback func()
	prefix   string
	selected bool
}

var _ Item = (*button)(nil)

// NewButton creates a new instance of button object
func NewButton(s string, callback func()) *button {
	return &button{
		s:        s,
		callback: callback,
	}
}

func (s *button) write() {
	cursor.MoveColumn(1)
	clearLine()
	if s.selected {
		write(fmt.Sprintf("[\u001b[7m%s%s\u001b[0m]", s.prefix, s.s))
	} else {
		write(fmt.Sprintf("[%s%s]", s.prefix, s.s))
	}
}

func (s *button) pick() {
	s.selected = true
	cursor.HideCursor()
}

func (s *button) unpick() {
	s.selected = false
	cursor.DisplayCursor()
}

func (s *button) handleInput(b []byte) {
	if len(b) > 0 && (b[0] == 10 || b[0] == 13) { // Enter
		s.callback()
	}
}

func (c *button) selectable() bool { return true }

func (c *button) setCursorPosition() {}

func (c *button) displayChildren() bool { return true }

func (c *button) setPrefix(prefix string) {
	c.prefix = prefix
}
