package s1_100

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	j := len(needle)
	for i := range haystack {
		if i+j > len(haystack) {
			break
		}
		if haystack[i:i+j] == needle {
			return i
		}
	}
	return -1
}

type arg_28 struct {
	haystack string
	needle   string
}

func Run_28() {
	input := []arg_28{
		{"sadbutsad", "sad"},
		{"leetcode", "leeto"},
	}

	for _, arg := range input {
		fmt.Println(strStr(arg.haystack, arg.needle))
	}
}
