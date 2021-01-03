package form

type Label struct {
	s string
}

func NewLabel(s string) *Label {
	return &Label{
		s: s,
	}
}

func (s *Label) write() {
	write(moveColumn(1))
	ClearLine()
	write(s.s)
}

func (s *Label) pick() {}

func (s *Label) unpick() {}

func (s *Label) handleInput(b []byte) {}

func (c *Label) selectable() bool { return false }
