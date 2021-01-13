package form

import (
	"fmt"
	"strings"

	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

type Item interface {
	// write writes the item where the cursor currently is
	write()

	// unpick tells the item that it is currently selected
	pick()

	// unpick tells the item that it is currently unselected
	unpick()

	// handleInput sniffs inputs byte by byte and process actions if needed
	handleInput(input.I)

	// selectable indicates if the item should be selectable of if it should
	// be skipped when navigating in the item list
	selectable() bool

	// setCursorPosition asks the item to set the cursor position on the x axis
	setCursorPosition()

	// displayChildren assert that, given the current item properties status, its
	// children can be display
	displayChildren() bool

	// setPrefix sets the item text prefix if relevant
	setPrefix(string)
}

type formItems []*formItem

type formItem struct {
	item     Item
	parent   *formItem
	children formItems
}

type form struct {
	items  formItems
	active bool
}

// NewForm creates a new instance of form object
func NewForm() *form {
	return &form{
		items:  make([]*formItem, 0),
		active: false,
	}
}

func (f *formItem) AddChildren(c ...Item) *formItem {
	for _, item := range c {
		formItem := &formItem{item: item, parent: f}
		formItem.setText()
		f.children = append(f.children, formItem)
	}
	return f
}

func (f *formItem) setText() {

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

// AddSubItem adds a subItem i dependant of the formItem f
// The rules applied to display the subItem are specific to
// each formItem
func (f *formItem) AddSubItem(c Item) *formItem {
	item := &formItem{item: c, parent: f}
	item.setText()
	f.children = append(f.children, item)
	return item
}

// AddItem adds one formItem to the form object
func (f *form) AddItem(p Item) *formItem {
	item := &formItem{
		item: p,
	}
	f.items = append(f.items, item)
	return item
}

// AddItems adds many formItems to the form object
func (f *form) AddItems(items ...Item) *form {
	for _, i := range items {
		item := &formItem{
			item: i,
		}
		f.items = append(f.items, item)
	}
	return f
}

func (f formItems) visibleItems() []Item {
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

	// Range over the form and first deselect then select the right one
	// to place the cursor at the right place.
	// Moving the cursor is handled by the pick() method.
	for n, p := range visibleItems {
		if n != i {
			p.unpick()
		}
	}

	// Move the cursor vertically at the right row and select it.
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
