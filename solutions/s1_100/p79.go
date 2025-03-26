package s1_100

import "fmt"

func exist(board [][]byte, word string) bool {
	board_count := make([]int, 52)
	for i := range board {
		for j := range board[i] {
			k := ToZeroIndex(board[i][j])
			board_count[k] += 1
		}
	}

	for i := range word {
		k := ToZeroIndex(word[i])
		board_count[k] -= 1
		if board_count[k] < 0 {
			return false
		}
	}

	path := make([][]bool, len(board))
	for i := range path {
		path[i] = make([]bool, len(board[0]))
	}

	for i := range board {
		for j := range board[i] {
			if helper_79(board, path, word, i, j) {
				return true
			}
		}
	}

	return false
}

func helper_79(board [][]byte, path [][]bool, word string, i, j int) bool {
	if board[i][j] != word[0] {
		return false
	}

	if len(word) == 1 {
		return true
	}

	path[i][j] = true
	if i < len(board)-1 && !path[i+1][j] && helper_79(board, path, word[1:], i+1, j) {
		return true
	}
	if i > 0 && !path[i-1][j] && helper_79(board, path, word[1:], i-1, j) {
		return true
	}
	if j < len(board[0])-1 && !path[i][j+1] && helper_79(board, path, word[1:], i, j+1) {
		return true
	}
	if j > 0 && !path[i][j-1] && helper_79(board, path, word[1:], i, j-1) {
		return true
	}
	path[i][j] = false

	return false
}

type arg_79 struct {
	board [][]byte
	word  string
}

func Run_79() {
	input := []arg_79{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"},
		{[][]byte{{'a'}}, "a"},
		{[][]byte{{'a', 'b'}}, "ab"},
	}

	for _, arg := range input {
		fmt.Println(exist(arg.board, arg.word))
	}
}
