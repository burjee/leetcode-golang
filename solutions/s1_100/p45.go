package s1_100

import (
	"fmt"
)

func jump(nums []int) int {
	ans := 0
	i := 0
	start := 1
	for i < len(nums)-1 {
		ans += 1
		if i+nums[i] >= len(nums)-1 {
			break
		}

		next := 0
		distance := 0
		for j := start; j < i+nums[i]+1; j += 1 {
			if j+nums[j] > distance {
				distance = j + nums[j]
				next = j
			}
		}
		start = i + nums[i] + 1
		i = next
	}
	return ans
}

type arg_45 struct {
	nums []int
}

func Run_45() {
	input := []arg_45{
		{[]int{2, 3, 1, 1, 4}},
		{[]int{2, 3, 0, 1, 4}},
		{[]int{0}},
	}

	for _, arg := range input {
		fmt.Println(jump(arg.nums))
	}
}
