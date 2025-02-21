package s1_100

import (
	"fmt"
)

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if target < nums[m] {
			if nums[l] <= target || nums[l] > nums[m] {
				r = m - 1
			} else if nums[l] > nums[r] {
				l = m + 1
			} else {
				break
			}
		} else if nums[m] < target {
			if target <= nums[r] || nums[m] > nums[r] {
				l = m + 1
			} else if nums[l] > nums[r] {
				r = m - 1
			} else {
				break
			}
		}
	}
	return -1
}

type arg_33 struct {
	nums   []int
	target int
}

func Run_33() {
	input := []arg_33{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
		{[]int{1}, 3},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8},
	}

	for _, arg := range input {
		fmt.Println(search(arg.nums, arg.target))
	}
}
