package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_53 struct {
	input arg_53
	ans   int
}

func Test_53(t *testing.T) {
	cases := []case_53{
		{arg_53{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{arg_53{[]int{1}}, 1},
		{arg_53{[]int{5, 4, -1, 7, 8}}, 23},
		{arg_53{[]int{-1}}, -1},
	}

	for _, c := range cases {
		ans := maxSubArray(c.input.nums)
		assert.Equal(t, c.ans, ans)
	}
}
