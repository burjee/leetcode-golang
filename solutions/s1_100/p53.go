package s1_100

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	sum := 0
	ans := nums[0]
	for _, num := range nums {
		sum = max(sum+num, num)
		ans = max(ans, sum)
	}
	return ans
}

type arg_53 struct {
	nums []int
}

func Run_53() {
	input := []arg_53{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		{[]int{1}},
		{[]int{5, 4, -1, 7, 8}},
		{[]int{-1}},
	}

	for _, arg := range input {
		fmt.Println(maxSubArray(arg.nums))
	}
}
