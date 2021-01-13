package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

// Label implements formItem interface
type Label struct {
	s      string
	prefix string
}

var _ Item = (*Label)(nil)

// NewLabel creates a new instance of Label object
func NewLabel(s string) *Label {
	return &Label{
		s: s,
	}
}

func (l *Label) write() {
	cursor.MoveColumn(1)
	clearLine()
	write(fmt.Sprintf("%s%s", l.prefix, l.s))
}

func (l *Label) pick() {}

func (l *Label) unpick() {}

func (l *Label) handleInput(i input.I) {}

func (l *Label) selectable() bool { return false }

func (l *Label) setCursorPosition() {}

func (l *Label) clearValue() {}

func (l *Label) displayChildren() bool { return true }

func (l *Label) setPrefix(prefix string) {
	l.prefix = prefix
}
