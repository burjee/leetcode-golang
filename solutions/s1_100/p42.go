package s1_100

import (
	"fmt"
)

func trap(height []int) int {
	ans := 0
	l := 0
	r := len(height) - 1
	h := 0
	for l < r {
		if height[l] <= height[r] {
			h = max(h, height[l])
			l += 1
			ans += max(h-height[l], 0)
		} else {
			h = max(h, height[r])
			r -= 1
			ans += max(h-height[r], 0)
		}
	}
	return ans
}

type arg_42 struct {
	height []int
}

func Run_42() {
	input := []arg_42{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{[]int{4, 2, 0, 3, 2, 5}},
	}

	for _, arg := range input {
		fmt.Println(trap(arg.height))
	}
}
