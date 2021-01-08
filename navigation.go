package form

import "fmt"

// func moveUp(amount int) string {
// 	if amount > 0 {
// 		return fmt.Sprintf("\u001b[%dA", amount)
// 	}
// 	return ""
// }

func moveColumn(amount int) {
	if amount > 0 {
		write(fmt.Sprintf("\u001b[%dG", amount))
	}
}

func movePrevLine(amount int) {
	if amount > 0 {
		write(fmt.Sprintf("\u001b[%dF", amount))
	}
}

func moveNextLine(amount int) {
	if amount > 0 {
		write(fmt.Sprintf("\u001b[%dE", amount))
	}
}
