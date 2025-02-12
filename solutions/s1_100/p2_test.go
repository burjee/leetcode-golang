package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_2 struct {
	input arg_2
	ans   []int
}

func Test_2(t *testing.T) {
	cases := []case_2{
		{arg_2{NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4})}, []int{7, 0, 8}},
		{arg_2{NewListNode([]int{0}), NewListNode([]int{0})}, []int{0}},
		{arg_2{NewListNode([]int{9, 9, 9, 9, 9, 9, 9}), NewListNode([]int{9, 9, 9, 9})}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, c := range cases {
		ans := addTwoNumbers(c.input.l1, c.input.l2)
		assert.Equal(t, c.ans, ans.ToArray())
	}
}
