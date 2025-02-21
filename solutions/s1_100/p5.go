package s1_100

import "fmt"

func longestPalindrome(s string) string {
	chars := []rune(s)
	length := len(chars)
	ans := chars[0:1]

	for i := range chars {
		if i < length-1 && chars[i] == chars[i+1] {
			helper_5(i, i+1, length, chars, &ans)
		}
		if i < length-2 && chars[i] == chars[i+2] {
			helper_5(i, i+2, length, chars, &ans)
		}
	}

	return string(ans)
}

func helper_5(l, r, length int, chars []rune, ans *[]rune) {
	for l > 0 && r < length-1 {
		if chars[l-1] != chars[r+1] {
			break
		}
		l -= 1
		r += 1
	}
	if r-l+1 > len(*ans) {
		*ans = chars[l : r+1]
	}
}

type arg_5 struct {
	s string
}

func Run_5() {
	input := []arg_5{
		{"babad"},
		{"cbbd"},
		{"aaaa"},
		{"aabaa"},
		{"bb"},
		{"abc"},
	}

	for _, arg := range input {
		fmt.Println(longestPalindrome(arg.s))
	}
}
