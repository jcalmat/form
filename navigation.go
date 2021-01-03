package form

import "fmt"

func moveUp(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dA", amount)
	}
	return ""
}

func moveDown(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dB", amount)
	}
	return ""
}

func moveRight(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dC", amount)
	}
	return ""
}

func moveLeft(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dD", amount)
	}
	return ""
}

func moveColumn(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dG", amount)
	}
	return ""
}

func moveNextLine(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dE", amount)
	}
	return ""
}

func movePrevLine(amount int) string {
	if amount > 0 {
		return fmt.Sprintf("\u001b[%dF", amount)
	}
	return ""
}
