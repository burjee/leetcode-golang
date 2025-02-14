package s1_100

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}

	return true
}

type arg_9 struct {
	x int
}

func Run_9() {
	input := []arg_9{
		{121},
		{-121},
		{10},
	}

	for _, arg := range input {
		fmt.Println(isPalindrome(arg.x))
	}
}
