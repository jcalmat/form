package form

import (
	"fmt"
	"strings"
)

// textField implements formItem interface
type textField struct {
	prefix            string
	input             string
	cursorPosition    int
	minCursorPosition int
	verticalSize      int
}

var _ formItem = (*textField)(nil)

// NewTextField creates a new instance of textField object
func NewTextField(prefix string) *textField {
	return &textField{
		prefix:            prefix,
		input:             "",
		minCursorPosition: len(prefix),
		cursorPosition:    len(prefix),
		verticalSize:      strings.Count("\n", prefix) + 1,
	}
}

// func (s *textField) moveCursor() {
// 	if s.cursorPosition > len(s.input) {
// 		s.cursorPosition = len(s.input)
// 	}
// 	if s.cursorPosition < 0 {
// 		s.cursorPosition = 0
// 	}

// 	moveColumn(s.minCursorPosition + s.cursorPosition + 1)
// }

func (s *textField) write() {
	if s.cursorPosition > len(s.input) {
		s.cursorPosition = len(s.input)
	}
	if s.cursorPosition < 0 {
		s.cursorPosition = 0
	}

	moveColumn(1)
	clearLine()
	write(s.prefix + "\u001b[37;1m" + s.input + "\u001b[0m")
	// s.moveCursor()
}

func (s *textField) pick() {
	// s.write()
}

func (s *textField) unpick() {}

func (s *textField) setCursorPosition() {
	if s.cursorPosition > len(s.input) {
		s.cursorPosition = len(s.input)
	}
	if s.cursorPosition < 0 {
		s.cursorPosition = 0
	}

	moveColumn(s.minCursorPosition + s.cursorPosition + 1)
}

func (s *textField) handleInput(b []byte) {
	if b[0] == 27 {
		if b[1] == 91 {
			switch b[2] {
			case 67: // Right
				s.cursorPosition++
				s.setCursorPosition()
			case 68: // Left
				s.cursorPosition--
				s.setCursorPosition()
			}
		}
		return
	}
	for _, c := range b {
		if c >= 32 && c <= 126 {
			s.input = fmt.Sprintf("%s%s%s", s.input[:s.cursorPosition], string(c), s.input[s.cursorPosition:])
			s.cursorPosition++
			s.write()
		} else if c == 127 {
			if s.cursorPosition > 0 {
				s.input = s.input[:s.cursorPosition-1] + s.input[s.cursorPosition:]
				s.cursorPosition--
				s.write()
			}
		}
	}
}

func (c *textField) isVisible() bool { return true }

func (c *textField) selectable() bool { return true }

func (c *textField) size() int { return c.verticalSize }

func (c *textField) Answer() string {
	return c.input
}

func (c *textField) displayChildren() bool { return c.input != "" }
