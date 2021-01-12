package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
)

// label implements formItem interface
type label struct {
	s      string
	prefix string
}

var _ Item = (*label)(nil)

// NewLabel creates a new instance of label object
func NewLabel(s string) *label {
	return &label{
		s: s,
	}
}

func (s *label) write() {
	cursor.MoveColumn(1)
	clearLine()
	write(fmt.Sprintf("%s%s", s.prefix, s.s))
}

func (s *label) pick() {}

func (s *label) unpick() {}

func (s *label) handleInput(b []byte) {}

func (c *label) selectable() bool { return false }

func (c *label) setCursorPosition() {}

func (c *label) displayChildren() bool { return true }

func (c *label) setPrefix(prefix string) {
	c.prefix = prefix
}
