package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_26 struct {
	input arg_26
	ans   []int
}

func Test_26(t *testing.T) {
	cases := []case_26{
		{arg_26{[]int{1, 1, 2}}, []int{1, 2}},
		{arg_26{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, []int{0, 1, 2, 3, 4}},
	}

	for _, c := range cases {
		ans := removeDuplicates(c.input.nums)
		assert.Equal(t, len(c.ans), ans)
		for i := range c.ans {
			assert.Equal(t, c.ans[i], c.input.nums[i])
		}
	}
}
