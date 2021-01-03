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

	write(moveUp(index))

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
		write(moveColumn(1))
		if n != i {
			p.unpick()
		}
		write("\n")
	}

	// Move the cursor vertically at the right row and select it.
	write(moveUp(len(ps.items) - i))
	if ps.items[i].selectable() {
		ps.items[i].pick()
	} else {
		return ps.pick(i, offset)
	}
	return i
}

func (ps *form) stop() {
	ps.active = false
}

func (ps *form) Ask() {
	ps.active = true

	// Do not process if there is no selectable formItem
	selectableItemCount := 0
	for _, i := range ps.items {
		if i.selectable() {
			selectableItemCount++
		}
	}
	if selectableItemCount == 0 {
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

	write(movePrevLine(len(ps.items)))
	selected := ps.pick(0, 0)

	var b []byte
	for {
		// Stop the formItem loop and clear the quit message
		if !ps.active {
			write(moveNextLine(len(ps.items) - selected + 1))
			ClearLine()
			write(moveNextLine(1))
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
