package input

type I []byte

var (
	UP    I = []byte{27, 91, 65}
	DOWN  I = []byte{27, 91, 66}
	RIGHT I = []byte{27, 91, 67}
	LEFT  I = []byte{27, 91, 68}
	ESC   I = []byte{27, 0, 0}
	ENTER I = []byte{10, 0, 0}
	TAB   I = []byte{9, 0, 0}
	DEL   I = []byte{127}
)

func (i I) Printable() bool {
	if len(i) > 0 && i[0] >= 32 && i[0] <= 126 {
		return true
	}
	return false
}
