package s1_100

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	helper_39(candidates, target, &[]int{}, &ans)
	return ans
}

func helper_39(candidates []int, target int, combinations *[]int, ans *[][]int) {
	if target == 0 {
		temp := make([]int, len(*combinations))
		copy(temp, *combinations)
		*ans = append(*ans, temp)
		return
	}

	if len(candidates) == 0 {
		return
	}

	if candidates[0] <= target {
		*combinations = append(*combinations, candidates[0])
		helper_39(candidates, target-candidates[0], combinations, ans)
		*combinations = (*combinations)[:len(*combinations)-1]
	}
	helper_39(candidates[1:], target, combinations, ans)
}

type arg_39 struct {
	candidates []int
	target     int
}

func Run_39() {
	input := []arg_39{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
		{[]int{2}, 1},
	}

	for _, arg := range input {
		fmt.Println(combinationSum(arg.candidates, arg.target))
	}
}
