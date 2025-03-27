package s1_100

import (
	"fmt"
)

func numTrees(n int) int {
	ans := make([]int, n+1)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i <= n; i += 1 {
		count := 0
		for j := 0; j < i; j += 1 {
			count += ans[j] * ans[i-j-1]
		}
		ans[i] = count
	}
	return ans[n]
}

type arg_96 struct {
	n int
}

func Run_96() {
	input := []arg_96{
		{6},
		{5},
		{4},
		{3},
		{2},
		{1},
	}

	for _, arg := range input {
		fmt.Println(numTrees(arg.n))
	}
}
