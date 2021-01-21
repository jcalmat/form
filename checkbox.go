package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

const (
	checkbox_uncheck string = "☐"
	checkbox_check   string = "☑"
)

// Checkbox implements Item interface
type Checkbox struct {
	question string
	prefix   string
	checked  bool
	selected bool
}

var _ Item = (*Checkbox)(nil)

// NewCheckbox creates a new instance of Checkbox object
func NewCheckbox(question string, checked bool) *Checkbox {
	return &Checkbox{
		prefix:   "",
		question: question,
		checked:  checked,
	}
}

func (c *Checkbox) write() {
	var question string

	Checkbox := checkbox_uncheck
	if c.checked {
		Checkbox = "\u001b[32;1m" + checkbox_check
	}
	if c.selected {
		Checkbox = "\u001b[7m" + Checkbox
	}

	question = fmt.Sprintf("%s%s %s\u001b[0m", c.prefix, Checkbox, c.question)

	clearLine()
	write(question)
	cursor.MoveColumn(1)
}

func (c *Checkbox) handleInput(i input.I) {
	if i.Is(input.ENTER) {
		c.toggle()
	}
}

func (c *Checkbox) setCursorPosition() {}

func (c *Checkbox) clearValue() {
	c.checked = false
}

func (c *Checkbox) pick() {
	c.selected = true
	cursor.HideCursor()
}

func (c *Checkbox) unpick() {
	c.selected = false
	cursor.DisplayCursor()
}

func (c *Checkbox) toggle() {
	c.checked = !c.checked
}

func (c *Checkbox) displayChildren() bool {
	return c.checked
}

func (c *Checkbox) selectable() bool { return true }

func (c *Checkbox) answer() interface{} {
	return c.Answer()
}

func (c *Checkbox) Answer() bool {
	return c.checked
}

func (c *Checkbox) setPrefix(prefix string) {
	c.prefix = prefix
}
