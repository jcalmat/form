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

// checkbox implements item interface
type checkbox struct {
	question string
	prefix   string
	checked  bool
	selected bool
}

var _ item = (*checkbox)(nil)

// NewCheckbox creates a new instance of checkbox object
func NewCheckbox(question string, checked bool) *FormItem {
	return NewFormItem(&checkbox{
		prefix:   "",
		question: question,
		checked:  checked,
	})
}

func (c *checkbox) write() {
	var question string

	checkbox := checkbox_uncheck
	if c.checked {
		checkbox = "\u001b[32;1m" + checkbox_check
	}
	if c.selected {
		checkbox = "\u001b[7m" + checkbox
	}

	question = fmt.Sprintf("%s%s %s\u001b[0m", c.prefix, checkbox, c.question)

	clearLine()
	write(question)
	cursor.MoveColumn(1)
}

func (c *checkbox) handleInput(i input.I) {
	if i.Is(input.ENTER) {
		c.toggle()
	}
}

func (c *checkbox) setCursorPosition() {}

func (c *checkbox) clearValue() {
	c.checked = false
}

func (c *checkbox) pick() {
	c.selected = true
	cursor.HideCursor()
}

func (c *checkbox) unpick() {
	c.selected = false
	cursor.DisplayCursor()
}

func (c *checkbox) toggle() {
	c.checked = !c.checked
}

func (c *checkbox) displayChildren() bool {
	return c.checked
}

func (c *checkbox) selectable() bool { return true }

func (c *checkbox) answer() interface{} {
	return c.checked
}

func (c *checkbox) setPrefix(prefix string) {
	c.prefix = prefix
}
