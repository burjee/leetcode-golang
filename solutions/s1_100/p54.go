package s1_100

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	top := 0
	left := 0
	bottom := len(matrix) - 1
	right := len(matrix[0]) - 1

	ans := make([]int, 0, bottom*right)
	for left <= right || top <= bottom {
		for h := left; h <= right && top <= bottom; h += 1 {
			ans = append(ans, matrix[top][h])
		}
		top += 1

		for v := top; v <= bottom && left <= right; v += 1 {
			ans = append(ans, matrix[v][right])
		}
		right -= 1

		for h := right; h >= left && top <= bottom; h -= 1 {
			ans = append(ans, matrix[bottom][h])
		}
		bottom -= 1

		for v := bottom; v >= top && left <= right; v -= 1 {
			ans = append(ans, matrix[v][left])
		}
		left += 1
	}
	return ans
}

type arg_54 struct {
	matrix [][]int
}

func Run_54() {
	input := []arg_54{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
	}

	for _, arg := range input {
		fmt.Println(spiralOrder(arg.matrix))
	}
}
