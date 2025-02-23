package s1_100

import (
	"fmt"
)

// Counterclockwise
func rotate(matrix [][]int) {
	length := len(matrix)
	for level := 0; level < length/2; level += 1 {
		for i := level; i < length-1-level; i += 1 {
			offset_level := length - 1 - level
			offset_end := length - 1 - i

			matrix[level][i], matrix[i][offset_level] = matrix[i][offset_level], matrix[level][i]

			matrix[level][i], matrix[offset_end][level] = matrix[offset_end][level], matrix[level][i]

			matrix[offset_end][level], matrix[offset_level][offset_end] = matrix[offset_level][offset_end], matrix[offset_end][level]
		}
	}
}

// Clockwise
// func rotate(matrix [][]int) {
// 	length := len(matrix)
// 	for level := 0; level < length/2; level += 1 {
// 		for i := level; i < length-1-level; i += 1 {
// 			v := i
// 			h := length - 1 - level
// 			temp := matrix[v][h]
// 			matrix[v][h] = matrix[level][i]

// 			v = h
// 			h = length - 1 - i
// 			temp, matrix[v][h] = matrix[v][h], temp

// 			v = h
// 			h = level
// 			temp, matrix[v][h] = matrix[v][h], temp

// 			matrix[level][i] = temp
// 		}
// 	}
// }

type arg_48 struct {
	matrix [][]int
}

func Run_48() {
	input := []arg_48{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}},
	}

	for _, arg := range input {
		rotate(arg.matrix)
		for _, arr := range arg.matrix {
			fmt.Printf("%2v\n", arr)
		}
	}
}
