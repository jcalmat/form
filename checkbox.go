package form

import (
	"fmt"
	"strings"
)

const (
	checkbox_uncheck string = "☐"
	checkbox_check   string = "☑"
)

// checkbox implements formItem interface
type checkbox struct {
	sentence     string
	checked      bool
	selected     bool
	verticalSize int
	parent       *checkbox
	children     []*checkbox
	// visible      bool
}

var _ formItem = (*checkbox)(nil)

// NewCheckbox creates a new instance of checkbox object
func NewCheckbox(s string, checked bool) *checkbox {
	return &checkbox{
		sentence:     s,
		checked:      checked,
		verticalSize: strings.Count("\n", s) + 1,
		// visible:      true,
	}
}

// AddDependance adds child to the current checkbox.
// If the value of the checkbox becomes true, the child checkbox will be
// displayed.
// If the value of the checkbox becomes false, the child checkbox will be
// disabled and its value will automatically be set to false.
func (c *checkbox) AddDependance(child *checkbox) {
	child.parent = c
	c.children = append(c.children, child)
	// if !c.visible || !c.checked {
	// 	child.visible = false
	// }

	// p := child.parent
	// parentsCount := 1
	// for {
	// 	if p.parent != nil {
	// 		parentsCount++
	// 		continue
	// 	}
	// 	break
	// }
	// child.sentence = fmt.Sprintf("%s╰─%s", strings.Repeat("  ", parentsCount), child.sentence)
	// child.minCursorPosition = parentsCount + 2
}

func (c *checkbox) Children() []*checkbox {
	ret := make([]*checkbox, 0)
	children := c.children
	ret = append(ret, children...)
	for i := 0; i < len(children); i++ {
		ret = append(ret, children[i].Children()...)
	}

	return ret
}

func (c *checkbox) write() {
	var s string

	parentsCount := 0
	parent := c.parent
	for {
		if parent != nil {
			parentsCount++
			parent = parent.parent
			continue
		}
		break
	}

	checkbox := checkbox_uncheck

	// s = checkbox_uncheck + " " + c.sentence
	if c.checked {
		checkbox = "\u001b[32;1m" + checkbox_check
	}
	if c.selected {
		checkbox = "\u001b[7m" + checkbox
	}

	if parentsCount > 0 {
		s = fmt.Sprintf("%s╰─%s %s\u001b[0m", strings.Repeat("  ", parentsCount-1), checkbox, c.sentence)
	} else {
		s = fmt.Sprintf("%s %s\u001b[0m", checkbox, c.sentence)
	}

	clearLine()
	write(s)
	moveColumn(1)
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

func (c *checkbox) setVisibility(b bool) {

}

// func (c *checkbox) displayChildren() {
// 	for _, v := range c.children {
// 		v.visible = c.checked && c.visible
// 		v.displayChildren()
// 	}
// }

func (c *checkbox) toggle() {
	c.checked = !c.checked
	// c.displayChildren()
}

func (c *checkbox) displayChildren() bool {
	if c.parent == nil {
		return true
	}
	return c.parent.checked
}

func (c *checkbox) isVisible() bool {
	visible := true
	parent := c.parent
	for {
		if parent != nil {
			if !parent.checked {
				visible = false
				break
			}
			parent = parent.parent
			continue
		}
		break
	}
	return visible
}

func (c *checkbox) selectable() bool { return true }

func (c *checkbox) size() int { return c.verticalSize }

func (c *checkbox) Answer() bool {
	return c.checked && c.isVisible()
}
