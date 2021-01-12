package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
)

// textField implements formItem interface
type textField struct {
	question          string
	prefix            string
	input             string
	cursorPosition    int
	minCursorPosition int
}

var _ Item = (*textField)(nil)

// NewTextField creates a new instance of textField object
func NewTextField(question string) *textField {
	return &textField{
		question:          question,
		input:             "",
		minCursorPosition: len(question),
		cursorPosition:    len(question),
	}
}

func (t *textField) write() {
	if t.cursorPosition > len(t.input) {
		t.cursorPosition = len(t.input)
	}
	if t.cursorPosition < 0 {
		t.cursorPosition = 0
	}

	cursor.MoveColumn(1)
	clearLine()
	write(t.question + "\u001b[37;1m" + t.input + "\u001b[0m")
}

func (t *textField) pick() {}

func (t *textField) unpick() {}

func (t *textField) setCursorPosition() {
	if t.cursorPosition > len(t.input) {
		t.cursorPosition = len(t.input)
	}
	if t.cursorPosition < 0 {
		t.cursorPosition = 0
	}

	cursor.MoveColumn(t.minCursorPosition + t.cursorPosition + 1)
}

func (t *textField) handleInput(b []byte) {
	if b[0] == 27 {
		if b[1] == 91 {
			switch b[2] {
			case 67: // Right
				t.cursorPosition++
				t.setCursorPosition()
			case 68: // Left
				t.cursorPosition--
				t.setCursorPosition()
			}
		}
		return
	}
	for _, c := range b {
		if c >= 32 && c <= 126 {
			t.input = fmt.Sprintf("%s%s%s", t.input[:t.cursorPosition], string(c), t.input[t.cursorPosition:])
			t.cursorPosition++
			t.write()
		} else if c == 127 {
			if t.cursorPosition > 0 {
				t.input = t.input[:t.cursorPosition-1] + t.input[t.cursorPosition:]
				t.cursorPosition--
				t.write()
			}
		}
	}
}

func (c *textField) selectable() bool { return true }

func (c *textField) Answer() string {
	return c.input
}

func (c *textField) displayChildren() bool { return c.input != "" }

func (c *textField) setPrefix(prefix string) {
	c.prefix = prefix
	c.minCursorPosition = len(fmt.Sprintf("%s%s", c.prefix, c.question))
}
