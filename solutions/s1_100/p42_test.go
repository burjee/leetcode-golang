package s1_100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type case_42 struct {
	input arg_42
	ans   int
}

func Test_42(t *testing.T) {
	cases := []case_42{
		{arg_42{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{arg_42{[]int{4, 2, 0, 3, 2, 5}}, 9},
	}

	for _, c := range cases {
		ans := trap(c.input.height)
		assert.Equal(t, c.ans, ans)
	}
}
