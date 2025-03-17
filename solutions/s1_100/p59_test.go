package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_59 struct {
	input arg_59
	ans   [][]int
}

func Test_59(t *testing.T) {
	cases := []case_59{
		{arg_59{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{arg_59{1}, [][]int{{1}}},
		{arg_59{4}, [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
	}

	for _, c := range cases {
		ans := generateMatrix(c.input.n)
		assert.Equal(t, c.ans, ans)
	}
}
