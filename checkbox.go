package form

const (
	CHECKBOX_UNCHECK string = "☐"
	CHECKBOX_CHECK   string = "☑"
)

type Checkbox struct {
	sentence string
	checked  bool
	selected bool
}

func NewCheckbox(s string, checked bool) *Checkbox {
	return &Checkbox{
		sentence: s,
		checked:  checked,
	}
}

func (c *Checkbox) write() {
	var s string
	s = CHECKBOX_UNCHECK + " " + c.sentence
	if c.checked {
		s = "\u001b[32;1m" + CHECKBOX_CHECK + " " + c.sentence + "\u001b[0m"
	}
	if c.selected {
		s = "\u001b[7m" + s + " \u001b[0m"
	}
	clearLine()
	write(s)
	moveColumn(1)
}

func (c *Checkbox) handleInput(b []byte) {
	if len(b) > 0 && (b[0] == 10 || b[0] == 13) { // Enter
		c.toggle()
	}
}

func (c *Checkbox) pick() {
	c.selected = true
	c.write()
}

func (c *Checkbox) unpick() {
	c.selected = false
	c.write()
}

func (c *Checkbox) toggle() {
	c.checked = !c.checked
	c.pick()
}

func (c *Checkbox) selectable() bool { return true }

func (c *Checkbox) Answer() bool {
	return c.checked
}
