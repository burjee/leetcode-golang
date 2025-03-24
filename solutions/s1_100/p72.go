package s1_100

import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
	word1 = "^" + word1
	word2 = "^" + word2

	distance := make([][]int, len(word1))
	for i := range distance {
		distance[i] = make([]int, len(word2))
	}

	for i := range distance {
		distance[i][0] = i
	}
	for j := range distance[0] {
		distance[0][j] = j
	}

	for i := 1; i < len(word1); i += 1 {
		for j := 1; j < len(word2); j += 1 {
			replace := 0
			if word1[i] != word2[j] {
				replace = 1
			}

			// distance[i][j] = min( "replace", "insert", "delete" )
			distance[i][j] = min(replace+distance[i-1][j-1], 1+distance[i][j-1], 1+distance[i-1][j])
		}
	}

	return distance[len(word1)-1][len(word2)-1]
}

type arg_72 struct {
	word1 string
	word2 string
}

func Run_72() {
	input := []arg_72{
		{"horse", "ros"},
		{"intention", "execution"},
		{"horse", "rose"},
		{"horse", "horse"},
		{"prosperity", "properties"},
	}

	for _, arg := range input {
		fmt.Println(minDistance(arg.word1, arg.word2))
	}
}
