package s1_100

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	i := len(s) - 1
	ans := 0

	for i >= 0 && s[i] != ' ' {
		i -= 1
		ans += 1
	}

	return ans
}

type arg_58 struct {
	s string
}

func Run_58() {
	input := []arg_58{
		{"Hello World"},
		{"   fly me   to   the moon  "},
		{"luffy is still joyboy"},
	}

	for _, arg := range input {
		fmt.Println(lengthOfLastWord(arg.s))
	}
}
