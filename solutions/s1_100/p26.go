package s1_100

import (
	"fmt"
)

func removeDuplicates_26(nums []int) int {
	l := 1
	r := 1
	for r < len(nums) {
		nums[l] = nums[r]
		if nums[l-1] != nums[l] {
			l += 1
		}
		r += 1
	}
	return l
}

type arg_26 struct {
	nums []int
}

func Run_26() {
	input := []arg_26{
		{[]int{1, 1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
	}

	for _, arg := range input {
		fmt.Print(removeDuplicates_26(arg.nums))
		fmt.Printf(" %v\n", arg.nums)
	}
}
