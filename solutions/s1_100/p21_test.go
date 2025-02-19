package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_21 struct {
	input arg_21
	ans   []int
}

func Test_21(t *testing.T) {
	cases := []case_21{
		{arg_21{NewListNode([]int{1, 2, 4}), NewListNode([]int{1, 3, 4})}, []int{1, 1, 2, 3, 4, 4}},
		{arg_21{NewListNode([]int{}), NewListNode([]int{})}, []int{}},
		{arg_21{NewListNode([]int{}), NewListNode([]int{0})}, []int{0}},
	}

	for _, c := range cases {
		ans := mergeTwoLists(c.input.list1, c.input.list2)
		assert.Equal(t, c.ans, ans.ToArray())
	}
}
