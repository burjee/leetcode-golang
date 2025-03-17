package s1_100

import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	for len(intervals) > 0 {
		if intervals[0][1] >= newInterval[0] {
			break
		}

		ans = append(ans, intervals[0])
		intervals = intervals[1:]
	}

	if len(intervals) > 0 {
		if newInterval[1] < intervals[0][0] {
			ans = append(ans, newInterval)
		} else {
			start := min(intervals[0][0], newInterval[0])
			end := max(intervals[0][1], newInterval[1])
			interval := []int{start, end}
			ans = append(ans, interval)
			intervals = intervals[1:]
		}

	} else {
		ans = append(ans, newInterval)
	}

	i := len(ans) - 1
	for _, interval := range intervals {
		if ans[i][1] < interval[0] {
			i += 1
			ans = append(ans, interval)
		} else {
			ans[i][1] = max(ans[i][1], interval[1])
		}
	}

	return ans
}

type arg_57 struct {
	intervals   [][]int
	newInterval []int
}

func Run_57() {
	input := []arg_57{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}},
		{[][]int{{1, 5}}, []int{0, 0}},
	}

	for _, arg := range input {
		fmt.Println(insert(arg.intervals, arg.newInterval))
	}
}
