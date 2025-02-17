package s1_100

import (
	"fmt"
)

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		m := min(height[l], height[r])
		store := (r - l) * m
		ans = max(ans, store)
		if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
		}
	}
	return ans
}

type arg_11 struct {
	height []int
}

func Run_11() {
	input := []arg_11{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		{[]int{1, 1}},
	}

	for _, arg := range input {
		fmt.Println(maxArea(arg.height))
	}
}
