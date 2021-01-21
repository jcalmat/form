package form

import (
	"fmt"
	"unicode/utf8"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

// TextField implements Item interface
type TextField struct {
	question          string
	prefix            string
	input             string
	cursorPosition    int
	minCursorPosition int
}

var _ Item = (*TextField)(nil)

// NewTextField creates a new instance of TextField object
func NewTextField(question string) *FormItem {
	return NewFormItem(&TextField{
		question:          question,
		input:             "",
		minCursorPosition: utf8.RuneCountInString(question),
		cursorPosition:    utf8.RuneCountInString(question),
	})
}

func (t *TextField) write() {
	cursor.MoveColumn(1)
	clearLine()
	write(t.prefix + t.question + "\u001b[37;1m" + t.input + "\u001b[0m")
	t.setCursorPosition()
}

func (t *TextField) pick() {}

func (t *TextField) unpick() {}

func (t *TextField) setCursorPosition() {
	if t.cursorPosition > utf8.RuneCountInString(t.input) {
		t.cursorPosition = utf8.RuneCountInString(t.input)
	}
	if t.cursorPosition < 0 {
		t.cursorPosition = 0
	}

	cursor.MoveColumn(t.minCursorPosition + t.cursorPosition + 1)
}

func (t *TextField) handleInput(i input.I) {
	if i.Is(input.RIGHT) {
		t.cursorPosition++
		t.setCursorPosition()
		return
	} else if i.Is(input.LEFT) {
		t.cursorPosition--
		t.setCursorPosition()
		return
	}

	for _, c := range i {
		if input.I([]byte{c}).Printable() {
			t.input = fmt.Sprintf("%s%s%s", t.input[:t.cursorPosition], string(c), t.input[t.cursorPosition:])
			t.cursorPosition++
			t.write()
		} else if input.I([]byte{c}).Is(input.DEL) {
			if t.cursorPosition > 0 {
				t.input = t.input[:t.cursorPosition-1] + t.input[t.cursorPosition:]
				t.cursorPosition--
				t.write()
			}
		}
	}
}

func (t *TextField) selectable() bool { return true }

func (t *TextField) Answer() string {
	return t.input
}

func (t *TextField) answer() interface{} {
	return t.Answer()
}

func (t *TextField) displayChildren() bool { return t.input != "" }

func (t *TextField) setPrefix(prefix string) {
	t.prefix = prefix
	t.minCursorPosition = utf8.RuneCountInString(t.prefix) + utf8.RuneCountInString(t.question)
	t.cursorPosition = t.minCursorPosition
}

func (t *TextField) clearValue() {
	t.input = ""
}
