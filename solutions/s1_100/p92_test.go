package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_92 struct {
	input arg_92
	ans   []int
}

func Test_92(t *testing.T) {
	cases := []case_92{
		{arg_92{NewListNode([]int{1, 2, 3, 4, 5}), 2, 4}, []int{1, 4, 3, 2, 5}},
		{arg_92{NewListNode([]int{5}), 1, 1}, []int{5}},
	}

	for _, c := range cases {
		ans := reverseBetween(c.input.head, c.input.left, c.input.right)
		assert.Equal(t, c.ans, ans.ToArray())
	}
}
