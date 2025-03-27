package s1_100

import (
	"fmt"
	"slices"
)

func merge_56(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })

	ans := [][]int{intervals[0]}
	i := 0
	for _, interval := range intervals[1:] {
		if ans[i][1] < interval[0] {
			i += 1
			ans = append(ans, interval)
		} else {
			ans[i][1] = max(ans[i][1], interval[1])
		}
	}
	return ans
}

type arg_56 struct {
	intervals [][]int
}

func Run_56() {
	input := []arg_56{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}},
	}

	for _, arg := range input {
		fmt.Println(merge_56(arg.intervals))
	}
}
