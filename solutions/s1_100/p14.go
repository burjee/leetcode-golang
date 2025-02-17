package s1_100

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	ans := ""
	var ch byte

outer:
	for {
		for i, str := range strs {
			if len(str) == 0 {
				break outer
			}
			if i == 0 {
				ch = str[0]
			} else if ch != str[0] {
				break outer
			}
			strs[i] = strs[i][1:]
		}
		ans += string(ch)
	}

	return ans
}

type arg_14 struct {
	strs []string
}

func Run_14() {
	input := []arg_14{
		{[]string{"flower", "flow", "flight"}},
		{[]string{"dog", "racecar", "car"}},
	}

	for _, arg := range input {
		fmt.Println(longestCommonPrefix(arg.strs))
	}
}
