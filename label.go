package form

// label implements formItem interface
type label struct {
	s string
}

var _ formItem = (*label)(nil)

// NewLabel creates a new instance of label object
func NewLabel(s string) *label {
	return &label{
		s: s,
	}
}

func (s *label) write() {
	moveColumn(1)
	clearLine()
	write(s.s)
}

func (s *label) pick() {}

func (s *label) unpick() {}

func (s *label) handleInput(b []byte) {}

func (c *label) selectable() bool { return false }

func (c *label) isVisible() bool { return true }

func (c *label) setCursorPosition() {}

func (c *label) displayChildren() bool { return true }
