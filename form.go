package form

import (
	"github.com/jcalmat/form/cursor"
	"github.com/jcalmat/form/input"
)

type Form struct {
	items  FormItems
	active bool
}

// NewForm creates a new instance of Form object
func NewForm() *Form {
	return &Form{
		items:  make([]*FormItem, 0),
		active: false,
	}
}

// AddItem adds one FormItem to the Form object
func (f *Form) AddItem(formItem *FormItem) *FormItem {
	f.items = append(f.items, formItem)
	return formItem
}

// // AddItems adds many FormItems to the Form object
// func (f *Form) AddItems(items ...Item) *Form {
// 	for _, i := range items {
// 		item := &FormItem{
// 			item: i,
// 		}
// 		f.items = append(f.items, item)
// 	}
// 	return f
// }

func (f *Form) visibleItems() []Item {
	return f.items.visibleItems()
}

func (f *Form) pick(index, offset int) int {
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

	// Range over the Form and first unpick then pick the right one
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

func (f *Form) stop() {
	f.active = false
}

func (f *Form) displayItems() {
	cursor.RestorePosition()

	cursor.ClearBelow()

	// Display all visible items.
	for _, p := range f.visibleItems() {
		p.write()
		write("\n")
	}
}

// Run displays the formItems and handles the user's inputs
func (f *Form) Run() {
	cursor.StartBufferedSession()
	defer cursor.RestoreSession()

	f.active = true
	visibleItems := f.visibleItems()

	for _, item := range f.items {
		item.setText()
	}

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
