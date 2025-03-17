package s1_100

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	top := 0
	left := 0
	bottom := n - 1
	right := n - 1

	count := 1
	for count <= n*n {
		for i := left; i <= right; i += 1 {
			ans[top][i] = count
			count += 1
		}
		top += 1

		for i := top; i <= bottom; i += 1 {
			ans[i][right] = count
			count += 1
		}
		right -= 1

		for i := right; i >= left; i -= 1 {
			ans[bottom][i] = count
			count += 1
		}
		bottom -= 1

		for i := bottom; i >= top; i -= 1 {
			ans[i][left] = count
			count += 1
		}
		left += 1
	}

	return ans
}

type arg_59 struct {
	n int
}

func Run_59() {
	input := []arg_59{
		{3},
		{1},
		{4},
	}

	for _, arg := range input {
		fmt.Println(generateMatrix(arg.n))
	}
}
