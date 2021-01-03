package form

type Button struct {
	text     string
	callback func()
}

func NewButton(s string, callback func()) *Button {
	return &Button{
		text:     s,
		callback: callback,
	}
}

func (c *Button) write() {
	ClearLine()
	write(">" + c.text + "<")
	write(moveColumn(1))
}

func (c *Button) handleInput(b []byte) {
	if len(b) > 0 && (b[0] == 10 || b[0] == 13) { // Enter
		c.callback()
	}
}

func (c *Button) pick() {
	c.write()
}

func (c *Button) unpick() {
	c.write()
}

func (c *Button) selectable() bool { return true }
