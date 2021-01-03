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
}

type form struct {
	items  []formItem
	active bool
}

func NewForm() *form {
	return &form{
		items:  make([]formItem, 0),
		active: false,
	}
}

func (ps *form) Add(p ...formItem) {
	ps.items = append(ps.items, p...)
}

func (ps *form) pick(index, offset int) int {

	movePrevLine(index)

	i := index + offset

	// Loop
	if i < 0 {
		i = len(ps.items) - 1
	}
	if i > len(ps.items)-1 {
		i = 0
	}

	// Range over the form and first deselect then select the right one
	// to place the cursor at the right place.
	// Moving the cursor is handled by the pick() method.
	for n, p := range ps.items {
		if n != i {
			p.unpick()
		}
		write("\n")
	}

	// Move the cursor vertically at the right row and select it.
	movePrevLine(len(ps.items) - i)
	if ps.items[i].selectable() {
		ps.items[i].pick()
	} else {
		return ps.pick(i, offset)
	}
	return i
}

func (ps *form) unpickAll(index int) {
	// Move the cursor to the top of the form and individually unpick them.
	movePrevLine(index)

	for _, p := range ps.items {
		moveColumn(1)
		p.unpick()
		write("\n")
	}
}

func (ps *form) stop() {
	ps.active = false
}

func (ps *form) Ask() {
	ps.active = true

	// Do not process if there is no selectable formItem
	var firstSelectable *int
	for index, i := range ps.items {
		if i.selectable() {
			firstSelectable = &index
			break
		}
	}
	if firstSelectable == nil {
		return
	}

	// Display all items.
	for _, p := range ps.items {
		p.write()
		write("\n")
	}

	// Write a quit message
	write(QUIT_KEY_MESSAGE)

	// Disable input buffering.
	_ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// Do not display entered characters on the screen.
	_ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// Restore the echoing state when exiting.
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run() //nolint

	movePrevLine(len(ps.items) - *firstSelectable)
	selected := ps.pick(*firstSelectable, 0)

	var b []byte
	for {
		// Stop the main loop and clear the quit message
		if !ps.active {
			ps.unpickAll(selected)
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
					selected = ps.pick(selected, -1)
					continue
				case 66: // Down
					selected = ps.pick(selected, 1)
					continue
				}
			}
			if b[1] == 0 {
				ps.stop()
			}
		}
		// Handle any other input and let the formItem process it.
		ps.items[selected].handleInput(b)

		// Handle Enter key to automatically select the next formItem.
		if len(b) > 0 && (b[0] == 10 || b[0] == 13) { // Enter
			// Stop if the selected formItem is the last one.
			if selected+1 == len(ps.items) {
				ps.stop()
				continue
			}
			selected = ps.pick(selected, 1)
		}
	}
}
