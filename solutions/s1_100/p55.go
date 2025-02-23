package s1_100

import (
	"fmt"
)

func canJump(nums []int) bool {
	i := 0
	end := 0
	for i <= end {
		end = max(end, i+nums[i])
		if end >= len(nums)-1 {
			return true
		}
		i += 1
	}
	return false
}

type arg_55 struct {
	nums []int
}

func Run_55() {
	input := []arg_55{
		{[]int{2, 3, 1, 1, 4}},
		{[]int{3, 2, 1, 0, 4}},
	}

	for _, arg := range input {
		fmt.Println(canJump(arg.nums))
	}
}
