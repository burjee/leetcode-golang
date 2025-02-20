package s1_100

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	l := 0
	for r := range nums {
		if nums[r] != val {
			nums[l] = nums[r]
			l += 1
		}
	}
	return l
}

// func removeElement(nums []int, val int) int {
// 	l := 0
// 	r := len(nums) - 1
// 	for l <= r {
// 		if nums[l] == val {
// 			nums[l] = nums[r]
// 			r -= 1
// 		} else {
// 			l += 1
// 		}
// 	}
// 	return l
// }

type arg_27 struct {
	nums []int
	val  int
}

func Run_27() {
	input := []arg_27{
		{[]int{3, 2, 2, 3}, 3},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
		{[]int{2}, 2},
		{[]int{2}, 3},
	}

	for _, arg := range input {
		fmt.Print(removeElement(arg.nums, arg.val))
		fmt.Printf(" %v\n", arg.nums)
	}
}
