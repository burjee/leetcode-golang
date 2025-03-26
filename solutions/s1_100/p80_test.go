package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_80 struct {
	input arg_80
	ans   []int
}

func Test_80(t *testing.T) {
	cases := []case_80{
		{arg_80{[]int{1, 1, 1, 2, 2, 3}}, []int{1, 1, 2, 2, 3}},
		{arg_80{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, []int{0, 0, 1, 1, 2, 3, 3}},
		{arg_80{[]int{0, 0, 1, 1, 2, 3, 3}}, []int{0, 0, 1, 1, 2, 3, 3}},
		{arg_80{[]int{1}}, []int{1}},
	}

	for _, c := range cases {
		k := removeDuplicatesII(c.input.nums)
		assert.Equal(t, len(c.ans), k)
		for i := 0; i < k; i += 1 {
			assert.Equal(t, c.ans[i], c.input.nums[i])
		}
	}
}
