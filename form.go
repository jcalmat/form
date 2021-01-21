package form

import (
	"fmt"
	"strings"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

type FormItems []*FormItem

type FormItem struct {
	item     Item
	parent   *FormItem
	children FormItems
}

type form struct {
	items  FormItems
	active bool
}

// NewForm creates a new instance of form object
func NewForm() *form {
	return &form{
		items:  make([]*FormItem, 0),
		active: false,
	}
}

func (f *FormItem) AddSubItems(c ...Item) *FormItem {
	for _, item := range c {
		formItem := &FormItem{item: item, parent: f}
		formItem.setText()
		f.children = append(f.children, formItem)
	}
	return f
}

func (f *FormItem) setText() {

	p := f.parent
	parentsCount := 0
	for {
		if p.parent != nil {
			parentsCount++
			p = p.parent
			continue
		}
		break
	}

	f.item.setPrefix(fmt.Sprintf("%s╰─", strings.Repeat("  ", parentsCount)))
}

// AddSubItem adds a subItem i dependant of the FormItem f
// The rules applied to display the subItem are specific to
// each FormItem
func (f *FormItem) AddSubItem(c Item) *FormItem {
	item := &FormItem{item: c, parent: f}
	item.setText()
	f.children = append(f.children, item)
	return item
}

// AddItem adds one FormItem to the form object
func (f *form) AddItem(p Item) *FormItem {
	item := &FormItem{
		item: p,
	}
	f.items = append(f.items, item)
	return item
}

// AddItems adds many FormItems to the form object
func (f *form) AddItems(items ...Item) *form {
	for _, i := range items {
		item := &FormItem{
			item: i,
		}
		f.items = append(f.items, item)
	}
	return f
}

func (f FormItems) visibleItems() []Item {
	items := make([]Item, 0)
	for _, v := range f {
		items = append(items, v.item)
		if v.children != nil && v.item.displayChildren() {
			items = append(items, v.children.visibleItems()...)
		}
	}
	return items
}

func (f *form) visibleItems() []Item {
	return f.items.visibleItems()
}

func (f *form) pick(index, offset int) int {
	cursor.MovePrevLine(index)

	i := index + offset
	visibleItems := f.visibleItems()

	// Loop
	if i < 0 {
		i = len(visibleItems) - 1
	}
	if i > len(visibleItems)-1 {
		i = 0
	}

	// Range over the form and first unpick then pick the right one
	// to move the cursor on y axis.
	// Moving the cursor on x axis is handled by the pick() method.
	for n, p := range visibleItems {
		if n != i {
			p.unpick()
		}
	}

	// Move the cursor on y axis at the right row and select it.
	if visibleItems[i].selectable() {
		visibleItems[i].pick()
	} else {
		return f.pick(i, offset)
	}

	f.displayItems()
	cursor.MovePrevLine(len(visibleItems) - i)
	visibleItems[i].setCursorPosition()

	return i
}

func (f *form) stop() {
	f.active = false
}

func (f *form) displayItems() {
	cursor.RestorePosition()

	cursor.ClearBelow()

	// Display all visible items.
	for _, p := range f.visibleItems() {
		p.write()
		write("\n")
	}
}

// clearHiddenItemsValues range over all items and subItems and reset the value
// of the hidden ones
func (f FormItems) clearHiddenItemsValues() {
	for _, formItem := range f {
		if formItem.parent != nil && !formItem.parent.item.displayChildren() {
			formItem.item.clearValue()
		}
		if formItem.children != nil {
			formItem.children.clearHiddenItemsValues()
		}
	}
}

// Run displays the formItems and handles the user's inputs
func (f *form) Run() {
	cursor.StartBufferedSession()
	defer cursor.RestoreSession()

	f.active = true
	visibleItems := f.visibleItems()

	// Save cursor position at first line.
	cursor.SavePosition()

	// Do not process if there is no selectable formItem
	var firstSelectable *int
	for index, i := range visibleItems {
		if i.selectable() {
			firstSelectable = &index
			break
		}
	}
	if firstSelectable == nil {
		return
	}

	f.AddItem(NewButton(done_button, func() { f.stop() }))
	f.AddItem(NewLabel(navigation_keys_message))
	f.displayItems()

	cursor.DisableInputBuffering()
	cursor.HideInputs()
	defer cursor.RestoreEchoingState()

	cursor.MovePrevLine(len(f.visibleItems()) - *firstSelectable)
	selected := f.pick(*firstSelectable, 0)

	for {
		visibleItems = f.visibleItems()

		// Stop the main loop and clear the quit message
		if !f.active {
			f.items.clearHiddenItemsValues()
			break
		}

		i := input.Handle()

		// Handle UP and DOWN arrow keys for vertical navigation.
		if i.Is(input.UP) {
			selected = f.pick(selected, -1)
			continue
		} else if i.Is(input.DOWN) {
			selected = f.pick(selected, 1)
			continue
		}
		if i.Is(input.ESC) {
			f.stop()
			continue
		}

		// Handle any other input and let the formItem process it.
		visibleItems[selected].handleInput(i)

		// Handle Enter key to automatically select the next formItem.
		if i.Is(input.ENTER) || i.Is(input.TAB) { // Enter or Tab keys
			selected = f.pick(selected, 1)
		}
	}
}
