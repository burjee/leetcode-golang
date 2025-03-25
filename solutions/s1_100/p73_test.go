package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_73 struct {
	input arg_73
	ans   [][]int
}

func Test_73(t *testing.T) {
	cases := []case_73{
		{arg_73{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{arg_73{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}

	for _, c := range cases {
		setZeroes(c.input.matrix)
		assert.Equal(t, c.ans, c.input.matrix)
	}
}
