package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

// label implements Item interface
type label struct {
	s      string
	prefix string
}

var _ Item = (*label)(nil)

// NewLabel creates a new instance of label object
func NewLabel(s string) *FormItem {
	return NewFormItem(&label{
		s: s,
	})
}

func (l *label) write() {
	cursor.MoveColumn(1)
	clearLine()
	write(fmt.Sprintf("%s%s", l.prefix, l.s))
}

func (l *label) pick() {}

func (l *label) unpick() {}

func (l *label) handleInput(i input.I) {}

func (l *label) selectable() bool { return false }

func (l *label) setCursorPosition() {}

func (l *label) clearValue() {}

func (l *label) displayChildren() bool { return true }

func (l *label) setPrefix(prefix string) {
	l.prefix = prefix
}

func (l *label) answer() interface{} { return nil }
