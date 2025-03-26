package collections

func Ptr(n int) *int {
	return &n
}

func ToZeroIndex(char byte) byte {
	if char < 97 {
		return char - 65
	} else {
		return char - 97 + 26
	}
}
