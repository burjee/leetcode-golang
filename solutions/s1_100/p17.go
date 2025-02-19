package s1_100

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letters := map[rune]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	ans := []string{""}

	for _, digit := range digits {
		temp := []string{}
		for _, s := range ans {
			for _, letter := range letters[digit] {
				temp = append(temp, s+string(letter))
			}
		}
		ans = temp
	}

	return ans
}

type arg_17 struct {
	digits string
}

func Run_17() {
	input := []arg_17{
		{"23"},
		{""},
		{"2"},
	}

	for _, arg := range input {
		fmt.Println(letterCombinations(arg.digits))
	}
}
