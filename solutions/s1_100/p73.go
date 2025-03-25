package s1_100

import "fmt"

func setZeroes(matrix [][]int) {
	h := false
	v := false
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0

				if i == 0 {
					h = true
				}
				if j == 0 {
					v = true
				}
			}
		}
	}

	for i := 1; i < len(matrix); i += 1 {
		for j := 1; j < len(matrix[0]); j += 1 {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if v {
		for i := range matrix {
			matrix[i][0] = 0
		}

	}
	if h {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}
}

type arg_73 struct {
	matrix [][]int
}

func Run_73() {
	input := []arg_73{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}},
	}

	for _, arg := range input {
		setZeroes(arg.matrix)
		for _, row := range arg.matrix {
			fmt.Printf("%v\n", row)
		}
		fmt.Println()
	}
}
