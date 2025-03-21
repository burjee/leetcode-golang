package s1_100

import (
	"fmt"
)

func climbStairs(n int) int {
	ans := make([]int, n+1)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i <= n; i += 1 {
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans[n]
}

type arg_70 struct {
	n int
}

func Run_70() {
	input := []arg_70{
		{2},
		{3},
	}

	for _, arg := range input {
		fmt.Println(climbStairs(arg.n))
	}
}
