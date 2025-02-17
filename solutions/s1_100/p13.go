package s1_100

import (
	"fmt"
)

func romanToInt(s string) int {
	m := map[rune]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}

	ans := 0
	pre := 0
	for _, ch := range s {
		value := m[ch]
		ans += value
		if value > pre {
			ans -= pre * 2
		}
		pre = value
	}

	return ans
}

// func romanToInt(s string) int {
// 	m := map[rune]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
// 	runes := []rune(s)
// 	ans := 0

// 	last := 0
// 	i := len(runes) - 1
// 	for i >= 0 {
// 		value := m[runes[i]]
// 		if value >= last {
// 			ans += value
// 		} else {
// 			ans -= value
// 		}
// 		last = value
// 		i -= 1
// 	}

// 	return ans
// }

type arg_13 struct {
	s string
}

func Run_13() {
	input := []arg_13{
		{"III"},
		{"LVIII"},
		{"MCMXCIV"},
	}

	for _, arg := range input {
		fmt.Println(romanToInt(arg.s))
	}
}
