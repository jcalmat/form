package form

import (
	"fmt"

	"github.com/jcalmat/form/cursor"
)

const (
	checkbox_uncheck string = "☐"
	checkbox_check   string = "☑"
)

// checkbox implements formItem interface
type checkbox struct {
	question string
	prefix   string
	checked  bool
	selected bool
}

var _ Item = (*checkbox)(nil)

// NewCheckbox creates a new instance of checkbox object
func NewCheckbox(question string, checked bool) *checkbox {
	return &checkbox{
		prefix:   "",
		question: question,
		checked:  checked,
	}
}

// AddDependance adds child to the current checkbox.
// If the value of the checkbox becomes true, the child checkbox will be
// displayed.
// If the value of the checkbox becomes false, the child checkbox will be
// disabled and its value will automatically be set to false.
// func (c *checkbox) AddDependance(child *checkbox) {
// 	child.parent = c
// 	c.children = append(c.children, child)
// 	// if !c.visible || !c.checked {
// 	// 	child.visible = false
// 	// }

// 	// p := child.parent
// 	// parentsCount := 1
// 	// for {
// 	// 	if p.parent != nil {
// 	// 		parentsCount++
// 	// 		continue
// 	// 	}
// 	// 	break
// 	// }
// 	// child.question = fmt.Sprintf("%s╰─%s", strings.Repeat("  ", parentsCount), child.question)
// 	// child.minCursorPosition = parentsCount + 2
// }

// func (c *checkbox) Children() []*checkbox {
// 	ret := make([]*checkbox, 0)
// 	// children := c.children
// 	// ret = append(ret, children...)
// 	// for i := 0; i < len(children); i++ {
// 	// 	ret = append(ret, children[i].Children()...)
// 	// }

// 	return ret
// }

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

func (c *checkbox) handleInput(b []byte) {
	if len(b) > 0 && (b[0] == 10 || b[0] == 13) { // Enter
		c.toggle()
	}
}

func (c *checkbox) setCursorPosition() {}

func (c *checkbox) pick() {
	c.selected = true
	// hide cursor
	write("\033[?25l")
	// c.write()
}

func (c *checkbox) unpick() {
	c.selected = false
	// show cursor
	write("\033[?25h")
	// c.write()
}

func (c *checkbox) toggle() {
	c.checked = !c.checked
	// c.displayChildren()
}

func (c *checkbox) displayChildren() bool {
	return c.checked
}

func (c *checkbox) selectable() bool { return true }

//TODO: Fix answer to return false if the item is not visible
func (c *checkbox) Answer() bool {
	return c.checked
}

func (c *checkbox) getQuestion() string {
	return fmt.Sprintf("%s%s", c.prefix, c.question)
}

func (c *checkbox) setPrefix(prefix string) {
	c.prefix = prefix
}
