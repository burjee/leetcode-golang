package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_57 struct {
	input arg_57
	ans   [][]int
}

func Test_57(t *testing.T) {
	cases := []case_57{
		{arg_57{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, [][]int{{1, 5}, {6, 9}}},
		{arg_57{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{arg_57{[][]int{{1, 5}}, []int{0, 0}}, [][]int{{0, 0}, {1, 5}}},
	}

	for _, c := range cases {
		ans := insert(c.input.intervals, c.input.newInterval)
		assert.Equal(t, c.ans, ans)
	}
}
