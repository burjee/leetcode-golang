package s1_100

import (
	"fmt"
	"slices"
)

func addBinary(a string, b string) string {
	b1 := []rune(a)
	b2 := []rune(b)
	slices.Reverse(b1)
	slices.Reverse(b2)

	ans := make([]rune, 0, len(a)+len(b))
	hold := 0
	for len(b1) > 0 || len(b2) > 0 {
		n := hold
		if len(b1) > 0 {
			if b1[0] == '1' {
				n += 1
			}
			b1 = b1[1:]
		}
		if len(b2) > 0 {
			if b2[0] == '1' {
				n += 1
			}
			b2 = b2[1:]
		}

		if n%2 == 0 {
			ans = append(ans, '0')
		} else {
			ans = append(ans, '1')
		}

		hold = (n & 0b10) >> 1
	}

	if hold == 1 {
		ans = append(ans, '1')
	}

	slices.Reverse(ans)
	return string(ans)
}

type arg_67 struct {
	a string
	b string
}

func Run_67() {
	input := []arg_67{
		{"11", "1"},
		{"1010", "1011"},
	}

	for _, arg := range input {
		fmt.Println(addBinary(arg.a, arg.b))
	}
}
