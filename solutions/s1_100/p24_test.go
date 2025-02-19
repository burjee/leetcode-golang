package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_24 struct {
	input arg_24
	ans   []int
}

func Test_24(t *testing.T) {
	cases := []case_24{
		{arg_24{NewListNode([]int{1, 2, 3, 4})}, []int{2, 1, 4, 3}},
		{arg_24{NewListNode([]int{})}, []int{}},
		{arg_24{NewListNode([]int{1})}, []int{1}},
		{arg_24{NewListNode([]int{1, 2, 3})}, []int{2, 1, 3}},
	}

	for _, c := range cases {
		ans := swapPairs(c.input.head)
		assert.Equal(t, c.ans, ans.ToArray())
	}
}
