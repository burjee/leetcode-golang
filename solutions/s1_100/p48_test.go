package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_48 struct {
	input arg_48
	ans   [][]int
}

func Test_48(t *testing.T) {
	cases := []case_48{
		{arg_48{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{arg_48{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}

	for _, c := range cases {
		rotate(c.input.matrix)
		assert.Equal(t, c.ans, c.input.matrix)
	}
}
