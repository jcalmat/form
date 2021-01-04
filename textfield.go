package form

import "fmt"

// TextField implements formItem interface
type TextField struct {
	prefix            string
	input             string
	cursorPosition    int
	minCursorPosition int
}

// NewTextField creates a new instance of TextField object
func NewTextField(prefix string) *TextField {
	return &TextField{
		prefix:            prefix,
		input:             "",
		minCursorPosition: len(prefix),
		cursorPosition:    len(prefix),
	}
}

func (s *TextField) moveCursor() {
	if s.cursorPosition > len(s.input) {
		s.cursorPosition = len(s.input)
	}
	if s.cursorPosition < 0 {
		s.cursorPosition = 0
	}

	moveColumn(s.minCursorPosition + s.cursorPosition + 1)
}

func (s *TextField) write() {
	if s.cursorPosition > len(s.input) {
		s.cursorPosition = len(s.input)
	}
	if s.cursorPosition < 0 {
		s.cursorPosition = 0
	}

	moveColumn(1)
	clearLine()
	write(s.prefix + "\u001b[37;1m" + s.input + "\u001b[0m")
	s.moveCursor()
}

func (s *TextField) pick() {
	s.write()
}

func (s *TextField) unpick() {}

func (s *TextField) handleInput(b []byte) {
	if b[0] == 27 {
		if b[1] == 91 {
			switch b[2] {
			case 67: // Right
				s.cursorPosition++
				s.moveCursor()
			case 68: // Left
				s.cursorPosition--
				s.moveCursor()
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

func (c *TextField) selectable() bool { return true }

func (c *TextField) Answer() string {
	return c.input
}
