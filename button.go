package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

// button implements item interface
type button struct {
	s        string
	callback func()
	prefix   string
	selected bool
}

var _ item = (*button)(nil)

// NewButton creates a new instance of button object
func NewButton(s string, callback func()) *FormItem {
	return NewFormItem(&button{
		s:        s,
		callback: callback,
	})
}

func (b *button) write() {
	cursor.MoveColumn(1)
	clearLine()
	if b.selected {
		write(fmt.Sprintf("[\u001b[7m%s%s\u001b[0m]", b.prefix, b.s))
	} else {
		write(fmt.Sprintf("[%s%s]", b.prefix, b.s))
	}
}

func (b *button) pick() {
	b.selected = true
	cursor.HideCursor()
}

func (b *button) unpick() {
	b.selected = false
	cursor.DisplayCursor()
}

func (b *button) handleInput(i input.I) {
	if i.Is(input.ENTER) {
		b.callback()
	}
}

func (b *button) selectable() bool { return true }

func (b *button) setCursorPosition() {}

func (b *button) clearValue() {}

func (b *button) displayChildren() bool { return true }

func (b *button) setPrefix(prefix string) {
	b.prefix = prefix
}

func (b *button) answer() interface{} { return nil }
