package form

// Label implements formItem interface
type Label struct {
	s string
}

// NewLabel creates a new instance of Label object
func NewLabel(s string) *Label {
	return &Label{
		s: s,
	}
}

func (s *Label) write() {
	moveColumn(1)
	clearLine()
	write(s.s)
}

func (s *Label) pick() {}

func (s *Label) unpick() {}

func (s *Label) handleInput(b []byte) {}

func (c *Label) selectable() bool { return false }
