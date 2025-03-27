package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_88 struct {
	input arg_88
	ans   []int
}

func Test_88(t *testing.T) {
	cases := []case_88{
		{arg_88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
		{arg_88{[]int{1}, 1, []int{}, 0}, []int{1}},
		{arg_88{[]int{0}, 0, []int{1}, 1}, []int{1}},
	}

	for _, c := range cases {
		merge_88(c.input.nums1, c.input.m, c.input.nums2, c.input.n)
		assert.Equal(t, c.ans, c.input.nums1)
	}
}
