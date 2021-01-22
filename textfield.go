package form

import (
	"fmt"
	"unicode/utf8"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

// textField implements item interface
type textField struct {
	question          string
	prefix            string
	input             string
	cursorPosition    int
	minCursorPosition int
}

var _ item = (*textField)(nil)

// NewTextField creates a new instance of textField object
func NewTextField(question string) *FormItem {
	return NewFormItem(&textField{
		question:          question,
		input:             "",
		minCursorPosition: utf8.RuneCountInString(question),
		cursorPosition:    utf8.RuneCountInString(question),
	})
}

func (t *textField) write() {
	cursor.MoveColumn(1)
	clearLine()
	write(t.prefix + t.question + "\u001b[37;1m" + t.input + "\u001b[0m")
	t.setCursorPosition()
}

func (t *textField) pick() {}

func (t *textField) unpick() {}

func (t *textField) setCursorPosition() {
	if t.cursorPosition > utf8.RuneCountInString(t.input) {
		t.cursorPosition = utf8.RuneCountInString(t.input)
	}
	if t.cursorPosition < 0 {
		t.cursorPosition = 0
	}

	cursor.MoveColumn(t.minCursorPosition + t.cursorPosition + 1)
}

func (t *textField) handleInput(i input.I) {
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

func (t *textField) selectable() bool { return true }

func (t *textField) answer() interface{} {
	return t.input
}

func (t *textField) displayChildren() bool { return t.input != "" }

func (t *textField) setPrefix(prefix string) {
	t.prefix = prefix
	t.minCursorPosition = utf8.RuneCountInString(t.prefix) + utf8.RuneCountInString(t.question)
	t.cursorPosition = t.minCursorPosition
}

func (t *textField) clearValue() {
	t.input = ""
}
