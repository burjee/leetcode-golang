package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_19 struct {
	input arg_19
	ans   []int
}

func Test_19(t *testing.T) {
	cases := []case_19{
		{arg_19{NewListNode([]int{1, 2, 3, 4, 5}), 2}, []int{1, 2, 3, 5}},
		{arg_19{NewListNode([]int{1}), 1}, []int{}},
		{arg_19{NewListNode([]int{1, 2}), 1}, []int{1}},
	}

	for _, c := range cases {
		ans := removeNthFromEnd(c.input.head, c.input.n)
		assert.Equal(t, c.ans, ans.ToArray())
	}
}
