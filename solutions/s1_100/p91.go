package s1_100

import (
	"fmt"
)

func numDecodings(s string) int {
	m := make(map[string]int)
	return helper_91(s, m)
}

func helper_91(s string, m map[string]int) int {
	if len(s) == 0 {
		return 1
	}

	if v, ok := m[s]; ok {
		return v
	}

	count := 0
	if len(s) > 0 {
		if s[0] != '0' {
			count += helper_91(s[1:], m)
		}
	}
	if len(s) > 1 {
		if s[0] == '1' {
			count += helper_91(s[2:], m)
		}
		if s[0] == '2' && s[1] != '7' && s[1] != '8' && s[1] != '9' {
			count += helper_91(s[2:], m)
		}
	}
	m[s] = count
	return count
}

type arg_91 struct {
	s string
}

func Run_91() {
	input := []arg_91{
		{"12"},
		{"226"},
		{"06"},
	}

	for _, arg := range input {
		fmt.Println(numDecodings(arg.s))
	}
}
