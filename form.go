package form

import (
	"os"
	"os/exec"
)

type formItem interface {
	write()
	pick()
	unpick()
	handleInput([]byte)
	selectable() bool
	isVisible() bool
	setCursorPosition()
	// displayChildren() bool
	// size() int
}

type form struct {
	items  []formItem
	active bool
}

// NewForm creates a new instance of form object
func NewForm() *form {
	return &form{
		items:  make([]formItem, 0),
		active: false,
	}
}

// Register adds formItems to the form object
func (f *form) Register(p ...formItem) {
	f.items = append(f.items, p...)
}

func (f *form) visibleItems() []formItem {
	items := make([]formItem, 0)
	for _, v := range f.items {
		if v.isVisible() {
			items = append(items, v)
		}
	}
	return items
}

func (f *form) pick(index, offset int) int {
	movePrevLine(index)

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
	// movePrevLine(len(visibleItems) - i)
	if visibleItems[i].selectable() {
		visibleItems[i].pick()
	} else {
		return f.pick(i, offset)
	}

	f.displayItems()
	movePrevLine(len(visibleItems) - i)
	visibleItems[i].setCursorPosition()

	return i
}

func (f *form) unpickAll(index int) {
	// Move the cursor to the top of the form and individually unpick them.
	movePrevLine(index)

	for _, p := range f.visibleItems() {
		moveColumn(1)
		p.unpick()
		write("\n")
	}
}

func (f *form) stop() {
	f.active = false
}

func (f *form) displayItems() {
	write("\x1b8")
	// Clear all from cursor until the end of the screen
	write("\u001b[0J")

	// Display all visible items.
	for _, p := range f.visibleItems() {
		p.write()
		write("\n")
	}

	// Write a quit message
	write(navigation_keys_message)
}

// Ask displays the formItems and handles the user's inputs
func (f *form) Ask() {
	f.active = true
	visibleItems := f.visibleItems()

	// Save cursor position at first line.
	write("\x1b7")

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

	f.displayItems()

	// Disable input buffering.
	_ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Do not display entered characters on the screen.
	_ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// Restore the echoing state when exiting.
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run() //nolint

	movePrevLine(len(f.visibleItems()) - *firstSelectable)
	selected := f.pick(*firstSelectable, 0)
	var b []byte
	for {
		visibleItems = f.visibleItems()

		// Stop the main loop and clear the quit message
		if !f.active {
			f.unpickAll(selected)
			clearLine()
			break
		}

		// Flush the byte array.
		b = make([]byte, 3)

		_, _ = os.Stdin.Read(b)
		// Handle UP and DOWN arrow keys for vertical navigation.
		if b[0] == 27 {
			if b[1] == 91 {
				switch b[2] {
				case 65: // Up
					selected = f.pick(selected, -1)
					continue
				case 66: // Down
					selected = f.pick(selected, 1)
					continue
				}
			}
			if b[1] == 0 {
				f.stop()
			}
		}
		// Handle any other input and let the formItem process it.
		visibleItems[selected].handleInput(b)

		// Handle Enter key to automatically select the next formItem.
		if len(b) > 0 && (b[0] == 10 || b[0] == 13 || b[0] == 9) { // Enter or Tab keys
			// Stop if the selected formItem is the last one.
			if b[0] == 9 && selected+1 == len(f.visibleItems()) {
				f.stop()
				continue
			}
			selected = f.pick(selected, 1)
		}
	}
}
