package s1_100

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_27 struct {
	input arg_27
	ans   []int
}

func Test_27(t *testing.T) {
	cases := []case_27{
		{arg_27{[]int{3, 2, 2, 3}, 3}, []int{2, 2}},
		{arg_27{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, []int{0, 0, 1, 3, 4}},
		{arg_27{[]int{2}, 2}, []int{}},
		{arg_27{[]int{2}, 3}, []int{2}},
	}

	for _, c := range cases {
		ans := removeElement(c.input.nums, c.input.val)
		c.input.nums = c.input.nums[:ans]
		slices.Sort(c.input.nums)

		assert.Equal(t, len(c.ans), ans)
		for i := range c.ans {
			assert.Equal(t, c.ans[i], c.input.nums[i])
		}
	}
}
