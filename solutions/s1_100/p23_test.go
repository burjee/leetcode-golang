package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_23 struct {
	input arg_23
	ans   []int
}

func Test_23(t *testing.T) {
	cases := []case_23{
		{arg_23{[]*ListNode{NewListNode([]int{1, 4, 5}), NewListNode([]int{1, 3, 4}), NewListNode([]int{2, 6})}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{arg_23{[]*ListNode{}}, []int{}},
		{arg_23{[]*ListNode{NewListNode([]int{})}}, []int{}},
	}

	for _, c := range cases {
		ans := mergeKLists(c.input.lists)
		assert.Equal(t, c.ans, ans.ToArray())
	}
}
