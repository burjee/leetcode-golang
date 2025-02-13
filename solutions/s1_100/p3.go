package s1_100

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	l := 0
	ans := 0

	chars := []rune(s)
	for r, char := range chars {
		if p, ok := m[char]; ok {
			if p >= l {
				l = p + 1
			}
		}
		m[char] = r
		ans = max(ans, r-l+1)
	}

	return ans
}

type arg_3 struct {
	s string
}

func Run_3() {
	input := []arg_3{
		{"abcabcbb"},
		{"bbbbb"},
		{"pwwkew"},
	}

	for _, arg := range input {
		fmt.Println(lengthOfLongestSubstring(arg.s))
	}
}
