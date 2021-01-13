package input

import "os"

func Handle() I {
	b := make([]byte, 3)

	_, _ = os.Stdin.Read(b)
	return I(b)
}

func (i I) Is(cmp I) bool {
	if len(i) != len(cmp) {
		return false
	}
	for index := range cmp {
		if cmp[index] != i[index] {
			return false
		}
	}
	return true
}
