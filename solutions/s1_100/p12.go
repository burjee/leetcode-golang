package s1_100

import (
	"fmt"
)

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := ""
	for num > 0 {
		if values[0] > num {
			values = values[1:]
			symbols = symbols[1:]
		} else {
			num -= values[0]
			ans += symbols[0]
		}
	}
	return ans
}

type arg_12 struct {
	num int
}

func Run_12() {
	input := []arg_12{
		{3749},
		{58},
		{1994},
	}

	for _, arg := range input {
		fmt.Println(intToRoman(arg.num))
	}
}
