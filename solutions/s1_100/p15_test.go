package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_15 struct {
	input arg_15
	ans   [][]int
}

func Test_15(t *testing.T) {
	cases := []case_15{
		{arg_15{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{arg_15{[]int{0, 1, 1}}, [][]int{}},
		{arg_15{[]int{0, 0, 0}}, [][]int{{0, 0, 0}}},
	}

	for _, c := range cases {
		ans := threeSum(c.input.nums)
		assert.Equal(t, c.ans, ans)
	}
}
