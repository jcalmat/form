package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

// Button implements Item interface
type Button struct {
	s        string
	callback func()
	prefix   string
	selected bool
}

var _ Item = (*Button)(nil)

// NewButton creates a new instance of Button object
func NewButton(s string, callback func()) *Button {
	return &Button{
		s:        s,
		callback: callback,
	}
}

func (b *Button) write() {
	cursor.MoveColumn(1)
	clearLine()
	if b.selected {
		write(fmt.Sprintf("[\u001b[7m%s%s\u001b[0m]", b.prefix, b.s))
	} else {
		write(fmt.Sprintf("[%s%s]", b.prefix, b.s))
	}
}

func (b *Button) pick() {
	b.selected = true
	cursor.HideCursor()
}

func (b *Button) unpick() {
	b.selected = false
	cursor.DisplayCursor()
}

func (b *Button) handleInput(i input.I) {
	if i.Is(input.ENTER) {
		b.callback()
	}
}

func (b *Button) selectable() bool { return true }

func (b *Button) setCursorPosition() {}

func (b *Button) clearValue() {}

func (b *Button) displayChildren() bool { return true }

func (b *Button) setPrefix(prefix string) {
	b.prefix = prefix
}
