package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_1 struct {
	input arg_1
	ans   []int
}

func Test_1(t *testing.T) {
	cases := []case_1{
		{arg_1{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{arg_1{[]int{2, 7, 11, 15}, 26}, []int{2, 3}},
		{arg_1{[]int{2, 7, 11, 15}, 13}, []int{0, 2}},
	}

	for _, c := range cases {
		ans := twoSum(c.input.nums, c.input.target)
		assert.Equal(t, c.ans, ans)
	}
}
