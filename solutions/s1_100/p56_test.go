package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_56 struct {
	input arg_56
	ans   [][]int
}

func Test_56(t *testing.T) {
	cases := []case_56{
		{arg_56{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{arg_56{[][]int{{1, 4}, {4, 5}}}, [][]int{{1, 5}}},
	}

	for _, c := range cases {
		ans := merge(c.input.intervals)
		assert.Equal(t, c.ans, ans)
	}
}
