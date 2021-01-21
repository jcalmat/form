package form

import "github.com/jcalmat/form/input"

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

	// clearValue reset the value to it's default state
	clearValue()

	answer() interface{}
}
